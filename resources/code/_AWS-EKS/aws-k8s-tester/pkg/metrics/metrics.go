// Package metrics defines metrics utilities.
package metrics

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"time"

	aws_s3 "github.com/aws/aws-k8s-tester/pkg/aws/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/olekukonko/tablewriter"
	dto "github.com/prometheus/client_model/go"
	"go.uber.org/zap"
)

// RequestsCompare compares two "RequestsSummary".
// Delta is computed with "A" as "before" and with "B" as "after".
type RequestsCompare struct {
	A RequestsSummary `json:"a" read-only:"true"`
	B RequestsSummary `json:"b" read-only:"true"`

	LantencyP50DeltaPercent   float64 `json:"latency-p50-delta-percent" read-only:"true"`
	LantencyP90DeltaPercent   float64 `json:"latency-p90-delta-percent" read-only:"true"`
	LantencyP99DeltaPercent   float64 `json:"latency-p99-delta-percent" read-only:"true"`
	LantencyP999DeltaPercent  float64 `json:"latency-p99.9-delta-percent" read-only:"true"`
	LantencyP9999DeltaPercent float64 `json:"latency-p99.99-delta-percent" read-only:"true"`
}

func (c RequestsCompare) JSON() string {
	b, _ := json.Marshal(c)
	return string(b)
}

func (c RequestsCompare) Table() string {
	buf := bytes.NewBuffer(nil)
	tb := tablewriter.NewWriter(buf)
	tb.SetAutoWrapText(false)
	tb.SetColWidth(1500)
	tb.SetCenterSeparator("*")
	tb.SetAlignment(tablewriter.ALIGN_CENTER)
	tb.SetCaption(true, "(% delta from 'A' to 'B')")
	tb.SetHeader([]string{"Percentile", fmt.Sprintf("A %q", c.A.TestID), fmt.Sprintf("B %q", c.B.TestID), "Delta"})

	tb.Append([]string{"50-pct Latency", c.A.LantencyP50.String(), c.B.LantencyP50.String(), toPercent(c.LantencyP50DeltaPercent)})
	tb.Append([]string{"90-pct Latency", c.A.LantencyP90.String(), c.B.LantencyP90.String(), toPercent(c.LantencyP90DeltaPercent)})
	tb.Append([]string{"99-pct Latency", c.A.LantencyP99.String(), c.B.LantencyP99.String(), toPercent(c.LantencyP99DeltaPercent)})
	tb.Append([]string{"99.9-pct Latency", c.A.LantencyP999.String(), c.B.LantencyP999.String(), toPercent(c.LantencyP999DeltaPercent)})
	tb.Append([]string{"99.99-pct Latency", c.A.LantencyP9999.String(), c.B.LantencyP9999.String(), toPercent(c.LantencyP9999DeltaPercent)})

	tb.Render()
	return buf.String()
}

func toPercent(f float64) string {
	sign := "+"
	if f < 0.0 {
		sign = ""
	}
	return fmt.Sprintf("%s%.3f %%", sign, f)
}

// CompareRequestsSummary compares two "RequestsSummary".
func CompareRequestsSummary(a RequestsSummary, b RequestsSummary) (c RequestsCompare, err error) {
	if len(a.LatencyHistogram) != len(b.LatencyHistogram) {
		return RequestsCompare{}, fmt.Errorf("len(a.LatencyHistogram) %d != len(b.LatencyHistogram) %d", len(a.LatencyHistogram), len(b.LatencyHistogram))
	}

	c = RequestsCompare{
		A: a,
		B: b,
	}

	// e.g. "A" 100, "B" 50 == -50%
	// e.g. "A" 50, "B" 100 == 100%
	deltaP50 := float64(b.LantencyP50) - float64(a.LantencyP50)
	deltaP50 /= float64(a.LantencyP50)
	deltaP50 *= 100.0
	deltaP50 = convertInvalid(deltaP50)

	deltaP90 := float64(b.LantencyP90) - float64(a.LantencyP90)
	deltaP90 /= float64(a.LantencyP90)
	deltaP90 *= 100.0
	deltaP90 = convertInvalid(deltaP90)

	deltaP99 := float64(b.LantencyP99) - float64(a.LantencyP99)
	deltaP99 /= float64(a.LantencyP99)
	deltaP99 *= 100.0
	deltaP99 = convertInvalid(deltaP99)

	deltaP999 := float64(b.LantencyP999) - float64(a.LantencyP999)
	deltaP999 /= float64(a.LantencyP999)
	deltaP999 *= 100.0
	deltaP999 = convertInvalid(deltaP999)

	deltaP9999 := float64(b.LantencyP9999) - float64(a.LantencyP9999)
	deltaP9999 /= float64(a.LantencyP9999)
	deltaP9999 *= 100.0
	deltaP9999 = convertInvalid(deltaP9999)

	c.LantencyP50DeltaPercent = deltaP50
	c.LantencyP90DeltaPercent = deltaP90
	c.LantencyP99DeltaPercent = deltaP99
	c.LantencyP999DeltaPercent = deltaP999
	c.LantencyP9999DeltaPercent = deltaP9999

	return c, nil
}

