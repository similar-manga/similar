# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAuthorId**](AuthorApi.md#DeleteAuthorId) | **Delete** /author/{id} | Delete Author
[**GetAuthor**](AuthorApi.md#GetAuthor) | **Get** /author | Author list
[**GetAuthorId**](AuthorApi.md#GetAuthorId) | **Get** /author/{id} | Get Author
[**PostAuthor**](AuthorApi.md#PostAuthor) | **Post** /author | Create Author
[**PutAuthorId**](AuthorApi.md#PutAuthorId) | **Put** /author/{id} | Update Author

# **DeleteAuthorId**

> Response DeleteAuthorId(ctx, id)
Delete Author

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Author ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthor**

> AuthorList GetAuthor(ctx, optional)
Author list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AuthorApiGetAuthorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorApiGetAuthorOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 10]
**offset** | **optional.Int32**| |
**ids** | [**optional.Interface of []string**](string.md)| Author ids (limited to 100 per request) |
**name** | **optional.String**| |

### Return type

[**AuthorList**](AuthorList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorId**

> AuthorResponse GetAuthorId(ctx, id)
Get Author

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Author ID |

### Return type

[**AuthorResponse**](AuthorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthor**

> AuthorResponse PostAuthor(ctx, optional)
Create Author

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AuthorApiPostAuthorOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorApiPostAuthorOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of AuthorCreate**](AuthorCreate.md)| This body is limited to 2kb max per call. |

### Return type

[**AuthorResponse**](AuthorResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutAuthorId**

> AuthorResponse PutAuthorId(ctx, id, optional)
Update Author

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Author ID |
**optional** | ***AuthorApiPutAuthorIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthorApiPutAuthorIdOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of AuthorEdit**](AuthorEdit.md)| This body is limited to 2kb max per call. |

### Return type

[**AuthorResponse**](AuthorResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

