# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteListId**](CustomListApi.md#DeleteListId) | **Delete** /list/{id} | Delete CustomList
[**DeleteMangaIdListListId**](CustomListApi.md#DeleteMangaIdListListId) | **Delete** /manga/{id}/list/{listId} | Remove Manga in CustomList
[**GetListId**](CustomListApi.md#GetListId) | **Get** /list/{id} | Get CustomList
[**GetListIdFeed**](CustomListApi.md#GetListIdFeed) | **Get** /list/{id}/feed | CustomList Manga feed
[**GetUserIdList**](CustomListApi.md#GetUserIdList) | **Get** /user/{id}/list | Get User&#x27;s CustomList list
[**GetUserList**](CustomListApi.md#GetUserList) | **Get** /user/list | Get logged User CustomList list
[**PostList**](CustomListApi.md#PostList) | **Post** /list | Create CustomList
[**PostMangaIdListListId**](CustomListApi.md#PostMangaIdListListId) | **Post** /manga/{id}/list/{listId} | Add Manga in CustomList
[**PutListId**](CustomListApi.md#PutListId) | **Put** /list/{id} | Update CustomList

# **DeleteListId**
> Response DeleteListId(ctx, id)
Delete CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| CustomList ID | 

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

# **GetListId**
> CustomListResponse GetListId(ctx, id)
Get CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| CustomList ID | 

### Return type

[**CustomListResponse**](CustomListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListIdFeed**
> ChapterList GetListIdFeed(ctx, id, optional)
CustomList Manga feed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)|  | 
 **optional** | ***CustomListApiGetListIdFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomListApiGetListIdFeedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 100]
 **offset** | **optional.Int32**|  | 
 **locales** | [**optional.Interface of []string**](string.md)|  | 
 **createdAtSince** | **optional.String**|  | 
 **updatedAtSince** | **optional.String**|  | 
 **publishAtSince** | **optional.String**|  | 

### Return type

[**ChapterList**](ChapterList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserIdList**
> CustomListList GetUserIdList(ctx, id, optional)
Get User's CustomList list

This will list only public CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| User ID | 
 **optional** | ***CustomListApiGetUserIdListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomListApiGetUserIdListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 

### Return type

[**CustomListList**](CustomListList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserList**
> CustomListList GetUserList(ctx, optional)
Get logged User CustomList list

This will list public and private CustomList 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomListApiGetUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomListApiGetUserListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**|  | [default to 10]
 **offset** | **optional.Int32**|  | 

### Return type

[**CustomListList**](CustomListList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostList**
> CustomListResponse PostList(ctx, optional)
Create CustomList

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomListApiPostListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomListApiPostListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CustomListCreate**](CustomListCreate.md)| This body is limited to 8kb max per call. | 

### Return type

[**CustomListResponse**](CustomListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
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

# **PutListId**
> CustomListResponse PutListId(ctx, id, optional)
Update CustomList

This body is limited to 8kb max per call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| CustomList ID | 
 **optional** | ***CustomListApiPutListIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomListApiPutListIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CustomListEdit**](CustomListEdit.md)|  | 

### Return type

[**CustomListResponse**](CustomListResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

