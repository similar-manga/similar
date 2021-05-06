# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLegacyMapping**](LegacyApi.md#PostLegacyMapping) | **Post** /legacy/mapping | Legacy ID mapping

# **PostLegacyMapping**
> []MappingIdResponse PostLegacyMapping(ctx, optional)
Legacy ID mapping

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LegacyApiPostLegacyMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LegacyApiPostLegacyMappingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of MappingIdBody**](MappingIdBody.md)| This body is limited to 10kb max per call. | 

### Return type

[**[]MappingIdResponse**](MappingIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

