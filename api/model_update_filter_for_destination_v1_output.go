/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateFilterForDestinationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilterForDestinationV1Output{}

// UpdateFilterForDestinationV1Output Output for UpdateDestinationFilterV1.
type UpdateFilterForDestinationV1Output struct {
	Filter DestinationFilterV1 `json:"filter"`
}

// NewUpdateFilterForDestinationV1Output instantiates a new UpdateFilterForDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterForDestinationV1Output(
	filter DestinationFilterV1,
) *UpdateFilterForDestinationV1Output {
	this := UpdateFilterForDestinationV1Output{}
	this.Filter = filter
	return &this
}

// NewUpdateFilterForDestinationV1OutputWithDefaults instantiates a new UpdateFilterForDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterForDestinationV1OutputWithDefaults() *UpdateFilterForDestinationV1Output {
	this := UpdateFilterForDestinationV1Output{}
	return &this
}

// GetFilter returns the Filter field value
func (o *UpdateFilterForDestinationV1Output) GetFilter() DestinationFilterV1 {
	if o == nil {
		var ret DestinationFilterV1
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestinationV1Output) GetFilterOk() (*DestinationFilterV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *UpdateFilterForDestinationV1Output) SetFilter(v DestinationFilterV1) {
	o.Filter = v
}

func (o UpdateFilterForDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilterForDestinationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	return toSerialize, nil
}

type NullableUpdateFilterForDestinationV1Output struct {
	value *UpdateFilterForDestinationV1Output
	isSet bool
}

func (v NullableUpdateFilterForDestinationV1Output) Get() *UpdateFilterForDestinationV1Output {
	return v.value
}

func (v *NullableUpdateFilterForDestinationV1Output) Set(val *UpdateFilterForDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterForDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterForDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterForDestinationV1Output(
	val *UpdateFilterForDestinationV1Output,
) *NullableUpdateFilterForDestinationV1Output {
	return &NullableUpdateFilterForDestinationV1Output{value: val, isSet: true}
}

func (v NullableUpdateFilterForDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterForDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
