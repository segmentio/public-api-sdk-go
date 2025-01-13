/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateFilterForSpaceInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFilterForSpaceInput{}

// CreateFilterForSpaceInput Input for CreateFilter.
type CreateFilterForSpaceInput struct {
	// The Space id to filter on.
	IntegrationId string `json:"integrationId"`
	// Whether the filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the filter.
	Name string `json:"name"`
	// The description of the filter.
	Description *string `json:"description,omitempty"`
	// The \"if\" statement for a filter.
	If string `json:"if"`
	// Whether the event is dropped.
	Drop *bool `json:"drop,omitempty"`
	// Describes the properties to be dropped on events that match the \"if\" statement.
	DropProperties []string `json:"dropProperties,omitempty"`
	// Describes the properties allowed on events that match the \"if\" statement.
	AllowProperties []string `json:"allowProperties,omitempty"`
}

// NewCreateFilterForSpaceInput instantiates a new CreateFilterForSpaceInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterForSpaceInput(
	integrationId string,
	name string,
	if_ string,
) *CreateFilterForSpaceInput {
	this := CreateFilterForSpaceInput{}
	this.IntegrationId = integrationId
	this.Name = name
	this.If = if_
	return &this
}

// NewCreateFilterForSpaceInputWithDefaults instantiates a new CreateFilterForSpaceInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFilterForSpaceInputWithDefaults() *CreateFilterForSpaceInput {
	this := CreateFilterForSpaceInput{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value
func (o *CreateFilterForSpaceInput) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value
func (o *CreateFilterForSpaceInput) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateFilterForSpaceInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateFilterForSpaceInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateFilterForSpaceInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value
func (o *CreateFilterForSpaceInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFilterForSpaceInput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFilterForSpaceInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFilterForSpaceInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFilterForSpaceInput) SetDescription(v string) {
	o.Description = &v
}

// GetIf returns the If field value
func (o *CreateFilterForSpaceInput) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *CreateFilterForSpaceInput) SetIf(v string) {
	o.If = v
}

// GetDrop returns the Drop field value if set, zero value otherwise.
func (o *CreateFilterForSpaceInput) GetDrop() bool {
	if o == nil || IsNil(o.Drop) {
		var ret bool
		return ret
	}
	return *o.Drop
}

// GetDropOk returns a tuple with the Drop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetDropOk() (*bool, bool) {
	if o == nil || IsNil(o.Drop) {
		return nil, false
	}
	return o.Drop, true
}

// HasDrop returns a boolean if a field has been set.
func (o *CreateFilterForSpaceInput) HasDrop() bool {
	if o != nil && !IsNil(o.Drop) {
		return true
	}

	return false
}

// SetDrop gets a reference to the given bool and assigns it to the Drop field.
func (o *CreateFilterForSpaceInput) SetDrop(v bool) {
	o.Drop = &v
}

// GetDropProperties returns the DropProperties field value if set, zero value otherwise.
func (o *CreateFilterForSpaceInput) GetDropProperties() []string {
	if o == nil || IsNil(o.DropProperties) {
		var ret []string
		return ret
	}
	return o.DropProperties
}

// GetDropPropertiesOk returns a tuple with the DropProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetDropPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.DropProperties) {
		return nil, false
	}
	return o.DropProperties, true
}

// HasDropProperties returns a boolean if a field has been set.
func (o *CreateFilterForSpaceInput) HasDropProperties() bool {
	if o != nil && !IsNil(o.DropProperties) {
		return true
	}

	return false
}

// SetDropProperties gets a reference to the given []string and assigns it to the DropProperties field.
func (o *CreateFilterForSpaceInput) SetDropProperties(v []string) {
	o.DropProperties = v
}

// GetAllowProperties returns the AllowProperties field value if set, zero value otherwise.
func (o *CreateFilterForSpaceInput) GetAllowProperties() []string {
	if o == nil || IsNil(o.AllowProperties) {
		var ret []string
		return ret
	}
	return o.AllowProperties
}

// GetAllowPropertiesOk returns a tuple with the AllowProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForSpaceInput) GetAllowPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowProperties) {
		return nil, false
	}
	return o.AllowProperties, true
}

// HasAllowProperties returns a boolean if a field has been set.
func (o *CreateFilterForSpaceInput) HasAllowProperties() bool {
	if o != nil && !IsNil(o.AllowProperties) {
		return true
	}

	return false
}

// SetAllowProperties gets a reference to the given []string and assigns it to the AllowProperties field.
func (o *CreateFilterForSpaceInput) SetAllowProperties(v []string) {
	o.AllowProperties = v
}

func (o CreateFilterForSpaceInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFilterForSpaceInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["integrationId"] = o.IntegrationId
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["if"] = o.If
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

type NullableCreateFilterForSpaceInput struct {
	value *CreateFilterForSpaceInput
	isSet bool
}

func (v NullableCreateFilterForSpaceInput) Get() *CreateFilterForSpaceInput {
	return v.value
}

func (v *NullableCreateFilterForSpaceInput) Set(val *CreateFilterForSpaceInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterForSpaceInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterForSpaceInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterForSpaceInput(
	val *CreateFilterForSpaceInput,
) *NullableCreateFilterForSpaceInput {
	return &NullableCreateFilterForSpaceInput{value: val, isSet: true}
}

func (v NullableCreateFilterForSpaceInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterForSpaceInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
