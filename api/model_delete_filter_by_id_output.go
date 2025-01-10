/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeleteFilterByIdOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFilterByIdOutput{}

// DeleteFilterByIdOutput Output for DeleteFilterById.
type DeleteFilterByIdOutput struct {
	// Filter deleted by filter id.
	Deleted bool `json:"deleted"`
}

// NewDeleteFilterByIdOutput instantiates a new DeleteFilterByIdOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFilterByIdOutput(deleted bool) *DeleteFilterByIdOutput {
	this := DeleteFilterByIdOutput{}
	this.Deleted = deleted
	return &this
}

// NewDeleteFilterByIdOutputWithDefaults instantiates a new DeleteFilterByIdOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFilterByIdOutputWithDefaults() *DeleteFilterByIdOutput {
	this := DeleteFilterByIdOutput{}
	return &this
}

// GetDeleted returns the Deleted field value
func (o *DeleteFilterByIdOutput) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *DeleteFilterByIdOutput) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *DeleteFilterByIdOutput) SetDeleted(v bool) {
	o.Deleted = v
}

func (o DeleteFilterByIdOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFilterByIdOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

type NullableDeleteFilterByIdOutput struct {
	value *DeleteFilterByIdOutput
	isSet bool
}

func (v NullableDeleteFilterByIdOutput) Get() *DeleteFilterByIdOutput {
	return v.value
}

func (v *NullableDeleteFilterByIdOutput) Set(val *DeleteFilterByIdOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFilterByIdOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFilterByIdOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFilterByIdOutput(
	val *DeleteFilterByIdOutput,
) *NullableDeleteFilterByIdOutput {
	return &NullableDeleteFilterByIdOutput{value: val, isSet: true}
}

func (v NullableDeleteFilterByIdOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFilterByIdOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
