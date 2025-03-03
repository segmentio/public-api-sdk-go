/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InsertFunctionInstanceAlpha type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsertFunctionInstanceAlpha{}

// InsertFunctionInstanceAlpha struct for InsertFunctionInstanceAlpha
type InsertFunctionInstanceAlpha struct {
	Id                string                 `json:"id"`
	Name              *string                `json:"name,omitempty"`
	IntegrationId     string                 `json:"integrationId"`
	ClassId           string                 `json:"classId"`
	Enabled           bool                   `json:"enabled"`
	CreatedAt         string                 `json:"createdAt"`
	UpdatedAt         string                 `json:"updatedAt"`
	Settings          map[string]interface{} `json:"settings"`
	EncryptedSettings map[string]interface{} `json:"encryptedSettings"`
}

// NewInsertFunctionInstanceAlpha instantiates a new InsertFunctionInstanceAlpha object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertFunctionInstanceAlpha(
	id string,
	integrationId string,
	classId string,
	enabled bool,
	createdAt string,
	updatedAt string,
	settings map[string]interface{},
	encryptedSettings map[string]interface{},
) *InsertFunctionInstanceAlpha {
	this := InsertFunctionInstanceAlpha{}
	this.Id = id
	this.IntegrationId = integrationId
	this.ClassId = classId
	this.Enabled = enabled
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Settings = settings
	this.EncryptedSettings = encryptedSettings
	return &this
}

// NewInsertFunctionInstanceAlphaWithDefaults instantiates a new InsertFunctionInstanceAlpha object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertFunctionInstanceAlphaWithDefaults() *InsertFunctionInstanceAlpha {
	this := InsertFunctionInstanceAlpha{}
	return &this
}

// GetId returns the Id field value
func (o *InsertFunctionInstanceAlpha) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InsertFunctionInstanceAlpha) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InsertFunctionInstanceAlpha) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InsertFunctionInstanceAlpha) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InsertFunctionInstanceAlpha) SetName(v string) {
	o.Name = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *InsertFunctionInstanceAlpha) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *InsertFunctionInstanceAlpha) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetClassId returns the ClassId field value
func (o *InsertFunctionInstanceAlpha) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InsertFunctionInstanceAlpha) SetClassId(v string) {
	o.ClassId = v
}

// GetEnabled returns the Enabled field value
func (o *InsertFunctionInstanceAlpha) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InsertFunctionInstanceAlpha) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InsertFunctionInstanceAlpha) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InsertFunctionInstanceAlpha) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InsertFunctionInstanceAlpha) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InsertFunctionInstanceAlpha) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetSettings returns the Settings field value
func (o *InsertFunctionInstanceAlpha) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *InsertFunctionInstanceAlpha) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetEncryptedSettings returns the EncryptedSettings field value
func (o *InsertFunctionInstanceAlpha) GetEncryptedSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.EncryptedSettings
}

// GetEncryptedSettingsOk returns a tuple with the EncryptedSettings field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstanceAlpha) GetEncryptedSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.EncryptedSettings, true
}

// SetEncryptedSettings sets field value
func (o *InsertFunctionInstanceAlpha) SetEncryptedSettings(v map[string]interface{}) {
	o.EncryptedSettings = v
}

func (o InsertFunctionInstanceAlpha) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsertFunctionInstanceAlpha) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["integrationId"] = o.IntegrationId
	toSerialize["classId"] = o.ClassId
	toSerialize["enabled"] = o.Enabled
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["settings"] = o.Settings
	toSerialize["encryptedSettings"] = o.EncryptedSettings
	return toSerialize, nil
}

type NullableInsertFunctionInstanceAlpha struct {
	value *InsertFunctionInstanceAlpha
	isSet bool
}

func (v NullableInsertFunctionInstanceAlpha) Get() *InsertFunctionInstanceAlpha {
	return v.value
}

func (v *NullableInsertFunctionInstanceAlpha) Set(val *InsertFunctionInstanceAlpha) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertFunctionInstanceAlpha) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertFunctionInstanceAlpha) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertFunctionInstanceAlpha(
	val *InsertFunctionInstanceAlpha,
) *NullableInsertFunctionInstanceAlpha {
	return &NullableInsertFunctionInstanceAlpha{value: val, isSet: true}
}

func (v NullableInsertFunctionInstanceAlpha) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertFunctionInstanceAlpha) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
