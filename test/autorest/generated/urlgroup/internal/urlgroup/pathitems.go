// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

type PathItemsOperations struct{}

// GetAllWithValuesCreateRequest creates the GetAllWithValues request.
func (PathItemsOperations) GetAllWithValuesCreateRequest(u url.URL, pathItemStringPath string, globalStringPath string, localStringPath string, globalStringQuery *string, options *PathItemsGetAllWithValuesOptions) (*azcore.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/pathItemStringQuery/localStringQuery"
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(globalStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	if options != nil && options.PathItemStringQuery != nil {
		query.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if globalStringQuery != nil {
		query.Set("globalStringQuery", *globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		query.Set("localStringQuery", *options.LocalStringQuery)
	}
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetAllWithValuesHandleResponse handles the GetAllWithValues response.
func (PathItemsOperations) GetAllWithValuesHandleResponse(resp *azcore.Response) (*PathItemsGetAllWithValuesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathItemsGetAllWithValuesResponse{RawResponse: resp.Response}, nil
}

// GetGlobalAndLocalQueryNullCreateRequest creates the GetGlobalAndLocalQueryNull request.
func (PathItemsOperations) GetGlobalAndLocalQueryNullCreateRequest(u url.URL, pathItemStringPath string, globalStringPath string, localStringPath string, globalStringQuery *string, options *PathItemsGetGlobalAndLocalQueryNullOptions) (*azcore.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/null"
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(globalStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	if options != nil && options.PathItemStringQuery != nil {
		query.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if globalStringQuery != nil {
		query.Set("globalStringQuery", *globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		query.Set("localStringQuery", *options.LocalStringQuery)
	}
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetGlobalAndLocalQueryNullHandleResponse handles the GetGlobalAndLocalQueryNull response.
func (PathItemsOperations) GetGlobalAndLocalQueryNullHandleResponse(resp *azcore.Response) (*PathItemsGetGlobalAndLocalQueryNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathItemsGetGlobalAndLocalQueryNullResponse{RawResponse: resp.Response}, nil
}

// GetGlobalQueryNullCreateRequest creates the GetGlobalQueryNull request.
func (PathItemsOperations) GetGlobalQueryNullCreateRequest(u url.URL, pathItemStringPath string, globalStringPath string, localStringPath string, globalStringQuery *string, options *PathItemsGetGlobalQueryNullOptions) (*azcore.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/null/pathItemStringQuery/localStringQuery"
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(globalStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	if options != nil && options.PathItemStringQuery != nil {
		query.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if globalStringQuery != nil {
		query.Set("globalStringQuery", *globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		query.Set("localStringQuery", *options.LocalStringQuery)
	}
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetGlobalQueryNullHandleResponse handles the GetGlobalQueryNull response.
func (PathItemsOperations) GetGlobalQueryNullHandleResponse(resp *azcore.Response) (*PathItemsGetGlobalQueryNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathItemsGetGlobalQueryNullResponse{RawResponse: resp.Response}, nil
}

// GetLocalPathItemQueryNullCreateRequest creates the GetLocalPathItemQueryNull request.
func (PathItemsOperations) GetLocalPathItemQueryNullCreateRequest(u url.URL, pathItemStringPath string, globalStringPath string, localStringPath string, globalStringQuery *string, options *PathItemsGetLocalPathItemQueryNullOptions) (*azcore.Request, error) {
	urlPath := "/pathitem/nullable/globalStringPath/{globalStringPath}/pathItemStringPath/{pathItemStringPath}/localStringPath/{localStringPath}/globalStringQuery/null/null"
	urlPath = strings.ReplaceAll(urlPath, "{pathItemStringPath}", url.PathEscape(pathItemStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{globalStringPath}", url.PathEscape(globalStringPath))
	urlPath = strings.ReplaceAll(urlPath, "{localStringPath}", url.PathEscape(localStringPath))
	u.Path = path.Join(u.Path, urlPath)
	query := u.Query()
	if options != nil && options.PathItemStringQuery != nil {
		query.Set("pathItemStringQuery", *options.PathItemStringQuery)
	}
	if globalStringQuery != nil {
		query.Set("globalStringQuery", *globalStringQuery)
	}
	if options != nil && options.LocalStringQuery != nil {
		query.Set("localStringQuery", *options.LocalStringQuery)
	}
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLocalPathItemQueryNullHandleResponse handles the GetLocalPathItemQueryNull response.
func (PathItemsOperations) GetLocalPathItemQueryNullHandleResponse(resp *azcore.Response) (*PathItemsGetLocalPathItemQueryNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathItemsGetLocalPathItemQueryNullResponse{RawResponse: resp.Response}, nil
}