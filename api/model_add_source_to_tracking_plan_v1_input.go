/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.7
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// AddSourceToTrackingPlanV1Input Connects a Source to a Tracking Plan.
type AddSourceToTrackingPlanV1Input struct {
	// The id of the Source associated with the Tracking Plan.  Config API note: analogous to `sourceName`.
	SourceId *string `json:"sourceId,omitempty"`
}

// NewAddSourceToTrackingPlanV1Input instantiates a new AddSourceToTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSourceToTrackingPlanV1Input() *AddSourceToTrackingPlanV1Input {
	this := AddSourceToTrackingPlanV1Input{}
	return &this
}

// NewAddSourceToTrackingPlanV1InputWithDefaults instantiates a new AddSourceToTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSourceToTrackingPlanV1InputWithDefaults() *AddSourceToTrackingPlanV1Input {
	this := AddSourceToTrackingPlanV1Input{}
	return &this
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *AddSourceToTrackingPlanV1Input) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSourceToTrackingPlanV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *AddSourceToTrackingPlanV1Input) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *AddSourceToTrackingPlanV1Input) SetSourceId(v string) {
	o.SourceId = &v
}

func (o AddSourceToTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	return json.Marshal(toSerialize)
}

type NullableAddSourceToTrackingPlanV1Input struct {
	value *AddSourceToTrackingPlanV1Input
	isSet bool
}

func (v NullableAddSourceToTrackingPlanV1Input) Get() *AddSourceToTrackingPlanV1Input {
	return v.value
}

func (v *NullableAddSourceToTrackingPlanV1Input) Set(val *AddSourceToTrackingPlanV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSourceToTrackingPlanV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSourceToTrackingPlanV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSourceToTrackingPlanV1Input(val *AddSourceToTrackingPlanV1Input) *NullableAddSourceToTrackingPlanV1Input {
	return &NullableAddSourceToTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableAddSourceToTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSourceToTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


