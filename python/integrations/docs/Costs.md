# Costs

Costs object contains an array of costs per day and a currency. Only if there is a date passed into an array, then the cost is being added/updated for this date. If you want to clear cost for a given date, you have to `null` directly into an array. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**currency** | **str** | Currency in ISO-4217 format. | 
**values** | [**[CostElement]**](CostElement.md) |  | 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


