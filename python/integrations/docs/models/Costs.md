# klarity_integrations.model.costs.Costs

Costs object contains an array of costs per day and a currency. Only if there is a date passed into an array, then the cost is being added/updated for this date. If you want to clear cost for a given date, you have to `null` directly into an array. 

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
dict, frozendict.frozendict,  | frozendict.frozendict,  | Costs object contains an array of costs per day and a currency. Only if there is a date passed into an array, then the cost is being added/updated for this date. If you want to clear cost for a given date, you have to &#x60;null&#x60; directly into an array.  | 

### Dictionary Keys
Key | Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | ------------- | -------------
**[values](#values)** | list, tuple,  | tuple,  |  | 
**currency** | str,  | str,  | Currency in ISO-4217 format. | 
**any_string_name** | dict, frozendict.frozendict, str, date, datetime, int, float, bool, decimal.Decimal, None, list, tuple, bytes, io.FileIO, io.BufferedReader | frozendict.frozendict, str, BoolClass, decimal.Decimal, NoneClass, tuple, bytes, FileIO | any string name can be used but the value must be the correct type | [optional]

# values

## Model Type Info
Input Type | Accessed Type | Description | Notes
------------ | ------------- | ------------- | -------------
list, tuple,  | tuple,  |  | 

### Tuple Items
Class Name | Input Type | Accessed Type | Description | Notes
------------- | ------------- | ------------- | ------------- | -------------
[**CostElement**](CostElement.md) | [**CostElement**](CostElement.md) | [**CostElement**](CostElement.md) |  | 

[[Back to Model list]](../../README.md#documentation-for-models) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to README]](../../README.md)

