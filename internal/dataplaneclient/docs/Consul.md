# Consul

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | [default to null]
**Defaults** | **string** | Name of the defaults section to be used in backends created by this service | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [default to null]
**HealthCheckPolicy** | **string** | Defines the health check conditions required for each node to be considered valid for the service.   none: all nodes are considered valid   any: a node is considered valid if any one health check is &#39;passing&#39;   all: a node is considered valid if all health checks are &#39;passing&#39;   min: a node is considered valid if the number of &#39;passing&#39; checks is greater or equal to the &#39;health_check_policy_min&#39; value.     If the node has less health checks configured then &#39;health_check_policy_min&#39; it is considered invalid. | [optional] [default to null]
**HealthCheckPolicyMin** | **int32** |  | [optional] [default to null]
**Id** | **string** | Auto generated ID. | [optional] [default to null]
**Mode** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Namespace** | **string** |  | [optional] [default to null]
**Port** | **int32** |  | [default to null]
**RetryTimeout** | **int32** | Duration in seconds in-between data pulling requests to the consul server | [default to null]
**ServerSlotsBase** | **int32** |  | [optional] [default to null]
**ServerSlotsGrowthIncrement** | **int32** |  | [optional] [default to null]
**ServerSlotsGrowthType** | **string** |  | [optional] [default to null]
**ServiceAllowlist** | **[]string** |  | [optional] [default to null]
**ServiceDenylist** | **[]string** |  | [optional] [default to null]
**ServiceNameRegexp** | **string** | Regular expression used to filter services by name. | [optional] [default to null]
**Token** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


