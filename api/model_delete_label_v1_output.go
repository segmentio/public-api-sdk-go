/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DeleteLabelV1Output Returns the status of a label deletion.
type DeleteLabelV1Output struct {
	// The status of the label deletion operation.
	Status string `json:"status"`
}

// NewDeleteLabelV1Output instantiates a new DeleteLabelV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteLabelV1Output(status string) *DeleteLabelV1Output {
	this := DeleteLabelV1Output{}
	this.Status = status
	return &this
}

// NewDeleteLabelV1OutputWithDefaults instantiates a new DeleteLabelV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteLabelV1OutputWithDefaults() *DeleteLabelV1Output {
	this := DeleteLabelV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeleteLabelV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteLabelV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteLabelV1Output) SetStatus(v string) {
	o.Status = v
}

func (o DeleteLabelV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteLabelV1Output struct {
	value *DeleteLabelV1Output
	isSet bool
}

func (v NullableDeleteLabelV1Output) Get() *DeleteLabelV1Output {
	return v.value
}

func (v *NullableDeleteLabelV1Output) Set(val *DeleteLabelV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteLabelV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteLabelV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteLabelV1Output(val *DeleteLabelV1Output) *NullableDeleteLabelV1Output {
	return &NullableDeleteLabelV1Output{value: val, isSet: true}
}

func (v NullableDeleteLabelV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteLabelV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
