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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
	"net/http"
	"net/url"
	"regexp"
)

// FileSharesServer is a fake server for instances of the armstorsimple1200series.FileSharesClient type.
type FileSharesServer struct {
	// BeginCreateOrUpdate is the fake for method FileSharesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, fileShare armstorsimple1200series.FileShare, options *armstorsimple1200series.FileSharesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armstorsimple1200series.FileSharesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FileSharesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientBeginDeleteOptions) (resp azfake.PollerResponder[armstorsimple1200series.FileSharesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FileSharesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientGetOptions) (resp azfake.Responder[armstorsimple1200series.FileSharesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDevicePager is the fake for method FileSharesClient.NewListByDevicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDevicePager func(deviceName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientListByDeviceOptions) (resp azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByDeviceResponse])

	// NewListByFileServerPager is the fake for method FileSharesClient.NewListByFileServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFileServerPager func(deviceName string, fileServerName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientListByFileServerOptions) (resp azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByFileServerResponse])

	// NewListMetricDefinitionPager is the fake for method FileSharesClient.NewListMetricDefinitionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricDefinitionPager func(deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientListMetricDefinitionOptions) (resp azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricDefinitionResponse])

	// NewListMetricsPager is the fake for method FileSharesClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(deviceName string, fileServerName string, shareName string, resourceGroupName string, managerName string, options *armstorsimple1200series.FileSharesClientListMetricsOptions) (resp azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricsResponse])
}

// NewFileSharesServerTransport creates a new instance of FileSharesServerTransport with the provided implementation.
// The returned FileSharesServerTransport instance is connected to an instance of armstorsimple1200series.FileSharesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFileSharesServerTransport(srv *FileSharesServer) *FileSharesServerTransport {
	return &FileSharesServerTransport{
		srv:                          srv,
		beginCreateOrUpdate:          newTracker[azfake.PollerResponder[armstorsimple1200series.FileSharesClientCreateOrUpdateResponse]](),
		beginDelete:                  newTracker[azfake.PollerResponder[armstorsimple1200series.FileSharesClientDeleteResponse]](),
		newListByDevicePager:         newTracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByDeviceResponse]](),
		newListByFileServerPager:     newTracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByFileServerResponse]](),
		newListMetricDefinitionPager: newTracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricDefinitionResponse]](),
		newListMetricsPager:          newTracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricsResponse]](),
	}
}

// FileSharesServerTransport connects instances of armstorsimple1200series.FileSharesClient to instances of FileSharesServer.
// Don't use this type directly, use NewFileSharesServerTransport instead.
type FileSharesServerTransport struct {
	srv                          *FileSharesServer
	beginCreateOrUpdate          *tracker[azfake.PollerResponder[armstorsimple1200series.FileSharesClientCreateOrUpdateResponse]]
	beginDelete                  *tracker[azfake.PollerResponder[armstorsimple1200series.FileSharesClientDeleteResponse]]
	newListByDevicePager         *tracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByDeviceResponse]]
	newListByFileServerPager     *tracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListByFileServerResponse]]
	newListMetricDefinitionPager *tracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricDefinitionResponse]]
	newListMetricsPager          *tracker[azfake.PagerResponder[armstorsimple1200series.FileSharesClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for FileSharesServerTransport.
func (f *FileSharesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FileSharesClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FileSharesClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FileSharesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FileSharesClient.NewListByDevicePager":
		resp, err = f.dispatchNewListByDevicePager(req)
	case "FileSharesClient.NewListByFileServerPager":
		resp, err = f.dispatchNewListByFileServerPager(req)
	case "FileSharesClient.NewListMetricDefinitionPager":
		resp, err = f.dispatchNewListMetricDefinitionPager(req)
	case "FileSharesClient.NewListMetricsPager":
		resp, err = f.dispatchNewListMetricsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FileSharesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple1200series.FileShare](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
		if err != nil {
			return nil, err
		}
		shareNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, fileServerNameParam, shareNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FileSharesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
		if err != nil {
			return nil, err
		}
		shareNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), deviceNameParam, fileServerNameParam, shareNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FileSharesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
	if err != nil {
		return nil, err
	}
	shareNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), deviceNameParam, fileServerNameParam, shareNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileShare, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchNewListByDevicePager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByDevicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDevicePager not implemented")}
	}
	newListByDevicePager := f.newListByDevicePager.get(req)
	if newListByDevicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByDevicePager(deviceNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByDevicePager = &resp
		f.newListByDevicePager.add(req, newListByDevicePager)
	}
	resp, err := server.PagerResponderNext(newListByDevicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByDevicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDevicePager) {
		f.newListByDevicePager.remove(req)
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchNewListByFileServerPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByFileServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFileServerPager not implemented")}
	}
	newListByFileServerPager := f.newListByFileServerPager.get(req)
	if newListByFileServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByFileServerPager(deviceNameParam, fileServerNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByFileServerPager = &resp
		f.newListByFileServerPager.add(req, newListByFileServerPager)
	}
	resp, err := server.PagerResponderNext(newListByFileServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByFileServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFileServerPager) {
		f.newListByFileServerPager.remove(req)
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchNewListMetricDefinitionPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListMetricDefinitionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricDefinitionPager not implemented")}
	}
	newListMetricDefinitionPager := f.newListMetricDefinitionPager.get(req)
	if newListMetricDefinitionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metricsDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
		if err != nil {
			return nil, err
		}
		shareNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListMetricDefinitionPager(deviceNameParam, fileServerNameParam, shareNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListMetricDefinitionPager = &resp
		f.newListMetricDefinitionPager.add(req, newListMetricDefinitionPager)
	}
	resp, err := server.PagerResponderNext(newListMetricDefinitionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListMetricDefinitionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricDefinitionPager) {
		f.newListMetricDefinitionPager.remove(req)
	}
	return resp, nil
}

func (f *FileSharesServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := f.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fileservers/(?P<fileServerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shares/(?P<shareName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		fileServerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileServerName")])
		if err != nil {
			return nil, err
		}
		shareNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armstorsimple1200series.FileSharesClientListMetricsOptions
		if filterParam != nil {
			options = &armstorsimple1200series.FileSharesClientListMetricsOptions{
				Filter: filterParam,
			}
		}
		resp := f.srv.NewListMetricsPager(deviceNameParam, fileServerNameParam, shareNameParam, resourceGroupNameParam, managerNameParam, options)
		newListMetricsPager = &resp
		f.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		f.newListMetricsPager.remove(req)
	}
	return resp, nil
}
