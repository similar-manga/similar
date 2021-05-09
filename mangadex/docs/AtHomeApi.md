# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAtHomeServerChapterId**](AtHomeApi.md#GetAtHomeServerChapterId) | **
Get** /at-home/server/{chapterId} | Get MD@Home node URL

# **GetAtHomeServerChapterId**

> InlineResponse2002 GetAtHomeServerChapterId(ctx, chapterId, optional)
Get MD@Home node URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**chapterId** | [**string**](.md)| Chapter ID |
**optional** | ***AtHomeApiGetAtHomeServerChapterIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AtHomeApiGetAtHomeServerChapterIdOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------

**ssl** | **optional.Bool**| | [default to false]

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

