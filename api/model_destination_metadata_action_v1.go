/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationMetadataActionV1 Represents an Action, a building block of behavior that can be performed by the Destination.
type DestinationMetadataActionV1 struct {
	// The primary key of the action.
	Id string `json:"id"`
	// A machine-readable key unique to the action definition.
	Slug string `json:"slug"`
	// A human-readable name for the action.
	Name string `json:"name"`
	// A human-readable description of the action. May include Markdown.
	Description string `json:"description"`
	// The platform on which this action runs.
	Platform string `json:"platform"`
	// Whether the action should be hidden.
	Hidden bool `json:"hidden"`
	// The default value used as the trigger when connecting this action.
	DefaultTrigger NullableString `json:"defaultTrigger"`
	// The fields expected in order to perform the action.
	Fields []DestinationMetadataActionFieldV1 `json:"fields"`
}

// NewDestinationMetadataActionV1 instantiates a new DestinationMetadataActionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataActionV1(
	id string,
	slug string,
	name string,
	description string,
	platform string,
	hidden bool,
	defaultTrigger NullableString,
	fields []DestinationMetadataActionFieldV1,
) *DestinationMetadataActionV1 {
	this := DestinationMetadataActionV1{}
	this.Id = id
	this.Slug = slug
	this.Name = name
	this.Description = description
	this.Platform = platform
	this.Hidden = hidden
	this.DefaultTrigger = defaultTrigger
	this.Fields = fields
	return &this
}

// NewDestinationMetadataActionV1WithDefaults instantiates a new DestinationMetadataActionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataActionV1WithDefaults() *DestinationMetadataActionV1 {
	this := DestinationMetadataActionV1{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationMetadataActionV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationMetadataActionV1) SetId(v string) {
	o.Id = v
}

// GetSlug returns the Slug field value
func (o *DestinationMetadataActionV1) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *DestinationMetadataActionV1) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *DestinationMetadataActionV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationMetadataActionV1) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *DestinationMetadataActionV1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DestinationMetadataActionV1) SetDescription(v string) {
	o.Description = v
}

// GetPlatform returns the Platform field value
func (o *DestinationMetadataActionV1) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value
func (o *DestinationMetadataActionV1) SetPlatform(v string) {
	o.Platform = v
}

// GetHidden returns the Hidden field value
func (o *DestinationMetadataActionV1) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *DestinationMetadataActionV1) SetHidden(v bool) {
	o.Hidden = v
}

// GetDefaultTrigger returns the DefaultTrigger field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DestinationMetadataActionV1) GetDefaultTrigger() string {
	if o == nil || o.DefaultTrigger.Get() == nil {
		var ret string
		return ret
	}

	return *o.DefaultTrigger.Get()
}

// GetDefaultTriggerOk returns a tuple with the DefaultTrigger field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestinationMetadataActionV1) GetDefaultTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultTrigger.Get(), o.DefaultTrigger.IsSet()
}

// SetDefaultTrigger sets field value
func (o *DestinationMetadataActionV1) SetDefaultTrigger(v string) {
	o.DefaultTrigger.Set(&v)
}

// GetFields returns the Fields field value
func (o *DestinationMetadataActionV1) GetFields() []DestinationMetadataActionFieldV1 {
	if o == nil {
		var ret []DestinationMetadataActionFieldV1
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionV1) GetFieldsOk() ([]DestinationMetadataActionFieldV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *DestinationMetadataActionV1) SetFields(v []DestinationMetadataActionFieldV1) {
	o.Fields = v
}

func (o DestinationMetadataActionV1) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["platform"] = o.Platform
	}
	if true {
		toSerialize["hidden"] = o.Hidden
	}
	if true {
		toSerialize["defaultTrigger"] = o.DefaultTrigger.Get()
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationMetadataActionV1 struct {
	value *DestinationMetadataActionV1
	isSet bool
}

func (v NullableDestinationMetadataActionV1) Get() *DestinationMetadataActionV1 {
	return v.value
}

func (v *NullableDestinationMetadataActionV1) Set(val *DestinationMetadataActionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataActionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataActionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataActionV1(
	val *DestinationMetadataActionV1,
) *NullableDestinationMetadataActionV1 {
	return &NullableDestinationMetadataActionV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataActionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataActionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
