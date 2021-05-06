# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChapterIdRead**](ChapterApi.md#ChapterIdRead) | **Post** /chapter/{id}/read | Mark Chapter read
[**ChapterIdUnread**](ChapterApi.md#ChapterIdUnread) | **Delete** /chapter/{id}/read | Mark Chapter unread
[**GetChapter**](ChapterApi.md#GetChapter) | **Get** /chapter | Chapter list
[**GetChapterId**](ChapterApi.md#GetChapterId) | **Get** /chapter/{id} | Get Chapter
[**PutChapterId**](ChapterApi.md#PutChapterId) | **Put** /chapter/{id} | Update Chapter

# **ChapterIdRead**
> InlineResponse2001 ChapterIdRead(ctx, id)
Mark Chapter read

Mark chapter as read for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChapterIdUnread**
> InlineResponse2001 ChapterIdUnread(ctx, id)
Mark Chapter unread

Mark chapter as unread for the current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChapter**
> ChapterList GetChapter(ctx, optional)
Chapter list

Chapter list, if you want Chapters for a given Manga, please check at feeds endpoints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChapterApiGetChapterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChapterApiGetChapterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 
 **ids** | [**optional.Interface of []string**](string.md)| Chapter ids (limited to 100 per request) | 
 **title** | **optional.String**|  | 
 **groups** | [**optional.Interface of []string**](string.md)|  | 
 **uploader** | [**optional.Interface of string**](.md)|  | 
 **manga** | [**optional.Interface of string**](.md)|  | 
 **volume** | **optional.String**|  | 
 **chapter** | **optional.String**|  | 
 **translatedLanguage** | **optional.String**|  | 
 **createdAtSince** | **optional.String**|  | 
 **updatedAtSince** | **optional.String**|  | 
 **publishAtSince** | **optional.String**|  | 
 **order** | [**optional.Interface of Order1**](.md)|  | 

### Return type

[**ChapterList**](ChapterList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChapterId**
> ChapterResponse GetChapterId(ctx, id)
Get Chapter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Chapter ID | 

### Return type

[**ChapterResponse**](ChapterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutChapterId**
> ChapterResponse PutChapterId(ctx, id, optional)
Update Chapter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| Chapter ID | 
 **optional** | ***ChapterApiPutChapterIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChapterApiPutChapterIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ChapterEdit**](ChapterEdit.md)| This body is limited to 32kb max per call. | 

### Return type

[**ChapterResponse**](ChapterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

