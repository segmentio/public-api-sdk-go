/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Space Space matching the given id.
type Space struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// NewSpace instantiates a new Space object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpace(id string, slug string, name string) *Space {
	this := Space{}
	this.Id = id
	this.Slug = slug
	this.Name = name
	return &this
}

// NewSpaceWithDefaults instantiates a new Space object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceWithDefaults() *Space {
	this := Space{}
	return &this
}

// GetId returns the Id field value
func (o *Space) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Space) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Space) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *Space) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Space) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Space) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *Space) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Space) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Space) SetName(v string) {
	o.Name = v
}

func (o Space) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSpace struct {
	value *Space
	isSet bool
}

func (v NullableSpace) Get() *Space {
	return v.value
}

func (v *NullableSpace) Set(val *Space) {
	v.value = val
	v.isSet = true
}

func (v NullableSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpace(val *Space) *NullableSpace {
	return &NullableSpace{value: val, isSet: true}
}

func (v NullableSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
