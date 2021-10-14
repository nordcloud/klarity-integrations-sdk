# Costs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | Currency in ISO-4217 format. | 
**Values** | [**[]CostElement**](CostElement.md) |  | 

## Methods

### NewCosts

`func NewCosts(currency string, values []CostElement, ) *Costs`

NewCosts instantiates a new Costs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostsWithDefaults

`func NewCostsWithDefaults() *Costs`

NewCostsWithDefaults instantiates a new Costs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Costs) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Costs) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Costs) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetValues

`func (o *Costs) GetValues() []CostElement`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Costs) GetValuesOk() (*[]CostElement, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Costs) SetValues(v []CostElement)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


