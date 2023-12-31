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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
	"net/http"
	"net/url"
	"regexp"
)

// HCRPAssignmentReportsServer is a fake server for instances of the armguestconfiguration.HCRPAssignmentReportsClient type.
type HCRPAssignmentReportsServer struct {
	// Get is the fake for method HCRPAssignmentReportsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, reportID string, machineName string, options *armguestconfiguration.HCRPAssignmentReportsClientGetOptions) (resp azfake.Responder[armguestconfiguration.HCRPAssignmentReportsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method HCRPAssignmentReportsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *armguestconfiguration.HCRPAssignmentReportsClientListOptions) (resp azfake.Responder[armguestconfiguration.HCRPAssignmentReportsClientListResponse], errResp azfake.ErrorResponder)
}

// NewHCRPAssignmentReportsServerTransport creates a new instance of HCRPAssignmentReportsServerTransport with the provided implementation.
// The returned HCRPAssignmentReportsServerTransport instance is connected to an instance of armguestconfiguration.HCRPAssignmentReportsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHCRPAssignmentReportsServerTransport(srv *HCRPAssignmentReportsServer) *HCRPAssignmentReportsServerTransport {
	return &HCRPAssignmentReportsServerTransport{srv: srv}
}

// HCRPAssignmentReportsServerTransport connects instances of armguestconfiguration.HCRPAssignmentReportsClient to instances of HCRPAssignmentReportsServer.
// Don't use this type directly, use NewHCRPAssignmentReportsServerTransport instead.
type HCRPAssignmentReportsServerTransport struct {
	srv *HCRPAssignmentReportsServer
}

// Do implements the policy.Transporter interface for HCRPAssignmentReportsServerTransport.
func (h *HCRPAssignmentReportsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HCRPAssignmentReportsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HCRPAssignmentReportsClient.List":
		resp, err = h.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HCRPAssignmentReportsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports/(?P<reportId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	guestConfigurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("guestConfigurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	reportIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportId")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, guestConfigurationAssignmentNameParam, reportIDParam, machineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentReport, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HCRPAssignmentReportsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if h.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	guestConfigurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("guestConfigurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.List(req.Context(), resourceGroupNameParam, guestConfigurationAssignmentNameParam, machineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentReportList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
