/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FQLDefinedPropertyV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FQLDefinedPropertyV1{}

// FQLDefinedPropertyV1 struct for FQLDefinedPropertyV1
type FQLDefinedPropertyV1 struct {
	// The FQL expression used to compute the property.
	Fql string `json:"fql"`
	// The new property name.
	PropertyName string `json:"propertyName"`
}

// NewFQLDefinedPropertyV1 instantiates a new FQLDefinedPropertyV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFQLDefinedPropertyV1(fql string, propertyName string) *FQLDefinedPropertyV1 {
	this := FQLDefinedPropertyV1{}
	this.Fql = fql
	this.PropertyName = propertyName
	return &this
}

// NewFQLDefinedPropertyV1WithDefaults instantiates a new FQLDefinedPropertyV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFQLDefinedPropertyV1WithDefaults() *FQLDefinedPropertyV1 {
	this := FQLDefinedPropertyV1{}
	return &this
}

// GetFql returns the Fql field value
func (o *FQLDefinedPropertyV1) GetFql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fql
}

// GetFqlOk returns a tuple with the Fql field value
// and a boolean to check if the value has been set.
func (o *FQLDefinedPropertyV1) GetFqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fql, true
}

// SetFql sets field value
func (o *FQLDefinedPropertyV1) SetFql(v string) {
	o.Fql = v
}

// GetPropertyName returns the PropertyName field value
func (o *FQLDefinedPropertyV1) GetPropertyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value
// and a boolean to check if the value has been set.
func (o *FQLDefinedPropertyV1) GetPropertyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyName, true
}

// SetPropertyName sets field value
func (o *FQLDefinedPropertyV1) SetPropertyName(v string) {
	o.PropertyName = v
}

func (o FQLDefinedPropertyV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FQLDefinedPropertyV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fql"] = o.Fql
	toSerialize["propertyName"] = o.PropertyName
	return toSerialize, nil
}

type NullableFQLDefinedPropertyV1 struct {
	value *FQLDefinedPropertyV1
	isSet bool
}

func (v NullableFQLDefinedPropertyV1) Get() *FQLDefinedPropertyV1 {
	return v.value
}

func (v *NullableFQLDefinedPropertyV1) Set(val *FQLDefinedPropertyV1) {
	v.value = val
	v.isSet = true
}

func (v NullableFQLDefinedPropertyV1) IsSet() bool {
	return v.isSet
}

func (v *NullableFQLDefinedPropertyV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFQLDefinedPropertyV1(val *FQLDefinedPropertyV1) *NullableFQLDefinedPropertyV1 {
	return &NullableFQLDefinedPropertyV1{value: val, isSet: true}
}

func (v NullableFQLDefinedPropertyV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFQLDefinedPropertyV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
