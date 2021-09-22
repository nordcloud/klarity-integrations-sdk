# EstateRecordsRequestBodyRecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal ID of a resource | 
**Name** | Pointer to **string** | Friendly name of a resource (id will be used, if not provided) | [optional] 
**Type** | **string** | Type of a resource (resourceType/extension) can contain just type, eg &#x60;jira&#x60; or with subtypes separated by \&quot;/\&quot; symbol | 
**Metadata** | Pointer to **map[string]interface{}** | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. Maximum object size is 1MB.  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are key &#x3D;&gt; value representing tags in Klarity. Maximum object size is 1MB. | [optional] 

## Methods

### NewEstateRecordsRequestBodyRecords

`func NewEstateRecordsRequestBodyRecords(id string, type_ string, ) *EstateRecordsRequestBodyRecords`

NewEstateRecordsRequestBodyRecords instantiates a new EstateRecordsRequestBodyRecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstateRecordsRequestBodyRecordsWithDefaults

`func NewEstateRecordsRequestBodyRecordsWithDefaults() *EstateRecordsRequestBodyRecords`

NewEstateRecordsRequestBodyRecordsWithDefaults instantiates a new EstateRecordsRequestBodyRecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EstateRecordsRequestBodyRecords) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EstateRecordsRequestBodyRecords) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EstateRecordsRequestBodyRecords) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EstateRecordsRequestBodyRecords) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EstateRecordsRequestBodyRecords) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EstateRecordsRequestBodyRecords) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EstateRecordsRequestBodyRecords) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EstateRecordsRequestBodyRecords) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EstateRecordsRequestBodyRecords) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EstateRecordsRequestBodyRecords) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *EstateRecordsRequestBodyRecords) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EstateRecordsRequestBodyRecords) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EstateRecordsRequestBodyRecords) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EstateRecordsRequestBodyRecords) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *EstateRecordsRequestBodyRecords) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EstateRecordsRequestBodyRecords) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EstateRecordsRequestBodyRecords) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EstateRecordsRequestBodyRecords) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


