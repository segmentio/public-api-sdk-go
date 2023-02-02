/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateFunctionV1Input Update a Function.
type UpdateFunctionV1Input struct {
	// The Function code.
	Code *string `json:"code,omitempty"`
	// The list of settings for this Function.
	Settings []FunctionSettingV1 `json:"settings,omitempty"`
	// A display name for this Function.
	DisplayName *string `json:"displayName,omitempty"`
	// A logo for this Function.
	LogoUrl *string `json:"logoUrl,omitempty"`
	// A description for this Function.
	Description *string `json:"description,omitempty"`
}

// NewUpdateFunctionV1Input instantiates a new UpdateFunctionV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFunctionV1Input() *UpdateFunctionV1Input {
	this := UpdateFunctionV1Input{}
	return &this
}

// NewUpdateFunctionV1InputWithDefaults instantiates a new UpdateFunctionV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFunctionV1InputWithDefaults() *UpdateFunctionV1Input {
	this := UpdateFunctionV1Input{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateFunctionV1Input) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Input) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateFunctionV1Input) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateFunctionV1Input) SetCode(v string) {
	o.Code = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *UpdateFunctionV1Input) GetSettings() []FunctionSettingV1 {
	if o == nil || o.Settings == nil {
		var ret []FunctionSettingV1
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Input) GetSettingsOk() ([]FunctionSettingV1, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *UpdateFunctionV1Input) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given []FunctionSettingV1 and assigns it to the Settings field.
func (o *UpdateFunctionV1Input) SetSettings(v []FunctionSettingV1) {
	o.Settings = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateFunctionV1Input) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Input) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateFunctionV1Input) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateFunctionV1Input) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *UpdateFunctionV1Input) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Input) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *UpdateFunctionV1Input) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *UpdateFunctionV1Input) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFunctionV1Input) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFunctionV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFunctionV1Input) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFunctionV1Input) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateFunctionV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.LogoUrl != nil {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateFunctionV1Input struct {
	value *UpdateFunctionV1Input
	isSet bool
}

func (v NullableUpdateFunctionV1Input) Get() *UpdateFunctionV1Input {
	return v.value
}

func (v *NullableUpdateFunctionV1Input) Set(val *UpdateFunctionV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFunctionV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFunctionV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFunctionV1Input(val *UpdateFunctionV1Input) *NullableUpdateFunctionV1Input {
	return &NullableUpdateFunctionV1Input{value: val, isSet: true}
}

func (v NullableUpdateFunctionV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFunctionV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
