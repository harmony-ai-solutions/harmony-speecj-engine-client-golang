# Go API client for Harmony Speech Engine

Simple Golang Client for [Harmony Speech Engine](https://github.com/harmony-ai-solutions/harmony-speech-engine).

!! PRE-RELEASE !!
Package Naming not final

(generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: v0.0.3
- Generator version: 7.9.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import harmonyspeech "github.com/harmony-ai-solutions/harmony-speech-engine-client-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `harmonyspeech.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), harmonyspeech.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `harmonyspeech.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), harmonyspeech.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `harmonyspeech.ContextOperationServerIndices` and `harmonyspeech.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), harmonyspeech.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), harmonyspeech.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**ConvertVoiceV1VoiceConvertPost**](docs/DefaultAPI.md#convertvoicev1voiceconvertpost) | **Post** /v1/voice/convert | Convert Voice
*DefaultAPI* | [**CreateEmbeddingV1EmbedSpeakerPost**](docs/DefaultAPI.md#createembeddingv1embedspeakerpost) | **Post** /v1/embed/speaker | Create Embedding
*DefaultAPI* | [**CreateSpeechV1AudioSpeechPost**](docs/DefaultAPI.md#createspeechv1audiospeechpost) | **Post** /v1/audio/speech | Create Speech
*DefaultAPI* | [**CreateTranscriptionV1AudioTranscriptionsPost**](docs/DefaultAPI.md#createtranscriptionv1audiotranscriptionspost) | **Post** /v1/audio/transcriptions | Create Transcription
*DefaultAPI* | [**HealthHealthGet**](docs/DefaultAPI.md#healthhealthget) | **Get** /health | Health
*DefaultAPI* | [**ShowAvailableEmbeddingModelsV1EmbedModelsGet**](docs/DefaultAPI.md#showavailableembeddingmodelsv1embedmodelsget) | **Get** /v1/embed/models | Show Available Embedding Models
*DefaultAPI* | [**ShowAvailableSpeechModelsV1AudioSpeechModelsGet**](docs/DefaultAPI.md#showavailablespeechmodelsv1audiospeechmodelsget) | **Get** /v1/audio/speech/models | Show Available Speech Models
*DefaultAPI* | [**ShowAvailableTranscriptionModelsV1AudioTranscriptionsModelsGet**](docs/DefaultAPI.md#showavailabletranscriptionmodelsv1audiotranscriptionsmodelsget) | **Get** /v1/audio/transcriptions/models | Show Available Transcription Models
*DefaultAPI* | [**ShowAvailableVoiceConversionModelsV1VoiceConvertModelsGet**](docs/DefaultAPI.md#showavailablevoiceconversionmodelsv1voiceconvertmodelsget) | **Get** /v1/voice/convert/models | Show Available Voice Conversion Models
*DefaultAPI* | [**ShowVersionVersionGet**](docs/DefaultAPI.md#showversionversionget) | **Get** /version | Show Version


## Documentation For Models

 - [AudioOutputOptions](docs/AudioOutputOptions.md)
 - [EmbedSpeakerRequest](docs/EmbedSpeakerRequest.md)
 - [EmbedSpeakerResponse](docs/EmbedSpeakerResponse.md)
 - [GenerationOptions](docs/GenerationOptions.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [ModelCard](docs/ModelCard.md)
 - [ModelList](docs/ModelList.md)
 - [SpeechToTextResponse](docs/SpeechToTextResponse.md)
 - [SpeechTranscribeRequest](docs/SpeechTranscribeRequest.md)
 - [TextToSpeechRequest](docs/TextToSpeechRequest.md)
 - [TextToSpeechResponse](docs/TextToSpeechResponse.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationErrorLocInner](docs/ValidationErrorLocInner.md)
 - [VoiceConversionRequest](docs/VoiceConversionRequest.md)
 - [VoiceConversionResponse](docs/VoiceConversionResponse.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



