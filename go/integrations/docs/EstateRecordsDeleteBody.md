# EstateRecordsDeleteBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | [**PeriodEnum**](PeriodEnum.md) |  | 
**Records** | [**[]EstateRecordsDeleteBodyRecordsInner**](EstateRecordsDeleteBodyRecordsInner.md) |  | 

## Methods

### NewEstateRecordsDeleteBody

`func NewEstateRecordsDeleteBody(period PeriodEnum, records []EstateRecordsDeleteBodyRecordsInner, ) *EstateRecordsDeleteBody`

NewEstateRecordsDeleteBody instantiates a new EstateRecordsDeleteBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstateRecordsDeleteBodyWithDefaults

`func NewEstateRecordsDeleteBodyWithDefaults() *EstateRecordsDeleteBody`

NewEstateRecordsDeleteBodyWithDefaults instantiates a new EstateRecordsDeleteBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *EstateRecordsDeleteBody) GetPeriod() PeriodEnum`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *EstateRecordsDeleteBody) GetPeriodOk() (*PeriodEnum, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *EstateRecordsDeleteBody) SetPeriod(v PeriodEnum)`

SetPeriod sets Period field to given value.


### GetRecords

`func (o *EstateRecordsDeleteBody) GetRecords() []EstateRecordsDeleteBodyRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *EstateRecordsDeleteBody) GetRecordsOk() (*[]EstateRecordsDeleteBodyRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *EstateRecordsDeleteBody) SetRecords(v []EstateRecordsDeleteBodyRecordsInner)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


