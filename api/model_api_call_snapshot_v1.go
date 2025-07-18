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

// checks if the APICallSnapshotV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APICallSnapshotV1{}

// APICallSnapshotV1 A snapshot of the number of API calls for a given Workspace.
type APICallSnapshotV1 struct {
	// A bigint of the number of API calls in this snapshot.
	ApiCalls string `json:"apiCalls"`
	// Timestamp of this snapshot within the billing cycle in the ISO-8601 format.
	Timestamp string `json:"timestamp"`
}

// NewAPICallSnapshotV1 instantiates a new APICallSnapshotV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPICallSnapshotV1(apiCalls string, timestamp string) *APICallSnapshotV1 {
	this := APICallSnapshotV1{}
	this.ApiCalls = apiCalls
	this.Timestamp = timestamp
	return &this
}

// NewAPICallSnapshotV1WithDefaults instantiates a new APICallSnapshotV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPICallSnapshotV1WithDefaults() *APICallSnapshotV1 {
	this := APICallSnapshotV1{}
	return &this
}

// GetApiCalls returns the ApiCalls field value
func (o *APICallSnapshotV1) GetApiCalls() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiCalls
}

// GetApiCallsOk returns a tuple with the ApiCalls field value
// and a boolean to check if the value has been set.
func (o *APICallSnapshotV1) GetApiCallsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiCalls, true
}

// SetApiCalls sets field value
func (o *APICallSnapshotV1) SetApiCalls(v string) {
	o.ApiCalls = v
}

// GetTimestamp returns the Timestamp field value
func (o *APICallSnapshotV1) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *APICallSnapshotV1) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *APICallSnapshotV1) SetTimestamp(v string) {
	o.Timestamp = v
}

func (o APICallSnapshotV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APICallSnapshotV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiCalls"] = o.ApiCalls
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableAPICallSnapshotV1 struct {
	value *APICallSnapshotV1
	isSet bool
}

func (v NullableAPICallSnapshotV1) Get() *APICallSnapshotV1 {
	return v.value
}

func (v *NullableAPICallSnapshotV1) Set(val *APICallSnapshotV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAPICallSnapshotV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAPICallSnapshotV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPICallSnapshotV1(val *APICallSnapshotV1) *NullableAPICallSnapshotV1 {
	return &NullableAPICallSnapshotV1{value: val, isSet: true}
}

func (v NullableAPICallSnapshotV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPICallSnapshotV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
