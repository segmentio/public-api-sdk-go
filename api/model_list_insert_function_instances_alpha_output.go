/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 47.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListInsertFunctionInstancesAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInsertFunctionInstancesAlphaOutput{}

// ListInsertFunctionInstancesAlphaOutput Returns a list of insert Function instances connected to the insert Function.
type ListInsertFunctionInstancesAlphaOutput struct {
	// All insert Function instances found.
	Instances  []InsertFunctionInstanceAlpha `json:"instances"`
	Pagination PaginationOutput              `json:"pagination"`
}

// NewListInsertFunctionInstancesAlphaOutput instantiates a new ListInsertFunctionInstancesAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInsertFunctionInstancesAlphaOutput(
	instances []InsertFunctionInstanceAlpha,
	pagination PaginationOutput,
) *ListInsertFunctionInstancesAlphaOutput {
	this := ListInsertFunctionInstancesAlphaOutput{}
	this.Instances = instances
	this.Pagination = pagination
	return &this
}

// NewListInsertFunctionInstancesAlphaOutputWithDefaults instantiates a new ListInsertFunctionInstancesAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInsertFunctionInstancesAlphaOutputWithDefaults() *ListInsertFunctionInstancesAlphaOutput {
	this := ListInsertFunctionInstancesAlphaOutput{}
	return &this
}

// GetInstances returns the Instances field value
func (o *ListInsertFunctionInstancesAlphaOutput) GetInstances() []InsertFunctionInstanceAlpha {
	if o == nil {
		var ret []InsertFunctionInstanceAlpha
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
func (o *ListInsertFunctionInstancesAlphaOutput) GetInstancesOk() ([]InsertFunctionInstanceAlpha, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *ListInsertFunctionInstancesAlphaOutput) SetInstances(v []InsertFunctionInstanceAlpha) {
	o.Instances = v
}

// GetPagination returns the Pagination field value
func (o *ListInsertFunctionInstancesAlphaOutput) GetPagination() PaginationOutput {
	if o == nil {
		var ret PaginationOutput
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ListInsertFunctionInstancesAlphaOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ListInsertFunctionInstancesAlphaOutput) SetPagination(v PaginationOutput) {
	o.Pagination = v
}

func (o ListInsertFunctionInstancesAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInsertFunctionInstancesAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instances"] = o.Instances
	toSerialize["pagination"] = o.Pagination
	return toSerialize, nil
}

type NullableListInsertFunctionInstancesAlphaOutput struct {
	value *ListInsertFunctionInstancesAlphaOutput
	isSet bool
}

func (v NullableListInsertFunctionInstancesAlphaOutput) Get() *ListInsertFunctionInstancesAlphaOutput {
	return v.value
}

func (v *NullableListInsertFunctionInstancesAlphaOutput) Set(
	val *ListInsertFunctionInstancesAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListInsertFunctionInstancesAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListInsertFunctionInstancesAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInsertFunctionInstancesAlphaOutput(
	val *ListInsertFunctionInstancesAlphaOutput,
) *NullableListInsertFunctionInstancesAlphaOutput {
	return &NullableListInsertFunctionInstancesAlphaOutput{value: val, isSet: true}
}

func (v NullableListInsertFunctionInstancesAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInsertFunctionInstancesAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
