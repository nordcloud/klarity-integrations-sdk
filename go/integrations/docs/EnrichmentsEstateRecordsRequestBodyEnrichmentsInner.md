# EnrichmentsEstateRecordsRequestBodyEnrichmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Record** | [**EnrichmentRecord**](EnrichmentRecord.md) |  | 
**Data** | **map[string]interface{}** | Data contains an arbitrary JSON object with enrichment of the record. Can not be empty object. Maximum object size is 1MB.  | 

## Methods

### NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner

`func NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner(record EnrichmentRecord, data map[string]interface{}, ) *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner`

NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner instantiates a new EnrichmentsEstateRecordsRequestBodyEnrichmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInnerWithDefaults

`func NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInnerWithDefaults() *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner`

NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInnerWithDefaults instantiates a new EnrichmentsEstateRecordsRequestBodyEnrichmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecord

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetRecord() EnrichmentRecord`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetRecordOk() (*EnrichmentRecord, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) SetRecord(v EnrichmentRecord)`

SetRecord sets Record field to given value.


### GetData

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnrichmentsEstateRecordsRequestBodyEnrichmentsInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


