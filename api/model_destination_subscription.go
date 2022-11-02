/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationSubscription The Destination subscription.
type DestinationSubscription struct {
	// The unique identifier for the subscription.
	Id string `json:"id"`
	// The name of the subscription.
	Name string `json:"name"`
	// The unique identifier for the Destination action to trigger.
	ActionId string `json:"actionId"`
	// The URL-friendly key for the associated Destination action.
	ActionSlug string `json:"actionSlug"`
	// The associated Destination instance id.
	DestinationId string `json:"destinationId"`
	// Is the subscription enabled.
	Enabled bool `json:"enabled"`
	// The customer settings for action fields.
	Settings ModelMap `json:"settings"`
	// FQL string that describes what events should trigger a Destination action.
	Trigger string `json:"trigger"`
}

// NewDestinationSubscription instantiates a new DestinationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationSubscription(id string, name string, actionId string, actionSlug string, destinationId string, enabled bool, settings ModelMap, trigger string) *DestinationSubscription {
	this := DestinationSubscription{}
	this.Id = id
	this.Name = name
	this.ActionId = actionId
	this.ActionSlug = actionSlug
	this.DestinationId = destinationId
	this.Enabled = enabled
	this.Settings = settings
	this.Trigger = trigger
	return &this
}

// NewDestinationSubscriptionWithDefaults instantiates a new DestinationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationSubscriptionWithDefaults() *DestinationSubscription {
	this := DestinationSubscription{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationSubscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationSubscription) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DestinationSubscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationSubscription) SetName(v string) {
	o.Name = v
}

// GetActionId returns the ActionId field value
func (o *DestinationSubscription) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *DestinationSubscription) SetActionId(v string) {
	o.ActionId = v
}

// GetActionSlug returns the ActionSlug field value
func (o *DestinationSubscription) GetActionSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionSlug
}

// GetActionSlugOk returns a tuple with the ActionSlug field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetActionSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionSlug, true
}

// SetActionSlug sets field value
func (o *DestinationSubscription) SetActionSlug(v string) {
	o.ActionSlug = v
}

// GetDestinationId returns the DestinationId field value
func (o *DestinationSubscription) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *DestinationSubscription) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetEnabled returns the Enabled field value
func (o *DestinationSubscription) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DestinationSubscription) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSettings returns the Settings field value
func (o *DestinationSubscription) GetSettings() ModelMap {
	if o == nil {
		var ret ModelMap
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetSettingsOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *DestinationSubscription) SetSettings(v ModelMap) {
	o.Settings = v
}

// GetTrigger returns the Trigger field value
func (o *DestinationSubscription) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *DestinationSubscription) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *DestinationSubscription) SetTrigger(v string) {
	o.Trigger = v
}

func (o DestinationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["actionId"] = o.ActionId
	}
	if true {
		toSerialize["actionSlug"] = o.ActionSlug
	}
	if true {
		toSerialize["destinationId"] = o.DestinationId
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationSubscription struct {
	value *DestinationSubscription
	isSet bool
}

func (v NullableDestinationSubscription) Get() *DestinationSubscription {
	return v.value
}

func (v *NullableDestinationSubscription) Set(val *DestinationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationSubscription(val *DestinationSubscription) *NullableDestinationSubscription {
	return &NullableDestinationSubscription{value: val, isSet: true}
}

func (v NullableDestinationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


