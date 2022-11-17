/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SourceMetadata The catalog item matched by id.
type SourceMetadata struct {
	// The id for this Source metadata in the Segment catalog.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The user-friendly name of this Source.  Config API note: equal to `displayName`.
	Name string `json:"name"`
	// The slug that identifies this Source in the Segment app.  Config API note: equal to `name`.
	Slug string `json:"slug"`
	// The description of this Source.
	Description string `json:"description"`
	Logos Logos1 `json:"logos"`
	// Options for this Source.
	Options []IntegrationOptionBeta `json:"options"`
	// A list of categories this Source belongs to.
	Categories []string `json:"categories"`
	// True if this is a Cloud Event Source.
	IsCloudEventSource bool `json:"isCloudEventSource"`
}

// NewSourceMetadata instantiates a new SourceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceMetadata(id string, name string, slug string, description string, logos Logos1, options []IntegrationOptionBeta, categories []string, isCloudEventSource bool) *SourceMetadata {
	this := SourceMetadata{}
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

// NewSourceMetadataWithDefaults instantiates a new SourceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceMetadataWithDefaults() *SourceMetadata {
	this := SourceMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *SourceMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SourceMetadata) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SourceMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceMetadata) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *SourceMetadata) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *SourceMetadata) SetSlug(v string) {
	o.Slug = v
}

// GetDescription returns the Description field value
func (o *SourceMetadata) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SourceMetadata) SetDescription(v string) {
	o.Description = v
}

// GetLogos returns the Logos field value
func (o *SourceMetadata) GetLogos() Logos1 {
	if o == nil {
		var ret Logos1
		return ret
	}

	return o.Logos
}

// GetLogosOk returns a tuple with the Logos field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetLogosOk() (*Logos1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logos, true
}

// SetLogos sets field value
func (o *SourceMetadata) SetLogos(v Logos1) {
	o.Logos = v
}

// GetOptions returns the Options field value
func (o *SourceMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		var ret []IntegrationOptionBeta
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetOptionsOk() ([]IntegrationOptionBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *SourceMetadata) SetOptions(v []IntegrationOptionBeta) {
	o.Options = v
}

// GetCategories returns the Categories field value
func (o *SourceMetadata) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *SourceMetadata) SetCategories(v []string) {
	o.Categories = v
}

// GetIsCloudEventSource returns the IsCloudEventSource field value
func (o *SourceMetadata) GetIsCloudEventSource() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCloudEventSource
}

// GetIsCloudEventSourceOk returns a tuple with the IsCloudEventSource field value
// and a boolean to check if the value has been set.
func (o *SourceMetadata) GetIsCloudEventSourceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCloudEventSource, true
}

// SetIsCloudEventSource sets field value
func (o *SourceMetadata) SetIsCloudEventSource(v bool) {
	o.IsCloudEventSource = v
}

func (o SourceMetadata) MarshalJSON() ([]byte, error) {
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

type NullableSourceMetadata struct {
	value *SourceMetadata
	isSet bool
}

func (v NullableSourceMetadata) Get() *SourceMetadata {
	return v.value
}

func (v *NullableSourceMetadata) Set(val *SourceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceMetadata(val *SourceMetadata) *NullableSourceMetadata {
	return &NullableSourceMetadata{value: val, isSet: true}
}

func (v NullableSourceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


