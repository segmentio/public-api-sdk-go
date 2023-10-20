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

// InsertFunctionInstance The created instance.
type InsertFunctionInstance struct {
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

// NewInsertFunctionInstance instantiates a new InsertFunctionInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertFunctionInstance(
	id string,
	integrationId string,
	classId string,
	enabled bool,
	createdAt string,
	updatedAt string,
	settings map[string]interface{},
	encryptedSettings map[string]interface{},
) *InsertFunctionInstance {
	this := InsertFunctionInstance{}
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

// NewInsertFunctionInstanceWithDefaults instantiates a new InsertFunctionInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertFunctionInstanceWithDefaults() *InsertFunctionInstance {
	this := InsertFunctionInstance{}
	return &this
}

// GetId returns the Id field value
func (o *InsertFunctionInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InsertFunctionInstance) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InsertFunctionInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InsertFunctionInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InsertFunctionInstance) SetName(v string) {
	o.Name = &v
}

// GetIntegrationId returns the IntegrationId field value
func (o *InsertFunctionInstance) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *InsertFunctionInstance) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetClassId returns the ClassId field value
func (o *InsertFunctionInstance) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *InsertFunctionInstance) SetClassId(v string) {
	o.ClassId = v
}

// GetEnabled returns the Enabled field value
func (o *InsertFunctionInstance) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InsertFunctionInstance) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InsertFunctionInstance) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InsertFunctionInstance) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InsertFunctionInstance) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InsertFunctionInstance) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetSettings returns the Settings field value
func (o *InsertFunctionInstance) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *InsertFunctionInstance) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetEncryptedSettings returns the EncryptedSettings field value
func (o *InsertFunctionInstance) GetEncryptedSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.EncryptedSettings
}

// GetEncryptedSettingsOk returns a tuple with the EncryptedSettings field value
// and a boolean to check if the value has been set.
func (o *InsertFunctionInstance) GetEncryptedSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptedSettings, true
}

// SetEncryptedSettings sets field value
func (o *InsertFunctionInstance) SetEncryptedSettings(v map[string]interface{}) {
	o.EncryptedSettings = v
}

func (o InsertFunctionInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["integrationId"] = o.IntegrationId
	}
	if true {
		toSerialize["classId"] = o.ClassId
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
	if true {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["encryptedSettings"] = o.EncryptedSettings
	}
	return json.Marshal(toSerialize)
}

type NullableInsertFunctionInstance struct {
	value *InsertFunctionInstance
	isSet bool
}

func (v NullableInsertFunctionInstance) Get() *InsertFunctionInstance {
	return v.value
}

func (v *NullableInsertFunctionInstance) Set(val *InsertFunctionInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertFunctionInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertFunctionInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertFunctionInstance(
	val *InsertFunctionInstance,
) *NullableInsertFunctionInstance {
	return &NullableInsertFunctionInstance{value: val, isSet: true}
}

func (v NullableInsertFunctionInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertFunctionInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
