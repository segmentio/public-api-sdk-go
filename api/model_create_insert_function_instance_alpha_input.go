/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 56.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateInsertFunctionInstanceAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInsertFunctionInstanceAlphaInput{}

// CreateInsertFunctionInstanceAlphaInput Creates an insert Function instance.
type CreateInsertFunctionInstanceAlphaInput struct {
	// Insert Function id to which this instance is associated.
	FunctionId string `json:"functionId"`
	// The Source or Destination id to be connected.
	IntegrationId string `json:"integrationId"`
	// Whether this insert Function instance should be enabled for the Destination.
	Enabled *bool `json:"enabled,omitempty"`
	// Defines the display name of the insert Function instance.
	Name string `json:"name"`
	// An object that contains settings for this insert Function instance based on the settings present in the insert Function class.
	Settings map[string]interface{} `json:"settings"`
}

// NewCreateInsertFunctionInstanceAlphaInput instantiates a new CreateInsertFunctionInstanceAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInsertFunctionInstanceAlphaInput(
	functionId string,
	integrationId string,
	name string,
	settings map[string]interface{},
) *CreateInsertFunctionInstanceAlphaInput {
	this := CreateInsertFunctionInstanceAlphaInput{}
	this.FunctionId = functionId
	this.IntegrationId = integrationId
	this.Name = name
	this.Settings = settings
	return &this
}

// NewCreateInsertFunctionInstanceAlphaInputWithDefaults instantiates a new CreateInsertFunctionInstanceAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInsertFunctionInstanceAlphaInputWithDefaults() *CreateInsertFunctionInstanceAlphaInput {
	this := CreateInsertFunctionInstanceAlphaInput{}
	return &this
}

// GetFunctionId returns the FunctionId field value
func (o *CreateInsertFunctionInstanceAlphaInput) GetFunctionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FunctionId
}

// GetFunctionIdOk returns a tuple with the FunctionId field value
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) GetFunctionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionId, true
}

// SetFunctionId sets field value
func (o *CreateInsertFunctionInstanceAlphaInput) SetFunctionId(v string) {
	o.FunctionId = v
}

// GetIntegrationId returns the IntegrationId field value
func (o *CreateInsertFunctionInstanceAlphaInput) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *CreateInsertFunctionInstanceAlphaInput) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateInsertFunctionInstanceAlphaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateInsertFunctionInstanceAlphaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value
func (o *CreateInsertFunctionInstanceAlphaInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateInsertFunctionInstanceAlphaInput) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *CreateInsertFunctionInstanceAlphaInput) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CreateInsertFunctionInstanceAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *CreateInsertFunctionInstanceAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o CreateInsertFunctionInstanceAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInsertFunctionInstanceAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["functionId"] = o.FunctionId
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["name"] = o.Name
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableCreateInsertFunctionInstanceAlphaInput struct {
	value *CreateInsertFunctionInstanceAlphaInput
	isSet bool
}

func (v NullableCreateInsertFunctionInstanceAlphaInput) Get() *CreateInsertFunctionInstanceAlphaInput {
	return v.value
}

func (v *NullableCreateInsertFunctionInstanceAlphaInput) Set(
	val *CreateInsertFunctionInstanceAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInsertFunctionInstanceAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInsertFunctionInstanceAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInsertFunctionInstanceAlphaInput(
	val *CreateInsertFunctionInstanceAlphaInput,
) *NullableCreateInsertFunctionInstanceAlphaInput {
	return &NullableCreateInsertFunctionInstanceAlphaInput{value: val, isSet: true}
}

func (v NullableCreateInsertFunctionInstanceAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInsertFunctionInstanceAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
