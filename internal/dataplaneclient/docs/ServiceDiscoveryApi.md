# \ServiceDiscoveryApi

All URIs are relative to *http://localhost/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAWSRegion**](ServiceDiscoveryApi.md#CreateAWSRegion) | **Post** /service_discovery/aws | Add a new AWS region
[**CreateConsul**](ServiceDiscoveryApi.md#CreateConsul) | **Post** /service_discovery/consul | Add a new Consul server
[**DeleteAWSRegion**](ServiceDiscoveryApi.md#DeleteAWSRegion) | **Delete** /service_discovery/aws/{id} | Delete an AWS region
[**DeleteConsul**](ServiceDiscoveryApi.md#DeleteConsul) | **Delete** /service_discovery/consul/{id} | Delete a Consul server
[**GetAWSRegion**](ServiceDiscoveryApi.md#GetAWSRegion) | **Get** /service_discovery/aws/{id} | Return an AWS region
[**GetAWSRegions**](ServiceDiscoveryApi.md#GetAWSRegions) | **Get** /service_discovery/aws | Return an array of all configured AWS regions
[**GetConsul**](ServiceDiscoveryApi.md#GetConsul) | **Get** /service_discovery/consul/{id} | Return one Consul server
[**GetConsuls**](ServiceDiscoveryApi.md#GetConsuls) | **Get** /service_discovery/consul | Return an array of all configured Consul servers
[**ReplaceAWSRegion**](ServiceDiscoveryApi.md#ReplaceAWSRegion) | **Put** /service_discovery/aws/{id} | Replace an AWS region
[**ReplaceConsul**](ServiceDiscoveryApi.md#ReplaceConsul) | **Put** /service_discovery/consul/{id} | Replace a Consul server


# **CreateAWSRegion**
> AwsRegion CreateAWSRegion(ctx, data)
Add a new AWS region

Add a new AWS region. Credentials are not required in case Dataplane API is running in an EC2 instance with proper IAM role attached.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AwsRegion**](AwsRegion.md)|  | 

### Return type

[**AwsRegion**](awsRegion.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateConsul**
> Consul CreateConsul(ctx, data)
Add a new Consul server

Adds a new Consul server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Consul**](Consul.md)|  | 

### Return type

[**Consul**](consul.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAWSRegion**
> DeleteAWSRegion(ctx, id)
Delete an AWS region

Delete an AWS region configuration by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| AWS region ID | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsul**
> DeleteConsul(ctx, id)
Delete a Consul server

Deletes a Consul server configuration by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Consul server Index | 

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAWSRegion**
> AwsRegion GetAWSRegion(ctx, id)
Return an AWS region

Return one AWS Region configuration by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| AWS region id | 

### Return type

[**AwsRegion**](awsRegion.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAWSRegions**
> AwsRegions GetAWSRegions(ctx, )
Return an array of all configured AWS regions

Return all configured AWS regions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AwsRegions**](awsRegions.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsul**
> Consul GetConsul(ctx, id)
Return one Consul server

Returns one Consul server configuration by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Consul server id | 

### Return type

[**Consul**](consul.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsuls**
> Consuls GetConsuls(ctx, )
Return an array of all configured Consul servers

Returns all configured Consul servers.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Consuls**](consuls.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAWSRegion**
> AwsRegion ReplaceAWSRegion(ctx, id, data)
Replace an AWS region

Replace an AWS region configuration by its id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| AWS Region ID | 
  **data** | [**AwsRegion**](AwsRegion.md)|  | 

### Return type

[**AwsRegion**](awsRegion.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceConsul**
> Consul ReplaceConsul(ctx, id, data)
Replace a Consul server

Replaces a Consul server configuration by it's id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Consul Index | 
  **data** | [**Consul**](Consul.md)|  | 

### Return type

[**Consul**](consul.md)

### Authorization

[basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

