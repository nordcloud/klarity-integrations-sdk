# EstateRecordsRequestBodyRecords


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | [**Id**](Id.md) |  | 
**type** | [**Type**](Type.md) |  | 
**name** | [**Name**](Name.md) |  | [optional] 
**metadata** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a &#x60;null&#x60; value. Maximum object size is 1MB.  | [optional] 
**tags** | [**Tags**](Tags.md) |  | [optional] 
**costs** | [**Costs**](Costs.md) |  | [optional] 
**valid_through** | [**ValidThrough**](ValidThrough.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


