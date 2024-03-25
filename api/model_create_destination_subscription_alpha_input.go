/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateDestinationSubscriptionAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDestinationSubscriptionAlphaInput{}

// CreateDestinationSubscriptionAlphaInput The basic input parameters for creating a Destination subscription.
type CreateDestinationSubscriptionAlphaInput struct {
	// A user-defined name for the subscription.
	Name string `json:"name"`
	// The associated action id the subscription should trigger.
	ActionId string `json:"actionId"`
	// The FQL statement.
	Trigger string `json:"trigger"`
	// Is the subscription enabled.
	Enabled bool `json:"enabled"`
	// Represents settings used to configure an action subscription.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// When creating a Reverse ETL connection, indicates the Model being used to extract data.
	ModelId *string `json:"modelId,omitempty"`
}

// NewCreateDestinationSubscriptionAlphaInput instantiates a new CreateDestinationSubscriptionAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationSubscriptionAlphaInput(
	name string,
	actionId string,
	trigger string,
	enabled bool,
) *CreateDestinationSubscriptionAlphaInput {
	this := CreateDestinationSubscriptionAlphaInput{}
	this.Name = name
	this.ActionId = actionId
	this.Trigger = trigger
	this.Enabled = enabled
	return &this
}

// NewCreateDestinationSubscriptionAlphaInputWithDefaults instantiates a new CreateDestinationSubscriptionAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationSubscriptionAlphaInputWithDefaults() *CreateDestinationSubscriptionAlphaInput {
	this := CreateDestinationSubscriptionAlphaInput{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDestinationSubscriptionAlphaInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDestinationSubscriptionAlphaInput) SetName(v string) {
	o.Name = v
}

// GetActionId returns the ActionId field value
func (o *CreateDestinationSubscriptionAlphaInput) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *CreateDestinationSubscriptionAlphaInput) SetActionId(v string) {
	o.ActionId = v
}

// GetTrigger returns the Trigger field value
func (o *CreateDestinationSubscriptionAlphaInput) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *CreateDestinationSubscriptionAlphaInput) SetTrigger(v string) {
	o.Trigger = v
}

// GetEnabled returns the Enabled field value
func (o *CreateDestinationSubscriptionAlphaInput) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateDestinationSubscriptionAlphaInput) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateDestinationSubscriptionAlphaInput) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateDestinationSubscriptionAlphaInput) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *CreateDestinationSubscriptionAlphaInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetModelId returns the ModelId field value if set, zero value otherwise.
func (o *CreateDestinationSubscriptionAlphaInput) GetModelId() string {
	if o == nil || IsNil(o.ModelId) {
		var ret string
		return ret
	}
	return *o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationSubscriptionAlphaInput) GetModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ModelId) {
		return nil, false
	}
	return o.ModelId, true
}

// HasModelId returns a boolean if a field has been set.
func (o *CreateDestinationSubscriptionAlphaInput) HasModelId() bool {
	if o != nil && !IsNil(o.ModelId) {
		return true
	}

	return false
}

// SetModelId gets a reference to the given string and assigns it to the ModelId field.
func (o *CreateDestinationSubscriptionAlphaInput) SetModelId(v string) {
	o.ModelId = &v
}

func (o CreateDestinationSubscriptionAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDestinationSubscriptionAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["actionId"] = o.ActionId
	toSerialize["trigger"] = o.Trigger
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.ModelId) {
		toSerialize["modelId"] = o.ModelId
	}
	return toSerialize, nil
}

type NullableCreateDestinationSubscriptionAlphaInput struct {
	value *CreateDestinationSubscriptionAlphaInput
	isSet bool
}

func (v NullableCreateDestinationSubscriptionAlphaInput) Get() *CreateDestinationSubscriptionAlphaInput {
	return v.value
}

func (v *NullableCreateDestinationSubscriptionAlphaInput) Set(
	val *CreateDestinationSubscriptionAlphaInput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationSubscriptionAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationSubscriptionAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationSubscriptionAlphaInput(
	val *CreateDestinationSubscriptionAlphaInput,
) *NullableCreateDestinationSubscriptionAlphaInput {
	return &NullableCreateDestinationSubscriptionAlphaInput{value: val, isSet: true}
}

func (v NullableCreateDestinationSubscriptionAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDestinationSubscriptionAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
