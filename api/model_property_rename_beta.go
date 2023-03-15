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

// PropertyRenameBeta struct for PropertyRenameBeta
type PropertyRenameBeta struct {
	// The old name of the property.
	OldName string `json:"oldName"`
	// The new name to rename the property.
	NewName string `json:"newName"`
}

// NewPropertyRenameBeta instantiates a new PropertyRenameBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyRenameBeta(oldName string, newName string) *PropertyRenameBeta {
	this := PropertyRenameBeta{}
	this.OldName = oldName
	this.NewName = newName
	return &this
}

// NewPropertyRenameBetaWithDefaults instantiates a new PropertyRenameBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyRenameBetaWithDefaults() *PropertyRenameBeta {
	this := PropertyRenameBeta{}
	return &this
}

// GetOldName returns the OldName field value
func (o *PropertyRenameBeta) GetOldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldName
}

// GetOldNameOk returns a tuple with the OldName field value
// and a boolean to check if the value has been set.
func (o *PropertyRenameBeta) GetOldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldName, true
}

// SetOldName sets field value
func (o *PropertyRenameBeta) SetOldName(v string) {
	o.OldName = v
}

// GetNewName returns the NewName field value
func (o *PropertyRenameBeta) GetNewName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value
// and a boolean to check if the value has been set.
func (o *PropertyRenameBeta) GetNewNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewName, true
}

// SetNewName sets field value
func (o *PropertyRenameBeta) SetNewName(v string) {
	o.NewName = v
}

func (o PropertyRenameBeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["oldName"] = o.OldName
	}
	if true {
		toSerialize["newName"] = o.NewName
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyRenameBeta struct {
	value *PropertyRenameBeta
	isSet bool
}

func (v NullablePropertyRenameBeta) Get() *PropertyRenameBeta {
	return v.value
}

func (v *NullablePropertyRenameBeta) Set(val *PropertyRenameBeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyRenameBeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyRenameBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyRenameBeta(val *PropertyRenameBeta) *NullablePropertyRenameBeta {
	return &NullablePropertyRenameBeta{value: val, isSet: true}
}

func (v NullablePropertyRenameBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyRenameBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
