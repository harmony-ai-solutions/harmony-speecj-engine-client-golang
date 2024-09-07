/*
Harmony Speech Engine API

API for the Harmony Speech Engine.

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harmonyspeech

import (
	"encoding/json"
)

// checks if the TextToSpeechRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextToSpeechRequest{}

// TextToSpeechRequest TextToSpeechRequest Used to trigger a speech generation for the provided text input using the specified model. Depending on model selection, the caller may need to provide additional params. Based on OpenAI TTS API; extended for Harmony Speech Engine features.
type TextToSpeechRequest struct {
	// the name of the model
	Model *string `json:"model,omitempty"`
	// the text to synthesize
	Input *string `json:"input,omitempty"`
	Language NullableString `json:"language,omitempty"`
	Voice NullableString `json:"voice,omitempty"`
	InputAudio NullableString `json:"input_audio,omitempty"`
	InputVadMode NullableString `json:"input_vad_mode,omitempty"`
	InputVadData NullableString `json:"input_vad_data,omitempty"`
	InputEmbedding NullableString `json:"input_embedding,omitempty"`
	GenerationOptions NullableGenerationOptions `json:"generation_options,omitempty"`
	OutputOptions NullableAudioOutputOptions `json:"output_options,omitempty"`
	PostGenerationFilters []VoiceConversionRequest `json:"post_generation_filters,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TextToSpeechRequest TextToSpeechRequest

// NewTextToSpeechRequest instantiates a new TextToSpeechRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextToSpeechRequest() *TextToSpeechRequest {
	this := TextToSpeechRequest{}
	var model string = ""
	this.Model = &model
	var input string = ""
	this.Input = &input
	return &this
}

// NewTextToSpeechRequestWithDefaults instantiates a new TextToSpeechRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextToSpeechRequestWithDefaults() *TextToSpeechRequest {
	this := TextToSpeechRequest{}
	var model string = ""
	this.Model = &model
	var input string = ""
	this.Input = &input
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *TextToSpeechRequest) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToSpeechRequest) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *TextToSpeechRequest) SetModel(v string) {
	o.Model = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *TextToSpeechRequest) GetInput() string {
	if o == nil || IsNil(o.Input) {
		var ret string
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextToSpeechRequest) GetInputOk() (*string, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given string and assigns it to the Input field.
func (o *TextToSpeechRequest) SetInput(v string) {
	o.Input = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetLanguage() string {
	if o == nil || IsNil(o.Language.Get()) {
		var ret string
		return ret
	}
	return *o.Language.Get()
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Language.Get(), o.Language.IsSet()
}

// HasLanguage returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasLanguage() bool {
	if o != nil && o.Language.IsSet() {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given NullableString and assigns it to the Language field.
func (o *TextToSpeechRequest) SetLanguage(v string) {
	o.Language.Set(&v)
}
// SetLanguageNil sets the value for Language to be an explicit nil
func (o *TextToSpeechRequest) SetLanguageNil() {
	o.Language.Set(nil)
}

// UnsetLanguage ensures that no value is present for Language, not even an explicit nil
func (o *TextToSpeechRequest) UnsetLanguage() {
	o.Language.Unset()
}

// GetVoice returns the Voice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetVoice() string {
	if o == nil || IsNil(o.Voice.Get()) {
		var ret string
		return ret
	}
	return *o.Voice.Get()
}

// GetVoiceOk returns a tuple with the Voice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetVoiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Voice.Get(), o.Voice.IsSet()
}

// HasVoice returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasVoice() bool {
	if o != nil && o.Voice.IsSet() {
		return true
	}

	return false
}

// SetVoice gets a reference to the given NullableString and assigns it to the Voice field.
func (o *TextToSpeechRequest) SetVoice(v string) {
	o.Voice.Set(&v)
}
// SetVoiceNil sets the value for Voice to be an explicit nil
func (o *TextToSpeechRequest) SetVoiceNil() {
	o.Voice.Set(nil)
}

// UnsetVoice ensures that no value is present for Voice, not even an explicit nil
func (o *TextToSpeechRequest) UnsetVoice() {
	o.Voice.Unset()
}

// GetInputAudio returns the InputAudio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetInputAudio() string {
	if o == nil || IsNil(o.InputAudio.Get()) {
		var ret string
		return ret
	}
	return *o.InputAudio.Get()
}

// GetInputAudioOk returns a tuple with the InputAudio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetInputAudioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputAudio.Get(), o.InputAudio.IsSet()
}

// HasInputAudio returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasInputAudio() bool {
	if o != nil && o.InputAudio.IsSet() {
		return true
	}

	return false
}

// SetInputAudio gets a reference to the given NullableString and assigns it to the InputAudio field.
func (o *TextToSpeechRequest) SetInputAudio(v string) {
	o.InputAudio.Set(&v)
}
// SetInputAudioNil sets the value for InputAudio to be an explicit nil
func (o *TextToSpeechRequest) SetInputAudioNil() {
	o.InputAudio.Set(nil)
}

// UnsetInputAudio ensures that no value is present for InputAudio, not even an explicit nil
func (o *TextToSpeechRequest) UnsetInputAudio() {
	o.InputAudio.Unset()
}

// GetInputVadMode returns the InputVadMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetInputVadMode() string {
	if o == nil || IsNil(o.InputVadMode.Get()) {
		var ret string
		return ret
	}
	return *o.InputVadMode.Get()
}

// GetInputVadModeOk returns a tuple with the InputVadMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetInputVadModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputVadMode.Get(), o.InputVadMode.IsSet()
}

// HasInputVadMode returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasInputVadMode() bool {
	if o != nil && o.InputVadMode.IsSet() {
		return true
	}

	return false
}

// SetInputVadMode gets a reference to the given NullableString and assigns it to the InputVadMode field.
func (o *TextToSpeechRequest) SetInputVadMode(v string) {
	o.InputVadMode.Set(&v)
}
// SetInputVadModeNil sets the value for InputVadMode to be an explicit nil
func (o *TextToSpeechRequest) SetInputVadModeNil() {
	o.InputVadMode.Set(nil)
}

// UnsetInputVadMode ensures that no value is present for InputVadMode, not even an explicit nil
func (o *TextToSpeechRequest) UnsetInputVadMode() {
	o.InputVadMode.Unset()
}

// GetInputVadData returns the InputVadData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetInputVadData() string {
	if o == nil || IsNil(o.InputVadData.Get()) {
		var ret string
		return ret
	}
	return *o.InputVadData.Get()
}

// GetInputVadDataOk returns a tuple with the InputVadData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetInputVadDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputVadData.Get(), o.InputVadData.IsSet()
}

// HasInputVadData returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasInputVadData() bool {
	if o != nil && o.InputVadData.IsSet() {
		return true
	}

	return false
}

// SetInputVadData gets a reference to the given NullableString and assigns it to the InputVadData field.
func (o *TextToSpeechRequest) SetInputVadData(v string) {
	o.InputVadData.Set(&v)
}
// SetInputVadDataNil sets the value for InputVadData to be an explicit nil
func (o *TextToSpeechRequest) SetInputVadDataNil() {
	o.InputVadData.Set(nil)
}

// UnsetInputVadData ensures that no value is present for InputVadData, not even an explicit nil
func (o *TextToSpeechRequest) UnsetInputVadData() {
	o.InputVadData.Unset()
}

// GetInputEmbedding returns the InputEmbedding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetInputEmbedding() string {
	if o == nil || IsNil(o.InputEmbedding.Get()) {
		var ret string
		return ret
	}
	return *o.InputEmbedding.Get()
}

// GetInputEmbeddingOk returns a tuple with the InputEmbedding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetInputEmbeddingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputEmbedding.Get(), o.InputEmbedding.IsSet()
}

// HasInputEmbedding returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasInputEmbedding() bool {
	if o != nil && o.InputEmbedding.IsSet() {
		return true
	}

	return false
}

// SetInputEmbedding gets a reference to the given NullableString and assigns it to the InputEmbedding field.
func (o *TextToSpeechRequest) SetInputEmbedding(v string) {
	o.InputEmbedding.Set(&v)
}
// SetInputEmbeddingNil sets the value for InputEmbedding to be an explicit nil
func (o *TextToSpeechRequest) SetInputEmbeddingNil() {
	o.InputEmbedding.Set(nil)
}

// UnsetInputEmbedding ensures that no value is present for InputEmbedding, not even an explicit nil
func (o *TextToSpeechRequest) UnsetInputEmbedding() {
	o.InputEmbedding.Unset()
}

// GetGenerationOptions returns the GenerationOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetGenerationOptions() GenerationOptions {
	if o == nil || IsNil(o.GenerationOptions.Get()) {
		var ret GenerationOptions
		return ret
	}
	return *o.GenerationOptions.Get()
}

// GetGenerationOptionsOk returns a tuple with the GenerationOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetGenerationOptionsOk() (*GenerationOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenerationOptions.Get(), o.GenerationOptions.IsSet()
}

// HasGenerationOptions returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasGenerationOptions() bool {
	if o != nil && o.GenerationOptions.IsSet() {
		return true
	}

	return false
}

// SetGenerationOptions gets a reference to the given NullableGenerationOptions and assigns it to the GenerationOptions field.
func (o *TextToSpeechRequest) SetGenerationOptions(v GenerationOptions) {
	o.GenerationOptions.Set(&v)
}
// SetGenerationOptionsNil sets the value for GenerationOptions to be an explicit nil
func (o *TextToSpeechRequest) SetGenerationOptionsNil() {
	o.GenerationOptions.Set(nil)
}

// UnsetGenerationOptions ensures that no value is present for GenerationOptions, not even an explicit nil
func (o *TextToSpeechRequest) UnsetGenerationOptions() {
	o.GenerationOptions.Unset()
}

// GetOutputOptions returns the OutputOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetOutputOptions() AudioOutputOptions {
	if o == nil || IsNil(o.OutputOptions.Get()) {
		var ret AudioOutputOptions
		return ret
	}
	return *o.OutputOptions.Get()
}

// GetOutputOptionsOk returns a tuple with the OutputOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetOutputOptionsOk() (*AudioOutputOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.OutputOptions.Get(), o.OutputOptions.IsSet()
}

// HasOutputOptions returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasOutputOptions() bool {
	if o != nil && o.OutputOptions.IsSet() {
		return true
	}

	return false
}

// SetOutputOptions gets a reference to the given NullableAudioOutputOptions and assigns it to the OutputOptions field.
func (o *TextToSpeechRequest) SetOutputOptions(v AudioOutputOptions) {
	o.OutputOptions.Set(&v)
}
// SetOutputOptionsNil sets the value for OutputOptions to be an explicit nil
func (o *TextToSpeechRequest) SetOutputOptionsNil() {
	o.OutputOptions.Set(nil)
}

// UnsetOutputOptions ensures that no value is present for OutputOptions, not even an explicit nil
func (o *TextToSpeechRequest) UnsetOutputOptions() {
	o.OutputOptions.Unset()
}

// GetPostGenerationFilters returns the PostGenerationFilters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TextToSpeechRequest) GetPostGenerationFilters() []VoiceConversionRequest {
	if o == nil {
		var ret []VoiceConversionRequest
		return ret
	}
	return o.PostGenerationFilters
}

// GetPostGenerationFiltersOk returns a tuple with the PostGenerationFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TextToSpeechRequest) GetPostGenerationFiltersOk() ([]VoiceConversionRequest, bool) {
	if o == nil || IsNil(o.PostGenerationFilters) {
		return nil, false
	}
	return o.PostGenerationFilters, true
}

// HasPostGenerationFilters returns a boolean if a field has been set.
func (o *TextToSpeechRequest) HasPostGenerationFilters() bool {
	if o != nil && !IsNil(o.PostGenerationFilters) {
		return true
	}

	return false
}

// SetPostGenerationFilters gets a reference to the given []VoiceConversionRequest and assigns it to the PostGenerationFilters field.
func (o *TextToSpeechRequest) SetPostGenerationFilters(v []VoiceConversionRequest) {
	o.PostGenerationFilters = v
}

func (o TextToSpeechRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextToSpeechRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if o.Language.IsSet() {
		toSerialize["language"] = o.Language.Get()
	}
	if o.Voice.IsSet() {
		toSerialize["voice"] = o.Voice.Get()
	}
	if o.InputAudio.IsSet() {
		toSerialize["input_audio"] = o.InputAudio.Get()
	}
	if o.InputVadMode.IsSet() {
		toSerialize["input_vad_mode"] = o.InputVadMode.Get()
	}
	if o.InputVadData.IsSet() {
		toSerialize["input_vad_data"] = o.InputVadData.Get()
	}
	if o.InputEmbedding.IsSet() {
		toSerialize["input_embedding"] = o.InputEmbedding.Get()
	}
	if o.GenerationOptions.IsSet() {
		toSerialize["generation_options"] = o.GenerationOptions.Get()
	}
	if o.OutputOptions.IsSet() {
		toSerialize["output_options"] = o.OutputOptions.Get()
	}
	if o.PostGenerationFilters != nil {
		toSerialize["post_generation_filters"] = o.PostGenerationFilters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TextToSpeechRequest) UnmarshalJSON(data []byte) (err error) {
	varTextToSpeechRequest := _TextToSpeechRequest{}

	err = json.Unmarshal(data, &varTextToSpeechRequest)

	if err != nil {
		return err
	}

	*o = TextToSpeechRequest(varTextToSpeechRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "model")
		delete(additionalProperties, "input")
		delete(additionalProperties, "language")
		delete(additionalProperties, "voice")
		delete(additionalProperties, "input_audio")
		delete(additionalProperties, "input_vad_mode")
		delete(additionalProperties, "input_vad_data")
		delete(additionalProperties, "input_embedding")
		delete(additionalProperties, "generation_options")
		delete(additionalProperties, "output_options")
		delete(additionalProperties, "post_generation_filters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTextToSpeechRequest struct {
	value *TextToSpeechRequest
	isSet bool
}

func (v NullableTextToSpeechRequest) Get() *TextToSpeechRequest {
	return v.value
}

func (v *NullableTextToSpeechRequest) Set(val *TextToSpeechRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTextToSpeechRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTextToSpeechRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextToSpeechRequest(val *TextToSpeechRequest) *NullableTextToSpeechRequest {
	return &NullableTextToSpeechRequest{value: val, isSet: true}
}

func (v NullableTextToSpeechRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextToSpeechRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


