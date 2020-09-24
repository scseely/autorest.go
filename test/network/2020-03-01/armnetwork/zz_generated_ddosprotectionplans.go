// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DdosProtectionPlansOperations contains the methods for the DdosProtectionPlans group.
type DdosProtectionPlansOperations interface {
	// BeginCreateOrUpdate - Creates or updates a DDoS protection plan.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansCreateOrUpdateOptions) (*DdosProtectionPlanPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (DdosProtectionPlanPoller, error)
	// BeginDelete - Deletes the specified DDoS protection plan.
	BeginDelete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansDeleteOptions) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified DDoS protection plan.
	Get(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansGetOptions) (*DdosProtectionPlanResponse, error)
	// List - Gets all DDoS protection plans in a subscription.
	List(options *DdosProtectionPlansListOptions) DdosProtectionPlanListResultPager
	// ListByResourceGroup - Gets all the DDoS protection plans in a resource group.
	ListByResourceGroup(resourceGroupName string, options *DdosProtectionPlansListByResourceGroupOptions) DdosProtectionPlanListResultPager
	// UpdateTags - Update a DDoS protection plan tags.
	UpdateTags(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansUpdateTagsOptions) (*DdosProtectionPlanResponse, error)
}

// DdosProtectionPlansClient implements the DdosProtectionPlansOperations interface.
// Don't use this type directly, use NewDdosProtectionPlansClient() instead.
type DdosProtectionPlansClient struct {
	*Client
	subscriptionID string
}

// NewDdosProtectionPlansClient creates a new instance of DdosProtectionPlansClient with the specified values.
func NewDdosProtectionPlansClient(c *Client, subscriptionID string) DdosProtectionPlansOperations {
	return &DdosProtectionPlansClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DdosProtectionPlansClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

func (client *DdosProtectionPlansClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansCreateOrUpdateOptions) (*DdosProtectionPlanPollerResponse, error) {
	resp, err := client.CreateOrUpdate(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return nil, err
	}
	result := &DdosProtectionPlanPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DdosProtectionPlansClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &ddosProtectionPlanPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DdosProtectionPlanResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *DdosProtectionPlansClient) ResumeCreateOrUpdate(token string) (DdosProtectionPlanPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DdosProtectionPlansClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &ddosProtectionPlanPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates a DDoS protection plan.
func (client *DdosProtectionPlansClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return resp, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DdosProtectionPlansClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters DdosProtectionPlan, options *DdosProtectionPlansCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DdosProtectionPlansClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*DdosProtectionPlanResponse, error) {
	result := DdosProtectionPlanResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosProtectionPlan)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DdosProtectionPlansClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

func (client *DdosProtectionPlansClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansDeleteOptions) (*HTTPPollerResponse, error) {
	resp, err := client.Delete(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return nil, err
	}
	result := &HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("DdosProtectionPlansClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *DdosProtectionPlansClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DdosProtectionPlansClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified DDoS protection plan.
func (client *DdosProtectionPlansClient) Delete(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansDeleteOptions) (*azcore.Response, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return resp, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *DdosProtectionPlansClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleError handles the Delete error response.
func (client *DdosProtectionPlansClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets information about the specified DDoS protection plan.
func (client *DdosProtectionPlansClient) Get(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansGetOptions) (*DdosProtectionPlanResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *DdosProtectionPlansClient) GetCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, options *DdosProtectionPlansGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *DdosProtectionPlansClient) GetHandleResponse(resp *azcore.Response) (*DdosProtectionPlanResponse, error) {
	result := DdosProtectionPlanResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosProtectionPlan)
}

// GetHandleError handles the Get error response.
func (client *DdosProtectionPlansClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all DDoS protection plans in a subscription.
func (client *DdosProtectionPlansClient) List(options *DdosProtectionPlansListOptions) DdosProtectionPlanListResultPager {
	return &ddosProtectionPlanListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, options)
		},
		responder: client.ListHandleResponse,
		errorer:   client.ListHandleError,
		advancer: func(ctx context.Context, resp *DdosProtectionPlanListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DdosProtectionPlanListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *DdosProtectionPlansClient) ListCreateRequest(ctx context.Context, options *DdosProtectionPlansListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ddosProtectionPlans"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *DdosProtectionPlansClient) ListHandleResponse(resp *azcore.Response) (*DdosProtectionPlanListResultResponse, error) {
	result := DdosProtectionPlanListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosProtectionPlanListResult)
}

// ListHandleError handles the List error response.
func (client *DdosProtectionPlansClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all the DDoS protection plans in a resource group.
func (client *DdosProtectionPlansClient) ListByResourceGroup(resourceGroupName string, options *DdosProtectionPlansListByResourceGroupOptions) DdosProtectionPlanListResultPager {
	return &ddosProtectionPlanListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.ListByResourceGroupHandleResponse,
		errorer:   client.ListByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp *DdosProtectionPlanListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DdosProtectionPlanListResult.NextLink)
		},
	}
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DdosProtectionPlansClient) ListByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DdosProtectionPlansListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DdosProtectionPlansClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*DdosProtectionPlanListResultResponse, error) {
	result := DdosProtectionPlanListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosProtectionPlanListResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DdosProtectionPlansClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Update a DDoS protection plan tags.
func (client *DdosProtectionPlansClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansUpdateTagsOptions) (*DdosProtectionPlanResponse, error) {
	req, err := client.UpdateTagsCreateRequest(ctx, resourceGroupName, ddosProtectionPlanName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *DdosProtectionPlansClient) UpdateTagsCreateRequest(ctx context.Context, resourceGroupName string, ddosProtectionPlanName string, parameters TagsObject, options *DdosProtectionPlansUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosProtectionPlans/{ddosProtectionPlanName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosProtectionPlanName}", url.PathEscape(ddosProtectionPlanName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *DdosProtectionPlansClient) UpdateTagsHandleResponse(resp *azcore.Response) (*DdosProtectionPlanResponse, error) {
	result := DdosProtectionPlanResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosProtectionPlan)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *DdosProtectionPlansClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}