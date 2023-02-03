/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationFilterV1 Represents a Destination filter.
type DestinationFilterV1 struct {
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

// NewDestinationFilterV1 instantiates a new DestinationFilterV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationFilterV1(
	id string,
	sourceId string,
	destinationId string,
	if_ string,
	actions []DestinationFilterActionV1,
	title string,
	enabled bool,
	createdAt string,
	updatedAt string,
) *DestinationFilterV1 {
	this := DestinationFilterV1{}
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

// NewDestinationFilterV1WithDefaults instantiates a new DestinationFilterV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationFilterV1WithDefaults() *DestinationFilterV1 {
	this := DestinationFilterV1{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationFilterV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationFilterV1) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value
func (o *DestinationFilterV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *DestinationFilterV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationId returns the DestinationId field value
func (o *DestinationFilterV1) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *DestinationFilterV1) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetIf returns the If field value
func (o *DestinationFilterV1) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *DestinationFilterV1) SetIf(v string) {
	o.If = v
}

// GetActions returns the Actions field value
func (o *DestinationFilterV1) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		var ret []DestinationFilterActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetActionsOk() ([]DestinationFilterActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *DestinationFilterV1) SetActions(v []DestinationFilterActionV1) {
	o.Actions = v
}

// GetTitle returns the Title field value
func (o *DestinationFilterV1) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DestinationFilterV1) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DestinationFilterV1) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DestinationFilterV1) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DestinationFilterV1) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DestinationFilterV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DestinationFilterV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *DestinationFilterV1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DestinationFilterV1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DestinationFilterV1) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DestinationFilterV1) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DestinationFilterV1) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o DestinationFilterV1) MarshalJSON() ([]byte, error) {
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

type NullableDestinationFilterV1 struct {
	value *DestinationFilterV1
	isSet bool
}

func (v NullableDestinationFilterV1) Get() *DestinationFilterV1 {
	return v.value
}

func (v *NullableDestinationFilterV1) Set(val *DestinationFilterV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationFilterV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationFilterV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationFilterV1(val *DestinationFilterV1) *NullableDestinationFilterV1 {
	return &NullableDestinationFilterV1{value: val, isSet: true}
}

func (v NullableDestinationFilterV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationFilterV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
