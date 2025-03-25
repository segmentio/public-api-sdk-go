/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveWriteKeyFromSourceAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveWriteKeyFromSourceAlphaOutput{}

// RemoveWriteKeyFromSourceAlphaOutput Returns the updated Source.
type RemoveWriteKeyFromSourceAlphaOutput struct {
	// The status of the operation.
	Status string `json:"status"`
}

// NewRemoveWriteKeyFromSourceAlphaOutput instantiates a new RemoveWriteKeyFromSourceAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveWriteKeyFromSourceAlphaOutput(status string) *RemoveWriteKeyFromSourceAlphaOutput {
	this := RemoveWriteKeyFromSourceAlphaOutput{}
	this.Status = status
	return &this
}

// NewRemoveWriteKeyFromSourceAlphaOutputWithDefaults instantiates a new RemoveWriteKeyFromSourceAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveWriteKeyFromSourceAlphaOutputWithDefaults() *RemoveWriteKeyFromSourceAlphaOutput {
	this := RemoveWriteKeyFromSourceAlphaOutput{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveWriteKeyFromSourceAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveWriteKeyFromSourceAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveWriteKeyFromSourceAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o RemoveWriteKeyFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveWriteKeyFromSourceAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableRemoveWriteKeyFromSourceAlphaOutput struct {
	value *RemoveWriteKeyFromSourceAlphaOutput
	isSet bool
}

func (v NullableRemoveWriteKeyFromSourceAlphaOutput) Get() *RemoveWriteKeyFromSourceAlphaOutput {
	return v.value
}

func (v *NullableRemoveWriteKeyFromSourceAlphaOutput) Set(
	val *RemoveWriteKeyFromSourceAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveWriteKeyFromSourceAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveWriteKeyFromSourceAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveWriteKeyFromSourceAlphaOutput(
	val *RemoveWriteKeyFromSourceAlphaOutput,
) *NullableRemoveWriteKeyFromSourceAlphaOutput {
	return &NullableRemoveWriteKeyFromSourceAlphaOutput{value: val, isSet: true}
}

func (v NullableRemoveWriteKeyFromSourceAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveWriteKeyFromSourceAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
