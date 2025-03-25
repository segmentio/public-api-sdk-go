/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the JourneyDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JourneyDefinition{}

// JourneyDefinition The journey definition.
type JourneyDefinition struct {
	// The initial state of the journey.
	InitialState string          `json:"initialState"`
	EntryRules   EntryRules      `json:"entryRules"`
	ExitRules    ExitRulesConfig `json:"exitRules"`
	// The states of the journey.
	States []JourneyState `json:"states"`
}

// NewJourneyDefinition instantiates a new JourneyDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJourneyDefinition(
	initialState string,
	entryRules EntryRules,
	exitRules ExitRulesConfig,
	states []JourneyState,
) *JourneyDefinition {
	this := JourneyDefinition{}
	this.InitialState = initialState
	this.EntryRules = entryRules
	this.ExitRules = exitRules
	this.States = states
	return &this
}

// NewJourneyDefinitionWithDefaults instantiates a new JourneyDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJourneyDefinitionWithDefaults() *JourneyDefinition {
	this := JourneyDefinition{}
	return &this
}

// GetInitialState returns the InitialState field value
func (o *JourneyDefinition) GetInitialState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialState
}

// GetInitialStateOk returns a tuple with the InitialState field value
// and a boolean to check if the value has been set.
func (o *JourneyDefinition) GetInitialStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialState, true
}

// SetInitialState sets field value
func (o *JourneyDefinition) SetInitialState(v string) {
	o.InitialState = v
}

// GetEntryRules returns the EntryRules field value
func (o *JourneyDefinition) GetEntryRules() EntryRules {
	if o == nil {
		var ret EntryRules
		return ret
	}

	return o.EntryRules
}

// GetEntryRulesOk returns a tuple with the EntryRules field value
// and a boolean to check if the value has been set.
func (o *JourneyDefinition) GetEntryRulesOk() (*EntryRules, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryRules, true
}

// SetEntryRules sets field value
func (o *JourneyDefinition) SetEntryRules(v EntryRules) {
	o.EntryRules = v
}

// GetExitRules returns the ExitRules field value
func (o *JourneyDefinition) GetExitRules() ExitRulesConfig {
	if o == nil {
		var ret ExitRulesConfig
		return ret
	}

	return o.ExitRules
}

// GetExitRulesOk returns a tuple with the ExitRules field value
// and a boolean to check if the value has been set.
func (o *JourneyDefinition) GetExitRulesOk() (*ExitRulesConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExitRules, true
}

// SetExitRules sets field value
func (o *JourneyDefinition) SetExitRules(v ExitRulesConfig) {
	o.ExitRules = v
}

// GetStates returns the States field value
func (o *JourneyDefinition) GetStates() []JourneyState {
	if o == nil {
		var ret []JourneyState
		return ret
	}

	return o.States
}

// GetStatesOk returns a tuple with the States field value
// and a boolean to check if the value has been set.
func (o *JourneyDefinition) GetStatesOk() ([]JourneyState, bool) {
	if o == nil {
		return nil, false
	}
	return o.States, true
}

// SetStates sets field value
func (o *JourneyDefinition) SetStates(v []JourneyState) {
	o.States = v
}

func (o JourneyDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JourneyDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["initialState"] = o.InitialState
	toSerialize["entryRules"] = o.EntryRules
	toSerialize["exitRules"] = o.ExitRules
	toSerialize["states"] = o.States
	return toSerialize, nil
}

type NullableJourneyDefinition struct {
	value *JourneyDefinition
	isSet bool
}

func (v NullableJourneyDefinition) Get() *JourneyDefinition {
	return v.value
}

func (v *NullableJourneyDefinition) Set(val *JourneyDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableJourneyDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableJourneyDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJourneyDefinition(val *JourneyDefinition) *NullableJourneyDefinition {
	return &NullableJourneyDefinition{value: val, isSet: true}
}

func (v NullableJourneyDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJourneyDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
