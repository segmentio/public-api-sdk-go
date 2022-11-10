/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationMetadataSubscriptionPresetV1 Represents a set of defaults for a Destination subscription.
type DestinationMetadataSubscriptionPresetV1 struct {
	// The unique identifier for the Destination Action to trigger.
	ActionId string `json:"actionId"`
	// The name of the subscription.
	Name string `json:"name"`
	// The default settings for action fields.
	Fields map[string]interface{} `json:"fields"`
	// FQL string that describes what events should trigger an action. See https://segment.com/docs/config-api/fql/ for more information regarding Segment's Filter Query Language (FQL).
	Trigger string `json:"trigger"`
}

// NewDestinationMetadataSubscriptionPresetV1 instantiates a new DestinationMetadataSubscriptionPresetV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataSubscriptionPresetV1(actionId string, name string, fields map[string]interface{}, trigger string) *DestinationMetadataSubscriptionPresetV1 {
	this := DestinationMetadataSubscriptionPresetV1{}
	this.ActionId = actionId
	this.Name = name
	this.Fields = fields
	this.Trigger = trigger
	return &this
}

// NewDestinationMetadataSubscriptionPresetV1WithDefaults instantiates a new DestinationMetadataSubscriptionPresetV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataSubscriptionPresetV1WithDefaults() *DestinationMetadataSubscriptionPresetV1 {
	this := DestinationMetadataSubscriptionPresetV1{}
	return &this
}

// GetActionId returns the ActionId field value
func (o *DestinationMetadataSubscriptionPresetV1) GetActionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionId
}

// GetActionIdOk returns a tuple with the ActionId field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataSubscriptionPresetV1) GetActionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionId, true
}

// SetActionId sets field value
func (o *DestinationMetadataSubscriptionPresetV1) SetActionId(v string) {
	o.ActionId = v
}

// GetName returns the Name field value
func (o *DestinationMetadataSubscriptionPresetV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataSubscriptionPresetV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationMetadataSubscriptionPresetV1) SetName(v string) {
	o.Name = v
}

// GetFields returns the Fields field value
func (o *DestinationMetadataSubscriptionPresetV1) GetFields() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataSubscriptionPresetV1) GetFieldsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *DestinationMetadataSubscriptionPresetV1) SetFields(v map[string]interface{}) {
	o.Fields = v
}

// GetTrigger returns the Trigger field value
func (o *DestinationMetadataSubscriptionPresetV1) GetTrigger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataSubscriptionPresetV1) GetTriggerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *DestinationMetadataSubscriptionPresetV1) SetTrigger(v string) {
	o.Trigger = v
}

func (o DestinationMetadataSubscriptionPresetV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["actionId"] = o.ActionId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationMetadataSubscriptionPresetV1 struct {
	value *DestinationMetadataSubscriptionPresetV1
	isSet bool
}

func (v NullableDestinationMetadataSubscriptionPresetV1) Get() *DestinationMetadataSubscriptionPresetV1 {
	return v.value
}

func (v *NullableDestinationMetadataSubscriptionPresetV1) Set(val *DestinationMetadataSubscriptionPresetV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataSubscriptionPresetV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataSubscriptionPresetV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataSubscriptionPresetV1(val *DestinationMetadataSubscriptionPresetV1) *NullableDestinationMetadataSubscriptionPresetV1 {
	return &NullableDestinationMetadataSubscriptionPresetV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataSubscriptionPresetV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataSubscriptionPresetV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


