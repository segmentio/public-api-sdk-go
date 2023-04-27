/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// FunctionSettingV1 struct for FunctionSettingV1
type FunctionSettingV1 struct {
	// The name of this Function Setting.
	Name string `json:"name"`
	// The label for this Function Setting.
	Label string `json:"label"`
	// A description of this Function Setting.
	Description string `json:"description"`
	// The type of this Function Setting.
	Type string `json:"type"`
	// Whether this Function Setting is required.
	Required bool `json:"required"`
	// Whether this Function Setting contains sensitive information.
	Sensitive bool `json:"sensitive"`
}

// NewFunctionSettingV1 instantiates a new FunctionSettingV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunctionSettingV1(
	name string,
	label string,
	description string,
	type_ string,
	required bool,
	sensitive bool,
) *FunctionSettingV1 {
	this := FunctionSettingV1{}
	this.Name = name
	this.Label = label
	this.Description = description
	this.Type = type_
	this.Required = required
	this.Sensitive = sensitive
	return &this
}

// NewFunctionSettingV1WithDefaults instantiates a new FunctionSettingV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunctionSettingV1WithDefaults() *FunctionSettingV1 {
	this := FunctionSettingV1{}
	return &this
}

// GetName returns the Name field value
func (o *FunctionSettingV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FunctionSettingV1) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value
func (o *FunctionSettingV1) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FunctionSettingV1) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value
func (o *FunctionSettingV1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *FunctionSettingV1) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *FunctionSettingV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FunctionSettingV1) SetType(v string) {
	o.Type = v
}

// GetRequired returns the Required field value
func (o *FunctionSettingV1) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FunctionSettingV1) SetRequired(v bool) {
	o.Required = v
}

// GetSensitive returns the Sensitive field value
func (o *FunctionSettingV1) GetSensitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value
// and a boolean to check if the value has been set.
func (o *FunctionSettingV1) GetSensitiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sensitive, true
}

// SetSensitive sets field value
func (o *FunctionSettingV1) SetSensitive(v bool) {
	o.Sensitive = v
}

func (o FunctionSettingV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["required"] = o.Required
	}
	if true {
		toSerialize["sensitive"] = o.Sensitive
	}
	return json.Marshal(toSerialize)
}

type NullableFunctionSettingV1 struct {
	value *FunctionSettingV1
	isSet bool
}

func (v NullableFunctionSettingV1) Get() *FunctionSettingV1 {
	return v.value
}

func (v *NullableFunctionSettingV1) Set(val *FunctionSettingV1) {
	v.value = val
	v.isSet = true
}

func (v NullableFunctionSettingV1) IsSet() bool {
	return v.isSet
}

func (v *NullableFunctionSettingV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunctionSettingV1(val *FunctionSettingV1) *NullableFunctionSettingV1 {
	return &NullableFunctionSettingV1{value: val, isSet: true}
}

func (v NullableFunctionSettingV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunctionSettingV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
