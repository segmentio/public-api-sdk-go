/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SourceEventVolumeDatapointV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceEventVolumeDatapointV1{}

// SourceEventVolumeDatapointV1 SourceEventVolumeDatapoint represents an individual timestamp/event count pair corresponding to a window of time designated by the granularity.
type SourceEventVolumeDatapointV1 struct {
	// The timestamp that corresponds to the beginning of the window given by the requested granularity.
	Time string `json:"time"`
	// The number of valid events Segment received in the given window, after the events completed the validation process.
	Count float32 `json:"count"`
}

// NewSourceEventVolumeDatapointV1 instantiates a new SourceEventVolumeDatapointV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceEventVolumeDatapointV1(time string, count float32) *SourceEventVolumeDatapointV1 {
	this := SourceEventVolumeDatapointV1{}
	this.Time = time
	this.Count = count
	return &this
}

// NewSourceEventVolumeDatapointV1WithDefaults instantiates a new SourceEventVolumeDatapointV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceEventVolumeDatapointV1WithDefaults() *SourceEventVolumeDatapointV1 {
	this := SourceEventVolumeDatapointV1{}
	return &this
}

// GetTime returns the Time field value
func (o *SourceEventVolumeDatapointV1) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeDatapointV1) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *SourceEventVolumeDatapointV1) SetTime(v string) {
	o.Time = v
}

// GetCount returns the Count field value
func (o *SourceEventVolumeDatapointV1) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeDatapointV1) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SourceEventVolumeDatapointV1) SetCount(v float32) {
	o.Count = v
}

func (o SourceEventVolumeDatapointV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceEventVolumeDatapointV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

type NullableSourceEventVolumeDatapointV1 struct {
	value *SourceEventVolumeDatapointV1
	isSet bool
}

func (v NullableSourceEventVolumeDatapointV1) Get() *SourceEventVolumeDatapointV1 {
	return v.value
}

func (v *NullableSourceEventVolumeDatapointV1) Set(val *SourceEventVolumeDatapointV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceEventVolumeDatapointV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceEventVolumeDatapointV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceEventVolumeDatapointV1(
	val *SourceEventVolumeDatapointV1,
) *NullableSourceEventVolumeDatapointV1 {
	return &NullableSourceEventVolumeDatapointV1{value: val, isSet: true}
}

func (v NullableSourceEventVolumeDatapointV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceEventVolumeDatapointV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
