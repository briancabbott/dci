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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"net/http"
	"net/url"
	"regexp"
)

// ConfigurationAssignmentsForResourceGroupServer is a fake server for instances of the armmaintenance.ConfigurationAssignmentsForResourceGroupClient type.
type ConfigurationAssignmentsForResourceGroupServer struct {
	// CreateOrUpdate is the fake for method ConfigurationAssignmentsForResourceGroupClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, configurationAssignmentName string, configurationAssignment armmaintenance.ConfigurationAssignment, options *armmaintenance.ConfigurationAssignmentsForResourceGroupClientCreateOrUpdateOptions) (resp azfake.Responder[armmaintenance.ConfigurationAssignmentsForResourceGroupClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ConfigurationAssignmentsForResourceGroupClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, configurationAssignmentName string, options *armmaintenance.ConfigurationAssignmentsForResourceGroupClientDeleteOptions) (resp azfake.Responder[armmaintenance.ConfigurationAssignmentsForResourceGroupClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConfigurationAssignmentsForResourceGroupClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, configurationAssignmentName string, options *armmaintenance.ConfigurationAssignmentsForResourceGroupClientGetOptions) (resp azfake.Responder[armmaintenance.ConfigurationAssignmentsForResourceGroupClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ConfigurationAssignmentsForResourceGroupClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, configurationAssignmentName string, configurationAssignment armmaintenance.ConfigurationAssignment, options *armmaintenance.ConfigurationAssignmentsForResourceGroupClientUpdateOptions) (resp azfake.Responder[armmaintenance.ConfigurationAssignmentsForResourceGroupClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewConfigurationAssignmentsForResourceGroupServerTransport creates a new instance of ConfigurationAssignmentsForResourceGroupServerTransport with the provided implementation.
// The returned ConfigurationAssignmentsForResourceGroupServerTransport instance is connected to an instance of armmaintenance.ConfigurationAssignmentsForResourceGroupClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConfigurationAssignmentsForResourceGroupServerTransport(srv *ConfigurationAssignmentsForResourceGroupServer) *ConfigurationAssignmentsForResourceGroupServerTransport {
	return &ConfigurationAssignmentsForResourceGroupServerTransport{srv: srv}
}

// ConfigurationAssignmentsForResourceGroupServerTransport connects instances of armmaintenance.ConfigurationAssignmentsForResourceGroupClient to instances of ConfigurationAssignmentsForResourceGroupServer.
// Don't use this type directly, use NewConfigurationAssignmentsForResourceGroupServerTransport instead.
type ConfigurationAssignmentsForResourceGroupServerTransport struct {
	srv *ConfigurationAssignmentsForResourceGroupServer
}

// Do implements the policy.Transporter interface for ConfigurationAssignmentsForResourceGroupServerTransport.
func (c *ConfigurationAssignmentsForResourceGroupServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConfigurationAssignmentsForResourceGroupClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConfigurationAssignmentsForResourceGroupClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ConfigurationAssignmentsForResourceGroupClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConfigurationAssignmentsForResourceGroupClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConfigurationAssignmentsForResourceGroupServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/configurationAssignments/(?P<configurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmaintenance.ConfigurationAssignment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, configurationAssignmentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationAssignmentsForResourceGroupServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/configurationAssignments/(?P<configurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, configurationAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationAssignmentsForResourceGroupServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/configurationAssignments/(?P<configurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, configurationAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationAssignmentsForResourceGroupServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Maintenance/configurationAssignments/(?P<configurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmaintenance.ConfigurationAssignment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	configurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, configurationAssignmentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
