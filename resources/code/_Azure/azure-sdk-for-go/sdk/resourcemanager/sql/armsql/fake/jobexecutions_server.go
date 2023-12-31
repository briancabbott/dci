//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// JobExecutionsServer is a fake server for instances of the armsql.JobExecutionsClient type.
type JobExecutionsServer struct {
	// Cancel is the fake for method JobExecutionsClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *armsql.JobExecutionsClientCancelOptions) (resp azfake.Responder[armsql.JobExecutionsClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method JobExecutionsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, options *armsql.JobExecutionsClientBeginCreateOptions) (resp azfake.PollerResponder[armsql.JobExecutionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method JobExecutionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *armsql.JobExecutionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.JobExecutionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobExecutionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobExecutionID string, options *armsql.JobExecutionsClientGetOptions) (resp azfake.Responder[armsql.JobExecutionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAgentPager is the fake for method JobExecutionsClient.NewListByAgentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAgentPager func(resourceGroupName string, serverName string, jobAgentName string, options *armsql.JobExecutionsClientListByAgentOptions) (resp azfake.PagerResponder[armsql.JobExecutionsClientListByAgentResponse])

	// NewListByJobPager is the fake for method JobExecutionsClient.NewListByJobPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByJobPager func(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *armsql.JobExecutionsClientListByJobOptions) (resp azfake.PagerResponder[armsql.JobExecutionsClientListByJobResponse])
}

// NewJobExecutionsServerTransport creates a new instance of JobExecutionsServerTransport with the provided implementation.
// The returned JobExecutionsServerTransport instance is connected to an instance of armsql.JobExecutionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobExecutionsServerTransport(srv *JobExecutionsServer) *JobExecutionsServerTransport {
	return &JobExecutionsServerTransport{
		srv:                 srv,
		beginCreate:         newTracker[azfake.PollerResponder[armsql.JobExecutionsClientCreateResponse]](),
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armsql.JobExecutionsClientCreateOrUpdateResponse]](),
		newListByAgentPager: newTracker[azfake.PagerResponder[armsql.JobExecutionsClientListByAgentResponse]](),
		newListByJobPager:   newTracker[azfake.PagerResponder[armsql.JobExecutionsClientListByJobResponse]](),
	}
}

// JobExecutionsServerTransport connects instances of armsql.JobExecutionsClient to instances of JobExecutionsServer.
// Don't use this type directly, use NewJobExecutionsServerTransport instead.
type JobExecutionsServerTransport struct {
	srv                 *JobExecutionsServer
	beginCreate         *tracker[azfake.PollerResponder[armsql.JobExecutionsClientCreateResponse]]
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armsql.JobExecutionsClientCreateOrUpdateResponse]]
	newListByAgentPager *tracker[azfake.PagerResponder[armsql.JobExecutionsClientListByAgentResponse]]
	newListByJobPager   *tracker[azfake.PagerResponder[armsql.JobExecutionsClientListByJobResponse]]
}

