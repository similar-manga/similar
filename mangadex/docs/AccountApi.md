# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountActivateCode**](AccountApi.md#GetAccountActivateCode) | **Get** /account/activate/{code} | Activate account
[**PostAccountActivateResend**](AccountApi.md#PostAccountActivateResend) | **
Post** /account/activate/resend | Resend Activation code
[**PostAccountCreate**](AccountApi.md#PostAccountCreate) | **Post** /account/create | Create Account
[**PostAccountRecover**](AccountApi.md#PostAccountRecover) | **Post** /account/recover | Recover given Account
[**PostAccountRecoverCode**](AccountApi.md#PostAccountRecoverCode) | **
Post** /account/recover/{code} | Complete Account recover

# **GetAccountActivateCode**

> AccountActivateResponse GetAccountActivateCode(ctx, code)
Activate account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  |

### Return type

[**AccountActivateResponse**](AccountActivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountActivateResend**

> AccountActivateResponse PostAccountActivateResend(ctx, optional)
Resend Activation code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AccountApiPostAccountActivateResendOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountApiPostAccountActivateResendOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of SendAccountActivationCode**](SendAccountActivationCode.md)| This body is limited to
1kb max per call. |

### Return type

[**AccountActivateResponse**](AccountActivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountCreate**

> UserResponse PostAccountCreate(ctx, optional)
Create Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AccountApiPostAccountCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountApiPostAccountCreateOpts struct Name | Type | Description |
Notes ------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of CreateAccount**](CreateAccount.md)| This body is limited to 4kb max per call. |

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountRecover**

> AccountActivateResponse PostAccountRecover(ctx, optional)
Recover given Account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**optional** | ***AccountApiPostAccountRecoverOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountApiPostAccountRecoverOpts struct Name | Type | Description
| Notes ------------- | ------------- | ------------- | -------------
**body** | [**optional.Interface of SendAccountActivationCode**](SendAccountActivationCode.md)| This body is limited to
1kb max per call. |

### Return type

[**AccountActivateResponse**](AccountActivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostAccountRecoverCode**

> AccountActivateResponse PostAccountRecoverCode(ctx, code, optional)
Complete Account recover

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**code** | **string**|  |
**optional** | ***AccountApiPostAccountRecoverCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AccountApiPostAccountRecoverCodeOpts struct Name | Type |
Description | Notes ------------- | ------------- | ------------- | -------------

**body** | [**optional.Interface of RecoverCompleteBody**](RecoverCompleteBody.md)| This body is limited to 2kb max per
call. |

### Return type

[**AccountActivateResponse**](AccountActivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

