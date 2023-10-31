/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateDestinationV1Output Creates a new Destination.
type CreateDestinationV1Output struct {
	Destination Destination2 `json:"destination"`
}

// NewCreateDestinationV1Output instantiates a new CreateDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationV1Output(destination Destination2) *CreateDestinationV1Output {
	this := CreateDestinationV1Output{}
	this.Destination = destination
	return &this
}

// NewCreateDestinationV1OutputWithDefaults instantiates a new CreateDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationV1OutputWithDefaults() *CreateDestinationV1Output {
	this := CreateDestinationV1Output{}
	return &this
}

// GetDestination returns the Destination field value
func (o *CreateDestinationV1Output) GetDestination() Destination2 {
	if o == nil {
		var ret Destination2
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Output) GetDestinationOk() (*Destination2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *CreateDestinationV1Output) SetDestination(v Destination2) {
	o.Destination = v
}

func (o CreateDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDestinationV1Output struct {
	value *CreateDestinationV1Output
	isSet bool
}

func (v NullableCreateDestinationV1Output) Get() *CreateDestinationV1Output {
	return v.value
}

func (v *NullableCreateDestinationV1Output) Set(val *CreateDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationV1Output(
	val *CreateDestinationV1Output,
) *NullableCreateDestinationV1Output {
	return &NullableCreateDestinationV1Output{value: val, isSet: true}
}

func (v NullableCreateDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
