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

// ListLabelsV1Output Returns all available labels for the current Workspace.
type ListLabelsV1Output struct {
	// All labels associated with the current Workspace.
	Labels []LabelV1 `json:"labels"`
}

// NewListLabelsV1Output instantiates a new ListLabelsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLabelsV1Output(labels []LabelV1) *ListLabelsV1Output {
	this := ListLabelsV1Output{}
	this.Labels = labels
	return &this
}

// NewListLabelsV1OutputWithDefaults instantiates a new ListLabelsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLabelsV1OutputWithDefaults() *ListLabelsV1Output {
	this := ListLabelsV1Output{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ListLabelsV1Output) GetLabels() []LabelV1 {
	if o == nil {
		var ret []LabelV1
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ListLabelsV1Output) GetLabelsOk() ([]LabelV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *ListLabelsV1Output) SetLabels(v []LabelV1) {
	o.Labels = v
}

func (o ListLabelsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableListLabelsV1Output struct {
	value *ListLabelsV1Output
	isSet bool
}

func (v NullableListLabelsV1Output) Get() *ListLabelsV1Output {
	return v.value
}

func (v *NullableListLabelsV1Output) Set(val *ListLabelsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListLabelsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListLabelsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLabelsV1Output(val *ListLabelsV1Output) *NullableListLabelsV1Output {
	return &NullableListLabelsV1Output{value: val, isSet: true}
}

func (v NullableListLabelsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLabelsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


