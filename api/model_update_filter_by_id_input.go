/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateFilterByIdInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilterByIdInput{}

// UpdateFilterByIdInput Input for UpdateFilterById.
type UpdateFilterByIdInput struct {
	// The Integration id of the resource.
	IntegrationId string `json:"integrationId"`
	// Whether the filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the filter.
	Name *string `json:"name,omitempty"`
	// The description of the filter.
	Description *string `json:"description,omitempty"`
	// The \"if\" statement for a filter.
	If *string `json:"if,omitempty"`
	// Whether the event is dropped.
	Drop *bool `json:"drop,omitempty"`
	// Describes the properties to be dropped on events that match the \"if\" statement.
	DropProperties []string `json:"dropProperties,omitempty"`
	// Describes the properties allowed on events that match the \"if\" statement.
	AllowProperties []string `json:"allowProperties,omitempty"`
}

// NewUpdateFilterByIdInput instantiates a new UpdateFilterByIdInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterByIdInput(integrationId string) *UpdateFilterByIdInput {
	this := UpdateFilterByIdInput{}
	this.IntegrationId = integrationId
	return &this
}

// NewUpdateFilterByIdInputWithDefaults instantiates a new UpdateFilterByIdInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterByIdInputWithDefaults() *UpdateFilterByIdInput {
	this := UpdateFilterByIdInput{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *UpdateFilterByIdInput) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *UpdateFilterByIdInput) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateFilterByIdInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFilterByIdInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateFilterByIdInput) SetDescription(v string) {
	o.Description = &v
}

// GetIf returns the If field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetIf() string {
	if o == nil || IsNil(o.If) {
		var ret string
		return ret
	}
	return *o.If
}

// GetIfOk returns a tuple with the If field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetIfOk() (*string, bool) {
	if o == nil || IsNil(o.If) {
		return nil, false
	}
	return o.If, true
}

// HasIf returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasIf() bool {
	if o != nil && !IsNil(o.If) {
		return true
	}

	return false
}

// SetIf gets a reference to the given string and assigns it to the If field.
func (o *UpdateFilterByIdInput) SetIf(v string) {
	o.If = &v
}

// GetDrop returns the Drop field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetDrop() bool {
	if o == nil || IsNil(o.Drop) {
		var ret bool
		return ret
	}
	return *o.Drop
}

// GetDropOk returns a tuple with the Drop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetDropOk() (*bool, bool) {
	if o == nil || IsNil(o.Drop) {
		return nil, false
	}
	return o.Drop, true
}

// HasDrop returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasDrop() bool {
	if o != nil && !IsNil(o.Drop) {
		return true
	}

	return false
}

// SetDrop gets a reference to the given bool and assigns it to the Drop field.
func (o *UpdateFilterByIdInput) SetDrop(v bool) {
	o.Drop = &v
}

// GetDropProperties returns the DropProperties field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetDropProperties() []string {
	if o == nil || IsNil(o.DropProperties) {
		var ret []string
		return ret
	}
	return o.DropProperties
}

// GetDropPropertiesOk returns a tuple with the DropProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetDropPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.DropProperties) {
		return nil, false
	}
	return o.DropProperties, true
}

// HasDropProperties returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasDropProperties() bool {
	if o != nil && !IsNil(o.DropProperties) {
		return true
	}

	return false
}

// SetDropProperties gets a reference to the given []string and assigns it to the DropProperties field.
func (o *UpdateFilterByIdInput) SetDropProperties(v []string) {
	o.DropProperties = v
}

// GetAllowProperties returns the AllowProperties field value if set, zero value otherwise.
func (o *UpdateFilterByIdInput) GetAllowProperties() []string {
	if o == nil || IsNil(o.AllowProperties) {
		var ret []string
		return ret
	}
	return o.AllowProperties
}

// GetAllowPropertiesOk returns a tuple with the AllowProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterByIdInput) GetAllowPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowProperties) {
		return nil, false
	}
	return o.AllowProperties, true
}

// HasAllowProperties returns a boolean if a field has been set.
func (o *UpdateFilterByIdInput) HasAllowProperties() bool {
	if o != nil && !IsNil(o.AllowProperties) {
		return true
	}

	return false
}

// SetAllowProperties gets a reference to the given []string and assigns it to the AllowProperties field.
func (o *UpdateFilterByIdInput) SetAllowProperties(v []string) {
	o.AllowProperties = v
}

func (o UpdateFilterByIdInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilterByIdInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.If) {
		toSerialize["if"] = o.If
	}
	if !IsNil(o.Drop) {
		toSerialize["drop"] = o.Drop
	}
	if !IsNil(o.DropProperties) {
		toSerialize["dropProperties"] = o.DropProperties
	}
	if !IsNil(o.AllowProperties) {
		toSerialize["allowProperties"] = o.AllowProperties
	}
	return toSerialize, nil
}

type NullableUpdateFilterByIdInput struct {
	value *UpdateFilterByIdInput
	isSet bool
}

func (v NullableUpdateFilterByIdInput) Get() *UpdateFilterByIdInput {
	return v.value
}

func (v *NullableUpdateFilterByIdInput) Set(val *UpdateFilterByIdInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterByIdInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterByIdInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterByIdInput(val *UpdateFilterByIdInput) *NullableUpdateFilterByIdInput {
	return &NullableUpdateFilterByIdInput{value: val, isSet: true}
}

func (v NullableUpdateFilterByIdInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterByIdInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
