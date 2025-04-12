/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationSubscriptionUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationSubscriptionUpdateInput{}

// DestinationSubscriptionUpdateInput The input parameters for updating a Destination subscription.
type DestinationSubscriptionUpdateInput struct {
	// The user-defined name for the subscription.
	Name *string `json:"name,omitempty"`
	// The fql statement.
	Trigger *string `json:"trigger,omitempty"`
	// Is the subscription enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Represents settings used to configure an action subscription.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// (Reverse ETL only) The reverse ETL model to attach this subscription to.
	ReverseETLModelId  *string                       `json:"reverseETLModelId,omitempty"`
	ReverseETLSchedule *ReverseEtlScheduleDefinition `json:"reverseETLSchedule,omitempty"`
}

// NewDestinationSubscriptionUpdateInput instantiates a new DestinationSubscriptionUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationSubscriptionUpdateInput() *DestinationSubscriptionUpdateInput {
	this := DestinationSubscriptionUpdateInput{}
	return &this
}

// NewDestinationSubscriptionUpdateInputWithDefaults instantiates a new DestinationSubscriptionUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationSubscriptionUpdateInputWithDefaults() *DestinationSubscriptionUpdateInput {
	this := DestinationSubscriptionUpdateInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DestinationSubscriptionUpdateInput) SetName(v string) {
	o.Name = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *DestinationSubscriptionUpdateInput) SetTrigger(v string) {
	o.Trigger = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DestinationSubscriptionUpdateInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetSettings() map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *DestinationSubscriptionUpdateInput) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetReverseETLModelId returns the ReverseETLModelId field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetReverseETLModelId() string {
	if o == nil || IsNil(o.ReverseETLModelId) {
		var ret string
		return ret
	}
	return *o.ReverseETLModelId
}

// GetReverseETLModelIdOk returns a tuple with the ReverseETLModelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetReverseETLModelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReverseETLModelId) {
		return nil, false
	}
	return o.ReverseETLModelId, true
}

// HasReverseETLModelId returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasReverseETLModelId() bool {
	if o != nil && !IsNil(o.ReverseETLModelId) {
		return true
	}

	return false
}

// SetReverseETLModelId gets a reference to the given string and assigns it to the ReverseETLModelId field.
func (o *DestinationSubscriptionUpdateInput) SetReverseETLModelId(v string) {
	o.ReverseETLModelId = &v
}

// GetReverseETLSchedule returns the ReverseETLSchedule field value if set, zero value otherwise.
func (o *DestinationSubscriptionUpdateInput) GetReverseETLSchedule() ReverseEtlScheduleDefinition {
	if o == nil || IsNil(o.ReverseETLSchedule) {
		var ret ReverseEtlScheduleDefinition
		return ret
	}
	return *o.ReverseETLSchedule
}

// GetReverseETLScheduleOk returns a tuple with the ReverseETLSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationSubscriptionUpdateInput) GetReverseETLScheduleOk() (*ReverseEtlScheduleDefinition, bool) {
	if o == nil || IsNil(o.ReverseETLSchedule) {
		return nil, false
	}
	return o.ReverseETLSchedule, true
}

// HasReverseETLSchedule returns a boolean if a field has been set.
func (o *DestinationSubscriptionUpdateInput) HasReverseETLSchedule() bool {
	if o != nil && !IsNil(o.ReverseETLSchedule) {
		return true
	}

	return false
}

// SetReverseETLSchedule gets a reference to the given ReverseEtlScheduleDefinition and assigns it to the ReverseETLSchedule field.
func (o *DestinationSubscriptionUpdateInput) SetReverseETLSchedule(v ReverseEtlScheduleDefinition) {
	o.ReverseETLSchedule = &v
}

func (o DestinationSubscriptionUpdateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationSubscriptionUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.ReverseETLModelId) {
		toSerialize["reverseETLModelId"] = o.ReverseETLModelId
	}
	if !IsNil(o.ReverseETLSchedule) {
		toSerialize["reverseETLSchedule"] = o.ReverseETLSchedule
	}
	return toSerialize, nil
}

type NullableDestinationSubscriptionUpdateInput struct {
	value *DestinationSubscriptionUpdateInput
	isSet bool
}

func (v NullableDestinationSubscriptionUpdateInput) Get() *DestinationSubscriptionUpdateInput {
	return v.value
}

func (v *NullableDestinationSubscriptionUpdateInput) Set(val *DestinationSubscriptionUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationSubscriptionUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationSubscriptionUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationSubscriptionUpdateInput(
	val *DestinationSubscriptionUpdateInput,
) *NullableDestinationSubscriptionUpdateInput {
	return &NullableDestinationSubscriptionUpdateInput{value: val, isSet: true}
}

func (v NullableDestinationSubscriptionUpdateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationSubscriptionUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
