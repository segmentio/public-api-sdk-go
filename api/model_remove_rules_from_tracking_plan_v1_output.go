/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveRulesFromTrackingPlanV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveRulesFromTrackingPlanV1Output{}

// RemoveRulesFromTrackingPlanV1Output Remove specified rules from a Tracking Plan.
type RemoveRulesFromTrackingPlanV1Output struct {
	// The status of the operation.
	Status string `json:"status"`
}

// NewRemoveRulesFromTrackingPlanV1Output instantiates a new RemoveRulesFromTrackingPlanV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveRulesFromTrackingPlanV1Output(status string) *RemoveRulesFromTrackingPlanV1Output {
	this := RemoveRulesFromTrackingPlanV1Output{}
	this.Status = status
	return &this
}

// NewRemoveRulesFromTrackingPlanV1OutputWithDefaults instantiates a new RemoveRulesFromTrackingPlanV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveRulesFromTrackingPlanV1OutputWithDefaults() *RemoveRulesFromTrackingPlanV1Output {
	this := RemoveRulesFromTrackingPlanV1Output{}
	return &this
}

// GetStatus returns the Status field value
func (o *RemoveRulesFromTrackingPlanV1Output) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *RemoveRulesFromTrackingPlanV1Output) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *RemoveRulesFromTrackingPlanV1Output) SetStatus(v string) {
	o.Status = v
}

func (o RemoveRulesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveRulesFromTrackingPlanV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableRemoveRulesFromTrackingPlanV1Output struct {
	value *RemoveRulesFromTrackingPlanV1Output
	isSet bool
}

func (v NullableRemoveRulesFromTrackingPlanV1Output) Get() *RemoveRulesFromTrackingPlanV1Output {
	return v.value
}

func (v *NullableRemoveRulesFromTrackingPlanV1Output) Set(
	val *RemoveRulesFromTrackingPlanV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveRulesFromTrackingPlanV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveRulesFromTrackingPlanV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveRulesFromTrackingPlanV1Output(
	val *RemoveRulesFromTrackingPlanV1Output,
) *NullableRemoveRulesFromTrackingPlanV1Output {
	return &NullableRemoveRulesFromTrackingPlanV1Output{value: val, isSet: true}
}

func (v NullableRemoveRulesFromTrackingPlanV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveRulesFromTrackingPlanV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
