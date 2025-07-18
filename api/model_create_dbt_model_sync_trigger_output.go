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

// checks if the CreateDbtModelSyncTriggerOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDbtModelSyncTriggerOutput{}

// CreateDbtModelSyncTriggerOutput Output for the createDbtModelSyncTriggerBySourceId endpoint.
type CreateDbtModelSyncTriggerOutput struct {
	DbtModelSyncTrigger DbtModelSyncTrigger `json:"dbtModelSyncTrigger"`
}

// NewCreateDbtModelSyncTriggerOutput instantiates a new CreateDbtModelSyncTriggerOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDbtModelSyncTriggerOutput(
	dbtModelSyncTrigger DbtModelSyncTrigger,
) *CreateDbtModelSyncTriggerOutput {
	this := CreateDbtModelSyncTriggerOutput{}
	this.DbtModelSyncTrigger = dbtModelSyncTrigger
	return &this
}

// NewCreateDbtModelSyncTriggerOutputWithDefaults instantiates a new CreateDbtModelSyncTriggerOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDbtModelSyncTriggerOutputWithDefaults() *CreateDbtModelSyncTriggerOutput {
	this := CreateDbtModelSyncTriggerOutput{}
	return &this
}

// GetDbtModelSyncTrigger returns the DbtModelSyncTrigger field value
func (o *CreateDbtModelSyncTriggerOutput) GetDbtModelSyncTrigger() DbtModelSyncTrigger {
	if o == nil {
		var ret DbtModelSyncTrigger
		return ret
	}

	return o.DbtModelSyncTrigger
}

// GetDbtModelSyncTriggerOk returns a tuple with the DbtModelSyncTrigger field value
// and a boolean to check if the value has been set.
func (o *CreateDbtModelSyncTriggerOutput) GetDbtModelSyncTriggerOk() (*DbtModelSyncTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DbtModelSyncTrigger, true
}

// SetDbtModelSyncTrigger sets field value
func (o *CreateDbtModelSyncTriggerOutput) SetDbtModelSyncTrigger(v DbtModelSyncTrigger) {
	o.DbtModelSyncTrigger = v
}

func (o CreateDbtModelSyncTriggerOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDbtModelSyncTriggerOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dbtModelSyncTrigger"] = o.DbtModelSyncTrigger
	return toSerialize, nil
}

type NullableCreateDbtModelSyncTriggerOutput struct {
	value *CreateDbtModelSyncTriggerOutput
	isSet bool
}

func (v NullableCreateDbtModelSyncTriggerOutput) Get() *CreateDbtModelSyncTriggerOutput {
	return v.value
}

func (v *NullableCreateDbtModelSyncTriggerOutput) Set(val *CreateDbtModelSyncTriggerOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDbtModelSyncTriggerOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDbtModelSyncTriggerOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDbtModelSyncTriggerOutput(
	val *CreateDbtModelSyncTriggerOutput,
) *NullableCreateDbtModelSyncTriggerOutput {
	return &NullableCreateDbtModelSyncTriggerOutput{value: val, isSet: true}
}

func (v NullableCreateDbtModelSyncTriggerOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDbtModelSyncTriggerOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
