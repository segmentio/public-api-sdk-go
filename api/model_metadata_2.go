/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Metadata2 The metadata for the Source.  Config API note: includes `catalogName` and `catalogId`.
type Metadata2 struct {
	// The id for this Source metadata in the Segment catalog.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The user-friendly name of this Source.  Config API note: equal to `displayName`.
	Name string `json:"name"`
	// The slug that identifies this Source in the Segment app.  Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The description of this Source.
	Description string `json:"description"`
	Logos       Logos1 `json:"logos"`
	// Options for this Source.
	Options []IntegrationOptionBeta `json:"options"`
	// A list of categories this Source belongs to.
	Categories []string `json:"categories"`
	// True if this is a Cloud Event Source.
	IsCloudEventSource bool `json:"isCloudEventSource"`
}

// NewMetadata2 instantiates a new Metadata2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata2(
	id string,
	name string,
	slug string,
	description string,
	logos Logos1,
	options []IntegrationOptionBeta,
	categories []string,
	isCloudEventSource bool,
) *Metadata2 {
	this := Metadata2{}
	this.Id = id
	this.Name = name
	this.Slug = slug
	this.Description = description
	this.Logos = logos
	this.Options = options
	this.Categories = categories
	this.IsCloudEventSource = isCloudEventSource
	return &this
}

// NewMetadata2WithDefaults instantiates a new Metadata2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadata2WithDefaults() *Metadata2 {
	this := Metadata2{}
	return &this
}

// GetId returns the Id field value
func (o *Metadata2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Metadata2) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Metadata2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Metadata2) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *Metadata2) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Metadata2) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value
func (o *Metadata2) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Metadata2) SetDescription(v string) {
	o.Description = v
}

// GetLogos returns the Logos field value
func (o *Metadata2) GetLogos() Logos1 {
	if o == nil {
		var ret Logos1
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetLogosOk() (*Logos1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *Metadata2) SetLogos(v Logos1) {
	o.Logos = v
}

// GetOptions returns the Options field value
func (o *Metadata2) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		var ret []IntegrationOptionBeta
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetOptionsOk() ([]IntegrationOptionBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *Metadata2) SetOptions(v []IntegrationOptionBeta) {
	o.Options = v
}

// GetCategories returns the Categories field value
func (o *Metadata2) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *Metadata2) SetCategories(v []string) {
	o.Categories = v
}

// GetIsCloudEventSource returns the IsCloudEventSource field value
func (o *Metadata2) GetIsCloudEventSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCloudEventSource
}

// GetIsCloudEventSourceOk returns a tuple with the IsCloudEventSource field value
// and a boolean to check if the value has been set.
func (o *Metadata2) GetIsCloudEventSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCloudEventSource, true
}

// SetIsCloudEventSource sets field value
func (o *Metadata2) SetIsCloudEventSource(v bool) {
	o.IsCloudEventSource = v
}

func (o Metadata2) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["isCloudEventSource"] = o.IsCloudEventSource
	}
	return json.Marshal(toSerialize)
}

type NullableMetadata2 struct {
	value *Metadata2
	isSet bool
}

func (v NullableMetadata2) Get() *Metadata2 {
	return v.value
}

func (v *NullableMetadata2) Set(val *Metadata2) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata2) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata2(val *Metadata2) *NullableMetadata2 {
	return &NullableMetadata2{value: val, isSet: true}
}

func (v NullableMetadata2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
