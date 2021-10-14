# klarity_integrations.KlarityIntegrationsApi

All URIs are relative to *https://integrations-api.klarity.nordcloudapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_estate_records_delete**](KlarityIntegrationsApi.md#v1_estate_records_delete) | **DELETE** /v1/estateRecords | Delete Klarity estate records
[**v1_estate_records_post**](KlarityIntegrationsApi.md#v1_estate_records_post) | **POST** /v1/estateRecords | Manage Klarity estate records


# **v1_estate_records_delete**
> ERRORUNKNOWN v1_estate_records_delete()

Delete Klarity estate records

In Klarity, to delete an estate record created from an External Integration, provide: either, nid; or, id, type, and name. Records can be deleted in current, previous or all periods. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import time
import klarity_integrations
from klarity_integrations.api import klarity_integrations_api
from klarity_integrations.model.estate_records_delete_body import EstateRecordsDeleteBody
from klarity_integrations.model.errorunknown import ERRORUNKNOWN
from pprint import pprint
# Defining the host is optional and defaults to https://integrations-api.klarity.nordcloudapp.com
# See configuration.py for a list of all supported configuration parameters.
configuration = klarity_integrations.Configuration(
    host = "https://integrations-api.klarity.nordcloudapp.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = klarity_integrations.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with klarity_integrations.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = klarity_integrations_api.KlarityIntegrationsApi(api_client)
    estate_records_delete_body = EstateRecordsDeleteBody(
        period=PeriodEnum("current"),
        records=[
            ,
        ],
    ) # EstateRecordsDeleteBody |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Delete Klarity estate records
        api_response = api_instance.v1_estate_records_delete(estate_records_delete_body=estate_records_delete_body)
        pprint(api_response)
    except klarity_integrations.ApiException as e:
        print("Exception when calling KlarityIntegrationsApi->v1_estate_records_delete: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estate_records_delete_body** | [**EstateRecordsDeleteBody**](EstateRecordsDeleteBody.md)|  | [optional]

### Return type

[**ERRORUNKNOWN**](ERRORUNKNOWN.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Status Accepted |  -  |
**400** | Bad request |  -  |
**401** | Unauthorized response |  -  |
**429** | Too many requests |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_estate_records_post**
> ERRORUNKNOWN v1_estate_records_post()

Manage Klarity estate records

In Klarity, create new estate records, or update the metadata, tags, and costs, in existing estate records. You update metadata, tags, and costs by data-upserting: If you provide a value for a given parameter, the relevant field in the estate record is updated. If you provide no value for a parameter, the given field is not updated in the estate record. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import time
import klarity_integrations
from klarity_integrations.api import klarity_integrations_api
from klarity_integrations.model.estate_records_request_body import EstateRecordsRequestBody
from klarity_integrations.model.errorunknown import ERRORUNKNOWN
from pprint import pprint
# Defining the host is optional and defaults to https://integrations-api.klarity.nordcloudapp.com
# See configuration.py for a list of all supported configuration parameters.
configuration = klarity_integrations.Configuration(
    host = "https://integrations-api.klarity.nordcloudapp.com"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = klarity_integrations.Configuration(
    access_token = 'YOUR_BEARER_TOKEN'
)

# Enter a context with an instance of the API client
with klarity_integrations.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = klarity_integrations_api.KlarityIntegrationsApi(api_client)
    estate_records_request_body = EstateRecordsRequestBody(
        records=[
            EstateRecordsRequestBodyRecords(
                id=Id("js002"),
                name=Name("Jira Subscription 002"),
                type=Type("jira/subscription/basic"),
                metadata={},
                tags=Tags(
                    key="key_example",
                ),
                costs=Costs(
                    currency="currency_example",
                    values=[
                        CostElement(
                            date=dateutil_parser('1970-01-01').date(),
                            value="4807288800152802179809",
                        ),
                    ],
                ),
                valid_through=ValidThrough("2048-09"),
            ),
        ],
    ) # EstateRecordsRequestBody |  (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Manage Klarity estate records
        api_response = api_instance.v1_estate_records_post(estate_records_request_body=estate_records_request_body)
        pprint(api_response)
    except klarity_integrations.ApiException as e:
        print("Exception when calling KlarityIntegrationsApi->v1_estate_records_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estate_records_request_body** | [**EstateRecordsRequestBody**](EstateRecordsRequestBody.md)|  | [optional]

### Return type

[**ERRORUNKNOWN**](ERRORUNKNOWN.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Status Accepted |  -  |
**400** | Bad request |  -  |
**401** | Unauthorized response |  -  |
**429** | Too many requests |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

