/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetDestinationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDestinationV1Output{}

// GetDestinationV1Output Returns a single Destination by its id.
type GetDestinationV1Output struct {
	Destination DestinationV1 `json:"destination"`
}

// NewGetDestinationV1Output instantiates a new GetDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestinationV1Output(destination DestinationV1) *GetDestinationV1Output {
	this := GetDestinationV1Output{}
	this.Destination = destination
	return &this
}

// NewGetDestinationV1OutputWithDefaults instantiates a new GetDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestinationV1OutputWithDefaults() *GetDestinationV1Output {
	this := GetDestinationV1Output{}
	return &this
}

// GetDestination returns the Destination field value
func (o *GetDestinationV1Output) GetDestination() DestinationV1 {
	if o == nil {
		var ret DestinationV1
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *GetDestinationV1Output) GetDestinationOk() (*DestinationV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *GetDestinationV1Output) SetDestination(v DestinationV1) {
	o.Destination = v
}

func (o GetDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDestinationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

type NullableGetDestinationV1Output struct {
	value *GetDestinationV1Output
	isSet bool
}

func (v NullableGetDestinationV1Output) Get() *GetDestinationV1Output {
	return v.value
}

func (v *NullableGetDestinationV1Output) Set(val *GetDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestinationV1Output(
	val *GetDestinationV1Output,
) *NullableGetDestinationV1Output {
	return &NullableGetDestinationV1Output{value: val, isSet: true}
}

func (v NullableGetDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