func convertInvalid(f float64) (v float64) {
	v = f
	if math.IsNaN(f) {
		v = 0
	}
	if math.IsInf(f, 1) {
		v = math.MaxFloat64
	}
	if math.IsInf(f, -1) {
		v = -math.MaxFloat64
	}
	return v
}

// RequestsSummary represents request results.
type RequestsSummary struct {
	// TestID is the test ID.
	TestID string `json:"test-id" read-only:"true"`

	// SuccessTotal is the number of successful client requests.
	SuccessTotal float64 `json:"success-total" read-only:"true"`
	// FailureTotal is the number of failed client requests.
	FailureTotal float64 `json:"failure-total" read-only:"true"`
	// LatencyHistogram is the client requests latency histogram.
	LatencyHistogram HistogramBuckets `json:"latency-histogram" read-only:"true"`

	// LantencyP50 is the 50-percentile latency.
	LantencyP50 time.Duration `json:"latency-p50" read-only:"true"`
	// LantencyP90 is the 90-percentile latency.
	LantencyP90 time.Duration `json:"latency-p90" read-only:"true"`
	// LantencyP99 is the 99-percentile latency.
	LantencyP99 time.Duration `json:"latency-p99" read-only:"true"`
	// LantencyP999 is the 99.9-percentile latency.
	LantencyP999 time.Duration `json:"latency-p99.9" read-only:"true"`
	// LantencyP9999 is the 99.99-percentile latency.
	LantencyP9999 time.Duration `json:"latency-p99.99" read-only:"true"`
}

func (rs RequestsSummary) JSON() string {
	b, _ := json.Marshal(rs)
	return string(b)
}

func (rs RequestsSummary) Table() string {
	return fmt.Sprintf(`
TEST ID: %q

        TOTAL: %.2f
SUCCESS TOTAL: %.2f
FAILURE TOTAL: %.2f

`,
		rs.TestID,
		rs.SuccessTotal+rs.FailureTotal,
		rs.SuccessTotal,
		rs.FailureTotal,
	) +
		rs.LatencyHistogram.Table() +
		fmt.Sprintf(`
   50-percentile Latency: %s
   90-percentile Latency: %s
   99-percentile Latency: %s
 99.9-percentile Latency: %s
99.99-percentile Latency: %s

`,
			rs.LantencyP50,
			rs.LantencyP90,
			rs.LantencyP99,
			rs.LantencyP999,
			rs.LantencyP9999,
		)
}

// DurationWithLabel is the duration with label.
// ref. https://en.wikipedia.org/wiki/Kolmogorov%E2%80%93Smirnov_test
type DurationWithLabel struct {
	time.Duration
	Label string
}

type DurationWithLabels []DurationWithLabel

func (ds DurationWithLabels) Len() int           { return len(ds) }
func (ds DurationWithLabels) Less(i, j int) bool { return ds[i].Duration < ds[j].Duration }
func (ds DurationWithLabels) Swap(i, j int)      { ds[i], ds[j] = ds[j], ds[i] }

// PickLantencyP50 returns the latency assuming durations are already sorted.
func (ds DurationWithLabels) PickLantencyP50() DurationWithLabel {
	n := len(ds)
	if n == 0 {
		return DurationWithLabel{}
	}
	if n == 1 {
		return ds[0]
	}

	idx := n / 2
	return ds[idx]
}

