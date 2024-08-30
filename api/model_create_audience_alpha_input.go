/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateAudienceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAudienceAlphaInput{}

// CreateAudienceAlphaInput Input to create an audience.
type CreateAudienceAlphaInput struct {
	// Name of the audience.
	Name string `json:"name"`
	// Determines whether a computation is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Description of the audience.
	Description *string                       `json:"description,omitempty"`
	Definition  AudienceComputationDefinition `json:"definition"`
	Options     *AudienceOptions              `json:"options,omitempty"`
}

// NewCreateAudienceAlphaInput instantiates a new CreateAudienceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAudienceAlphaInput(
	name string,
	definition AudienceComputationDefinition,
) *CreateAudienceAlphaInput {
	this := CreateAudienceAlphaInput{}
	this.Name = name
	this.Definition = definition
	return &this
}

// NewCreateAudienceAlphaInputWithDefaults instantiates a new CreateAudienceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAudienceAlphaInputWithDefaults() *CreateAudienceAlphaInput {
	this := CreateAudienceAlphaInput{}
	return &this
}

// GetName returns the Name field value
func (o *CreateAudienceAlphaInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateAudienceAlphaInput) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateAudienceAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateAudienceAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateAudienceAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAudienceAlphaInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAudienceAlphaInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAudienceAlphaInput) SetDescription(v string) {
	o.Description = &v
}

// GetDefinition returns the Definition field value
func (o *CreateAudienceAlphaInput) GetDefinition() AudienceComputationDefinition {
	if o == nil {
		var ret AudienceComputationDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetDefinitionOk() (*AudienceComputationDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *CreateAudienceAlphaInput) SetDefinition(v AudienceComputationDefinition) {
	o.Definition = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateAudienceAlphaInput) GetOptions() AudienceOptions {
	if o == nil || IsNil(o.Options) {
		var ret AudienceOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetOptionsOk() (*AudienceOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateAudienceAlphaInput) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AudienceOptions and assigns it to the Options field.
func (o *CreateAudienceAlphaInput) SetOptions(v AudienceOptions) {
	o.Options = &v
}

func (o CreateAudienceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAudienceAlphaInput) ToMap() (map[string]interface{}, error) {
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

type NullableCreateAudienceAlphaInput struct {
	value *CreateAudienceAlphaInput
	isSet bool
}

func (v NullableCreateAudienceAlphaInput) Get() *CreateAudienceAlphaInput {
	return v.value
}

func (v *NullableCreateAudienceAlphaInput) Set(val *CreateAudienceAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAudienceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAudienceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAudienceAlphaInput(
	val *CreateAudienceAlphaInput,
) *NullableCreateAudienceAlphaInput {
	return &NullableCreateAudienceAlphaInput{value: val, isSet: true}
}

func (v NullableCreateAudienceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAudienceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
