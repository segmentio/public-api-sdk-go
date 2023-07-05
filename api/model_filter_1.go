/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Filter1 The requested Destination filter.
type Filter1 struct {
	// The unique id of this filter.
	Id string `json:"id"`
	// The id of the Source associated with this filter.
	SourceId string `json:"sourceId"`
	// The id of the Destination associated with this filter.
	DestinationId string `json:"destinationId"`
	// A condition that defines whether to apply this filter to a payload.
	If string `json:"if"`
	// A list of actions this filter performs.
	Actions []DestinationFilterActionV1 `json:"actions"`
	// A title for this filter.
	Title string `json:"title"`
	// A description for this filter.
	Description *string `json:"description,omitempty"`
	// When set to true, this filter is active.
	Enabled bool `json:"enabled"`
	// The timestamp of this filter's creation.
	CreatedAt string `json:"createdAt"`
	// The timestamp of this filter's last change.
	UpdatedAt string `json:"updatedAt"`
}

// NewFilter1 instantiates a new Filter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter1(
	id string,
	sourceId string,
	destinationId string,
	if_ string,
	actions []DestinationFilterActionV1,
	title string,
	enabled bool,
	createdAt string,
	updatedAt string,
) *Filter1 {
	this := Filter1{}
	this.Id = id
	this.SourceId = sourceId
	this.DestinationId = destinationId
	this.If = if_
	this.Actions = actions
	this.Title = title
	this.Enabled = enabled
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewFilter1WithDefaults instantiates a new Filter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilter1WithDefaults() *Filter1 {
	this := Filter1{}
	return &this
}

// GetId returns the Id field value
func (o *Filter1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Filter1) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value
func (o *Filter1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *Filter1) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationId returns the DestinationId field value
func (o *Filter1) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *Filter1) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetIf returns the If field value
func (o *Filter1) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *Filter1) SetIf(v string) {
	o.If = v
}

// GetActions returns the Actions field value
func (o *Filter1) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		var ret []DestinationFilterActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetActionsOk() ([]DestinationFilterActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *Filter1) SetActions(v []DestinationFilterActionV1) {
	o.Actions = v
}

// GetTitle returns the Title field value
func (o *Filter1) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Filter1) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Filter1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Filter1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Filter1) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Filter1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Filter1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Filter1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Filter1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Filter1) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Filter1) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Filter1) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o Filter1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["destinationId"] = o.DestinationId
	}
	if true {
		toSerialize["if"] = o.If
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableFilter1 struct {
	value *Filter1
	isSet bool
}

func (v NullableFilter1) Get() *Filter1 {
	return v.value
}

func (v *NullableFilter1) Set(val *Filter1) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter1) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter1(val *Filter1) *NullableFilter1 {
	return &NullableFilter1{value: val, isSet: true}
}

func (v NullableFilter1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
