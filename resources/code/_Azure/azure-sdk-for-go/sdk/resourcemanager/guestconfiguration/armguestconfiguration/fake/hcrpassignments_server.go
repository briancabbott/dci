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

// HCRPAssignmentsServer is a fake server for instances of the armguestconfiguration.HCRPAssignmentsClient type.
type HCRPAssignmentsServer struct {
	// CreateOrUpdate is the fake for method HCRPAssignmentsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, guestConfigurationAssignmentName string, resourceGroupName string, machineName string, parameters armguestconfiguration.Assignment, options *armguestconfiguration.HCRPAssignmentsClientCreateOrUpdateOptions) (resp azfake.Responder[armguestconfiguration.HCRPAssignmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HCRPAssignmentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *armguestconfiguration.HCRPAssignmentsClientDeleteOptions) (resp azfake.Responder[armguestconfiguration.HCRPAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HCRPAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, guestConfigurationAssignmentName string, machineName string, options *armguestconfiguration.HCRPAssignmentsClientGetOptions) (resp azfake.Responder[armguestconfiguration.HCRPAssignmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HCRPAssignmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, machineName string, options *armguestconfiguration.HCRPAssignmentsClientListOptions) (resp azfake.PagerResponder[armguestconfiguration.HCRPAssignmentsClientListResponse])
}

// NewHCRPAssignmentsServerTransport creates a new instance of HCRPAssignmentsServerTransport with the provided implementation.
// The returned HCRPAssignmentsServerTransport instance is connected to an instance of armguestconfiguration.HCRPAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHCRPAssignmentsServerTransport(srv *HCRPAssignmentsServer) *HCRPAssignmentsServerTransport {
	return &HCRPAssignmentsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armguestconfiguration.HCRPAssignmentsClientListResponse]](),
	}
}

// HCRPAssignmentsServerTransport connects instances of armguestconfiguration.HCRPAssignmentsClient to instances of HCRPAssignmentsServer.
// Don't use this type directly, use NewHCRPAssignmentsServerTransport instead.
type HCRPAssignmentsServerTransport struct {
	srv          *HCRPAssignmentsServer
	newListPager *tracker[azfake.PagerResponder[armguestconfiguration.HCRPAssignmentsClientListResponse]]
}

// Do implements the policy.Transporter interface for HCRPAssignmentsServerTransport.
func (h *HCRPAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HCRPAssignmentsClient.CreateOrUpdate":
		resp, err = h.dispatchCreateOrUpdate(req)
	case "HCRPAssignmentsClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HCRPAssignmentsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HCRPAssignmentsClient.NewListPager":
		resp, err = h.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HCRPAssignmentsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armguestconfiguration.Assignment](req)
	if err != nil {
		return nil, err
	}
	guestConfigurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("guestConfigurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), guestConfigurationAssignmentNameParam, resourceGroupNameParam, machineNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Assignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HCRPAssignmentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, guestConfigurationAssignmentNameParam, machineNameParam, nil)
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

func (h *HCRPAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, guestConfigurationAssignmentNameParam, machineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Assignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HCRPAssignmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := h.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListPager(resourceGroupNameParam, machineNameParam, nil)
		newListPager = &resp
		h.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		h.newListPager.remove(req)
	}
	return resp, nil
}
