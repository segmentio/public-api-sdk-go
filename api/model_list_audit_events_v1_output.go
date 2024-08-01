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

// checks if the ListAuditEventsV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListAuditEventsV1Output{}

// ListAuditEventsV1Output Returns a list of Audit Trail events for the current Workspace.
type ListAuditEventsV1Output struct {
	// Audit trail events for the current Workspace.
	Events     []AuditEventV1    `json:"events"`
	Pagination *PaginationOutput `json:"pagination,omitempty"`
}

// NewListAuditEventsV1Output instantiates a new ListAuditEventsV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAuditEventsV1Output(events []AuditEventV1) *ListAuditEventsV1Output {
	this := ListAuditEventsV1Output{}
	this.Events = events
	return &this
}

// NewListAuditEventsV1OutputWithDefaults instantiates a new ListAuditEventsV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAuditEventsV1OutputWithDefaults() *ListAuditEventsV1Output {
	this := ListAuditEventsV1Output{}
	return &this
}

// GetEvents returns the Events field value
func (o *ListAuditEventsV1Output) GetEvents() []AuditEventV1 {
	if o == nil {
		var ret []AuditEventV1
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ListAuditEventsV1Output) GetEventsOk() ([]AuditEventV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ListAuditEventsV1Output) SetEvents(v []AuditEventV1) {
	o.Events = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ListAuditEventsV1Output) GetPagination() PaginationOutput {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationOutput
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAuditEventsV1Output) GetPaginationOk() (*PaginationOutput, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ListAuditEventsV1Output) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationOutput and assigns it to the Pagination field.
func (o *ListAuditEventsV1Output) SetPagination(v PaginationOutput) {
	o.Pagination = &v
}

func (o ListAuditEventsV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAuditEventsV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableListAuditEventsV1Output struct {
	value *ListAuditEventsV1Output
	isSet bool
}

func (v NullableListAuditEventsV1Output) Get() *ListAuditEventsV1Output {
	return v.value
}

func (v *NullableListAuditEventsV1Output) Set(val *ListAuditEventsV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListAuditEventsV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListAuditEventsV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAuditEventsV1Output(
	val *ListAuditEventsV1Output,
) *NullableListAuditEventsV1Output {
	return &NullableListAuditEventsV1Output{value: val, isSet: true}
}

func (v NullableListAuditEventsV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAuditEventsV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
