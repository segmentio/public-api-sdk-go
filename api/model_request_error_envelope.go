/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RequestErrorEnvelope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestErrorEnvelope{}

// RequestErrorEnvelope Envelope type for RequestErrors.
type RequestErrorEnvelope struct {
	Errors []RequestError `json:"errors"`
}

// NewRequestErrorEnvelope instantiates a new RequestErrorEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestErrorEnvelope(errors []RequestError) *RequestErrorEnvelope {
	this := RequestErrorEnvelope{}
	this.Errors = errors
	return &this
}

// NewRequestErrorEnvelopeWithDefaults instantiates a new RequestErrorEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestErrorEnvelopeWithDefaults() *RequestErrorEnvelope {
	this := RequestErrorEnvelope{}
	return &this
}

// GetErrors returns the Errors field value
func (o *RequestErrorEnvelope) GetErrors() []RequestError {
	if o == nil {
		var ret []RequestError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *RequestErrorEnvelope) GetErrorsOk() ([]RequestError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *RequestErrorEnvelope) SetErrors(v []RequestError) {
	o.Errors = v
}

func (o RequestErrorEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestErrorEnvelope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableRequestErrorEnvelope struct {
	value *RequestErrorEnvelope
	isSet bool
}

func (v NullableRequestErrorEnvelope) Get() *RequestErrorEnvelope {
	return v.value
}

func (v *NullableRequestErrorEnvelope) Set(val *RequestErrorEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestErrorEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestErrorEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestErrorEnvelope(val *RequestErrorEnvelope) *NullableRequestErrorEnvelope {
	return &NullableRequestErrorEnvelope{value: val, isSet: true}
}

func (v NullableRequestErrorEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestErrorEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