// Do implements the policy.Transporter interface for JobExecutionsServerTransport.
func (j *JobExecutionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobExecutionsClient.Cancel":
		resp, err = j.dispatchCancel(req)
	case "JobExecutionsClient.BeginCreate":
		resp, err = j.dispatchBeginCreate(req)
	case "JobExecutionsClient.BeginCreateOrUpdate":
		resp, err = j.dispatchBeginCreateOrUpdate(req)
	case "JobExecutionsClient.Get":
		resp, err = j.dispatchGet(req)
	case "JobExecutionsClient.NewListByAgentPager":
		resp, err = j.dispatchNewListByAgentPager(req)
	case "JobExecutionsClient.NewListByJobPager":
		resp, err = j.dispatchNewListByJobPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if j.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	jobExecutionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Cancel(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobExecutionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if j.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := j.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginCreate(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		j.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		j.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		j.beginCreate.remove(req)
	}

	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := j.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		jobExecutionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobExecutionIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		j.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		j.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		j.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	jobExecutionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobExecutionIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobExecution, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchNewListByAgentPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByAgentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAgentPager not implemented")}
	}
	newListByAgentPager := j.newListByAgentPager.get(req)
	if newListByAgentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		createTimeMinUnescaped, err := url.QueryUnescape(qp.Get("createTimeMin"))
		if err != nil {
			return nil, err
		}
		createTimeMinParam, err := parseOptional(createTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		createTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("createTimeMax"))
		if err != nil {
			return nil, err
		}
		createTimeMaxParam, err := parseOptional(createTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMinUnescaped, err := url.QueryUnescape(qp.Get("endTimeMin"))
		if err != nil {
			return nil, err
		}
		endTimeMinParam, err := parseOptional(endTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("endTimeMax"))
		if err != nil {
			return nil, err
		}
		endTimeMaxParam, err := parseOptional(endTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		isActiveUnescaped, err := url.QueryUnescape(qp.Get("isActive"))
		if err != nil {
			return nil, err
		}
		isActiveParam, err := parseOptional(isActiveUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsql.JobExecutionsClientListByAgentOptions
		if createTimeMinParam != nil || createTimeMaxParam != nil || endTimeMinParam != nil || endTimeMaxParam != nil || isActiveParam != nil || skipParam != nil || topParam != nil {
			options = &armsql.JobExecutionsClientListByAgentOptions{
				CreateTimeMin: createTimeMinParam,
				CreateTimeMax: createTimeMaxParam,
				EndTimeMin:    endTimeMinParam,
				EndTimeMax:    endTimeMaxParam,
				IsActive:      isActiveParam,
				Skip:          skipParam,
				Top:           topParam,
			}
		}
		resp := j.srv.NewListByAgentPager(resourceGroupNameParam, serverNameParam, jobAgentNameParam, options)
		newListByAgentPager = &resp
		j.newListByAgentPager.add(req, newListByAgentPager)
		server.PagerResponderInjectNextLinks(newListByAgentPager, req, func(page *armsql.JobExecutionsClientListByAgentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAgentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByAgentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAgentPager) {
		j.newListByAgentPager.remove(req)
	}
	return resp, nil
}

func (j *JobExecutionsServerTransport) dispatchNewListByJobPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByJobPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByJobPager not implemented")}
	}
	newListByJobPager := j.newListByJobPager.get(req)
	if newListByJobPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		createTimeMinUnescaped, err := url.QueryUnescape(qp.Get("createTimeMin"))
		if err != nil {
			return nil, err
		}
		createTimeMinParam, err := parseOptional(createTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		createTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("createTimeMax"))
		if err != nil {
			return nil, err
		}
		createTimeMaxParam, err := parseOptional(createTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMinUnescaped, err := url.QueryUnescape(qp.Get("endTimeMin"))
		if err != nil {
			return nil, err
		}
		endTimeMinParam, err := parseOptional(endTimeMinUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		endTimeMaxUnescaped, err := url.QueryUnescape(qp.Get("endTimeMax"))
		if err != nil {
			return nil, err
		}
		endTimeMaxParam, err := parseOptional(endTimeMaxUnescaped, func(v string) (time.Time, error) { return time.Parse(time.RFC3339Nano, v) })
		if err != nil {
			return nil, err
		}
		isActiveUnescaped, err := url.QueryUnescape(qp.Get("isActive"))
		if err != nil {
			return nil, err
		}
		isActiveParam, err := parseOptional(isActiveUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsql.JobExecutionsClientListByJobOptions
		if createTimeMinParam != nil || createTimeMaxParam != nil || endTimeMinParam != nil || endTimeMaxParam != nil || isActiveParam != nil || skipParam != nil || topParam != nil {
			options = &armsql.JobExecutionsClientListByJobOptions{
				CreateTimeMin: createTimeMinParam,
				CreateTimeMax: createTimeMaxParam,
				EndTimeMin:    endTimeMinParam,
				EndTimeMax:    endTimeMaxParam,
				IsActive:      isActiveParam,
				Skip:          skipParam,
				Top:           topParam,
			}
		}
		resp := j.srv.NewListByJobPager(resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, options)
		newListByJobPager = &resp
		j.newListByJobPager.add(req, newListByJobPager)
		server.PagerResponderInjectNextLinks(newListByJobPager, req, func(page *armsql.JobExecutionsClientListByJobResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByJobPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByJobPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByJobPager) {
		j.newListByJobPager.remove(req)
	}
	return resp, nil
}
