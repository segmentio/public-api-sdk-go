/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// StreamStatusV1 StreamStatus represents status of each stream including all the Destinations corresponding to the stream.
type StreamStatusV1 struct {
	Id                string                `json:"id"`
	DestinationStatus []DestinationStatusV1 `json:"destinationStatus"`
}

// NewStreamStatusV1 instantiates a new StreamStatusV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamStatusV1(id string, destinationStatus []DestinationStatusV1) *StreamStatusV1 {
	this := StreamStatusV1{}
	this.Id = id
	this.DestinationStatus = destinationStatus
	return &this
}

// NewStreamStatusV1WithDefaults instantiates a new StreamStatusV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamStatusV1WithDefaults() *StreamStatusV1 {
	this := StreamStatusV1{}
	return &this
}

// GetId returns the Id field value
func (o *StreamStatusV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamStatusV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamStatusV1) SetId(v string) {
	o.Id = v
}

// GetDestinationStatus returns the DestinationStatus field value
func (o *StreamStatusV1) GetDestinationStatus() []DestinationStatusV1 {
	if o == nil {
		var ret []DestinationStatusV1
		return ret
	}

	return o.DestinationStatus
}

// GetDestinationStatusOk returns a tuple with the DestinationStatus field value
// and a boolean to check if the value has been set.
func (o *StreamStatusV1) GetDestinationStatusOk() ([]DestinationStatusV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationStatus, true
}

// SetDestinationStatus sets field value
func (o *StreamStatusV1) SetDestinationStatus(v []DestinationStatusV1) {
	o.DestinationStatus = v
}

func (o StreamStatusV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["destinationStatus"] = o.DestinationStatus
	}
	return json.Marshal(toSerialize)
}

type NullableStreamStatusV1 struct {
	value *StreamStatusV1
	isSet bool
}

func (v NullableStreamStatusV1) Get() *StreamStatusV1 {
	return v.value
}

func (v *NullableStreamStatusV1) Set(val *StreamStatusV1) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamStatusV1) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamStatusV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamStatusV1(val *StreamStatusV1) *NullableStreamStatusV1 {
	return &NullableStreamStatusV1{value: val, isSet: true}
}

func (v NullableStreamStatusV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamStatusV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
