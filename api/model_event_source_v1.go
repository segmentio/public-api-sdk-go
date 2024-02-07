/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 40.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the EventSourceV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSourceV1{}

// EventSourceV1 Source represents a Segment Source.
type EventSourceV1 struct {
	// The id of the Source where the events came from.
	Id string `json:"id"`
	// The name of the Source, if applicable.
	Name *string `json:"name,omitempty"`
	// The slug of the Source, if applicable.
	Slug *string `json:"slug,omitempty"`
}

// NewEventSourceV1 instantiates a new EventSourceV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSourceV1(id string) *EventSourceV1 {
	this := EventSourceV1{}
	this.Id = id
	return &this
}

// NewEventSourceV1WithDefaults instantiates a new EventSourceV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSourceV1WithDefaults() *EventSourceV1 {
	this := EventSourceV1{}
	return &this
}

// GetId returns the Id field value
func (o *EventSourceV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventSourceV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventSourceV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EventSourceV1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSourceV1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EventSourceV1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EventSourceV1) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *EventSourceV1) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSourceV1) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *EventSourceV1) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *EventSourceV1) SetSlug(v string) {
	o.Slug = &v
}

func (o EventSourceV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSourceV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	return toSerialize, nil
}

type NullableEventSourceV1 struct {
	value *EventSourceV1
	isSet bool
}

func (v NullableEventSourceV1) Get() *EventSourceV1 {
	return v.value
}

func (v *NullableEventSourceV1) Set(val *EventSourceV1) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSourceV1) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSourceV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSourceV1(val *EventSourceV1) *NullableEventSourceV1 {
	return &NullableEventSourceV1{value: val, isSet: true}
}

func (v NullableEventSourceV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSourceV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
