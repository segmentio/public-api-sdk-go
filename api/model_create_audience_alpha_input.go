/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.0.0
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
	// The name of the computation .
	Name string `json:"name"`
	// The description of the computation.
	Description string                        `json:"description"`
	Definition  AudienceComputationDefinition `json:"definition"`
}

// NewCreateAudienceAlphaInput instantiates a new CreateAudienceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAudienceAlphaInput(
	name string,
	description string,
	definition AudienceComputationDefinition,
) *CreateAudienceAlphaInput {
	this := CreateAudienceAlphaInput{}
	this.Name = name
	this.Description = description
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

// GetDescription returns the Description field value
func (o *CreateAudienceAlphaInput) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateAudienceAlphaInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateAudienceAlphaInput) SetDescription(v string) {
	o.Description = v
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
	toSerialize["description"] = o.Description
	toSerialize["definition"] = o.Definition
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