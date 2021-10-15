# EstateRecordsDeleteBodyRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nid** | Pointer to **string** | Nordcloud ID of a resource in Klarity. | [optional] 
**Id** | Pointer to **string** | Internal ID of a resource. | [optional] 
**Name** | Pointer to **string** | Friendly name of a resource (id will be used, if not provided). | [optional] 
**Type** | Pointer to **string** | Type of a resource (resourceType/extension) can contain just type, eg &#x60;jira&#x60; or with subtypes separated by \&quot;/\&quot; symbol. | [optional] 

## Methods

### NewEstateRecordsDeleteBodyRecords

`func NewEstateRecordsDeleteBodyRecords() *EstateRecordsDeleteBodyRecords`

NewEstateRecordsDeleteBodyRecords instantiates a new EstateRecordsDeleteBodyRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstateRecordsDeleteBodyRecordsWithDefaults

`func NewEstateRecordsDeleteBodyRecordsWithDefaults() *EstateRecordsDeleteBodyRecords`

NewEstateRecordsDeleteBodyRecordsWithDefaults instantiates a new EstateRecordsDeleteBodyRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNid

`func (o *EstateRecordsDeleteBodyRecords) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *EstateRecordsDeleteBodyRecords) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *EstateRecordsDeleteBodyRecords) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *EstateRecordsDeleteBodyRecords) HasNid() bool`

HasNid returns a boolean if a field has been set.

### GetId

`func (o *EstateRecordsDeleteBodyRecords) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EstateRecordsDeleteBodyRecords) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EstateRecordsDeleteBodyRecords) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EstateRecordsDeleteBodyRecords) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EstateRecordsDeleteBodyRecords) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EstateRecordsDeleteBodyRecords) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EstateRecordsDeleteBodyRecords) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EstateRecordsDeleteBodyRecords) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EstateRecordsDeleteBodyRecords) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EstateRecordsDeleteBodyRecords) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EstateRecordsDeleteBodyRecords) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EstateRecordsDeleteBodyRecords) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


