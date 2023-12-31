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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
	"net/http"
	"net/url"
	"regexp"
)

// ExperimentsServer is a fake server for instances of the armfrontdoor.ExperimentsClient type.
type ExperimentsServer struct {
	// BeginCreateOrUpdate is the fake for method ExperimentsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters armfrontdoor.Experiment, options *armfrontdoor.ExperimentsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armfrontdoor.ExperimentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ExperimentsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *armfrontdoor.ExperimentsClientBeginDeleteOptions) (resp azfake.PollerResponder[armfrontdoor.ExperimentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExperimentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, experimentName string, options *armfrontdoor.ExperimentsClientGetOptions) (resp azfake.Responder[armfrontdoor.ExperimentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProfilePager is the fake for method ExperimentsClient.NewListByProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProfilePager func(resourceGroupName string, profileName string, options *armfrontdoor.ExperimentsClientListByProfileOptions) (resp azfake.PagerResponder[armfrontdoor.ExperimentsClientListByProfileResponse])

	// BeginUpdate is the fake for method ExperimentsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, profileName string, experimentName string, parameters armfrontdoor.ExperimentUpdateModel, options *armfrontdoor.ExperimentsClientBeginUpdateOptions) (resp azfake.PollerResponder[armfrontdoor.ExperimentsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewExperimentsServerTransport creates a new instance of ExperimentsServerTransport with the provided implementation.
// The returned ExperimentsServerTransport instance is connected to an instance of armfrontdoor.ExperimentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExperimentsServerTransport(srv *ExperimentsServer) *ExperimentsServerTransport {
	return &ExperimentsServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientDeleteResponse]](),
		newListByProfilePager: newTracker[azfake.PagerResponder[armfrontdoor.ExperimentsClientListByProfileResponse]](),
		beginUpdate:           newTracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientUpdateResponse]](),
	}
}

// ExperimentsServerTransport connects instances of armfrontdoor.ExperimentsClient to instances of ExperimentsServer.
// Don't use this type directly, use NewExperimentsServerTransport instead.
type ExperimentsServerTransport struct {
	srv                   *ExperimentsServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientDeleteResponse]]
	newListByProfilePager *tracker[azfake.PagerResponder[armfrontdoor.ExperimentsClientListByProfileResponse]]
	beginUpdate           *tracker[azfake.PollerResponder[armfrontdoor.ExperimentsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ExperimentsServerTransport.
func (e *ExperimentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExperimentsClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "ExperimentsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "ExperimentsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExperimentsClient.NewListByProfilePager":
		resp, err = e.dispatchNewListByProfilePager(req)
	case "ExperimentsClient.BeginUpdate":
		resp, err = e.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExperimentsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Experiments/(?P<experimentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.Experiment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		experimentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("experimentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, profileNameParam, experimentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *ExperimentsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Experiments/(?P<experimentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		experimentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("experimentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, experimentNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *ExperimentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Experiments/(?P<experimentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	experimentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("experimentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, experimentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Experiment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExperimentsServerTransport) dispatchNewListByProfilePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProfilePager not implemented")}
	}
	newListByProfilePager := e.newListByProfilePager.get(req)
	if newListByProfilePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Experiments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByProfilePager(resourceGroupNameParam, profileNameParam, nil)
		newListByProfilePager = &resp
		e.newListByProfilePager.add(req, newListByProfilePager)
		server.PagerResponderInjectNextLinks(newListByProfilePager, req, func(page *armfrontdoor.ExperimentsClientListByProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProfilePager) {
		e.newListByProfilePager.remove(req)
	}
	return resp, nil
}

func (e *ExperimentsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := e.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/NetworkExperimentProfiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Experiments/(?P<experimentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armfrontdoor.ExperimentUpdateModel](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		experimentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("experimentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdate(req.Context(), resourceGroupNameParam, profileNameParam, experimentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		e.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		e.beginUpdate.remove(req)
	}

	return resp, nil
}
