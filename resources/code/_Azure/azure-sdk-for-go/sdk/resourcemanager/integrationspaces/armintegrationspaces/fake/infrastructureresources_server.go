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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/integrationspaces/armintegrationspaces"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// InfrastructureResourcesServer is a fake server for instances of the armintegrationspaces.InfrastructureResourcesClient type.
type InfrastructureResourcesServer struct {
	// CreateOrUpdate is the fake for method InfrastructureResourcesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, spaceName string, infrastructureResourceName string, resource armintegrationspaces.InfrastructureResource, options *armintegrationspaces.InfrastructureResourcesClientCreateOrUpdateOptions) (resp azfake.Responder[armintegrationspaces.InfrastructureResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method InfrastructureResourcesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, spaceName string, infrastructureResourceName string, options *armintegrationspaces.InfrastructureResourcesClientDeleteOptions) (resp azfake.Responder[armintegrationspaces.InfrastructureResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InfrastructureResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, spaceName string, infrastructureResourceName string, options *armintegrationspaces.InfrastructureResourcesClientGetOptions) (resp azfake.Responder[armintegrationspaces.InfrastructureResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListBySpacePager is the fake for method InfrastructureResourcesClient.NewListBySpacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySpacePager func(resourceGroupName string, spaceName string, options *armintegrationspaces.InfrastructureResourcesClientListBySpaceOptions) (resp azfake.PagerResponder[armintegrationspaces.InfrastructureResourcesClientListBySpaceResponse])

	// Patch is the fake for method InfrastructureResourcesClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceGroupName string, spaceName string, infrastructureResourceName string, properties armintegrationspaces.InfrastructureResourceUpdate, options *armintegrationspaces.InfrastructureResourcesClientPatchOptions) (resp azfake.Responder[armintegrationspaces.InfrastructureResourcesClientPatchResponse], errResp azfake.ErrorResponder)
}

// NewInfrastructureResourcesServerTransport creates a new instance of InfrastructureResourcesServerTransport with the provided implementation.
// The returned InfrastructureResourcesServerTransport instance is connected to an instance of armintegrationspaces.InfrastructureResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInfrastructureResourcesServerTransport(srv *InfrastructureResourcesServer) *InfrastructureResourcesServerTransport {
	return &InfrastructureResourcesServerTransport{
		srv:                 srv,
		newListBySpacePager: newTracker[azfake.PagerResponder[armintegrationspaces.InfrastructureResourcesClientListBySpaceResponse]](),
	}
}

// InfrastructureResourcesServerTransport connects instances of armintegrationspaces.InfrastructureResourcesClient to instances of InfrastructureResourcesServer.
// Don't use this type directly, use NewInfrastructureResourcesServerTransport instead.
type InfrastructureResourcesServerTransport struct {
	srv                 *InfrastructureResourcesServer
	newListBySpacePager *tracker[azfake.PagerResponder[armintegrationspaces.InfrastructureResourcesClientListBySpaceResponse]]
}

// Do implements the policy.Transporter interface for InfrastructureResourcesServerTransport.
func (i *InfrastructureResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InfrastructureResourcesClient.CreateOrUpdate":
		resp, err = i.dispatchCreateOrUpdate(req)
	case "InfrastructureResourcesClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "InfrastructureResourcesClient.Get":
		resp, err = i.dispatchGet(req)
	case "InfrastructureResourcesClient.NewListBySpacePager":
		resp, err = i.dispatchNewListBySpacePager(req)
	case "InfrastructureResourcesClient.Patch":
		resp, err = i.dispatchPatch(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InfrastructureResourcesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/infrastructureResources/(?P<infrastructureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.InfrastructureResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	infrastructureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("infrastructureResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, spaceNameParam, infrastructureResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InfrastructureResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InfrastructureResourcesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/infrastructureResources/(?P<infrastructureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	infrastructureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("infrastructureResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, spaceNameParam, infrastructureResourceNameParam, nil)
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

func (i *InfrastructureResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/infrastructureResources/(?P<infrastructureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	infrastructureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("infrastructureResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, spaceNameParam, infrastructureResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InfrastructureResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InfrastructureResourcesServerTransport) dispatchNewListBySpacePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListBySpacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySpacePager not implemented")}
	}
	newListBySpacePager := i.newListBySpacePager.get(req)
	if newListBySpacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/infrastructureResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
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
		maxpagesizeUnescaped, err := url.QueryUnescape(qp.Get("maxpagesize"))
		if err != nil {
			return nil, err
		}
		maxpagesizeParam, err := parseOptional(maxpagesizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		selectEscaped := qp["select"]
		selectParam := make([]string, len(selectEscaped))
		for i, v := range selectEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			selectParam[i] = u
		}
		expandEscaped := qp["expand"]
		expandParam := make([]string, len(expandEscaped))
		for i, v := range expandEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			expandParam[i] = u
		}
		orderbyEscaped := qp["orderby"]
		orderbyParam := make([]string, len(orderbyEscaped))
		for i, v := range orderbyEscaped {
			u, unescapeErr := url.QueryUnescape(v)
			if unescapeErr != nil {
				return nil, unescapeErr
			}
			orderbyParam[i] = u
		}
		spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
		if err != nil {
			return nil, err
		}
		var options *armintegrationspaces.InfrastructureResourcesClientListBySpaceOptions
		if topParam != nil || skipParam != nil || maxpagesizeParam != nil || filterParam != nil || len(selectParam) > 0 || len(expandParam) > 0 || len(orderbyParam) > 0 {
			options = &armintegrationspaces.InfrastructureResourcesClientListBySpaceOptions{
				Top:         topParam,
				Skip:        skipParam,
				Maxpagesize: maxpagesizeParam,
				Filter:      filterParam,
				Select:      selectParam,
				Expand:      expandParam,
				Orderby:     orderbyParam,
			}
		}
		resp := i.srv.NewListBySpacePager(resourceGroupNameParam, spaceNameParam, options)
		newListBySpacePager = &resp
		i.newListBySpacePager.add(req, newListBySpacePager)
		server.PagerResponderInjectNextLinks(newListBySpacePager, req, func(page *armintegrationspaces.InfrastructureResourcesClientListBySpaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySpacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListBySpacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySpacePager) {
		i.newListBySpacePager.remove(req)
	}
	return resp, nil
}

func (i *InfrastructureResourcesServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if i.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.IntegrationSpaces/spaces/(?P<spaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/infrastructureResources/(?P<infrastructureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armintegrationspaces.InfrastructureResourceUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spaceName")])
	if err != nil {
		return nil, err
	}
	infrastructureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("infrastructureResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Patch(req.Context(), resourceGroupNameParam, spaceNameParam, infrastructureResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InfrastructureResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