// PickLantencyP90 returns the latency assuming durations are already sorted.
func (ds DurationWithLabels) PickLantencyP90() DurationWithLabel {
	n := len(ds)
	if n == 0 {
		return DurationWithLabel{}
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 90 / 100
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP99 returns the latency assuming durations are already sorted.
func (ds DurationWithLabels) PickLantencyP99() DurationWithLabel {
	n := len(ds)
	if n == 0 {
		return DurationWithLabel{}
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 99 / 100
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP999 returns the latency assuming durations are already sorted.
func (ds DurationWithLabels) PickLantencyP999() DurationWithLabel {
	n := len(ds)
	if n == 0 {
		return DurationWithLabel{}
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 999 / 1000
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP9999 returns the latency assuming durations are already sorted.
func (ds DurationWithLabels) PickLantencyP9999() DurationWithLabel {
	n := len(ds)
	if n == 0 {
		return DurationWithLabel{}
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 9999 / 10000
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

func (ds DurationWithLabels) CSV(path string) error {
	csvFile, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		csvFile, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	csvWriter.Write([]string{"label", "duration-ms"})

	rows := make([][]string, len(ds))
	for idx := range ds {
		rows[idx] = []string{ds[idx].Label, fmt.Sprintf("%d", ds[idx].Milliseconds())}
	}
	return csvWriter.WriteAll(rows)
}

// LabelDurations labels durations.
func LabelDurations(ds1 Durations, label string) (ds2 DurationWithLabels) {
	ds2 = make(DurationWithLabels, len(ds1), len(ds1))
	for idx := range ds1 {
		ds2[idx] = DurationWithLabel{
			Duration: ds1[idx],
			Label:    label,
		}
	}
	return ds2
}

type Durations []time.Duration

func (ds Durations) Len() int           { return len(ds) }
func (ds Durations) Less(i, j int) bool { return ds[i] < ds[j] }
func (ds Durations) Swap(i, j int)      { ds[i], ds[j] = ds[j], ds[i] }

// PickLantencyP50 returns the latency assuming durations are already sorted.
func (ds Durations) PickLantencyP50() time.Duration {
	n := len(ds)
	if n == 0 {
		return time.Duration(0)
	}
	if n == 1 {
		return ds[0]
	}

	idx := n / 2
	return ds[idx]
}

// PickLantencyP90 returns the latency assuming durations are already sorted.
func (ds Durations) PickLantencyP90() time.Duration {
	n := len(ds)
	if n == 0 {
		return time.Duration(0)
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 90 / 100
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP99 returns the latency assuming durations are already sorted.
func (ds Durations) PickLantencyP99() time.Duration {
	n := len(ds)
	if n == 0 {
		return time.Duration(0)
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 99 / 100
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP999 returns the latency assuming durations are already sorted.
func (ds Durations) PickLantencyP999() time.Duration {
	n := len(ds)
	if n == 0 {
		return time.Duration(0)
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 999 / 1000
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// PickLantencyP9999 returns the latency assuming durations are already sorted.
func (ds Durations) PickLantencyP9999() time.Duration {
	n := len(ds)
	if n == 0 {
		return time.Duration(0)
	}
	if n == 1 {
		return ds[0]
	}

	idx := n * 9999 / 10000
	if idx >= n {
		return ds[n-1]
	}
	return ds[idx]
}

// HistogramBucket represents metrics latency bucket.
type HistogramBucket struct {
	Scale      string  `json:"scale"`
	LowerBound float64 `json:"lower-bound"`
	UpperBound float64 `json:"upper-bound"`
	Count      uint64  `json:"count"`
}

func (bucket HistogramBucket) String() string {
	b, _ := json.Marshal(bucket)
	return string(b)
}

type HistogramBuckets []HistogramBucket

func (buckets HistogramBuckets) Len() int { return len(buckets) }

func (buckets HistogramBuckets) Less(i, j int) bool {
	return buckets[i].LowerBound < buckets[j].LowerBound
}

func (buckets HistogramBuckets) Swap(i, j int) {
	t := buckets[i]
	buckets[i] = buckets[j]
	buckets[j] = t
}

// ParseHistogram parses Prometheus histogram.
func ParseHistogram(scale string, histo *dto.Histogram) (buckets HistogramBuckets, err error) {
	if histo == nil {
		return nil, errors.New("nil Histogram")
	}

	total := *histo.SampleCount
	n := len(histo.Bucket)

	buckets = make(HistogramBuckets, n+1)
	buckets[n] = HistogramBucket{
		Scale:      scale,
		UpperBound: math.MaxFloat64,
		Count:      total,
	}
	for idx, bv := range histo.Bucket {
		buckets[idx] = HistogramBucket{
			Scale:      scale,
			UpperBound: *bv.UpperBound,
			Count:      *bv.CumulativeCount,
		}
	}
	for idx := n; idx > 0; idx-- { // start from last, end at second to last
		// convert cumulative count to per-bucket count
		buckets[idx].Count = buckets[idx].Count - buckets[idx-1].Count
		// use previous bucket upper bound as lower bound
		buckets[idx].LowerBound = buckets[idx-1].UpperBound
	}

	sort.Sort(HistogramBuckets(buckets))
	return buckets, nil
}

// MergeHistograms merges two histograms.
func MergeHistograms(a HistogramBuckets, b HistogramBuckets) (HistogramBuckets, error) {
	counts := make(map[HistogramBucket]uint64)
	for _, cur := range a {
		key := HistogramBucket{
			Scale:      cur.Scale,
			LowerBound: cur.LowerBound,
			UpperBound: cur.UpperBound,
		}
		counts[key] = cur.Count
	}
	for _, cur := range b {
		key := HistogramBucket{
			Scale:      cur.Scale,
			LowerBound: cur.LowerBound,
			UpperBound: cur.UpperBound,
		}
		prevCount, ok := counts[key]
		if !ok {
			return nil, fmt.Errorf("[%f, %f] not found in previous counts", cur.LowerBound, cur.UpperBound)
		}
		counts[key] = prevCount + cur.Count
	}
	hs := make(HistogramBuckets, 0, len(counts))
	for k, v := range counts {
		hs = append(hs, HistogramBucket{
			Scale:      k.Scale,
			LowerBound: k.LowerBound,
			UpperBound: k.UpperBound,
			Count:      v,
		})
	}
	sort.Sort(HistogramBuckets(hs))
	return hs, nil
}

// Table converts "HistogramBuckets" to table.
func (buckets HistogramBuckets) Table() string {
	if len(buckets) == 0 {
		return ""
	}
	buf := bytes.NewBuffer(nil)
	tb := tablewriter.NewWriter(buf)
	tb.SetAutoWrapText(false)
	tb.SetColWidth(1500)
	tb.SetCenterSeparator("*")
	tb.SetAlignment(tablewriter.ALIGN_CENTER)
	tb.SetCaption(true, fmt.Sprintf("	(%q scale)", buckets[0].Scale))
	tb.SetHeader([]string{"lower bound", "upper bound", "count"})
	for _, v := range buckets {
		lo := fmt.Sprintf("%f", v.LowerBound)
		if v.Scale == "milliseconds" {
			lo = fmt.Sprintf("%.3f", v.LowerBound)
		}
		hi := fmt.Sprintf("%f", v.UpperBound)
		if v.Scale == "milliseconds" {
			hi = fmt.Sprintf("%.3f", v.UpperBound)
		}
		if v.UpperBound == math.MaxFloat64 {
			hi = "math.MaxFloat64"
		}
		tb.Append([]string{lo, hi, fmt.Sprintf("%d", v.Count)})
	}
	tb.Render()
	return buf.String()
}

// DownloaDurationsRFromS3 downloads the file from S3 bucket, and parses "Durations".
func DownloadDurationsFromS3(lg *zap.Logger, s3API s3iface.S3API, bucketName string, s3Key string) (rs Durations, err error) {
	var localPath string
	localPath, err = aws_s3.DownloadToTempFile(
		lg,
		s3API,
		bucketName,
		s3Key,
	)
	if err != nil {
		return Durations{}, fmt.Errorf("failed to download downloads summary %v", err)
	}
	defer os.RemoveAll(localPath)

	rf, err := os.OpenFile(localPath, os.O_RDONLY, 0444)
	if err != nil {
		lg.Warn("failed to read a file", zap.Error(err))
		return Durations{}, err
	}
	defer rf.Close()

	if err = json.NewDecoder(rf).Decode(&rs); err != nil {
		lg.Warn("failed to decode a JSON file", zap.Error(err))
		return Durations{}, err
	}
	return rs, nil
}

// DownloadRequestsSummaryFromS3 downloads the file from S3 bucket, and parses "RequestsSummary".
func DownloadRequestsSummaryFromS3(lg *zap.Logger, s3API s3iface.S3API, bucketName string, s3Key string) (rs RequestsSummary, err error) {
	var localPath string
	localPath, err = aws_s3.DownloadToTempFile(
		lg,
		s3API,
		bucketName,
		s3Key,
	)
	if err != nil {
		return RequestsSummary{}, fmt.Errorf("failed to download requests summary %v", err)
	}
	defer os.RemoveAll(localPath)

	rf, err := os.OpenFile(localPath, os.O_RDONLY, 0444)
	if err != nil {
		lg.Warn("failed to read a file", zap.Error(err))
		return RequestsSummary{}, err
	}
	defer rf.Close()

	if err = json.NewDecoder(rf).Decode(&rs); err != nil {
		lg.Warn("failed to decode a JSON file", zap.Error(err))
		return RequestsSummary{}, err
	}
	return rs, nil
}
