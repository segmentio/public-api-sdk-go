/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 48.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddSourceToTrackingPlanV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSourceToTrackingPlanV1Output{}

// AddSourceToTrackingPlanV1Output Connects a Source to a Tracking Plan.
type AddSourceToTrackingPlanV1Output struct {
	// The operation status.
	Status string `json:"status"`
}

// NewAddSourceToTrackingPlanV1Output instantiates a new AddSourceToTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSourceToTrackingPlanV1Output(status string) *AddSourceToTrackingPlanV1Output {
	this := AddSourceToTrackingPlanV1Output{}
	this.Status = status
	return &this
}

// NewAddSourceToTrackingPlanV1OutputWithDefaults instantiates a new AddSourceToTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSourceToTrackingPlanV1OutputWithDefaults() *AddSourceToTrackingPlanV1Output {
	this := AddSourceToTrackingPlanV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *AddSourceToTrackingPlanV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AddSourceToTrackingPlanV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AddSourceToTrackingPlanV1Output) SetStatus(v string) {
	o.Status = v
}

func (o AddSourceToTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSourceToTrackingPlanV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableAddSourceToTrackingPlanV1Output struct {
	value *AddSourceToTrackingPlanV1Output
	isSet bool
}

func (v NullableAddSourceToTrackingPlanV1Output) Get() *AddSourceToTrackingPlanV1Output {
	return v.value
}

func (v *NullableAddSourceToTrackingPlanV1Output) Set(val *AddSourceToTrackingPlanV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSourceToTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSourceToTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSourceToTrackingPlanV1Output(
	val *AddSourceToTrackingPlanV1Output,
) *NullableAddSourceToTrackingPlanV1Output {
	return &NullableAddSourceToTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableAddSourceToTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSourceToTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
