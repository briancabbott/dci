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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// DeviceSecurityGroupsServer is a fake server for instances of the armsecurity.DeviceSecurityGroupsClient type.
type DeviceSecurityGroupsServer struct {
	// CreateOrUpdate is the fake for method DeviceSecurityGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceID string, deviceSecurityGroupName string, deviceSecurityGroup armsecurity.DeviceSecurityGroup, options *armsecurity.DeviceSecurityGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurity.DeviceSecurityGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DeviceSecurityGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceID string, deviceSecurityGroupName string, options *armsecurity.DeviceSecurityGroupsClientDeleteOptions) (resp azfake.Responder[armsecurity.DeviceSecurityGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DeviceSecurityGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceID string, deviceSecurityGroupName string, options *armsecurity.DeviceSecurityGroupsClientGetOptions) (resp azfake.Responder[armsecurity.DeviceSecurityGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeviceSecurityGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceID string, options *armsecurity.DeviceSecurityGroupsClientListOptions) (resp azfake.PagerResponder[armsecurity.DeviceSecurityGroupsClientListResponse])
}

// NewDeviceSecurityGroupsServerTransport creates a new instance of DeviceSecurityGroupsServerTransport with the provided implementation.
// The returned DeviceSecurityGroupsServerTransport instance is connected to an instance of armsecurity.DeviceSecurityGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeviceSecurityGroupsServerTransport(srv *DeviceSecurityGroupsServer) *DeviceSecurityGroupsServerTransport {
	return &DeviceSecurityGroupsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurity.DeviceSecurityGroupsClientListResponse]](),
	}
}

// DeviceSecurityGroupsServerTransport connects instances of armsecurity.DeviceSecurityGroupsClient to instances of DeviceSecurityGroupsServer.
// Don't use this type directly, use NewDeviceSecurityGroupsServerTransport instead.
type DeviceSecurityGroupsServerTransport struct {
	srv          *DeviceSecurityGroupsServer
	newListPager *tracker[azfake.PagerResponder[armsecurity.DeviceSecurityGroupsClientListResponse]]
}

// Do implements the policy.Transporter interface for DeviceSecurityGroupsServerTransport.
func (d *DeviceSecurityGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DeviceSecurityGroupsClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DeviceSecurityGroupsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DeviceSecurityGroupsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DeviceSecurityGroupsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DeviceSecurityGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/deviceSecurityGroups/(?P<deviceSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurity.DeviceSecurityGroup](req)
	if err != nil {
		return nil, err
	}
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	deviceSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceIDParam, deviceSecurityGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeviceSecurityGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSecurityGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/deviceSecurityGroups/(?P<deviceSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	deviceSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceIDParam, deviceSecurityGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSecurityGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/deviceSecurityGroups/(?P<deviceSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	deviceSecurityGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceIDParam, deviceSecurityGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeviceSecurityGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeviceSecurityGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/deviceSecurityGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceIDParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurity.DeviceSecurityGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}
