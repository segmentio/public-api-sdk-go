/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetReverseETLSyncStatusOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReverseETLSyncStatusOutput{}

// GetReverseETLSyncStatusOutput Output for triggering a manual sync for a RETL connection.
type GetReverseETLSyncStatusOutput struct {
	ReverseETLSyncStatus ReverseETLSyncStatus `json:"reverseETLSyncStatus"`
}

// NewGetReverseETLSyncStatusOutput instantiates a new GetReverseETLSyncStatusOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReverseETLSyncStatusOutput(
	reverseETLSyncStatus ReverseETLSyncStatus,
) *GetReverseETLSyncStatusOutput {
	this := GetReverseETLSyncStatusOutput{}
	this.ReverseETLSyncStatus = reverseETLSyncStatus
	return &this
}

// NewGetReverseETLSyncStatusOutputWithDefaults instantiates a new GetReverseETLSyncStatusOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReverseETLSyncStatusOutputWithDefaults() *GetReverseETLSyncStatusOutput {
	this := GetReverseETLSyncStatusOutput{}
	return &this
}

// GetReverseETLSyncStatus returns the ReverseETLSyncStatus field value
func (o *GetReverseETLSyncStatusOutput) GetReverseETLSyncStatus() ReverseETLSyncStatus {
	if o == nil {
		var ret ReverseETLSyncStatus
		return ret
	}

	return o.ReverseETLSyncStatus
}

// GetReverseETLSyncStatusOk returns a tuple with the ReverseETLSyncStatus field value
// and a boolean to check if the value has been set.
func (o *GetReverseETLSyncStatusOutput) GetReverseETLSyncStatusOk() (*ReverseETLSyncStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReverseETLSyncStatus, true
}

// SetReverseETLSyncStatus sets field value
func (o *GetReverseETLSyncStatusOutput) SetReverseETLSyncStatus(v ReverseETLSyncStatus) {
	o.ReverseETLSyncStatus = v
}

func (o GetReverseETLSyncStatusOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReverseETLSyncStatusOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reverseETLSyncStatus"] = o.ReverseETLSyncStatus
	return toSerialize, nil
}

type NullableGetReverseETLSyncStatusOutput struct {
	value *GetReverseETLSyncStatusOutput
	isSet bool
}

func (v NullableGetReverseETLSyncStatusOutput) Get() *GetReverseETLSyncStatusOutput {
	return v.value
}

func (v *NullableGetReverseETLSyncStatusOutput) Set(val *GetReverseETLSyncStatusOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReverseETLSyncStatusOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReverseETLSyncStatusOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReverseETLSyncStatusOutput(
	val *GetReverseETLSyncStatusOutput,
) *NullableGetReverseETLSyncStatusOutput {
	return &NullableGetReverseETLSyncStatusOutput{value: val, isSet: true}
}

func (v NullableGetReverseETLSyncStatusOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReverseETLSyncStatusOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
