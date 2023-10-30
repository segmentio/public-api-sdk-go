/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateFilterForDestinationV1Output Output for CreateDestinationFiltersV1.
type CreateFilterForDestinationV1Output struct {
	Filter Filter2 `json:"filter"`
}

// NewCreateFilterForDestinationV1Output instantiates a new CreateFilterForDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterForDestinationV1Output(filter Filter2) *CreateFilterForDestinationV1Output {
	this := CreateFilterForDestinationV1Output{}
	this.Filter = filter
	return &this
}

// NewCreateFilterForDestinationV1OutputWithDefaults instantiates a new CreateFilterForDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFilterForDestinationV1OutputWithDefaults() *CreateFilterForDestinationV1Output {
	this := CreateFilterForDestinationV1Output{}
	return &this
}

// GetFilter returns the Filter field value
func (o *CreateFilterForDestinationV1Output) GetFilter() Filter2 {
	if o == nil {
		var ret Filter2
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Output) GetFilterOk() (*Filter2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *CreateFilterForDestinationV1Output) SetFilter(v Filter2) {
	o.Filter = v
}

func (o CreateFilterForDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFilterForDestinationV1Output struct {
	value *CreateFilterForDestinationV1Output
	isSet bool
}

func (v NullableCreateFilterForDestinationV1Output) Get() *CreateFilterForDestinationV1Output {
	return v.value
}

func (v *NullableCreateFilterForDestinationV1Output) Set(val *CreateFilterForDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterForDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterForDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterForDestinationV1Output(
	val *CreateFilterForDestinationV1Output,
) *NullableCreateFilterForDestinationV1Output {
	return &NullableCreateFilterForDestinationV1Output{value: val, isSet: true}
}

func (v NullableCreateFilterForDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterForDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
