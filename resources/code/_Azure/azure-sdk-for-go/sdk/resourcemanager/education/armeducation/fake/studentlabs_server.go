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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/education/armeducation"
	"net/http"
	"net/url"
	"regexp"
)

// StudentLabsServer is a fake server for instances of the armeducation.StudentLabsClient type.
type StudentLabsServer struct {
	// Get is the fake for method StudentLabsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, studentLabName string, options *armeducation.StudentLabsClientGetOptions) (resp azfake.Responder[armeducation.StudentLabsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListAllPager is the fake for method StudentLabsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armeducation.StudentLabsClientListAllOptions) (resp azfake.PagerResponder[armeducation.StudentLabsClientListAllResponse])
}

// NewStudentLabsServerTransport creates a new instance of StudentLabsServerTransport with the provided implementation.
// The returned StudentLabsServerTransport instance is connected to an instance of armeducation.StudentLabsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStudentLabsServerTransport(srv *StudentLabsServer) *StudentLabsServerTransport {
	return &StudentLabsServerTransport{
		srv:             srv,
		newListAllPager: newTracker[azfake.PagerResponder[armeducation.StudentLabsClientListAllResponse]](),
	}
}

// StudentLabsServerTransport connects instances of armeducation.StudentLabsClient to instances of StudentLabsServer.
// Don't use this type directly, use NewStudentLabsServerTransport instead.
type StudentLabsServerTransport struct {
	srv             *StudentLabsServer
	newListAllPager *tracker[azfake.PagerResponder[armeducation.StudentLabsClientListAllResponse]]
}

// Do implements the policy.Transporter interface for StudentLabsServerTransport.
func (s *StudentLabsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StudentLabsClient.Get":
		resp, err = s.dispatchGet(req)
	case "StudentLabsClient.NewListAllPager":
		resp, err = s.dispatchNewListAllPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StudentLabsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Education/studentLabs/(?P<studentLabName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	studentLabNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("studentLabName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), studentLabNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StudentLabDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StudentLabsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := s.newListAllPager.get(req)
	if newListAllPager == nil {
		resp := s.srv.NewListAllPager(nil)
		newListAllPager = &resp
		s.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armeducation.StudentLabsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		s.newListAllPager.remove(req)
	}
	return resp, nil
}
