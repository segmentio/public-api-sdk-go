/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateComputedTraitForSpaceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateComputedTraitForSpaceAlphaInput{}

// UpdateComputedTraitForSpaceAlphaInput Input to update a computed trait.
type UpdateComputedTraitForSpaceAlphaInput struct {
	// Enabled/disabled status for the computed trait.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the computation.
	Name *string `json:"name,omitempty"`
	// The description of the computation.
	Description *string          `json:"description,omitempty"`
	Definition  *TraitDefinition `json:"definition,omitempty"`
}

// NewUpdateComputedTraitForSpaceAlphaInput instantiates a new UpdateComputedTraitForSpaceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateComputedTraitForSpaceAlphaInput() *UpdateComputedTraitForSpaceAlphaInput {
	this := UpdateComputedTraitForSpaceAlphaInput{}
	return &this
}

// NewUpdateComputedTraitForSpaceAlphaInputWithDefaults instantiates a new UpdateComputedTraitForSpaceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateComputedTraitForSpaceAlphaInputWithDefaults() *UpdateComputedTraitForSpaceAlphaInput {
	this := UpdateComputedTraitForSpaceAlphaInput{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateComputedTraitForSpaceAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateComputedTraitForSpaceAlphaInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateComputedTraitForSpaceAlphaInput) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetDefinition() TraitDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret TraitDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) GetDefinitionOk() (*TraitDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *UpdateComputedTraitForSpaceAlphaInput) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given TraitDefinition and assigns it to the Definition field.
func (o *UpdateComputedTraitForSpaceAlphaInput) SetDefinition(v TraitDefinition) {
	o.Definition = &v
}

func (o UpdateComputedTraitForSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateComputedTraitForSpaceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	return toSerialize, nil
}

type NullableUpdateComputedTraitForSpaceAlphaInput struct {
	value *UpdateComputedTraitForSpaceAlphaInput
	isSet bool
}

func (v NullableUpdateComputedTraitForSpaceAlphaInput) Get() *UpdateComputedTraitForSpaceAlphaInput {
	return v.value
}

func (v *NullableUpdateComputedTraitForSpaceAlphaInput) Set(
	val *UpdateComputedTraitForSpaceAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateComputedTraitForSpaceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateComputedTraitForSpaceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateComputedTraitForSpaceAlphaInput(
	val *UpdateComputedTraitForSpaceAlphaInput,
) *NullableUpdateComputedTraitForSpaceAlphaInput {
	return &NullableUpdateComputedTraitForSpaceAlphaInput{value: val, isSet: true}
}

func (v NullableUpdateComputedTraitForSpaceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateComputedTraitForSpaceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
