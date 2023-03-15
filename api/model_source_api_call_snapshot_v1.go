/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SourceAPICallSnapshotV1 A snapshot of the number of API calls for a given Source in a Workspace.
type SourceAPICallSnapshotV1 struct {
	// The Source id.
	SourceId string `json:"sourceId"`
	// A bigint of the number of API calls in this snapshot.
	ApiCalls string `json:"apiCalls"`
	// Timestamp of this snapshot within the billing cycle in the ISO-8601 format.
	Timestamp string `json:"timestamp"`
}

// NewSourceAPICallSnapshotV1 instantiates a new SourceAPICallSnapshotV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAPICallSnapshotV1(
	sourceId string,
	apiCalls string,
	timestamp string,
) *SourceAPICallSnapshotV1 {
	this := SourceAPICallSnapshotV1{}
	this.SourceId = sourceId
	this.ApiCalls = apiCalls
	this.Timestamp = timestamp
	return &this
}

// NewSourceAPICallSnapshotV1WithDefaults instantiates a new SourceAPICallSnapshotV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAPICallSnapshotV1WithDefaults() *SourceAPICallSnapshotV1 {
	this := SourceAPICallSnapshotV1{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *SourceAPICallSnapshotV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *SourceAPICallSnapshotV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *SourceAPICallSnapshotV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetApiCalls returns the ApiCalls field value
func (o *SourceAPICallSnapshotV1) GetApiCalls() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiCalls
}

// GetApiCallsOk returns a tuple with the ApiCalls field value
// and a boolean to check if the value has been set.
func (o *SourceAPICallSnapshotV1) GetApiCallsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiCalls, true
}

// SetApiCalls sets field value
func (o *SourceAPICallSnapshotV1) SetApiCalls(v string) {
	o.ApiCalls = v
}

// GetTimestamp returns the Timestamp field value
func (o *SourceAPICallSnapshotV1) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *SourceAPICallSnapshotV1) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *SourceAPICallSnapshotV1) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o SourceAPICallSnapshotV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["apiCalls"] = o.ApiCalls
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableSourceAPICallSnapshotV1 struct {
	value *SourceAPICallSnapshotV1
	isSet bool
}

func (v NullableSourceAPICallSnapshotV1) Get() *SourceAPICallSnapshotV1 {
	return v.value
}

func (v *NullableSourceAPICallSnapshotV1) Set(val *SourceAPICallSnapshotV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAPICallSnapshotV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAPICallSnapshotV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAPICallSnapshotV1(
	val *SourceAPICallSnapshotV1,
) *NullableSourceAPICallSnapshotV1 {
	return &NullableSourceAPICallSnapshotV1{value: val, isSet: true}
}

func (v NullableSourceAPICallSnapshotV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAPICallSnapshotV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
