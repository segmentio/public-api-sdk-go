/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReplaceStepsForJourneyAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceStepsForJourneyAlphaInput{}

// ReplaceStepsForJourneyAlphaInput Replace steps for a journey input.
type ReplaceStepsForJourneyAlphaInput struct {
	// The Initial state of the journey.
	InitialState string `json:"initialState"`
	// States of the journey.
	States []JourneyState `json:"states"`
}

// NewReplaceStepsForJourneyAlphaInput instantiates a new ReplaceStepsForJourneyAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceStepsForJourneyAlphaInput(
	initialState string,
	states []JourneyState,
) *ReplaceStepsForJourneyAlphaInput {
	this := ReplaceStepsForJourneyAlphaInput{}
	this.InitialState = initialState
	this.States = states
	return &this
}

// NewReplaceStepsForJourneyAlphaInputWithDefaults instantiates a new ReplaceStepsForJourneyAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceStepsForJourneyAlphaInputWithDefaults() *ReplaceStepsForJourneyAlphaInput {
	this := ReplaceStepsForJourneyAlphaInput{}
	return &this
}

// GetInitialState returns the InitialState field value
func (o *ReplaceStepsForJourneyAlphaInput) GetInitialState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialState
}

// GetInitialStateOk returns a tuple with the InitialState field value
// and a boolean to check if the value has been set.
func (o *ReplaceStepsForJourneyAlphaInput) GetInitialStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialState, true
}

// SetInitialState sets field value
func (o *ReplaceStepsForJourneyAlphaInput) SetInitialState(v string) {
	o.InitialState = v
}

// GetStates returns the States field value
func (o *ReplaceStepsForJourneyAlphaInput) GetStates() []JourneyState {
	if o == nil {
		var ret []JourneyState
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *ReplaceStepsForJourneyAlphaInput) GetStatesOk() ([]JourneyState, bool) {
	if o == nil {
		return nil, false
	}
	return o.States, true
}

// SetStates sets field value
func (o *ReplaceStepsForJourneyAlphaInput) SetStates(v []JourneyState) {
	o.States = v
}

func (o ReplaceStepsForJourneyAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceStepsForJourneyAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["initialState"] = o.InitialState
	toSerialize["states"] = o.States
	return toSerialize, nil
}

type NullableReplaceStepsForJourneyAlphaInput struct {
	value *ReplaceStepsForJourneyAlphaInput
	isSet bool
}

func (v NullableReplaceStepsForJourneyAlphaInput) Get() *ReplaceStepsForJourneyAlphaInput {
	return v.value
}

func (v *NullableReplaceStepsForJourneyAlphaInput) Set(val *ReplaceStepsForJourneyAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceStepsForJourneyAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceStepsForJourneyAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceStepsForJourneyAlphaInput(
	val *ReplaceStepsForJourneyAlphaInput,
) *NullableReplaceStepsForJourneyAlphaInput {
	return &NullableReplaceStepsForJourneyAlphaInput{value: val, isSet: true}
}

func (v NullableReplaceStepsForJourneyAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceStepsForJourneyAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
