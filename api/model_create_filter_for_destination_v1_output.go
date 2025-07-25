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

// checks if the CreateFilterForDestinationV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFilterForDestinationV1Output{}

// CreateFilterForDestinationV1Output Output for CreateDestinationFiltersV1.
type CreateFilterForDestinationV1Output struct {
	Filter DestinationFilterV1 `json:"filter"`
}

// NewCreateFilterForDestinationV1Output instantiates a new CreateFilterForDestinationV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterForDestinationV1Output(
	filter DestinationFilterV1,
) *CreateFilterForDestinationV1Output {
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
func (o *CreateFilterForDestinationV1Output) GetFilter() DestinationFilterV1 {
	if o == nil {
		var ret DestinationFilterV1
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Output) GetFilterOk() (*DestinationFilterV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *CreateFilterForDestinationV1Output) SetFilter(v DestinationFilterV1) {
	o.Filter = v
}

func (o CreateFilterForDestinationV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFilterForDestinationV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filter"] = o.Filter
	return toSerialize, nil
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
