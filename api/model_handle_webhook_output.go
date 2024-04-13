/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 48.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the HandleWebhookOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandleWebhookOutput{}

// HandleWebhookOutput Function webhook output status.
type HandleWebhookOutput struct {
	// The http status code.
	StatusCode float32 `json:"statusCode"`
	// The status of the operation.
	Success bool `json:"success"`
}

// NewHandleWebhookOutput instantiates a new HandleWebhookOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandleWebhookOutput(statusCode float32, success bool) *HandleWebhookOutput {
	this := HandleWebhookOutput{}
	this.StatusCode = statusCode
	this.Success = success
	return &this
}

// NewHandleWebhookOutputWithDefaults instantiates a new HandleWebhookOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandleWebhookOutputWithDefaults() *HandleWebhookOutput {
	this := HandleWebhookOutput{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *HandleWebhookOutput) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookOutput) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *HandleWebhookOutput) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetSuccess returns the Success field value
func (o *HandleWebhookOutput) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookOutput) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *HandleWebhookOutput) SetSuccess(v bool) {
	o.Success = v
}

func (o HandleWebhookOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandleWebhookOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusCode"] = o.StatusCode
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

type NullableHandleWebhookOutput struct {
	value *HandleWebhookOutput
	isSet bool
}

func (v NullableHandleWebhookOutput) Get() *HandleWebhookOutput {
	return v.value
}

func (v *NullableHandleWebhookOutput) Set(val *HandleWebhookOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableHandleWebhookOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableHandleWebhookOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandleWebhookOutput(val *HandleWebhookOutput) *NullableHandleWebhookOutput {
	return &NullableHandleWebhookOutput{value: val, isSet: true}
}

func (v NullableHandleWebhookOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandleWebhookOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
