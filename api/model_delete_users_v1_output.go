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

// checks if the DeleteUsersV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUsersV1Output{}

// DeleteUsersV1Output Returns the status of the removal operation.
type DeleteUsersV1Output struct {
	// A flag that indicates the status of a successful deletion operation.
	Status string `json:"status"`
}

// NewDeleteUsersV1Output instantiates a new DeleteUsersV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUsersV1Output(status string) *DeleteUsersV1Output {
	this := DeleteUsersV1Output{}
	this.Status = status
	return &this
}

// NewDeleteUsersV1OutputWithDefaults instantiates a new DeleteUsersV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUsersV1OutputWithDefaults() *DeleteUsersV1Output {
	this := DeleteUsersV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeleteUsersV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteUsersV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteUsersV1Output) SetStatus(v string) {
	o.Status = v
}

func (o DeleteUsersV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUsersV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableDeleteUsersV1Output struct {
	value *DeleteUsersV1Output
	isSet bool
}

func (v NullableDeleteUsersV1Output) Get() *DeleteUsersV1Output {
	return v.value
}

func (v *NullableDeleteUsersV1Output) Set(val *DeleteUsersV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUsersV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUsersV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUsersV1Output(val *DeleteUsersV1Output) *NullableDeleteUsersV1Output {
	return &NullableDeleteUsersV1Output{value: val, isSet: true}
}

func (v NullableDeleteUsersV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUsersV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
