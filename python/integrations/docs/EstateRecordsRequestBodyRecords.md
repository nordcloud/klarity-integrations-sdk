# EstateRecordsRequestBodyRecords


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Internal ID of a resource | 
**type** | **str** | Type of a resource (resourceType/extension) can contain just type, eg &#x60;jira&#x60; or with subtypes separated by \&quot;/\&quot; symbol | 
**name** | **str** | Friendly name of a resource (id will be used, if not provided) | [optional] 
**metadata** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Metadata contains a JSON object with information about the record. Keys of the object can by any string value, excluding the &#x60;tags&#x60; key which is reserved for the tags property. Maximum object size is 1MB.  | [optional] 
**tags** | **{str: (str,)}** | Tags are key &#x3D;&gt; value representing tags in Klarity. Maximum object size is 1MB. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


