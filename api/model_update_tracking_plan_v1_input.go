/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateTrackingPlanV1Input Updates the Workspace's Tracking Plan.
type UpdateTrackingPlanV1Input struct {
	// The Tracking Plan's name.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// The Tracking Plan's description.
	Description *string `json:"description,omitempty"`
}

// NewUpdateTrackingPlanV1Input instantiates a new UpdateTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTrackingPlanV1Input() *UpdateTrackingPlanV1Input {
	this := UpdateTrackingPlanV1Input{}
	return &this
}

// NewUpdateTrackingPlanV1InputWithDefaults instantiates a new UpdateTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTrackingPlanV1InputWithDefaults() *UpdateTrackingPlanV1Input {
	this := UpdateTrackingPlanV1Input{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTrackingPlanV1Input) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTrackingPlanV1Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTrackingPlanV1Input) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTrackingPlanV1Input) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateTrackingPlanV1Input) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTrackingPlanV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateTrackingPlanV1Input) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateTrackingPlanV1Input) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTrackingPlanV1Input struct {
	value *UpdateTrackingPlanV1Input
	isSet bool
}

func (v NullableUpdateTrackingPlanV1Input) Get() *UpdateTrackingPlanV1Input {
	return v.value
}

func (v *NullableUpdateTrackingPlanV1Input) Set(val *UpdateTrackingPlanV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTrackingPlanV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTrackingPlanV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTrackingPlanV1Input(val *UpdateTrackingPlanV1Input) *NullableUpdateTrackingPlanV1Input {
	return &NullableUpdateTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableUpdateTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


