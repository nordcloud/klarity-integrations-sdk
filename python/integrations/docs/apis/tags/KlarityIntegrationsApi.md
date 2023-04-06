<a name="__pageTop"></a>
# klarity_integrations.apis.tags.klarity_integrations_api.KlarityIntegrationsApi

All URIs are relative to *https://integrations-api.klarity.nordcloudapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_enrichments_estate_records_post**](#v1_enrichments_estate_records_post) | **post** /v1/enrichments/estateRecords | Enrich Klarity estate records
[**v1_estate_records_delete**](#v1_estate_records_delete) | **delete** /v1/estateRecords | Delete Klarity estate records
[**v1_estate_records_post**](#v1_estate_records_post) | **post** /v1/estateRecords | Manage Klarity estate records

# **v1_enrichments_estate_records_post**
<a name="v1_enrichments_estate_records_post"></a>
> AcceptedResponseBody v1_enrichments_estate_records_post()

Enrich Klarity estate records

Enriching works as upsert per integration, meaning new enrichment will be created if there was none for given integration and if enrichment already existed for given integration, it will be replaced by new one. Cost records can not be enriched. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import klarity_integrations
from klarity_integrations.apis.tags import klarity_integrations_api
from klarity_integrations.model.enrichments_estate_records_request_body import EnrichmentsEstateRecordsRequestBody
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

    # example passing only optional values
    body = EnrichmentsEstateRecordsRequestBody(
        enrichments=[
            dict(
                record=EnrichmentRecord(
                    nid="nid_example",
                ),
                data=dict(),
            )
        ],
    )
    try:
        # Enrich Klarity estate records
        api_response = api_instance.v1_enrichments_estate_records_post(
            body=body,
        )
        pprint(api_response)
    except klarity_integrations.ApiException as e:
        print("Exception when calling KlarityIntegrationsApi->v1_enrichments_estate_records_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, Unset] | optional, default is unset |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**EnrichmentsEstateRecordsRequestBody**](../../models/EnrichmentsEstateRecordsRequestBody.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
202 | [ApiResponseFor202](#v1_enrichments_estate_records_post.ApiResponseFor202) | Status Accepted
400 | [ApiResponseFor400](#v1_enrichments_estate_records_post.ApiResponseFor400) | Bad request
401 | [ApiResponseFor401](#v1_enrichments_estate_records_post.ApiResponseFor401) | Unauthorized response
429 | [ApiResponseFor429](#v1_enrichments_estate_records_post.ApiResponseFor429) | Too many requests
500 | [ApiResponseFor500](#v1_enrichments_estate_records_post.ApiResponseFor500) | Internal server error

#### v1_enrichments_estate_records_post.ApiResponseFor202
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor202ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor202ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**AcceptedResponseBody**](../../models/AcceptedResponseBody.md) |  | 


#### v1_enrichments_estate_records_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor400ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor400ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_enrichments_estate_records_post.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor401ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor401ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_enrichments_estate_records_post.ApiResponseFor429
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor429ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor429ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_enrichments_estate_records_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor500ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor500ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **v1_estate_records_delete**
<a name="v1_estate_records_delete"></a>
> AcceptedResponseBody v1_estate_records_delete()

Delete Klarity estate records

In Klarity, to delete an estate record created from an External Integration, provide: either, nid; or, id, type, and name. Records can be deleted in current, previous or all periods. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import klarity_integrations
from klarity_integrations.apis.tags import klarity_integrations_api
from klarity_integrations.model.error_response import ErrorResponse
from klarity_integrations.model.estate_records_delete_body import EstateRecordsDeleteBody
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

    # example passing only optional values
    body = EstateRecordsDeleteBody(
        period=PeriodEnum("current"),
        records=[
            dict(
                nid="nid_example",
                id=Id("js002"),
                name=Name("Jira Subscription 002"),
                type=Type("jira/subscription/basic"),
            )
        ],
    )
    try:
        # Delete Klarity estate records
        api_response = api_instance.v1_estate_records_delete(
            body=body,
        )
        pprint(api_response)
    except klarity_integrations.ApiException as e:
        print("Exception when calling KlarityIntegrationsApi->v1_estate_records_delete: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, Unset] | optional, default is unset |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**EstateRecordsDeleteBody**](../../models/EstateRecordsDeleteBody.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
202 | [ApiResponseFor202](#v1_estate_records_delete.ApiResponseFor202) | Status Accepted
400 | [ApiResponseFor400](#v1_estate_records_delete.ApiResponseFor400) | Bad request
401 | [ApiResponseFor401](#v1_estate_records_delete.ApiResponseFor401) | Unauthorized response
429 | [ApiResponseFor429](#v1_estate_records_delete.ApiResponseFor429) | Too many requests
500 | [ApiResponseFor500](#v1_estate_records_delete.ApiResponseFor500) | Internal server error

#### v1_estate_records_delete.ApiResponseFor202
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor202ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor202ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**AcceptedResponseBody**](../../models/AcceptedResponseBody.md) |  | 


#### v1_estate_records_delete.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor400ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor400ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_delete.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor401ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor401ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_delete.ApiResponseFor429
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor429ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor429ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_delete.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor500ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor500ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

# **v1_estate_records_post**
<a name="v1_estate_records_post"></a>
> AcceptedResponseBody v1_estate_records_post()

Manage Klarity estate records

In Klarity, create new estate records, or update the metadata, tags, and costs, in existing estate records. You update metadata, tags, and costs by data-upserting: If you provide a value for a given parameter, the relevant field in the estate record is updated. If you provide no value for a parameter, the given field is not updated in the estate record. 

### Example

* Bearer (JWT) Authentication (bearerAuth):
```python
import klarity_integrations
from klarity_integrations.apis.tags import klarity_integrations_api
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

    # example passing only optional values
    body = EstateRecordsRequestBody(
        insert_in_period=InsertInPeriodEnum("current"),
        records=[
            dict(
                id=Id("js002"),
                name=Name("Jira Subscription 002"),
                type=Type("jira/subscription/basic"),
                metadata=dict(),
                tags=Tags(
                    key="key_example",
                ),
                costs=Costs(
                    currency="currency_example",
                    values=[
                        CostElement(
                            date="1970-01-01",
                            value="4807288800152802179809",
                        )
                    ],
                ),
                valid_through=ValidThrough("2048-09"),
            )
        ],
    )
    try:
        # Manage Klarity estate records
        api_response = api_instance.v1_estate_records_post(
            body=body,
        )
        pprint(api_response)
    except klarity_integrations.ApiException as e:
        print("Exception when calling KlarityIntegrationsApi->v1_estate_records_post: %s\n" % e)
```
### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
body | typing.Union[SchemaForRequestBodyApplicationJson, Unset] | optional, default is unset |
content_type | str | optional, default is 'application/json' | Selects the schema and serialization of the request body
accept_content_types | typing.Tuple[str] | default is ('application/json', ) | Tells the server the content type(s) that are accepted by the client
stream | bool | default is False | if True then the response.content will be streamed and loaded from a file like object. When downloading a file, set this to True to force the code to deserialize the content to a FileSchema file
timeout | typing.Optional[typing.Union[int, typing.Tuple]] | default is None | the timeout used by the rest client
skip_deserialization | bool | default is False | when True, headers and body will be unset and an instance of api_client.ApiResponseWithoutDeserialization will be returned

### body

# SchemaForRequestBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**EstateRecordsRequestBody**](../../models/EstateRecordsRequestBody.md) |  | 


### Return Types, Responses

Code | Class | Description
------------- | ------------- | -------------
n/a | api_client.ApiResponseWithoutDeserialization | When skip_deserialization is True this response is returned
202 | [ApiResponseFor202](#v1_estate_records_post.ApiResponseFor202) | Status Accepted
400 | [ApiResponseFor400](#v1_estate_records_post.ApiResponseFor400) | Bad request
401 | [ApiResponseFor401](#v1_estate_records_post.ApiResponseFor401) | Unauthorized response
429 | [ApiResponseFor429](#v1_estate_records_post.ApiResponseFor429) | Too many requests
500 | [ApiResponseFor500](#v1_estate_records_post.ApiResponseFor500) | Internal server error

#### v1_estate_records_post.ApiResponseFor202
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor202ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor202ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**AcceptedResponseBody**](../../models/AcceptedResponseBody.md) |  | 


#### v1_estate_records_post.ApiResponseFor400
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor400ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor400ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_post.ApiResponseFor401
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor401ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor401ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_post.ApiResponseFor429
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor429ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor429ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


#### v1_estate_records_post.ApiResponseFor500
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
response | urllib3.HTTPResponse | Raw response |
body | typing.Union[SchemaFor500ResponseBodyApplicationJson, ] |  |
headers | Unset | headers were not defined |

# SchemaFor500ResponseBodyApplicationJson
Type | Description  | Notes
------------- | ------------- | -------------
[**ErrorResponse**](../../models/ErrorResponse.md) |  | 


### Authorization

[bearerAuth](../../../README.md#bearerAuth)

[[Back to top]](#__pageTop) [[Back to API list]](../../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../../README.md#documentation-for-models) [[Back to README]](../../../README.md)

