/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListTransformationsV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTransformationsV1Output{}

// ListTransformationsV1Output Lists the Transformations associated with the current Workspace.
type ListTransformationsV1Output struct {
	// A paginated list of Transformations.
	Transformations []TransformationV1 `json:"transformations"`
	Pagination      PaginationOutput   `json:"pagination"`
}

// NewListTransformationsV1Output instantiates a new ListTransformationsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransformationsV1Output(
	transformations []TransformationV1,
	pagination PaginationOutput,
) *ListTransformationsV1Output {
	this := ListTransformationsV1Output{}
	this.Transformations = transformations
	this.Pagination = pagination
	return &this
}

// NewListTransformationsV1OutputWithDefaults instantiates a new ListTransformationsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransformationsV1OutputWithDefaults() *ListTransformationsV1Output {
	this := ListTransformationsV1Output{}
	return &this
}

// GetTransformations returns the Transformations field value
func (o *ListTransformationsV1Output) GetTransformations() []TransformationV1 {
	if o == nil {
		var ret []TransformationV1
		return ret
	}

	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value
// and a boolean to check if the value has been set.
func (o *ListTransformationsV1Output) GetTransformationsOk() ([]TransformationV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transformations, true
}

// SetTransformations sets field value
func (o *ListTransformationsV1Output) SetTransformations(v []TransformationV1) {
	o.Transformations = v
}

// GetPagination returns the Pagination field value
func (o *ListTransformationsV1Output) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListTransformationsV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListTransformationsV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListTransformationsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTransformationsV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformations"] = o.Transformations
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListTransformationsV1Output struct {
	value *ListTransformationsV1Output
	isSet bool
}

func (v NullableListTransformationsV1Output) Get() *ListTransformationsV1Output {
	return v.value
}

func (v *NullableListTransformationsV1Output) Set(val *ListTransformationsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransformationsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransformationsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransformationsV1Output(
	val *ListTransformationsV1Output,
) *NullableListTransformationsV1Output {
	return &NullableListTransformationsV1Output{value: val, isSet: true}
}

func (v NullableListTransformationsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransformationsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
