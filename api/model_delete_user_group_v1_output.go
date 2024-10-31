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

// checks if the DeleteUserGroupV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteUserGroupV1Output{}

// DeleteUserGroupV1Output Returns the status of the completed deletion operation.
type DeleteUserGroupV1Output struct {
	// A flag indicating the status of a successful deletion operation.
	Status string `json:"status"`
}

// NewDeleteUserGroupV1Output instantiates a new DeleteUserGroupV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteUserGroupV1Output(status string) *DeleteUserGroupV1Output {
	this := DeleteUserGroupV1Output{}
	this.Status = status
	return &this
}

// NewDeleteUserGroupV1OutputWithDefaults instantiates a new DeleteUserGroupV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteUserGroupV1OutputWithDefaults() *DeleteUserGroupV1Output {
	this := DeleteUserGroupV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *DeleteUserGroupV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DeleteUserGroupV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DeleteUserGroupV1Output) SetStatus(v string) {
	o.Status = v
}

func (o DeleteUserGroupV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteUserGroupV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableDeleteUserGroupV1Output struct {
	value *DeleteUserGroupV1Output
	isSet bool
}

func (v NullableDeleteUserGroupV1Output) Get() *DeleteUserGroupV1Output {
	return v.value
}

func (v *NullableDeleteUserGroupV1Output) Set(val *DeleteUserGroupV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteUserGroupV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteUserGroupV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteUserGroupV1Output(
	val *DeleteUserGroupV1Output,
) *NullableDeleteUserGroupV1Output {
	return &NullableDeleteUserGroupV1Output{value: val, isSet: true}
}

func (v NullableDeleteUserGroupV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteUserGroupV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
