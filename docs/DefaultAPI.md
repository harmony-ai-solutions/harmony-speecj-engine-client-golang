# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertVoiceV1VoiceConvertPost**](DefaultAPI.md#ConvertVoiceV1VoiceConvertPost) | **Post** /v1/voice/convert | Convert Voice
[**CreateEmbeddingV1EmbedSpeakerPost**](DefaultAPI.md#CreateEmbeddingV1EmbedSpeakerPost) | **Post** /v1/embed/speaker | Create Embedding
[**CreateSpeechV1AudioSpeechPost**](DefaultAPI.md#CreateSpeechV1AudioSpeechPost) | **Post** /v1/audio/speech | Create Speech
[**CreateTranscriptionV1AudioTranscriptionsPost**](DefaultAPI.md#CreateTranscriptionV1AudioTranscriptionsPost) | **Post** /v1/audio/transcriptions | Create Transcription
[**CreateVadV1AudioVadPost**](DefaultAPI.md#CreateVadV1AudioVadPost) | **Post** /v1/audio/vad | Create Vad
[**HealthHealthGet**](DefaultAPI.md#HealthHealthGet) | **Get** /health | Health
[**ShowAvailableEmbeddingModelsV1EmbedModelsGet**](DefaultAPI.md#ShowAvailableEmbeddingModelsV1EmbedModelsGet) | **Get** /v1/embed/models | Show Available Embedding Models
[**ShowAvailableSpeechModelsV1AudioSpeechModelsGet**](DefaultAPI.md#ShowAvailableSpeechModelsV1AudioSpeechModelsGet) | **Get** /v1/audio/speech/models | Show Available Speech Models
[**ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet**](DefaultAPI.md#ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet) | **Get** /v1/audio/transcriptions/models | Show Available Transcription Models
[**ShowAvailableVadModelsV1AudioVadModelsGet**](DefaultAPI.md#ShowAvailableVadModelsV1AudioVadModelsGet) | **Get** /v1/audio/vad/models | Show Available Vad Models
[**ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet**](DefaultAPI.md#ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet) | **Get** /v1/voice/convert/models | Show Available Voice Conversion Models
[**ShowVersionVersionGet**](DefaultAPI.md#ShowVersionVersionGet) | **Get** /version | Show Version



## ConvertVoiceV1VoiceConvertPost

> VoiceConversionResponse ConvertVoiceV1VoiceConvertPost(ctx).VoiceConversionRequest(voiceConversionRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Convert Voice



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	voiceConversionRequest := *openapiclient.NewVoiceConversionRequest() // VoiceConversionRequest | 
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ConvertVoiceV1VoiceConvertPost(context.Background()).VoiceConversionRequest(voiceConversionRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ConvertVoiceV1VoiceConvertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConvertVoiceV1VoiceConvertPost`: VoiceConversionResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ConvertVoiceV1VoiceConvertPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConvertVoiceV1VoiceConvertPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceConversionRequest** | [**VoiceConversionRequest**](VoiceConversionRequest.md) |  | 
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**VoiceConversionResponse**](VoiceConversionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmbeddingV1EmbedSpeakerPost

> EmbedSpeakerResponse CreateEmbeddingV1EmbedSpeakerPost(ctx).EmbedSpeakerRequest(embedSpeakerRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Create Embedding



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	embedSpeakerRequest := *openapiclient.NewEmbedSpeakerRequest() // EmbedSpeakerRequest | 
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateEmbeddingV1EmbedSpeakerPost(context.Background()).EmbedSpeakerRequest(embedSpeakerRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateEmbeddingV1EmbedSpeakerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmbeddingV1EmbedSpeakerPost`: EmbedSpeakerResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateEmbeddingV1EmbedSpeakerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmbeddingV1EmbedSpeakerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embedSpeakerRequest** | [**EmbedSpeakerRequest**](EmbedSpeakerRequest.md) |  | 
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**EmbedSpeakerResponse**](EmbedSpeakerResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSpeechV1AudioSpeechPost

> TextToSpeechResponse CreateSpeechV1AudioSpeechPost(ctx).TextToSpeechRequest(textToSpeechRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Create Speech



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	textToSpeechRequest := *openapiclient.NewTextToSpeechRequest() // TextToSpeechRequest | 
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateSpeechV1AudioSpeechPost(context.Background()).TextToSpeechRequest(textToSpeechRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateSpeechV1AudioSpeechPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSpeechV1AudioSpeechPost`: TextToSpeechResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateSpeechV1AudioSpeechPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSpeechV1AudioSpeechPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **textToSpeechRequest** | [**TextToSpeechRequest**](TextToSpeechRequest.md) |  | 
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**TextToSpeechResponse**](TextToSpeechResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTranscriptionV1AudioTranscriptionsPost

> SpeechToTextResponse CreateTranscriptionV1AudioTranscriptionsPost(ctx).SpeechTranscribeRequest(speechTranscribeRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Create Transcription



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	speechTranscribeRequest := *openapiclient.NewSpeechTranscribeRequest() // SpeechTranscribeRequest | 
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTranscriptionV1AudioTranscriptionsPost(context.Background()).SpeechTranscribeRequest(speechTranscribeRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTranscriptionV1AudioTranscriptionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTranscriptionV1AudioTranscriptionsPost`: SpeechToTextResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTranscriptionV1AudioTranscriptionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTranscriptionV1AudioTranscriptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **speechTranscribeRequest** | [**SpeechTranscribeRequest**](SpeechTranscribeRequest.md) |  | 
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**SpeechToTextResponse**](SpeechToTextResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVadV1AudioVadPost

> DetectVoiceActivityResponse CreateVadV1AudioVadPost(ctx).DetectVoiceActivityRequest(detectVoiceActivityRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Create Vad



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	detectVoiceActivityRequest := *openapiclient.NewDetectVoiceActivityRequest() // DetectVoiceActivityRequest | 
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateVadV1AudioVadPost(context.Background()).DetectVoiceActivityRequest(detectVoiceActivityRequest).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateVadV1AudioVadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVadV1AudioVadPost`: DetectVoiceActivityResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateVadV1AudioVadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVadV1AudioVadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detectVoiceActivityRequest** | [**DetectVoiceActivityRequest**](DetectVoiceActivityRequest.md) |  | 
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**DetectVoiceActivityResponse**](DetectVoiceActivityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthHealthGet

> map[string]interface{} HealthHealthGet(ctx).Execute()

Health



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HealthHealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthHealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthHealthGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthHealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthHealthGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAvailableEmbeddingModelsV1EmbedModelsGet

> ModelList ShowAvailableEmbeddingModelsV1EmbedModelsGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Available Embedding Models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowAvailableEmbeddingModelsV1EmbedModelsGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowAvailableEmbeddingModelsV1EmbedModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAvailableEmbeddingModelsV1EmbedModelsGet`: ModelList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowAvailableEmbeddingModelsV1EmbedModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowAvailableEmbeddingModelsV1EmbedModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**ModelList**](ModelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAvailableSpeechModelsV1AudioSpeechModelsGet

> ModelList ShowAvailableSpeechModelsV1AudioSpeechModelsGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Available Speech Models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowAvailableSpeechModelsV1AudioSpeechModelsGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowAvailableSpeechModelsV1AudioSpeechModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAvailableSpeechModelsV1AudioSpeechModelsGet`: ModelList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowAvailableSpeechModelsV1AudioSpeechModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowAvailableSpeechModelsV1AudioSpeechModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**ModelList**](ModelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet

> ModelList ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Available Transcription Models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet`: ModelList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**ModelList**](ModelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAvailableVadModelsV1AudioVadModelsGet

> ModelList ShowAvailableVadModelsV1AudioVadModelsGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Available Vad Models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowAvailableVadModelsV1AudioVadModelsGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowAvailableVadModelsV1AudioVadModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAvailableVadModelsV1AudioVadModelsGet`: ModelList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowAvailableVadModelsV1AudioVadModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowAvailableVadModelsV1AudioVadModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

[**ModelList**](ModelList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet

> interface{} ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Available Voice Conversion Models



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowAvailableVoiceConversionModelsV1VoiceConvertModelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowVersionVersionGet

> map[string]interface{} ShowVersionVersionGet(ctx).XApiKey(xApiKey).ApiKey(apiKey).Execute()

Show Version



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
)

func main() {
	xApiKey := "xApiKey_example" // string |  (optional)
	apiKey := "apiKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ShowVersionVersionGet(context.Background()).XApiKey(xApiKey).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ShowVersionVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowVersionVersionGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ShowVersionVersionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowVersionVersionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xApiKey** | **string** |  | 
 **apiKey** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

