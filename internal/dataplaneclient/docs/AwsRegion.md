# AwsRegion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | **string** | AWS Access Key ID. | [optional] [default to null]
**Allowlist** | [**[]AwsFilters**](awsFilters.md) | Specify the AWS filters used to filter the EC2 instances to add | [optional] [default to null]
**Denylist** | [**[]AwsFilters**](awsFilters.md) | Specify the AWS filters used to filter the EC2 instances to ignore | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [default to null]
**Id** | **string** | Auto generated ID. | [optional] [default to null]
**Ipv4Address** | **string** | Select which IPv4 address the Service Discovery has to use for the backend server entry | [default to null]
**Name** | **string** |  | [default to null]
**Region** | **string** |  | [default to null]
**RetryTimeout** | **int32** | Duration in seconds in-between data pulling requests to the AWS region | [default to null]
**SecretAccessKey** | **string** | AWS Secret Access Key. | [optional] [default to null]
**ServerSlotsBase** | **int32** |  | [optional] [default to null]
**ServerSlotsGrowthIncrement** | **int32** |  | [optional] [default to null]
**ServerSlotsGrowthType** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


