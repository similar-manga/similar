# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGroupId**](ScanlationGroupApi.md#DeleteGroupId) | **Delete** /group/{id} | Delete Scanlation Group
[**DeleteGroupIdFollow**](ScanlationGroupApi.md#DeleteGroupIdFollow) | **
Delete** /group/{id}/follow | Unfollow Scanlation Group
[**GetGroupId**](ScanlationGroupApi.md#GetGroupId) | **Get** /group/{id} | View Scanlation Group
[**GetSearchGroup**](ScanlationGroupApi.md#GetSearchGroup) | **Get** /group | Scanlation Group list
[**GetUserFollowsGroup**](ScanlationGroupApi.md#GetUserFollowsGroup) | **
Get** /user/follows/group | Get logged User followed Groups
[**PostGroup**](ScanlationGroupApi.md#PostGroup) | **Post** /group | Create Scanlation Group
[**PostGroupIdFollow**](ScanlationGroupApi.md#PostGroupIdFollow) | **Post** /group/{id}/follow | Follow Scanlation Group
[**PutGroupId**](ScanlationGroupApi.md#PutGroupId) | **Put** /group/{id} | Update Scanlation Group

# **DeleteGroupId**

> Response DeleteGroupId(ctx, id)
Delete Scanlation Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Scanlation Group ID |

### Return type

[**Response**](Response.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteGroupIdFollow**

> Response DeleteGroupIdFollow(ctx, id)
Unfollow Scanlation Group

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

# **GetGroupId**

> ScanlationGroupResponse GetGroupId(ctx, id)
View Scanlation Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Scanlation Group ID |

### Return type

[**ScanlationGroupResponse**](ScanlationGroupResponse.md)

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
**optional** | ***ScanlationGroupApiGetSearchGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScanlationGroupApiGetSearchGroupOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 10]
**offset** | **optional.Int32**| |
**ids** | [**optional.Interface of []string**](string.md)| ScanlationGroup ids (limited to 100 per request) |
**name** | **optional.String**| |

### Return type

[**ScanlationGroupList**](ScanlationGroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserFollowsGroup**

> ScanlationGroupList GetUserFollowsGroup(ctx, optional)
Get logged User followed Groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***ScanlationGroupApiGetUserFollowsGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScanlationGroupApiGetUserFollowsGroupOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------
**limit** | **optional.Int32**| | [default to 10]
**offset** | **optional.Int32**| |

### Return type

[**ScanlationGroupList**](ScanlationGroupList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGroup**

> ScanlationGroupResponse PostGroup(ctx, optional)
Create Scanlation Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***ScanlationGroupApiPostGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScanlationGroupApiPostGroupOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of CreateScanlationGroup**](CreateScanlationGroup.md)| This body is limited to 8kb max
per call. |

### Return type

[**ScanlationGroupResponse**](ScanlationGroupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostGroupIdFollow**

> Response PostGroupIdFollow(ctx, id)
Follow Scanlation Group

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

# **PutGroupId**

> ScanlationGroupResponse PutGroupId(ctx, id, optional)
Update Scanlation Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md)| Scanlation Group ID |
**optional** | ***ScanlationGroupApiPutGroupIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ScanlationGroupApiPutGroupIdOpts struct Name | Type | Description
| Notes ------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of ScanlationGroupEdit**](ScanlationGroupEdit.md)| This body is limited to 8kb max per
call. |

### Return type

[**ScanlationGroupResponse**](ScanlationGroupResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

