/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ResourceV1 Represents a permission's resource.
type ResourceV1 struct {
	// The id of this resource.
	Id string `json:"id"`
	// The kind of resource this permission applies to.
	Type string `json:"type"`
}

// NewResourceV1 instantiates a new ResourceV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceV1(id string, type_ string) *ResourceV1 {
	this := ResourceV1{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewResourceV1WithDefaults instantiates a new ResourceV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceV1WithDefaults() *ResourceV1 {
	this := ResourceV1{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceV1) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ResourceV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ResourceV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ResourceV1) SetType(v string) {
	o.Type = v
}

func (o ResourceV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableResourceV1 struct {
	value *ResourceV1
	isSet bool
}

func (v NullableResourceV1) Get() *ResourceV1 {
	return v.value
}

func (v *NullableResourceV1) Set(val *ResourceV1) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceV1) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceV1(val *ResourceV1) *NullableResourceV1 {
	return &NullableResourceV1{value: val, isSet: true}
}

func (v NullableResourceV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


