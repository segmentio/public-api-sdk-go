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

// checks if the UpdateFilterByIdOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilterByIdOutput{}

// UpdateFilterByIdOutput Output for UpdateFilterById.
type UpdateFilterByIdOutput struct {
	Filter Filter `json:"filter"`
}

// NewUpdateFilterByIdOutput instantiates a new UpdateFilterByIdOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterByIdOutput(filter Filter) *UpdateFilterByIdOutput {
	this := UpdateFilterByIdOutput{}
	this.Filter = filter
	return &this
}

// NewUpdateFilterByIdOutputWithDefaults instantiates a new UpdateFilterByIdOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterByIdOutputWithDefaults() *UpdateFilterByIdOutput {
	this := UpdateFilterByIdOutput{}
	return &this
}

// GetFilter returns the Filter field value
func (o *UpdateFilterByIdOutput) GetFilter() Filter {
	if o == nil {
		var ret Filter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdOutput) GetFilterOk() (*Filter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *UpdateFilterByIdOutput) SetFilter(v Filter) {
	o.Filter = v
}

func (o UpdateFilterByIdOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilterByIdOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	return toSerialize, nil
}

type NullableUpdateFilterByIdOutput struct {
	value *UpdateFilterByIdOutput
	isSet bool
}

func (v NullableUpdateFilterByIdOutput) Get() *UpdateFilterByIdOutput {
	return v.value
}

func (v *NullableUpdateFilterByIdOutput) Set(val *UpdateFilterByIdOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterByIdOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterByIdOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterByIdOutput(
	val *UpdateFilterByIdOutput,
) *NullableUpdateFilterByIdOutput {
	return &NullableUpdateFilterByIdOutput{value: val, isSet: true}
}

func (v NullableUpdateFilterByIdOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterByIdOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
