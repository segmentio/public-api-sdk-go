/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateComputedTraitAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateComputedTraitAlphaInput{}

// CreateComputedTraitAlphaInput Input to create a trait.
type CreateComputedTraitAlphaInput struct {
	// The name of the computation.
	Name string `json:"name"`
	// Determines whether a computation is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The description of the computation.
	Description *string         `json:"description,omitempty"`
	Definition  TraitDefinition `json:"definition"`
	Options     *TraitOptions   `json:"options,omitempty"`
}

// NewCreateComputedTraitAlphaInput instantiates a new CreateComputedTraitAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateComputedTraitAlphaInput(
	name string,
	definition TraitDefinition,
) *CreateComputedTraitAlphaInput {
	this := CreateComputedTraitAlphaInput{}
	this.Name = name
	this.Definition = definition
	return &this
}

// NewCreateComputedTraitAlphaInputWithDefaults instantiates a new CreateComputedTraitAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateComputedTraitAlphaInputWithDefaults() *CreateComputedTraitAlphaInput {
	this := CreateComputedTraitAlphaInput{}
	return &this
}

// GetName returns the Name field value
func (o *CreateComputedTraitAlphaInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateComputedTraitAlphaInput) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateComputedTraitAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateComputedTraitAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateComputedTraitAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateComputedTraitAlphaInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateComputedTraitAlphaInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateComputedTraitAlphaInput) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value
func (o *CreateComputedTraitAlphaInput) GetDefinition() TraitDefinition {
	if o == nil {
		var ret TraitDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaInput) GetDefinitionOk() (*TraitDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *CreateComputedTraitAlphaInput) SetDefinition(v TraitDefinition) {
	o.Definition = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateComputedTraitAlphaInput) GetOptions() TraitOptions {
	if o == nil || IsNil(o.Options) {
		var ret TraitOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateComputedTraitAlphaInput) GetOptionsOk() (*TraitOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateComputedTraitAlphaInput) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given TraitOptions and assigns it to the Options field.
func (o *CreateComputedTraitAlphaInput) SetOptions(v TraitOptions) {
	o.Options = &v
}

func (o CreateComputedTraitAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateComputedTraitAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["definition"] = o.Definition
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableCreateComputedTraitAlphaInput struct {
	value *CreateComputedTraitAlphaInput
	isSet bool
}

func (v NullableCreateComputedTraitAlphaInput) Get() *CreateComputedTraitAlphaInput {
	return v.value
}

func (v *NullableCreateComputedTraitAlphaInput) Set(val *CreateComputedTraitAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateComputedTraitAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateComputedTraitAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateComputedTraitAlphaInput(
	val *CreateComputedTraitAlphaInput,
) *NullableCreateComputedTraitAlphaInput {
	return &NullableCreateComputedTraitAlphaInput{value: val, isSet: true}
}

func (v NullableCreateComputedTraitAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateComputedTraitAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
