/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 41.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddSourceToTrackingPlanV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSourceToTrackingPlanV1Input{}

// AddSourceToTrackingPlanV1Input Connects a Source to a Tracking Plan.
type AddSourceToTrackingPlanV1Input struct {
	// The id of the Source associated with the Tracking Plan.  Config API note: analogous to `sourceName`.
	SourceId string `json:"sourceId"`
}

// NewAddSourceToTrackingPlanV1Input instantiates a new AddSourceToTrackingPlanV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSourceToTrackingPlanV1Input(sourceId string) *AddSourceToTrackingPlanV1Input {
	this := AddSourceToTrackingPlanV1Input{}
	this.SourceId = sourceId
	return &this
}

// NewAddSourceToTrackingPlanV1InputWithDefaults instantiates a new AddSourceToTrackingPlanV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSourceToTrackingPlanV1InputWithDefaults() *AddSourceToTrackingPlanV1Input {
	this := AddSourceToTrackingPlanV1Input{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *AddSourceToTrackingPlanV1Input) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *AddSourceToTrackingPlanV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *AddSourceToTrackingPlanV1Input) SetSourceId(v string) {
	o.SourceId = v
}

func (o AddSourceToTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSourceToTrackingPlanV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	return toSerialize, nil
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

func NewNullableAddSourceToTrackingPlanV1Input(
	val *AddSourceToTrackingPlanV1Input,
) *NullableAddSourceToTrackingPlanV1Input {
	return &NullableAddSourceToTrackingPlanV1Input{value: val, isSet: true}
}

func (v NullableAddSourceToTrackingPlanV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSourceToTrackingPlanV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
