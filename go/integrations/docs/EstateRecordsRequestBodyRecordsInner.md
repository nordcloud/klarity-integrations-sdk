# EstateRecordsRequestBodyRecordsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal ID of a resource. | 
**Name** | Pointer to **string** | Friendly name of a resource (id will be used, if not provided). | [optional] 
**Type** | **string** | Type of a resource (resourceType/extension) can contain just type, eg &#x60;jira&#x60; or with subtypes separated by \&quot;/\&quot; symbol. | 
**Metadata** | Pointer to **map[string]interface{}** | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. If you want to clear the metadata key, you have to provide it directly in an object with a &#x60;null&#x60; value. Maximum object size is 1MB.  | [optional] 
**Tags** | Pointer to **map[string]string** | Tags are key &#x3D;&gt; value representing tags in Klarity. If you want to clear the tag key, you have to provide it directly in an object with a &#x60;null&#x60; value. Maximum object size is 1MB.  | [optional] 
**Costs** | Pointer to [**Costs**](Costs.md) |  | [optional] 
**ValidThrough** | Pointer to **string** | Determines in which period the resource will be closed. If specified, it has to be in &#x60;YYYY-MM&#x60; format and be at least current period - it can not be past period. If you need to insert Estate Records in previous period, use insertInPeriod field - when set it will treat previous period as current. If not specified, the record will exist indefinitely. E.g. a resource with &#x60;validThrough&#x60; set to &#x60;2022-05&#x60; will be active till May 2022 and start being inactive in June 2022.  | [optional] 

## Methods

### NewEstateRecordsRequestBodyRecordsInner

`func NewEstateRecordsRequestBodyRecordsInner(id string, type_ string, ) *EstateRecordsRequestBodyRecordsInner`

NewEstateRecordsRequestBodyRecordsInner instantiates a new EstateRecordsRequestBodyRecordsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstateRecordsRequestBodyRecordsInnerWithDefaults

`func NewEstateRecordsRequestBodyRecordsInnerWithDefaults() *EstateRecordsRequestBodyRecordsInner`

NewEstateRecordsRequestBodyRecordsInnerWithDefaults instantiates a new EstateRecordsRequestBodyRecordsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EstateRecordsRequestBodyRecordsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EstateRecordsRequestBodyRecordsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EstateRecordsRequestBodyRecordsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EstateRecordsRequestBodyRecordsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EstateRecordsRequestBodyRecordsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EstateRecordsRequestBodyRecordsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EstateRecordsRequestBodyRecordsInner) SetType(v string)`

SetType sets Type field to given value.


### GetMetadata

`func (o *EstateRecordsRequestBodyRecordsInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EstateRecordsRequestBodyRecordsInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EstateRecordsRequestBodyRecordsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *EstateRecordsRequestBodyRecordsInner) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EstateRecordsRequestBodyRecordsInner) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EstateRecordsRequestBodyRecordsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCosts

`func (o *EstateRecordsRequestBodyRecordsInner) GetCosts() Costs`

GetCosts returns the Costs field if non-nil, zero value otherwise.

### GetCostsOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetCostsOk() (*Costs, bool)`

GetCostsOk returns a tuple with the Costs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosts

`func (o *EstateRecordsRequestBodyRecordsInner) SetCosts(v Costs)`

SetCosts sets Costs field to given value.

### HasCosts

`func (o *EstateRecordsRequestBodyRecordsInner) HasCosts() bool`

HasCosts returns a boolean if a field has been set.

### GetValidThrough

`func (o *EstateRecordsRequestBodyRecordsInner) GetValidThrough() string`

GetValidThrough returns the ValidThrough field if non-nil, zero value otherwise.

### GetValidThroughOk

`func (o *EstateRecordsRequestBodyRecordsInner) GetValidThroughOk() (*string, bool)`

GetValidThroughOk returns a tuple with the ValidThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidThrough

`func (o *EstateRecordsRequestBodyRecordsInner) SetValidThrough(v string)`

SetValidThrough sets ValidThrough field to given value.

### HasValidThrough

`func (o *EstateRecordsRequestBodyRecordsInner) HasValidThrough() bool`

HasValidThrough returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


