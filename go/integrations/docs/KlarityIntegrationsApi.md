# \KlarityIntegrationsApi

All URIs are relative to *https://integrations-api.klarity.nordcloudapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EstateRecordsPost**](KlarityIntegrationsApi.md#V1EstateRecordsPost) | **Post** /v1/estateRecords | Manage Klarity estate records



## V1EstateRecordsPost

> AcceptedResponseBody V1EstateRecordsPost(ctx).EstateRecordsRequestBody(estateRecordsRequestBody).Execute()

Manage Klarity estate records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    estateRecordsRequestBody := *openapiclient.NewEstateRecordsRequestBody() // EstateRecordsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.KlarityIntegrationsApi.V1EstateRecordsPost(context.Background()).EstateRecordsRequestBody(estateRecordsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarityIntegrationsApi.V1EstateRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EstateRecordsPost`: AcceptedResponseBody
    fmt.Fprintf(os.Stdout, "Response from `KlarityIntegrationsApi.V1EstateRecordsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EstateRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estateRecordsRequestBody** | [**EstateRecordsRequestBody**](EstateRecordsRequestBody.md) |  | 

### Return type

[**AcceptedResponseBody**](AcceptedResponseBody.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

