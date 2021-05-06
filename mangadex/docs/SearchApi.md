# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthor**](SearchApi.md#GetAuthor) | **Get** /author | Author list
[**GetChapter**](SearchApi.md#GetChapter) | **Get** /chapter | Chapter list
[**GetSearchGroup**](SearchApi.md#GetSearchGroup) | **Get** /group | Scanlation Group list
[**GetSearchManga**](SearchApi.md#GetSearchManga) | **Get** /manga | Manga list

# **GetAuthor**
> AuthorList GetAuthor(ctx, optional)
Author list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiGetAuthorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetAuthorOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 
 **ids** | [**optional.Interface of []string**](string.md)| Author ids (limited to 100 per request) | 
 **name** | **optional.String**|  | 

### Return type

[**AuthorList**](AuthorList.md)

### Authorization

No authorization required

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
 **optional** | ***SearchApiGetChapterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetChapterOpts struct
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

# **GetSearchGroup**
> ScanlationGroupList GetSearchGroup(ctx, optional)
Scanlation Group list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiGetSearchGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetSearchGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 
 **ids** | [**optional.Interface of []string**](string.md)| ScanlationGroup ids (limited to 100 per request) | 
 **name** | **optional.String**|  | 

### Return type

[**ScanlationGroupList**](ScanlationGroupList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearchManga**
> MangaList GetSearchManga(ctx, optional)
Manga list

Search a list of Manga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiGetSearchMangaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetSearchMangaOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 
 **title** | **optional.String**|  | 
 **authors** | [**optional.Interface of []string**](string.md)|  | 
 **artists** | [**optional.Interface of []string**](string.md)|  | 
 **year** | **optional.Int32**|  | 
 **includedTags** | [**optional.Interface of []string**](string.md)|  | 
 **includedTagsMode** | **optional.String**|  | [default to AND]
 **excludedTags** | [**optional.Interface of []string**](string.md)|  | 
 **excludedTagsMode** | **optional.String**|  | [default to OR]
 **status** | [**optional.Interface of []string**](string.md)|  | 
 **originalLanguage** | [**optional.Interface of []string**](string.md)|  | 
 **publicationDemographic** | [**optional.Interface of []string**](string.md)|  | 
 **ids** | [**optional.Interface of []string**](string.md)| Manga ids (limited to 100 per request) | 
 **contentRating** | [**optional.Interface of []string**](string.md)|  | 
 **createdAtSince** | **optional.String**|  | 
 **updatedAtSince** | **optional.String**|  | 
 **order** | [**optional.Interface of Order**](.md)|  | 

### Return type

[**MangaList**](MangaList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

