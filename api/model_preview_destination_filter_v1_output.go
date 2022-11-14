/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.5
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PreviewDestinationFilterV1Output Preview output from applying the Destination filter. Segment modifies or nullifies payloads depending on the provided filter actions.
type PreviewDestinationFilterV1Output struct {
	// The pre-filter JSON input.
	InputPayload map[string]interface{} `json:"inputPayload"`
	// The filtered JSON output.
	Result NullableModelMap `json:"result"`
}

// NewPreviewDestinationFilterV1Output instantiates a new PreviewDestinationFilterV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewDestinationFilterV1Output(inputPayload map[string]interface{}, result NullableModelMap) *PreviewDestinationFilterV1Output {
	this := PreviewDestinationFilterV1Output{}
	this.InputPayload = inputPayload
	this.Result = result
	return &this
}

// NewPreviewDestinationFilterV1OutputWithDefaults instantiates a new PreviewDestinationFilterV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewDestinationFilterV1OutputWithDefaults() *PreviewDestinationFilterV1Output {
	this := PreviewDestinationFilterV1Output{}
	return &this
}

// GetInputPayload returns the InputPayload field value
func (o *PreviewDestinationFilterV1Output) GetInputPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.InputPayload
}

// GetInputPayloadOk returns a tuple with the InputPayload field value
// and a boolean to check if the value has been set.
func (o *PreviewDestinationFilterV1Output) GetInputPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.InputPayload, true
}

// SetInputPayload sets field value
func (o *PreviewDestinationFilterV1Output) SetInputPayload(v map[string]interface{}) {
	o.InputPayload = v
}

// GetResult returns the Result field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *PreviewDestinationFilterV1Output) GetResult() ModelMap {
	if o == nil || o.Result.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PreviewDestinationFilterV1Output) GetResultOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// SetResult sets field value
func (o *PreviewDestinationFilterV1Output) SetResult(v ModelMap) {
	o.Result.Set(&v)
}

func (o PreviewDestinationFilterV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputPayload"] = o.InputPayload
	}
	if true {
		toSerialize["result"] = o.Result.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePreviewDestinationFilterV1Output struct {
	value *PreviewDestinationFilterV1Output
	isSet bool
}

func (v NullablePreviewDestinationFilterV1Output) Get() *PreviewDestinationFilterV1Output {
	return v.value
}

func (v *NullablePreviewDestinationFilterV1Output) Set(val *PreviewDestinationFilterV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewDestinationFilterV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewDestinationFilterV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewDestinationFilterV1Output(val *PreviewDestinationFilterV1Output) *NullablePreviewDestinationFilterV1Output {
	return &NullablePreviewDestinationFilterV1Output{value: val, isSet: true}
}

func (v NullablePreviewDestinationFilterV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewDestinationFilterV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


