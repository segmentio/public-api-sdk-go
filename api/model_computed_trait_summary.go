/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 43.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ComputedTraitSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComputedTraitSummary{}

// ComputedTraitSummary Defines a Computed trait.
type ComputedTraitSummary struct {
	// Computed trait id.
	Id string `json:"id"`
	// Space id for the computed trait.
	SpaceId string `json:"spaceId"`
	// Name of the computed trait.
	Name string `json:"name"`
	// Description of the computed trait.
	Description string `json:"description"`
	// Key for the computed trait.
	Key string `json:"key"`
	// Enabled/disabled status for the computed trait.
	Enabled    bool               `json:"enabled"`
	Definition NullableDefinition `json:"definition"`
	// Status for the computed trait.  Possible values: Backfilling, Computing, Failed, Live, Awaiting Destinations, Disabled.
	Status *string `json:"status,omitempty"`
	// User id who created the computed trait.
	CreatedBy string `json:"createdBy"`
	// User id who last updated the computed trait.
	UpdatedBy string `json:"updatedBy"`
	// The timestamp of the computed trait's creation.
	CreatedAt string `json:"createdAt"`
	// The timestamp of the computed trait's last change.
	UpdatedAt string `json:"updatedAt"`
}

// NewComputedTraitSummary instantiates a new ComputedTraitSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputedTraitSummary(
	id string,
	spaceId string,
	name string,
	description string,
	key string,
	enabled bool,
	definition NullableDefinition,
	createdBy string,
	updatedBy string,
	createdAt string,
	updatedAt string,
) *ComputedTraitSummary {
	this := ComputedTraitSummary{}
	this.Id = id
	this.SpaceId = spaceId
	this.Name = name
	this.Description = description
	this.Key = key
	this.Enabled = enabled
	this.Definition = definition
	this.CreatedBy = createdBy
	this.UpdatedBy = updatedBy
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewComputedTraitSummaryWithDefaults instantiates a new ComputedTraitSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputedTraitSummaryWithDefaults() *ComputedTraitSummary {
	this := ComputedTraitSummary{}
	return &this
}

// GetId returns the Id field value
func (o *ComputedTraitSummary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ComputedTraitSummary) SetId(v string) {
	o.Id = v
}

// GetSpaceId returns the SpaceId field value
func (o *ComputedTraitSummary) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *ComputedTraitSummary) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetName returns the Name field value
func (o *ComputedTraitSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ComputedTraitSummary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ComputedTraitSummary) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ComputedTraitSummary) SetDescription(v string) {
	o.Description = v
}

// GetKey returns the Key field value
func (o *ComputedTraitSummary) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ComputedTraitSummary) SetKey(v string) {
	o.Key = v
}

// GetEnabled returns the Enabled field value
func (o *ComputedTraitSummary) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ComputedTraitSummary) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDefinition returns the Definition field value
// If the value is explicit nil, the zero value for Definition will be returned
func (o *ComputedTraitSummary) GetDefinition() Definition {
	if o == nil || o.Definition.Get() == nil {
		var ret Definition
		return ret
	}

	return *o.Definition.Get()
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputedTraitSummary) GetDefinitionOk() (*Definition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Definition.Get(), o.Definition.IsSet()
}

// SetDefinition sets field value
func (o *ComputedTraitSummary) SetDefinition(v Definition) {
	o.Definition.Set(&v)
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComputedTraitSummary) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComputedTraitSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ComputedTraitSummary) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ComputedTraitSummary) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ComputedTraitSummary) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *ComputedTraitSummary) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *ComputedTraitSummary) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ComputedTraitSummary) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ComputedTraitSummary) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ComputedTraitSummary) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ComputedTraitSummary) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ComputedTraitSummary) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o ComputedTraitSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComputedTraitSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["spaceId"] = o.SpaceId
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["key"] = o.Key
	toSerialize["enabled"] = o.Enabled
	toSerialize["definition"] = o.Definition.Get()
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["updatedBy"] = o.UpdatedBy
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableComputedTraitSummary struct {
	value *ComputedTraitSummary
	isSet bool
}

func (v NullableComputedTraitSummary) Get() *ComputedTraitSummary {
	return v.value
}

func (v *NullableComputedTraitSummary) Set(val *ComputedTraitSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableComputedTraitSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableComputedTraitSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputedTraitSummary(val *ComputedTraitSummary) *NullableComputedTraitSummary {
	return &NullableComputedTraitSummary{value: val, isSet: true}
}

func (v NullableComputedTraitSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputedTraitSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
