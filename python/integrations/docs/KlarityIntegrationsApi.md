# klarity_integrations.KlarityIntegrationsApi

All URIs are relative to *https://integrations-api.klarity.nordcloudapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_estate_records_post**](KlarityIntegrationsApi.md#v1_estate_records_post) | **POST** /v1/estateRecords | Manage Klarity estate records


# **v1_estate_records_post**
> AcceptedResponseBody v1_estate_records_post()

Manage Klarity estate records

Create Klarity estate records or update their metadata, tags and costs. Metadata, tags and costs updating works as a data upserting. If you will not provide one of them, it will not be updated. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import time
import klarity_integrations
from klarity_integrations.api import klarity_integrations_api
from klarity_integrations.model.estate_records_request_body import EstateRecordsRequestBody
from klarity_integrations.model.error_response import ErrorResponse
from klarity_integrations.model.accepted_response_body import AcceptedResponseBody
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
                id="js002",
                name="Jira Subscription 002",
                type="jira/subscription/basic",
                metadata={},
                tags={
                    "key": "key_example",
                },
                costs=Costs(
                    currency="currency_example",
                    values=[
                        CostElement(
                            date=dateutil_parser('1970-01-01').date(),
                            value="4807288800152802179809",
                        ),
                    ],
                ),
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

[**AcceptedResponseBody**](AcceptedResponseBody.md)

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

