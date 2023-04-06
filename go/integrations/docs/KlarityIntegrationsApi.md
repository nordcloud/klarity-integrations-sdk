# \KlarityIntegrationsApi

All URIs are relative to *https://integrations-api.klarity.nordcloudapp.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnrichmentsEstateRecordsPost**](KlarityIntegrationsApi.md#V1EnrichmentsEstateRecordsPost) | **Post** /v1/enrichments/estateRecords | Enrich Klarity estate records
[**V1EstateRecordsDelete**](KlarityIntegrationsApi.md#V1EstateRecordsDelete) | **Delete** /v1/estateRecords | Delete Klarity estate records
[**V1EstateRecordsPost**](KlarityIntegrationsApi.md#V1EstateRecordsPost) | **Post** /v1/estateRecords | Manage Klarity estate records



## V1EnrichmentsEstateRecordsPost

> AcceptedResponseBody V1EnrichmentsEstateRecordsPost(ctx).EnrichmentsEstateRecordsRequestBody(enrichmentsEstateRecordsRequestBody).Execute()

Enrich Klarity estate records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nordcloud/klarity-integrations-sdk/go/integrations"
)

func main() {
    enrichmentsEstateRecordsRequestBody := *openapiclient.NewEnrichmentsEstateRecordsRequestBody([]openapiclient.EnrichmentsEstateRecordsRequestBodyEnrichmentsInner{*openapiclient.NewEnrichmentsEstateRecordsRequestBodyEnrichmentsInner(*openapiclient.NewEnrichmentRecord(), map[string]interface{}({"lastBackupTime":"2021-10-19T07:10:38.147Z","details":{"succeeded":true,"durationSeconds":1234}}))}) // EnrichmentsEstateRecordsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarityIntegrationsApi.V1EnrichmentsEstateRecordsPost(context.Background()).EnrichmentsEstateRecordsRequestBody(enrichmentsEstateRecordsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarityIntegrationsApi.V1EnrichmentsEstateRecordsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnrichmentsEstateRecordsPost`: AcceptedResponseBody
    fmt.Fprintf(os.Stdout, "Response from `KlarityIntegrationsApi.V1EnrichmentsEstateRecordsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrichmentsEstateRecordsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enrichmentsEstateRecordsRequestBody** | [**EnrichmentsEstateRecordsRequestBody**](EnrichmentsEstateRecordsRequestBody.md) |  | 

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


## V1EstateRecordsDelete

> AcceptedResponseBody V1EstateRecordsDelete(ctx).EstateRecordsDeleteBody(estateRecordsDeleteBody).Execute()

Delete Klarity estate records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nordcloud/klarity-integrations-sdk/go/integrations"
)

func main() {
    estateRecordsDeleteBody := *openapiclient.NewEstateRecordsDeleteBody(openapiclient.PeriodEnum("current"), []openapiclient.EstateRecordsDeleteBodyRecordsInner{*openapiclient.NewEstateRecordsDeleteBodyRecordsInner()}) // EstateRecordsDeleteBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarityIntegrationsApi.V1EstateRecordsDelete(context.Background()).EstateRecordsDeleteBody(estateRecordsDeleteBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarityIntegrationsApi.V1EstateRecordsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EstateRecordsDelete`: AcceptedResponseBody
    fmt.Fprintf(os.Stdout, "Response from `KlarityIntegrationsApi.V1EstateRecordsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EstateRecordsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **estateRecordsDeleteBody** | [**EstateRecordsDeleteBody**](EstateRecordsDeleteBody.md) |  | 

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
    openapiclient "github.com/nordcloud/klarity-integrations-sdk/go/integrations"
)

func main() {
    estateRecordsRequestBody := *openapiclient.NewEstateRecordsRequestBody([]openapiclient.EstateRecordsRequestBodyRecordsInner{*openapiclient.NewEstateRecordsRequestBodyRecordsInner("js002", "jira/subscription/basic")}) // EstateRecordsRequestBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarityIntegrationsApi.V1EstateRecordsPost(context.Background()).EstateRecordsRequestBody(estateRecordsRequestBody).Execute()
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

