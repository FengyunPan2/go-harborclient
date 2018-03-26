# \DefaultApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesScanAllPost**](DefaultApi.md#RepositoriesScanAllPost) | **Post** /repositories/scanAll | Scan all images of the registry.


# **RepositoriesScanAllPost**
> RepositoriesScanAllPost($projectId)

Scan all images of the registry.

The server will launch different jobs to scan each image on the regsitry, so this is equivalent to calling  the API to scan the image one by one in background, so there's no way to track the overall status of the \"scan all\" action.  Only system adim has permission to call this API.   


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int32**| When this parm is set only the images under the project identified by the project_id will be scanned. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

