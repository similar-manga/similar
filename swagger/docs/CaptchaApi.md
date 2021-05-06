# {{classname}}

All URIs are relative to *https://api.mangadex.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCaptchaSolve**](CaptchaApi.md#PostCaptchaSolve) | **Post** /captcha/solve | Solve Captcha

# **PostCaptchaSolve**
> InlineResponse2001 PostCaptchaSolve(ctx, optional)
Solve Captcha

Captchas can be solved explicitly through this endpoint, another way is to add a `X-Captcha-Result` header to any request. The same logic will verify the captcha and is probably more convenient because it takes one less request.  Authentication is optional. Captchas are tracked for both the client ip and for the user id, if you are logged in you want to send your session token but that is not required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CaptchaApiPostCaptchaSolveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CaptchaApiPostCaptchaSolveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Body**](Body.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

