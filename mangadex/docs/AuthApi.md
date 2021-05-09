# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthCheck**](AuthApi.md#GetAuthCheck) | **Get** /auth/check | Check token
[**PostAuthLogin**](AuthApi.md#PostAuthLogin) | **Post** /auth/login | Login
[**PostAuthLogout**](AuthApi.md#PostAuthLogout) | **Post** /auth/logout | Logout
[**PostAuthRefresh**](AuthApi.md#PostAuthRefresh) | **Post** /auth/refresh | Refresh token

# **GetAuthCheck**

> CheckResponse GetAuthCheck(ctx, )
Check token

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**CheckResponse**](CheckResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthLogin**

> LoginResponse PostAuthLogin(ctx, optional)
Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AuthApiPostAuthLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthApiPostAuthLoginOpts struct Name | Type | Description | Notes
------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of Login**](Login.md)| This body is limited to 2kb max per call. |

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthLogout**

> LogoutResponse PostAuthLogout(ctx, )
Logout

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**LogoutResponse**](LogoutResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAuthRefresh**

> RefreshResponse PostAuthRefresh(ctx, optional)
Refresh token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AuthApiPostAuthRefreshOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthApiPostAuthRefreshOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of RefreshToken**](RefreshToken.md)| This body is limited to 2kb max per call. |

### Return type

[**RefreshResponse**](RefreshResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

