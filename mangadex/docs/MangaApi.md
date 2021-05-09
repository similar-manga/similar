# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteChapterId**](MangaApi.md#DeleteChapterId) | **Delete** /chapter/{id} | Delete Chapter
[**DeleteMangaId**](MangaApi.md#DeleteMangaId) | **Delete** /manga/{id} | Delete Manga
[**DeleteMangaIdFollow**](MangaApi.md#DeleteMangaIdFollow) | **Delete** /manga/{id}/follow | Unfollow Manga
[**DeleteMangaIdListListId**](MangaApi.md#DeleteMangaIdListListId) | **
Delete** /manga/{id}/list/{listId} | Remove Manga in CustomList
[**GetMangaChapterReadmarkers**](MangaApi.md#GetMangaChapterReadmarkers) | **Get** /manga/{id}/read | Manga read markers
[**GetMangaId**](MangaApi.md#GetMangaId) | **Get** /manga/{id} | View Manga
[**GetMangaIdFeed**](MangaApi.md#GetMangaIdFeed) | **Get** /manga/{id}/feed | Manga feed
[**GetMangaRandom**](MangaApi.md#GetMangaRandom) | **Get** /manga/random | Get a random Manga
[**GetMangaStatus**](MangaApi.md#GetMangaStatus) | **Get** /manga/status | Get all Manga reading status for logged User
[**GetMangaTag**](MangaApi.md#GetMangaTag) | **Get** /manga/tag | Tag list
[**GetSearchManga**](MangaApi.md#GetSearchManga) | **Get** /manga | Manga list
[**GetUserFollowsManga**](MangaApi.md#GetUserFollowsManga) | **
Get** /user/follows/manga | Get logged User followed Manga list
[**GetUserFollowsMangaFeed**](MangaApi.md#GetUserFollowsMangaFeed) | **
Get** /user/follows/manga/feed | Get logged User followed Manga feed
[**PostManga**](MangaApi.md#PostManga) | **Post** /manga | Create Manga
[**PostMangaIdFollow**](MangaApi.md#PostMangaIdFollow) | **Post** /manga/{id}/follow | Follow Manga
[**PostMangaIdListListId**](MangaApi.md#PostMangaIdListListId) | **
Post** /manga/{id}/list/{listId} | Add Manga in CustomList
[**PostMangaIdStatus**](MangaApi.md#PostMangaIdStatus) | **Post** /manga/{id}/status | Update Manga reading status
[**PutMangaId**](MangaApi.md#PutMangaId) | **Put** /manga/{id} | Update Manga

# **DeleteChapterId**

> Response DeleteChapterId(ctx, id)
Delete Chapter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Chapter ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMangaId**

> Response DeleteMangaId(ctx, id)
Delete Manga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMangaIdFollow**

> Response DeleteMangaIdFollow(ctx, id)
Unfollow Manga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)|  |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMangaIdListListId**

> Response DeleteMangaIdListListId(ctx, id, listId)
Remove Manga in CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |
**listId** | [**string**](.md)| CustomList ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaChapterReadmarkers**

> InlineResponse200 GetMangaChapterReadmarkers(ctx, id)
Manga read markers

A list of chapter ids that are marked as read for the specified manga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)|  |

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaId**

> MangaResponse GetMangaId(ctx, id)
View Manga

View Manga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |

### Return type

[**MangaResponse**](MangaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaIdFeed**

> ChapterList GetMangaIdFeed(ctx, id, optional)
Manga feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |
**optional** | ***MangaApiGetMangaIdFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiGetMangaIdFeedOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------

**limit** | **optional.Int32**| | [default to 100]
**offset** | **optional.Int32**| |
**locales** | [**optional.Interface of []string**](string.md)| |
**createdAtSince** | **optional.String**| |
**updatedAtSince** | **optional.String**| |
**publishAtSince** | **optional.String**| |

### Return type

[**ChapterList**](ChapterList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaRandom**

> MangaResponse GetMangaRandom(ctx, )
Get a random Manga

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**MangaResponse**](MangaResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaStatus**

> InlineResponse2003 GetMangaStatus(ctx, )
Get all Manga reading status for logged User

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMangaTag**

> []TagResponse GetMangaTag(ctx, )
Tag list

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TagResponse**](TagResponse.md)

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
**optional** | ***MangaApiGetSearchMangaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiGetSearchMangaOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 10]
**offset** | **optional.Int32**| |
**title** | **optional.String**| |
**authors** | [**optional.Interface of []string**](string.md)| |
**artists** | [**optional.Interface of []string**](string.md)| |
**year** | **optional.Int32**| |
**includedTags** | [**optional.Interface of []string**](string.md)| |
**includedTagsMode** | **optional.String**| | [default to AND]
**excludedTags** | [**optional.Interface of []string**](string.md)| |
**excludedTagsMode** | **optional.String**| | [default to OR]
**status** | [**optional.Interface of []string**](string.md)| |
**originalLanguage** | [**optional.Interface of []string**](string.md)| |
**publicationDemographic** | [**optional.Interface of []string**](string.md)| |
**ids** | [**optional.Interface of []string**](string.md)| Manga ids (limited to 100 per request) |
**contentRating** | [**optional.Interface of []string**](string.md)| |
**createdAtSince** | **optional.String**| |
**updatedAtSince** | **optional.String**| |
**order** | [**optional.Interface of Order**](.md)| |

### Return type

[**MangaList**](MangaList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserFollowsManga**

> MangaList GetUserFollowsManga(ctx, optional)
Get logged User followed Manga list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***MangaApiGetUserFollowsMangaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiGetUserFollowsMangaOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 10]
**offset** | **optional.Int32**| |

### Return type

[**MangaList**](MangaList.md)

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
**optional** | ***MangaApiGetUserFollowsMangaFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiGetUserFollowsMangaFeedOpts struct Name | Type |
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

# **PostManga**

> MangaResponse PostManga(ctx, optional)
Create Manga

Create a new Manga.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***MangaApiPostMangaOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiPostMangaOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of MangaCreate**](MangaCreate.md)| This body is limited to 16kb max per call. |

### Return type

[**MangaResponse**](MangaResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMangaIdFollow**

> Response PostMangaIdFollow(ctx, id)
Follow Manga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)|  |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMangaIdListListId**

> Response PostMangaIdListListId(ctx, id, listId)
Add Manga in CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |
**listId** | [**string**](.md)| CustomList ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostMangaIdStatus**

> Response PostMangaIdStatus(ctx, id, optional)
Update Manga reading status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)|  |
**optional** | ***MangaApiPostMangaIdStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiPostMangaIdStatusOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of UpdateMangaStatus**](UpdateMangaStatus.md)| Using a &#x60;null&#x60; value in
&#x60;status&#x60; field will remove the Manga reading status. This body is limited to 2kb max per call. |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutMangaId**

> MangaResponse PutMangaId(ctx, id, optional)
Update Manga

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Manga ID |
**optional** | ***MangaApiPutMangaIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MangaApiPutMangaIdOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of MangaEdit**](MangaEdit.md)| This body is limited to 16kb max per call. |

### Return type

[**MangaResponse**](MangaResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

