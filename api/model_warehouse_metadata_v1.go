/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WarehouseMetadataV1 The metadata for an instance of Segment’s data Warehouse product.
type WarehouseMetadataV1 struct {
	// The id of this object.
	Id string `json:"id"`
	// The name of this object.
	Name string `json:"name"`
	// A human-readable, unique identifier for object.
	Slug string `json:"slug"`
	// A description, in English, of this object.
	Description string `json:"description"`
	Logos       Logos2 `json:"logos"`
	// The Integration options for this object.
	Options []IntegrationOptionBeta `json:"options"`
}

// NewWarehouseMetadataV1 instantiates a new WarehouseMetadataV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseMetadataV1(
	id string,
	name string,
	slug string,
	description string,
	logos Logos2,
	options []IntegrationOptionBeta,
) *WarehouseMetadataV1 {
	this := WarehouseMetadataV1{}
	this.Id = id
	this.Name = name
	this.Slug = slug
	this.Description = description
	this.Logos = logos
	this.Options = options
	return &this
}

// NewWarehouseMetadataV1WithDefaults instantiates a new WarehouseMetadataV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseMetadataV1WithDefaults() *WarehouseMetadataV1 {
	this := WarehouseMetadataV1{}
	return &this
}

// GetId returns the Id field value
func (o *WarehouseMetadataV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WarehouseMetadataV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *WarehouseMetadataV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WarehouseMetadataV1) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *WarehouseMetadataV1) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *WarehouseMetadataV1) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value
func (o *WarehouseMetadataV1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WarehouseMetadataV1) SetDescription(v string) {
	o.Description = v
}

// GetLogos returns the Logos field value
func (o *WarehouseMetadataV1) GetLogos() Logos2 {
	if o == nil {
		var ret Logos2
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetLogosOk() (*Logos2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *WarehouseMetadataV1) SetLogos(v Logos2) {
	o.Logos = v
}

// GetOptions returns the Options field value
func (o *WarehouseMetadataV1) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		var ret []IntegrationOptionBeta
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *WarehouseMetadataV1) GetOptionsOk() ([]IntegrationOptionBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *WarehouseMetadataV1) SetOptions(v []IntegrationOptionBeta) {
	o.Options = v
}

func (o WarehouseMetadataV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["logos"] = o.Logos
	}
	if true {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseMetadataV1 struct {
	value *WarehouseMetadataV1
	isSet bool
}

func (v NullableWarehouseMetadataV1) Get() *WarehouseMetadataV1 {
	return v.value
}

func (v *NullableWarehouseMetadataV1) Set(val *WarehouseMetadataV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseMetadataV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseMetadataV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseMetadataV1(val *WarehouseMetadataV1) *NullableWarehouseMetadataV1 {
	return &NullableWarehouseMetadataV1{value: val, isSet: true}
}

func (v NullableWarehouseMetadataV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseMetadataV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
