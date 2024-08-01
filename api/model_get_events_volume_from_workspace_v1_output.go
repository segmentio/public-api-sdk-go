/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 51.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetEventsVolumeFromWorkspaceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEventsVolumeFromWorkspaceV1Output{}

// GetEventsVolumeFromWorkspaceV1Output GetEventsVolumeFromWorkspaceV1Output represents the results given the input query.
type GetEventsVolumeFromWorkspaceV1Output struct {
	// Observability event volume path.
	Path  string                              `json:"path"`
	Query GetEventsVolumeFromWorkspaceV1Query `json:"query"`
	// The resultant list of series broken down by the dimensions requested over the time frame requested and ordered by the total count of events in all series. Note: The limit of entries returned is 5000.
	Result     []SourceEventVolumeV1 `json:"result"`
	Pagination *PaginationOutput     `json:"pagination,omitempty"`
}

// NewGetEventsVolumeFromWorkspaceV1Output instantiates a new GetEventsVolumeFromWorkspaceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEventsVolumeFromWorkspaceV1Output(
	path string,
	query GetEventsVolumeFromWorkspaceV1Query,
	result []SourceEventVolumeV1,
) *GetEventsVolumeFromWorkspaceV1Output {
	this := GetEventsVolumeFromWorkspaceV1Output{}
	this.Path = path
	this.Query = query
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

// GetPath returns the Path field value
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *GetEventsVolumeFromWorkspaceV1Output) SetPath(v string) {
	o.Path = v
}

// GetQuery returns the Query field value
func (o *GetEventsVolumeFromWorkspaceV1Output) GetQuery() GetEventsVolumeFromWorkspaceV1Query {
	if o == nil {
		var ret GetEventsVolumeFromWorkspaceV1Query
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetQueryOk() (*GetEventsVolumeFromWorkspaceV1Query, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *GetEventsVolumeFromWorkspaceV1Output) SetQuery(v GetEventsVolumeFromWorkspaceV1Query) {
	o.Query = v
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
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPagination() PaginationOutput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationOutput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetEventsVolumeFromWorkspaceV1Output) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationOutput and assigns it to the Pagination field.
func (o *GetEventsVolumeFromWorkspaceV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = &v
}

func (o GetEventsVolumeFromWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEventsVolumeFromWorkspaceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["query"] = o.Query
	toSerialize["result"] = o.Result
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetEventsVolumeFromWorkspaceV1Output struct {
	value *GetEventsVolumeFromWorkspaceV1Output
	isSet bool
}

func (v NullableGetEventsVolumeFromWorkspaceV1Output) Get() *GetEventsVolumeFromWorkspaceV1Output {
	return v.value
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Output) Set(
	val *GetEventsVolumeFromWorkspaceV1Output,
) {
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

func NewNullableGetEventsVolumeFromWorkspaceV1Output(
	val *GetEventsVolumeFromWorkspaceV1Output,
) *NullableGetEventsVolumeFromWorkspaceV1Output {
	return &NullableGetEventsVolumeFromWorkspaceV1Output{value: val, isSet: true}
}

func (v NullableGetEventsVolumeFromWorkspaceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEventsVolumeFromWorkspaceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
