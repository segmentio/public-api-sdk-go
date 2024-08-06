/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 52.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateAudienceForSpaceInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAudienceForSpaceInput{}

// UpdateAudienceForSpaceInput Input to update an audience.
type UpdateAudienceForSpaceInput struct {
	// Enabled/disabled status for the audience.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the computation.
	Name *string `json:"name,omitempty"`
	// The description of the computation.
	Description *string                        `json:"description,omitempty"`
	Definition  *AudienceComputationDefinition `json:"definition,omitempty"`
}

// NewUpdateAudienceForSpaceInput instantiates a new UpdateAudienceForSpaceInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAudienceForSpaceInput() *UpdateAudienceForSpaceInput {
	this := UpdateAudienceForSpaceInput{}
	return &this
}

// NewUpdateAudienceForSpaceInputWithDefaults instantiates a new UpdateAudienceForSpaceInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAudienceForSpaceInputWithDefaults() *UpdateAudienceForSpaceInput {
	this := UpdateAudienceForSpaceInput{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateAudienceForSpaceInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpaceInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateAudienceForSpaceInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateAudienceForSpaceInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAudienceForSpaceInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpaceInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAudienceForSpaceInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAudienceForSpaceInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateAudienceForSpaceInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpaceInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateAudienceForSpaceInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateAudienceForSpaceInput) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *UpdateAudienceForSpaceInput) GetDefinition() AudienceComputationDefinition {
	if o == nil || IsNil(o.Definition) {
		var ret AudienceComputationDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAudienceForSpaceInput) GetDefinitionOk() (*AudienceComputationDefinition, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *UpdateAudienceForSpaceInput) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given AudienceComputationDefinition and assigns it to the Definition field.
func (o *UpdateAudienceForSpaceInput) SetDefinition(v AudienceComputationDefinition) {
	o.Definition = &v
}

func (o UpdateAudienceForSpaceInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAudienceForSpaceInput) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateAudienceForSpaceInput struct {
	value *UpdateAudienceForSpaceInput
	isSet bool
}

func (v NullableUpdateAudienceForSpaceInput) Get() *UpdateAudienceForSpaceInput {
	return v.value
}

func (v *NullableUpdateAudienceForSpaceInput) Set(val *UpdateAudienceForSpaceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAudienceForSpaceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAudienceForSpaceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAudienceForSpaceInput(
	val *UpdateAudienceForSpaceInput,
) *NullableUpdateAudienceForSpaceInput {
	return &NullableUpdateAudienceForSpaceInput{value: val, isSet: true}
}

func (v NullableUpdateAudienceForSpaceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAudienceForSpaceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
