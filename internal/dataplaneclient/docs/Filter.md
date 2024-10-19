# Filter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | Name of the fcgi-app section this filter will use. | [optional] [default to null]
**BandwidthLimitName** | **string** | Filter name that will be used by &#39;set-bandwidth-limit&#39; actions to reference a specific bandwidth limitation filter | [optional] [default to null]
**CacheName** | **string** |  | [optional] [default to null]
**DefaultLimit** | **int32** | The max number of bytes that can be forwarded over the period. The value must be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes. | [optional] [default to null]
**DefaultPeriod** | **int32** | The default time period used to evaluate the bandwidth limitation rate. It can be specified for per-stream bandwidth limitation filters only. It follows the HAProxy time format and is expressed in milliseconds. | [optional] [default to null]
**Key** | **string** | A sample expression rule. It describes what elements will be analyzed, extracted, combined, and used to select which table entry to update the counters. It must be specified for shared bandwidth limitation filters only. | [optional] [default to null]
**Limit** | **int32** | The max number of bytes that can be forwarded over the period. The value must be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes. | [optional] [default to null]
**MinSize** | **int32** | The optional minimum number of bytes forwarded at a time by a stream excluding the last packet that may be smaller. This value can be specified for per-stream and shared bandwidth limitation filters. It follows the HAProxy size format and is expressed in bytes. | [optional] [default to null]
**SpoeConfig** | **string** |  | [optional] [default to null]
**SpoeEngine** | **string** |  | [optional] [default to null]
**Table** | **string** | An optional table to be used instead of the default one, which is the stick-table declared in the current proxy. It can be specified for shared bandwidth limitation filters only. | [optional] [default to null]
**TraceHexdump** | **bool** |  | [optional] [default to null]
**TraceName** | **string** |  | [optional] [default to null]
**TraceRndForwarding** | **bool** |  | [optional] [default to null]
**TraceRndParsing** | **bool** |  | [optional] [default to null]
**Type_** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


