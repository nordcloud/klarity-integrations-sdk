# EstateRecordsRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsertInPeriod** | Pointer to [**InsertInPeriodEnum**](InsertInPeriodEnum.md) |  | [optional] 
**Records** | [**[]EstateRecordsRequestBodyRecordsInner**](EstateRecordsRequestBodyRecordsInner.md) |  | 

## Methods

### NewEstateRecordsRequestBody

`func NewEstateRecordsRequestBody(records []EstateRecordsRequestBodyRecordsInner, ) *EstateRecordsRequestBody`

NewEstateRecordsRequestBody instantiates a new EstateRecordsRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstateRecordsRequestBodyWithDefaults

`func NewEstateRecordsRequestBodyWithDefaults() *EstateRecordsRequestBody`

NewEstateRecordsRequestBodyWithDefaults instantiates a new EstateRecordsRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsertInPeriod

`func (o *EstateRecordsRequestBody) GetInsertInPeriod() InsertInPeriodEnum`

GetInsertInPeriod returns the InsertInPeriod field if non-nil, zero value otherwise.

### GetInsertInPeriodOk

`func (o *EstateRecordsRequestBody) GetInsertInPeriodOk() (*InsertInPeriodEnum, bool)`

GetInsertInPeriodOk returns a tuple with the InsertInPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertInPeriod

`func (o *EstateRecordsRequestBody) SetInsertInPeriod(v InsertInPeriodEnum)`

SetInsertInPeriod sets InsertInPeriod field to given value.

### HasInsertInPeriod

`func (o *EstateRecordsRequestBody) HasInsertInPeriod() bool`

HasInsertInPeriod returns a boolean if a field has been set.

### GetRecords

`func (o *EstateRecordsRequestBody) GetRecords() []EstateRecordsRequestBodyRecordsInner`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *EstateRecordsRequestBody) GetRecordsOk() (*[]EstateRecordsRequestBodyRecordsInner, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *EstateRecordsRequestBody) SetRecords(v []EstateRecordsRequestBodyRecordsInner)`

SetRecords sets Records field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


