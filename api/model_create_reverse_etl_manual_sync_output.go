/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 56.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateReverseETLManualSyncOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateReverseETLManualSyncOutput{}

// CreateReverseETLManualSyncOutput Output for triggering a manual sync for a RETL connection.
type CreateReverseETLManualSyncOutput struct {
	ReverseETLManualSync ReverseETLManualSyncJobOutput `json:"reverseETLManualSync"`
}

// NewCreateReverseETLManualSyncOutput instantiates a new CreateReverseETLManualSyncOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReverseETLManualSyncOutput(
	reverseETLManualSync ReverseETLManualSyncJobOutput,
) *CreateReverseETLManualSyncOutput {
	this := CreateReverseETLManualSyncOutput{}
	this.ReverseETLManualSync = reverseETLManualSync
	return &this
}

// NewCreateReverseETLManualSyncOutputWithDefaults instantiates a new CreateReverseETLManualSyncOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReverseETLManualSyncOutputWithDefaults() *CreateReverseETLManualSyncOutput {
	this := CreateReverseETLManualSyncOutput{}
	return &this
}

// GetReverseETLManualSync returns the ReverseETLManualSync field value
func (o *CreateReverseETLManualSyncOutput) GetReverseETLManualSync() ReverseETLManualSyncJobOutput {
	if o == nil {
		var ret ReverseETLManualSyncJobOutput
		return ret
	}

	return o.ReverseETLManualSync
}

// GetReverseETLManualSyncOk returns a tuple with the ReverseETLManualSync field value
// and a boolean to check if the value has been set.
func (o *CreateReverseETLManualSyncOutput) GetReverseETLManualSyncOk() (*ReverseETLManualSyncJobOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseETLManualSync, true
}

// SetReverseETLManualSync sets field value
func (o *CreateReverseETLManualSyncOutput) SetReverseETLManualSync(
	v ReverseETLManualSyncJobOutput,
) {
	o.ReverseETLManualSync = v
}

func (o CreateReverseETLManualSyncOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateReverseETLManualSyncOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseETLManualSync"] = o.ReverseETLManualSync
	return toSerialize, nil
}

type NullableCreateReverseETLManualSyncOutput struct {
	value *CreateReverseETLManualSyncOutput
	isSet bool
}

func (v NullableCreateReverseETLManualSyncOutput) Get() *CreateReverseETLManualSyncOutput {
	return v.value
}

func (v *NullableCreateReverseETLManualSyncOutput) Set(val *CreateReverseETLManualSyncOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReverseETLManualSyncOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReverseETLManualSyncOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReverseETLManualSyncOutput(
	val *CreateReverseETLManualSyncOutput,
) *NullableCreateReverseETLManualSyncOutput {
	return &NullableCreateReverseETLManualSyncOutput{value: val, isSet: true}
}

func (v NullableCreateReverseETLManualSyncOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReverseETLManualSyncOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
