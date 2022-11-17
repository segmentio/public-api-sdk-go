/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeleteSourceV1Output Returns the status of a Source deletion.
type DeleteSourceV1Output struct {
	// The status of the Source deletion operation.
	Status string `json:"status"`
}

// NewDeleteSourceV1Output instantiates a new DeleteSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSourceV1Output(status string) *DeleteSourceV1Output {
	this := DeleteSourceV1Output{}
	this.Status = status
	return &this
}

// NewDeleteSourceV1OutputWithDefaults instantiates a new DeleteSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSourceV1OutputWithDefaults() *DeleteSourceV1Output {
	this := DeleteSourceV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeleteSourceV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteSourceV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteSourceV1Output) SetStatus(v string) {
	o.Status = v
}

func (o DeleteSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteSourceV1Output struct {
	value *DeleteSourceV1Output
	isSet bool
}

func (v NullableDeleteSourceV1Output) Get() *DeleteSourceV1Output {
	return v.value
}

func (v *NullableDeleteSourceV1Output) Set(val *DeleteSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSourceV1Output(val *DeleteSourceV1Output) *NullableDeleteSourceV1Output {
	return &NullableDeleteSourceV1Output{value: val, isSet: true}
}

func (v NullableDeleteSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


