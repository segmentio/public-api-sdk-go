/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.6.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the Preview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Preview{}

// Preview struct for Preview
type Preview struct {
	// Unique identifier for tracking and retrieving results of the preview.
	Id string `json:"id"`
}

// NewPreview instantiates a new Preview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreview(id string) *Preview {
	this := Preview{}
	this.Id = id
	return &this
}

// NewPreviewWithDefaults instantiates a new Preview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewWithDefaults() *Preview {
	this := Preview{}
	return &this
}

// GetId returns the Id field value
func (o *Preview) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Preview) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Preview) SetId(v string) {
	o.Id = v
}

func (o Preview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Preview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullablePreview struct {
	value *Preview
	isSet bool
}

func (v NullablePreview) Get() *Preview {
	return v.value
}

func (v *NullablePreview) Set(val *Preview) {
	v.value = val
	v.isSet = true
}

func (v NullablePreview) IsSet() bool {
	return v.isSet
}

func (v *NullablePreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreview(val *Preview) *NullablePreview {
	return &NullablePreview{value: val, isSet: true}
}

func (v NullablePreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
