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

// checks if the PropertyRenameV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyRenameV1{}

// PropertyRenameV1 struct for PropertyRenameV1
type PropertyRenameV1 struct {
	// The old name of the property.
	OldName string `json:"oldName"`
	// The new name to rename the property.
	NewName string `json:"newName"`
}

// NewPropertyRenameV1 instantiates a new PropertyRenameV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyRenameV1(oldName string, newName string) *PropertyRenameV1 {
	this := PropertyRenameV1{}
	this.OldName = oldName
	this.NewName = newName
	return &this
}

// NewPropertyRenameV1WithDefaults instantiates a new PropertyRenameV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyRenameV1WithDefaults() *PropertyRenameV1 {
	this := PropertyRenameV1{}
	return &this
}

// GetOldName returns the OldName field value
func (o *PropertyRenameV1) GetOldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldName
}

// GetOldNameOk returns a tuple with the OldName field value
// and a boolean to check if the value has been set.
func (o *PropertyRenameV1) GetOldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldName, true
}

// SetOldName sets field value
func (o *PropertyRenameV1) SetOldName(v string) {
	o.OldName = v
}

// GetNewName returns the NewName field value
func (o *PropertyRenameV1) GetNewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value
// and a boolean to check if the value has been set.
func (o *PropertyRenameV1) GetNewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewName, true
}

// SetNewName sets field value
func (o *PropertyRenameV1) SetNewName(v string) {
	o.NewName = v
}

func (o PropertyRenameV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyRenameV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldName"] = o.OldName
	toSerialize["newName"] = o.NewName
	return toSerialize, nil
}

type NullablePropertyRenameV1 struct {
	value *PropertyRenameV1
	isSet bool
}

func (v NullablePropertyRenameV1) Get() *PropertyRenameV1 {
	return v.value
}

func (v *NullablePropertyRenameV1) Set(val *PropertyRenameV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyRenameV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyRenameV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyRenameV1(val *PropertyRenameV1) *NullablePropertyRenameV1 {
	return &NullablePropertyRenameV1{value: val, isSet: true}
}

func (v NullablePropertyRenameV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyRenameV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
