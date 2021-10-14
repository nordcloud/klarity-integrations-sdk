# ValidThrough

Determines in which period the resource will be closed. If specified, it has to be in `YYYY-MM` format and be at least current period - it can not be past period. If not specified, the record will exist indefinitely. E.g. a resource with `validThrough` set to `2022-05` will be active till May 2022 and start being inactive in June 2022. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**value** | **str** | Determines in which period the resource will be closed. If specified, it has to be in &#x60;YYYY-MM&#x60; format and be at least current period - it can not be past period. If not specified, the record will exist indefinitely. E.g. a resource with &#x60;validThrough&#x60; set to &#x60;2022-05&#x60; will be active till May 2022 and start being inactive in June 2022.  | 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


