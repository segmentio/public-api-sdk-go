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

// checks if the CreateTrackingPlanV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTrackingPlanV1Input{}

// CreateTrackingPlanV1Input Creates a Tracking Plan in the Workspace.
type CreateTrackingPlanV1Input struct {
	// The Tracking Plan's name.  Config API note: equal to `displayName`.
	Name string `json:"name"`
	// The Tracking Plan's description.
	Description *string `json:"description,omitempty"`
	// The Tracking Plan's type.
	Type string `json:"type"`
}

// NewCreateTrackingPlanV1Input instantiates a new CreateTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTrackingPlanV1Input(name string, type_ string) *CreateTrackingPlanV1Input {
	this := CreateTrackingPlanV1Input{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewCreateTrackingPlanV1InputWithDefaults instantiates a new CreateTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTrackingPlanV1InputWithDefaults() *CreateTrackingPlanV1Input {
	this := CreateTrackingPlanV1Input{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTrackingPlanV1Input) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTrackingPlanV1Input) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTrackingPlanV1Input) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateTrackingPlanV1Input) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTrackingPlanV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateTrackingPlanV1Input) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateTrackingPlanV1Input) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *CreateTrackingPlanV1Input) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateTrackingPlanV1Input) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateTrackingPlanV1Input) SetType(v string) {
	o.Type = v
}

func (o CreateTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTrackingPlanV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCreateTrackingPlanV1Input struct {
	value *CreateTrackingPlanV1Input
	isSet bool
}

func (v NullableCreateTrackingPlanV1Input) Get() *CreateTrackingPlanV1Input {
	return v.value
}

func (v *NullableCreateTrackingPlanV1Input) Set(val *CreateTrackingPlanV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTrackingPlanV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTrackingPlanV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTrackingPlanV1Input(
	val *CreateTrackingPlanV1Input,
) *NullableCreateTrackingPlanV1Input {
	return &NullableCreateTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableCreateTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
