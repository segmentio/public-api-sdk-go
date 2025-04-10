/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput{}

// ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput The reverse ETL sync statuses that were looked up.
type ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput struct {
	// The reverse ETL sync statuses that were looked up of the subscription id.
	SyncStatuses []ReverseETLSyncStatus `json:"syncStatuses"`
	Pagination   *PaginationOutput      `json:"pagination,omitempty"`
}

// NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput(
	syncStatuses []ReverseETLSyncStatus,
) *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput {
	this := ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput{}
	this.SyncStatuses = syncStatuses
	return &this
}

// NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutputWithDefaults instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutputWithDefaults() *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput {
	this := ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput{}
	return &this
}

// GetSyncStatuses returns the SyncStatuses field value
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetSyncStatuses() []ReverseETLSyncStatus {
	if o == nil {
		var ret []ReverseETLSyncStatus
		return ret
	}

	return o.SyncStatuses
}

// GetSyncStatusesOk returns a tuple with the SyncStatuses field value
// and a boolean to check if the value has been set.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetSyncStatusesOk() ([]ReverseETLSyncStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyncStatuses, true
}

// SetSyncStatuses sets field value
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) SetSyncStatuses(
	v []ReverseETLSyncStatus,
) {
	o.SyncStatuses = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetPagination() PaginationOutput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationOutput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationOutput and assigns it to the Pagination field.
func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) SetPagination(
	v PaginationOutput,
) {
	o.Pagination = &v
}

func (o ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["syncStatuses"] = o.SyncStatuses
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput struct {
	value *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput
	isSet bool
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) Get() *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput {
	return v.value
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) Set(
	val *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput(
	val *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput,
) *NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput {
	return &NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput{
		value: val,
		isSet: true,
	}
}

func (v NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
