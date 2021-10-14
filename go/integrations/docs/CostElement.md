# CostElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date in YYYY-MM-DD format for which cost is charged. | 
**Value** | Pointer to **string** | Value of the cost. A dot (&#x60;.&#x60;) is used as a decimal separator. Use &#x60;null&#x60; to clear cost for given date. | [optional] 

## Methods

### NewCostElement

`func NewCostElement(date string, ) *CostElement`

NewCostElement instantiates a new CostElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostElementWithDefaults

`func NewCostElementWithDefaults() *CostElement`

NewCostElementWithDefaults instantiates a new CostElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CostElement) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CostElement) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CostElement) SetDate(v string)`

SetDate sets Date field to given value.


### GetValue

`func (o *CostElement) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CostElement) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CostElement) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CostElement) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


