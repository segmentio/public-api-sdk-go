/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetFilterInDestinationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFilterInDestinationV1Output{}

// GetFilterInDestinationV1Output Output for GetDestinationFiltersV1.
type GetFilterInDestinationV1Output struct {
	Filter DestinationFilterV1 `json:"filter"`
}

// NewGetFilterInDestinationV1Output instantiates a new GetFilterInDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFilterInDestinationV1Output(filter DestinationFilterV1) *GetFilterInDestinationV1Output {
	this := GetFilterInDestinationV1Output{}
	this.Filter = filter
	return &this
}

// NewGetFilterInDestinationV1OutputWithDefaults instantiates a new GetFilterInDestinationV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFilterInDestinationV1OutputWithDefaults() *GetFilterInDestinationV1Output {
	this := GetFilterInDestinationV1Output{}
	return &this
}

// GetFilter returns the Filter field value
func (o *GetFilterInDestinationV1Output) GetFilter() DestinationFilterV1 {
	if o == nil {
		var ret DestinationFilterV1
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *GetFilterInDestinationV1Output) GetFilterOk() (*DestinationFilterV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *GetFilterInDestinationV1Output) SetFilter(v DestinationFilterV1) {
	o.Filter = v
}

func (o GetFilterInDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFilterInDestinationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	return toSerialize, nil
}

type NullableGetFilterInDestinationV1Output struct {
	value *GetFilterInDestinationV1Output
	isSet bool
}

func (v NullableGetFilterInDestinationV1Output) Get() *GetFilterInDestinationV1Output {
	return v.value
}

func (v *NullableGetFilterInDestinationV1Output) Set(val *GetFilterInDestinationV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFilterInDestinationV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFilterInDestinationV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFilterInDestinationV1Output(
	val *GetFilterInDestinationV1Output,
) *NullableGetFilterInDestinationV1Output {
	return &NullableGetFilterInDestinationV1Output{value: val, isSet: true}
}

func (v NullableGetFilterInDestinationV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFilterInDestinationV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
