/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateFunctionV1Input Creates a Function.
type CreateFunctionV1Input struct {
	// The Function code.
	Code string `json:"code"`
	// The list of settings for this Function.
	Settings []FunctionSettingV1 `json:"settings,omitempty"`
	// A display name for this Function.  Note that Destination Functions append the Workspace to the display name.
	DisplayName string `json:"displayName"`
	// The Function type.  Config API note: equal to `type`.
	ResourceType string `json:"resourceType"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
}

// NewCreateFunctionV1Input instantiates a new CreateFunctionV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFunctionV1Input(
	code string,
	displayName string,
	resourceType string,
) *CreateFunctionV1Input {
	this := CreateFunctionV1Input{}
	this.Code = code
	this.DisplayName = displayName
	this.ResourceType = resourceType
	return &this
}

// NewCreateFunctionV1InputWithDefaults instantiates a new CreateFunctionV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFunctionV1InputWithDefaults() *CreateFunctionV1Input {
	this := CreateFunctionV1Input{}
	return &this
}

// GetCode returns the Code field value
func (o *CreateFunctionV1Input) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Input) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CreateFunctionV1Input) SetCode(v string) {
	o.Code = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateFunctionV1Input) GetSettings() []FunctionSettingV1 {
	if o == nil || o.Settings == nil {
		var ret []FunctionSettingV1
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Input) GetSettingsOk() ([]FunctionSettingV1, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateFunctionV1Input) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []FunctionSettingV1 and assigns it to the Settings field.
func (o *CreateFunctionV1Input) SetSettings(v []FunctionSettingV1) {
	o.Settings = v
}

// GetDisplayName returns the DisplayName field value
func (o *CreateFunctionV1Input) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Input) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateFunctionV1Input) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetResourceType returns the ResourceType field value
func (o *CreateFunctionV1Input) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Input) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CreateFunctionV1Input) SetResourceType(v string) {
	o.ResourceType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFunctionV1Input) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFunctionV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFunctionV1Input) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFunctionV1Input) SetDescription(v string) {
	o.Description = &v
}

func (o CreateFunctionV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFunctionV1Input struct {
	value *CreateFunctionV1Input
	isSet bool
}

func (v NullableCreateFunctionV1Input) Get() *CreateFunctionV1Input {
	return v.value
}

func (v *NullableCreateFunctionV1Input) Set(val *CreateFunctionV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFunctionV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFunctionV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFunctionV1Input(val *CreateFunctionV1Input) *NullableCreateFunctionV1Input {
	return &NullableCreateFunctionV1Input{value: val, isSet: true}
}

func (v NullableCreateFunctionV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFunctionV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
