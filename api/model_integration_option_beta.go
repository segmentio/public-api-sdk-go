/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 54.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the IntegrationOptionBeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationOptionBeta{}

// IntegrationOptionBeta Describes an Integration option field required to set up a Segment Integration such as Sources, Destinations, or warehouses.
type IntegrationOptionBeta struct {
	// The name identifying this option in the context of a Segment Integration.
	Name string `json:"name"`
	// Defines the type for this option in the schema. Types are most commonly strings, but may also represent other primitive types, such as booleans, and numbers, as well as complex types, such as objects and arrays.
	Type string `json:"type"`
	// Whether this is a required option when setting up the Integration.
	Required bool `json:"required"`
	// An optional short text description of the field.
	Description *string `json:"description,omitempty"`
	// An optional default value for the field.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// An optional label for this field.
	Label *string `json:"label,omitempty"`
}

// NewIntegrationOptionBeta instantiates a new IntegrationOptionBeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationOptionBeta(name string, type_ string, required bool) *IntegrationOptionBeta {
	this := IntegrationOptionBeta{}
	this.Name = name
	this.Type = type_
	this.Required = required
	return &this
}

// NewIntegrationOptionBetaWithDefaults instantiates a new IntegrationOptionBeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationOptionBetaWithDefaults() *IntegrationOptionBeta {
	this := IntegrationOptionBeta{}
	return &this
}

// GetName returns the Name field value
func (o *IntegrationOptionBeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationOptionBeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationOptionBeta) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *IntegrationOptionBeta) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IntegrationOptionBeta) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IntegrationOptionBeta) SetType(v string) {
	o.Type = v
}

// GetRequired returns the Required field value
func (o *IntegrationOptionBeta) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *IntegrationOptionBeta) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *IntegrationOptionBeta) SetRequired(v bool) {
	o.Required = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationOptionBeta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOptionBeta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationOptionBeta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationOptionBeta) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IntegrationOptionBeta) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IntegrationOptionBeta) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *IntegrationOptionBeta) HasDefaultValue() bool {
	if o != nil && IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *IntegrationOptionBeta) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IntegrationOptionBeta) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOptionBeta) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IntegrationOptionBeta) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *IntegrationOptionBeta) SetLabel(v string) {
	o.Label = &v
}

func (o IntegrationOptionBeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationOptionBeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["required"] = o.Required
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableIntegrationOptionBeta struct {
	value *IntegrationOptionBeta
	isSet bool
}

func (v NullableIntegrationOptionBeta) Get() *IntegrationOptionBeta {
	return v.value
}

func (v *NullableIntegrationOptionBeta) Set(val *IntegrationOptionBeta) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationOptionBeta) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationOptionBeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationOptionBeta(val *IntegrationOptionBeta) *NullableIntegrationOptionBeta {
	return &NullableIntegrationOptionBeta{value: val, isSet: true}
}

func (v NullableIntegrationOptionBeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationOptionBeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
