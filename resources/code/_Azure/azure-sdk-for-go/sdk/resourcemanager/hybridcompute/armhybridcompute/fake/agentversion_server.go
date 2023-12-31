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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AgentVersionServer is a fake server for instances of the armhybridcompute.AgentVersionClient type.
type AgentVersionServer struct {
	// Get is the fake for method AgentVersionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, osType string, version string, options *armhybridcompute.AgentVersionClientGetOptions) (resp azfake.Responder[armhybridcompute.AgentVersionClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method AgentVersionClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, osType string, options *armhybridcompute.AgentVersionClientListOptions) (resp azfake.Responder[armhybridcompute.AgentVersionClientListResponse], errResp azfake.ErrorResponder)
}

// NewAgentVersionServerTransport creates a new instance of AgentVersionServerTransport with the provided implementation.
// The returned AgentVersionServerTransport instance is connected to an instance of armhybridcompute.AgentVersionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAgentVersionServerTransport(srv *AgentVersionServer) *AgentVersionServerTransport {
	return &AgentVersionServerTransport{srv: srv}
}

// AgentVersionServerTransport connects instances of armhybridcompute.AgentVersionClient to instances of AgentVersionServer.
// Don't use this type directly, use NewAgentVersionServerTransport instead.
type AgentVersionServerTransport struct {
	srv *AgentVersionServer
}

// Do implements the policy.Transporter interface for AgentVersionServerTransport.
func (a *AgentVersionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AgentVersionClient.Get":
		resp, err = a.dispatchGet(req)
	case "AgentVersionClient.List":
		resp, err = a.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AgentVersionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.HybridCompute/osType/(?P<osType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentVersions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	osTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("osType")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), osTypeParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AgentVersionServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if a.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/providers/Microsoft\.HybridCompute/osType/(?P<osType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/agentVersions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	osTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("osType")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.List(req.Context(), osTypeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AgentVersionsList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
