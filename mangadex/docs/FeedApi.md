# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetListIdFeed**](FeedApi.md#GetListIdFeed) | **Get** /list/{id}/feed | CustomList Manga feed
[**GetUserFollowsMangaFeed**](FeedApi.md#GetUserFollowsMangaFeed) | **
Get** /user/follows/manga/feed | Get logged User followed Manga feed

# **GetListIdFeed**

> ChapterList GetListIdFeed(ctx, id, optional)
CustomList Manga feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)|  |
**optional** | ***FeedApiGetListIdFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FeedApiGetListIdFeedOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**limit** | **optional.Int32**| | [default to 100]
**offset** | **optional.Int32**| |
**locales** | [**optional.Interface of []string**](string.md)| |
**createdAtSince** | **optional.String**| |
**updatedAtSince** | **optional.String**| |
**publishAtSince** | **optional.String**| |

### Return type

[**ChapterList**](ChapterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserFollowsMangaFeed**

> ChapterList GetUserFollowsMangaFeed(ctx, optional)
Get logged User followed Manga feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***FeedApiGetUserFollowsMangaFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FeedApiGetUserFollowsMangaFeedOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 100]
**offset** | **optional.Int32**| |
**locales** | [**optional.Interface of []string**](string.md)| |
**createdAtSince** | **optional.String**| |
**updatedAtSince** | **optional.String**| |
**publishAtSince** | **optional.String**| |

### Return type

[**ChapterList**](ChapterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

