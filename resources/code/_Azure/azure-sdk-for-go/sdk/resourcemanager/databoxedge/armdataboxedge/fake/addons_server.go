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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
	"net/http"
	"net/url"
	"regexp"
)

// AddonsServer is a fake server for instances of the armdataboxedge.AddonsClient type.
type AddonsServer struct {
	// BeginCreateOrUpdate is the fake for method AddonsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, addon armdataboxedge.AddonClassification, options *armdataboxedge.AddonsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdataboxedge.AddonsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AddonsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *armdataboxedge.AddonsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdataboxedge.AddonsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AddonsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, roleName string, addonName string, resourceGroupName string, options *armdataboxedge.AddonsClientGetOptions) (resp azfake.Responder[armdataboxedge.AddonsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByRolePager is the fake for method AddonsClient.NewListByRolePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByRolePager func(deviceName string, roleName string, resourceGroupName string, options *armdataboxedge.AddonsClientListByRoleOptions) (resp azfake.PagerResponder[armdataboxedge.AddonsClientListByRoleResponse])
}

// NewAddonsServerTransport creates a new instance of AddonsServerTransport with the provided implementation.
// The returned AddonsServerTransport instance is connected to an instance of armdataboxedge.AddonsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAddonsServerTransport(srv *AddonsServer) *AddonsServerTransport {
	return &AddonsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdataboxedge.AddonsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdataboxedge.AddonsClientDeleteResponse]](),
		newListByRolePager:  newTracker[azfake.PagerResponder[armdataboxedge.AddonsClientListByRoleResponse]](),
	}
}

// AddonsServerTransport connects instances of armdataboxedge.AddonsClient to instances of AddonsServer.
// Don't use this type directly, use NewAddonsServerTransport instead.
type AddonsServerTransport struct {
	srv                 *AddonsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdataboxedge.AddonsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdataboxedge.AddonsClientDeleteResponse]]
	newListByRolePager  *tracker[azfake.PagerResponder[armdataboxedge.AddonsClientListByRoleResponse]]
}

// Do implements the policy.Transporter interface for AddonsServerTransport.
func (a *AddonsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AddonsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AddonsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AddonsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AddonsClient.NewListByRolePager":
		resp, err = a.dispatchNewListByRolePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AddonsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles/(?P<roleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addons/(?P<addonName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		raw, err := readRequestBody(req)
		if err != nil {
			return nil, err
		}
		body, err := unmarshalAddonClassification(raw)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		roleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleName")])
		if err != nil {
			return nil, err
		}
		addonNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("addonName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, roleNameParam, addonNameParam, resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AddonsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles/(?P<roleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addons/(?P<addonName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		roleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleName")])
		if err != nil {
			return nil, err
		}
		addonNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("addonName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), deviceNameParam, roleNameParam, addonNameParam, resourceGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AddonsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles/(?P<roleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addons/(?P<addonName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	roleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleName")])
	if err != nil {
		return nil, err
	}
	addonNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("addonName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), deviceNameParam, roleNameParam, addonNameParam, resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AddonClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AddonsServerTransport) dispatchNewListByRolePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByRolePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByRolePager not implemented")}
	}
	newListByRolePager := a.newListByRolePager.get(req)
	if newListByRolePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/roles/(?P<roleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addons`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		roleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("roleName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByRolePager(deviceNameParam, roleNameParam, resourceGroupNameParam, nil)
		newListByRolePager = &resp
		a.newListByRolePager.add(req, newListByRolePager)
		server.PagerResponderInjectNextLinks(newListByRolePager, req, func(page *armdataboxedge.AddonsClientListByRoleResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByRolePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByRolePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByRolePager) {
		a.newListByRolePager.remove(req)
	}
	return resp, nil
}
