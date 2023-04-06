# klarity_integrations.model.estate_records_request_body.EstateRecordsRequestBody

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[records](#records)** | list, tuple,  | tuple,  |  | 
**insertInPeriod** | [**InsertInPeriodEnum**](InsertInPeriodEnum.md) | [**InsertInPeriodEnum**](InsertInPeriodEnum.md) |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# records

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[items](#items) | dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

# items

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  |  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**id** | [**Id**](Id.md) | [**Id**](Id.md) |  | 
**type** | [**Type**](Type.md) | [**Type**](Type.md) |  | 
**name** | [**Name**](Name.md) | [**Name**](Name.md) |  | [optional] 
**[metadata](#metadata)** | dict, frozendict.frozendict,  | frozendict.frozendict,  | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a &#x60;null&#x60; value. Maximum object size is 1MB.  | [optional] 
**tags** | [**Tags**](Tags.md) | [**Tags**](Tags.md) |  | [optional] 
**costs** | [**Costs**](Costs.md) | [**Costs**](Costs.md) |  | [optional] 
**validThrough** | [**ValidThrough**](ValidThrough.md) | [**ValidThrough**](ValidThrough.md) |  | [optional] 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# metadata

Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the `tags` key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a `null` value. Maximum object size is 1MB. 

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a &#x60;null&#x60; value. Maximum object size is 1MB.  | 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

