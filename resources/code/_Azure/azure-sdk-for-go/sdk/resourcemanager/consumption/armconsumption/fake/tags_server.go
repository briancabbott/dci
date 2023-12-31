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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"net/http"
	"net/url"
	"regexp"
)

// TagsServer is a fake server for instances of the armconsumption.TagsClient type.
type TagsServer struct {
	// Get is the fake for method TagsClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Get func(ctx context.Context, scope string, options *armconsumption.TagsClientGetOptions) (resp azfake.Responder[armconsumption.TagsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewTagsServerTransport creates a new instance of TagsServerTransport with the provided implementation.
// The returned TagsServerTransport instance is connected to an instance of armconsumption.TagsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTagsServerTransport(srv *TagsServer) *TagsServerTransport {
	return &TagsServerTransport{srv: srv}
}

// TagsServerTransport connects instances of armconsumption.TagsClient to instances of TagsServer.
// Don't use this type directly, use NewTagsServerTransport instead.
type TagsServerTransport struct {
	srv *TagsServer
}

// Do implements the policy.Transporter interface for TagsServerTransport.
func (t *TagsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TagsClient.Get":
		resp, err = t.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TagsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Consumption/tags`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), scopeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagsResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
