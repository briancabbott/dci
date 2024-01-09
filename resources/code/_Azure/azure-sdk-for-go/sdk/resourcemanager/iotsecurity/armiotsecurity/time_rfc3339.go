//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(`(Z|z|\+|-)(\d+:\d+)*"*$`)

const (
	utcDateTimeJSON = `"2006-01-02T15:04:05.999999999"`
	utcDateTime     = "2006-01-02T15:04:05.999999999"
	dateTimeJSON    = `"` + time.RFC3339Nano + `"`
)

type dateTimeRFC3339 time.Time

func (t dateTimeRFC3339) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalJSON()
}

func (t dateTimeRFC3339) MarshalText() ([]byte, error) {
	tt := time.Time(t)
	return tt.MarshalText()
}

func (t *dateTimeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcDateTimeJSON
	if tzOffsetRegex.Match(data) {
		layout = dateTimeJSON
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) UnmarshalText(data []byte) error {
	layout := utcDateTime
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *dateTimeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = dateTimeRFC3339(p)
	return err
}

func populateDateTimeRFC3339(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateTimeRFC3339)(t)
}

func unpopulateDateTimeRFC3339(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux dateTimeRFC3339
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
