/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Audience The audience summary output.
type Audience struct {
	// Audience id.
	Id string `json:"id"`
	// Space id for the audience.
	SpaceId string `json:"spaceId"`
	// Name of the audience.
	Name string `json:"name"`
	// Description of the audience.
	Description string `json:"description"`
	// Key for the audience.
	Key string `json:"key"`
	// Enabled/disabled status for the audience.
	Enabled bool `json:"enabled"`
	// User id who created the audience.
	CreatedBy string `json:"createdBy"`
	// User id who last updated the audience.
	UpdatedBy string `json:"updatedBy"`
	// Date the audience was created.
	CreatedAt string `json:"createdAt"`
	// Date the audience was last updated.
	UpdatedAt string `json:"updatedAt"`
}

// NewAudience instantiates a new Audience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudience(
	id string,
	spaceId string,
	name string,
	description string,
	key string,
	enabled bool,
	createdBy string,
	updatedBy string,
	createdAt string,
	updatedAt string,
) *Audience {
	this := Audience{}
	this.Id = id
	this.SpaceId = spaceId
	this.Name = name
	this.Description = description
	this.Key = key
	this.Enabled = enabled
	this.CreatedBy = createdBy
	this.UpdatedBy = updatedBy
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAudienceWithDefaults instantiates a new Audience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceWithDefaults() *Audience {
	this := Audience{}
	return &this
}

// GetId returns the Id field value
func (o *Audience) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Audience) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Audience) SetId(v string) {
	o.Id = v
}

// GetSpaceId returns the SpaceId field value
func (o *Audience) GetSpaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value
// and a boolean to check if the value has been set.
func (o *Audience) GetSpaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpaceId, true
}

// SetSpaceId sets field value
func (o *Audience) SetSpaceId(v string) {
	o.SpaceId = v
}

// GetName returns the Name field value
func (o *Audience) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Audience) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Audience) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Audience) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Audience) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Audience) SetDescription(v string) {
	o.Description = v
}

// GetKey returns the Key field value
func (o *Audience) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Audience) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Audience) SetKey(v string) {
	o.Key = v
}

// GetEnabled returns the Enabled field value
func (o *Audience) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Audience) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Audience) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Audience) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Audience) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Audience) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedBy returns the UpdatedBy field value
func (o *Audience) GetUpdatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *Audience) GetUpdatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value
func (o *Audience) SetUpdatedBy(v string) {
	o.UpdatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Audience) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Audience) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Audience) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Audience) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Audience) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Audience) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

func (o Audience) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["spaceId"] = o.SpaceId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableAudience struct {
	value *Audience
	isSet bool
}

func (v NullableAudience) Get() *Audience {
	return v.value
}

func (v *NullableAudience) Set(val *Audience) {
	v.value = val
	v.isSet = true
}

func (v NullableAudience) IsSet() bool {
	return v.isSet
}

func (v *NullableAudience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudience(val *Audience) *NullableAudience {
	return &NullableAudience{value: val, isSet: true}
}

func (v NullableAudience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
