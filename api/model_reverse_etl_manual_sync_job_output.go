/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReverseETLManualSyncJobOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseETLManualSyncJobOutput{}

// ReverseETLManualSyncJobOutput StartJobOutput returns the result of triggering the specified job.
type ReverseETLManualSyncJobOutput struct {
	// A datetime that indicates when the job was triggered.
	StartedAt string `json:"startedAt"`
	// The unique id for the sync that was triggered.
	SyncId string `json:"syncId"`
}

// NewReverseETLManualSyncJobOutput instantiates a new ReverseETLManualSyncJobOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseETLManualSyncJobOutput(
	startedAt string,
	syncId string,
) *ReverseETLManualSyncJobOutput {
	this := ReverseETLManualSyncJobOutput{}
	this.StartedAt = startedAt
	this.SyncId = syncId
	return &this
}

// NewReverseETLManualSyncJobOutputWithDefaults instantiates a new ReverseETLManualSyncJobOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseETLManualSyncJobOutputWithDefaults() *ReverseETLManualSyncJobOutput {
	this := ReverseETLManualSyncJobOutput{}
	return &this
}

// GetStartedAt returns the StartedAt field value
func (o *ReverseETLManualSyncJobOutput) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ReverseETLManualSyncJobOutput) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *ReverseETLManualSyncJobOutput) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetSyncId returns the SyncId field value
func (o *ReverseETLManualSyncJobOutput) GetSyncId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncId
}

// GetSyncIdOk returns a tuple with the SyncId field value
// and a boolean to check if the value has been set.
func (o *ReverseETLManualSyncJobOutput) GetSyncIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncId, true
}

// SetSyncId sets field value
func (o *ReverseETLManualSyncJobOutput) SetSyncId(v string) {
	o.SyncId = v
}

func (o ReverseETLManualSyncJobOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseETLManualSyncJobOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["syncId"] = o.SyncId
	return toSerialize, nil
}

type NullableReverseETLManualSyncJobOutput struct {
	value *ReverseETLManualSyncJobOutput
	isSet bool
}

func (v NullableReverseETLManualSyncJobOutput) Get() *ReverseETLManualSyncJobOutput {
	return v.value
}

func (v *NullableReverseETLManualSyncJobOutput) Set(val *ReverseETLManualSyncJobOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseETLManualSyncJobOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseETLManualSyncJobOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseETLManualSyncJobOutput(
	val *ReverseETLManualSyncJobOutput,
) *NullableReverseETLManualSyncJobOutput {
	return &NullableReverseETLManualSyncJobOutput{value: val, isSet: true}
}

func (v NullableReverseETLManualSyncJobOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseETLManualSyncJobOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
