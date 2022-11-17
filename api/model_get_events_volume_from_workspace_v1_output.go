/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.6
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetEventsVolumeFromWorkspaceV1Output GetEventsVolumeFromWorkspaceV1Output represents the results given the input query.
type GetEventsVolumeFromWorkspaceV1Output struct {
	// The resultant list of series broken down by the dimensions requested over the time frame requested and ordered by the total count of events in all series. Note: The limit of entries returned is 5000.
	Result []SourceEventVolumeV1 `json:"result"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewGetEventsVolumeFromWorkspaceV1Output instantiates a new GetEventsVolumeFromWorkspaceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventsVolumeFromWorkspaceV1Output(result []SourceEventVolumeV1) *GetEventsVolumeFromWorkspaceV1Output {
	this := GetEventsVolumeFromWorkspaceV1Output{}
	this.Result = result
	return &this
}

// NewGetEventsVolumeFromWorkspaceV1OutputWithDefaults instantiates a new GetEventsVolumeFromWorkspaceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEventsVolumeFromWorkspaceV1OutputWithDefaults() *GetEventsVolumeFromWorkspaceV1Output {
	this := GetEventsVolumeFromWorkspaceV1Output{}
	return &this
}

// GetResult returns the Result field value
func (o *GetEventsVolumeFromWorkspaceV1Output) GetResult() []SourceEventVolumeV1 {
	if o == nil {
		var ret []SourceEventVolumeV1
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetResultOk() ([]SourceEventVolumeV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *GetEventsVolumeFromWorkspaceV1Output) SetResult(v []SourceEventVolumeV1) {
	o.Result = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPagination() Pagination {
	if o == nil || o.Pagination == nil {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPaginationOk() (*Pagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) HasPagination() bool {
	if o != nil && o.Pagination != nil {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetEventsVolumeFromWorkspaceV1Output) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o GetEventsVolumeFromWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableGetEventsVolumeFromWorkspaceV1Output struct {
	value *GetEventsVolumeFromWorkspaceV1Output
	isSet bool
}

func (v NullableGetEventsVolumeFromWorkspaceV1Output) Get() *GetEventsVolumeFromWorkspaceV1Output {
	return v.value
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Output) Set(val *GetEventsVolumeFromWorkspaceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEventsVolumeFromWorkspaceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEventsVolumeFromWorkspaceV1Output(val *GetEventsVolumeFromWorkspaceV1Output) *NullableGetEventsVolumeFromWorkspaceV1Output {
	return &NullableGetEventsVolumeFromWorkspaceV1Output{value: val, isSet: true}
}

func (v NullableGetEventsVolumeFromWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


