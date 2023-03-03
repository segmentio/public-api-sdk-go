/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListTransformationsBetaOutput Lists the Transformations associated with the current Workspace.
type ListTransformationsBetaOutput struct {
	// A paginated list of Transformations.
	Transformations []TransformationBeta `json:"transformations"`
	Pagination      Pagination           `json:"pagination"`
}

// NewListTransformationsBetaOutput instantiates a new ListTransformationsBetaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransformationsBetaOutput(
	transformations []TransformationBeta,
	pagination Pagination,
) *ListTransformationsBetaOutput {
	this := ListTransformationsBetaOutput{}
	this.Transformations = transformations
	this.Pagination = pagination
	return &this
}

// NewListTransformationsBetaOutputWithDefaults instantiates a new ListTransformationsBetaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransformationsBetaOutputWithDefaults() *ListTransformationsBetaOutput {
	this := ListTransformationsBetaOutput{}
	return &this
}

// GetTransformations returns the Transformations field value
func (o *ListTransformationsBetaOutput) GetTransformations() []TransformationBeta {
	if o == nil {
		var ret []TransformationBeta
		return ret
	}

	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value
// and a boolean to check if the value has been set.
func (o *ListTransformationsBetaOutput) GetTransformationsOk() ([]TransformationBeta, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transformations, true
}

// SetTransformations sets field value
func (o *ListTransformationsBetaOutput) SetTransformations(v []TransformationBeta) {
	o.Transformations = v
}

// GetPagination returns the Pagination field value
func (o *ListTransformationsBetaOutput) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListTransformationsBetaOutput) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListTransformationsBetaOutput) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o ListTransformationsBetaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transformations"] = o.Transformations
	}
	if true {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableListTransformationsBetaOutput struct {
	value *ListTransformationsBetaOutput
	isSet bool
}

func (v NullableListTransformationsBetaOutput) Get() *ListTransformationsBetaOutput {
	return v.value
}

func (v *NullableListTransformationsBetaOutput) Set(val *ListTransformationsBetaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransformationsBetaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransformationsBetaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransformationsBetaOutput(
	val *ListTransformationsBetaOutput,
) *NullableListTransformationsBetaOutput {
	return &NullableListTransformationsBetaOutput{value: val, isSet: true}
}

func (v NullableListTransformationsBetaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransformationsBetaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
