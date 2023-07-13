/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TrackingPlan The requested Tracking Plan.
type TrackingPlan struct {
	// The Tracking Plan's identifier.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The Tracking Plan's name.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// URL-friendly slug of this Tracking Plan.  Config API note: equal to `name`.
	Slug *string `json:"slug,omitempty"`
	// The Tracking Plan's description.
	Description *string `json:"description,omitempty"`
	// The Tracking Plan's type.
	Type string `json:"type"`
	// The timestamp of the last change to the Tracking Plan.  Config API note: equal to `updateTime`.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// The timestamp of this Tracking Plan's creation.  Config API note: equal to `createTime`.
	CreatedAt *string `json:"createdAt,omitempty"`
}

// NewTrackingPlan instantiates a new TrackingPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingPlan(id string, type_ string) *TrackingPlan {
	this := TrackingPlan{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewTrackingPlanWithDefaults instantiates a new TrackingPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingPlanWithDefaults() *TrackingPlan {
	this := TrackingPlan{}
	return &this
}

// GetId returns the Id field value
func (o *TrackingPlan) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TrackingPlan) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrackingPlan) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrackingPlan) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrackingPlan) SetName(v string) {
	o.Name = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *TrackingPlan) GetSlug() string {
	if o == nil || o.Slug == nil {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetSlugOk() (*string, bool) {
	if o == nil || o.Slug == nil {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *TrackingPlan) HasSlug() bool {
	if o != nil && o.Slug != nil {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *TrackingPlan) SetSlug(v string) {
	o.Slug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TrackingPlan) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TrackingPlan) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TrackingPlan) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *TrackingPlan) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TrackingPlan) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TrackingPlan) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TrackingPlan) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TrackingPlan) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrackingPlan) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingPlan) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrackingPlan) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TrackingPlan) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o TrackingPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Slug != nil {
		toSerialize["slug"] = o.Slug
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTrackingPlan struct {
	value *TrackingPlan
	isSet bool
}

func (v NullableTrackingPlan) Get() *TrackingPlan {
	return v.value
}

func (v *NullableTrackingPlan) Set(val *TrackingPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingPlan(val *TrackingPlan) *NullableTrackingPlan {
	return &NullableTrackingPlan{value: val, isSet: true}
}

func (v NullableTrackingPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
