/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 45.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TransformationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationV1{}

// TransformationV1 Represents a Transformation.
type TransformationV1 struct {
	// The id of the Transformation.
	Id string `json:"id"`
	// The name of the Transformation.
	Name string `json:"name"`
	// The Source associated with the Transformation.
	SourceId string `json:"sourceId"`
	// The optional Destination metadata associated with the Transformation.
	DestinationMetadataId *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation is enabled.
	Enabled bool `json:"enabled"`
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:  Track -\\> \"event='\\<eventName\\>'\"  Identify -\\> \"type='identify'\"  Group -\\> \"type='group'\"
	If string `json:"if"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// Optional array for defining new properties in FQL. Limited to 1 property right now.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
	// Optional array for allowing properties from your events.
	AllowProperties             []string                     `json:"allowProperties,omitempty"`
	HashPropertiesConfiguration *HashPropertiesConfiguration `json:"hashPropertiesConfiguration,omitempty"`
}

// NewTransformationV1 instantiates a new TransformationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationV1(
	id string,
	name string,
	sourceId string,
	enabled bool,
	if_ string,
) *TransformationV1 {
	this := TransformationV1{}
	this.Id = id
	this.Name = name
	this.SourceId = sourceId
	this.Enabled = enabled
	this.If = if_
	return &this
}

// NewTransformationV1WithDefaults instantiates a new TransformationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationV1WithDefaults() *TransformationV1 {
	this := TransformationV1{}
	return &this
}

// GetId returns the Id field value
func (o *TransformationV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransformationV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *TransformationV1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TransformationV1) SetName(v string) {
	o.Name = v
}

// GetSourceId returns the SourceId field value
func (o *TransformationV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *TransformationV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value if set, zero value otherwise.
func (o *TransformationV1) GetDestinationMetadataId() string {
	if o == nil || IsNil(o.DestinationMetadataId) {
		var ret string
		return ret
	}
	return *o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationMetadataId) {
		return nil, false
	}
	return o.DestinationMetadataId, true
}

// HasDestinationMetadataId returns a boolean if a field has been set.
func (o *TransformationV1) HasDestinationMetadataId() bool {
	if o != nil && !IsNil(o.DestinationMetadataId) {
		return true
	}

	return false
}

// SetDestinationMetadataId gets a reference to the given string and assigns it to the DestinationMetadataId field.
func (o *TransformationV1) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = &v
}

// GetEnabled returns the Enabled field value
func (o *TransformationV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TransformationV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIf returns the If field value
func (o *TransformationV1) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *TransformationV1) SetIf(v string) {
	o.If = v
}

// GetNewEventName returns the NewEventName field value if set, zero value otherwise.
func (o *TransformationV1) GetNewEventName() string {
	if o == nil || IsNil(o.NewEventName) {
		var ret string
		return ret
	}
	return *o.NewEventName
}

// GetNewEventNameOk returns a tuple with the NewEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetNewEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewEventName) {
		return nil, false
	}
	return o.NewEventName, true
}

// HasNewEventName returns a boolean if a field has been set.
func (o *TransformationV1) HasNewEventName() bool {
	if o != nil && !IsNil(o.NewEventName) {
		return true
	}

	return false
}

// SetNewEventName gets a reference to the given string and assigns it to the NewEventName field.
func (o *TransformationV1) SetNewEventName(v string) {
	o.NewEventName = &v
}

// GetPropertyRenames returns the PropertyRenames field value if set, zero value otherwise.
func (o *TransformationV1) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil || IsNil(o.PropertyRenames) {
		var ret []PropertyRenameV1
		return ret
	}
	return o.PropertyRenames
}

// GetPropertyRenamesOk returns a tuple with the PropertyRenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetPropertyRenamesOk() ([]PropertyRenameV1, bool) {
	if o == nil || IsNil(o.PropertyRenames) {
		return nil, false
	}
	return o.PropertyRenames, true
}

// HasPropertyRenames returns a boolean if a field has been set.
func (o *TransformationV1) HasPropertyRenames() bool {
	if o != nil && !IsNil(o.PropertyRenames) {
		return true
	}

	return false
}

// SetPropertyRenames gets a reference to the given []PropertyRenameV1 and assigns it to the PropertyRenames field.
func (o *TransformationV1) SetPropertyRenames(v []PropertyRenameV1) {
	o.PropertyRenames = v
}

// GetPropertyValueTransformations returns the PropertyValueTransformations field value if set, zero value otherwise.
func (o *TransformationV1) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil || IsNil(o.PropertyValueTransformations) {
		var ret []PropertyValueTransformationV1
		return ret
	}
	return o.PropertyValueTransformations
}

// GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetPropertyValueTransformationsOk() ([]PropertyValueTransformationV1, bool) {
	if o == nil || IsNil(o.PropertyValueTransformations) {
		return nil, false
	}
	return o.PropertyValueTransformations, true
}

// HasPropertyValueTransformations returns a boolean if a field has been set.
func (o *TransformationV1) HasPropertyValueTransformations() bool {
	if o != nil && !IsNil(o.PropertyValueTransformations) {
		return true
	}

	return false
}

// SetPropertyValueTransformations gets a reference to the given []PropertyValueTransformationV1 and assigns it to the PropertyValueTransformations field.
func (o *TransformationV1) SetPropertyValueTransformations(v []PropertyValueTransformationV1) {
	o.PropertyValueTransformations = v
}

// GetFqlDefinedProperties returns the FqlDefinedProperties field value if set, zero value otherwise.
func (o *TransformationV1) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil || IsNil(o.FqlDefinedProperties) {
		var ret []FQLDefinedPropertyV1
		return ret
	}
	return o.FqlDefinedProperties
}

// GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetFqlDefinedPropertiesOk() ([]FQLDefinedPropertyV1, bool) {
	if o == nil || IsNil(o.FqlDefinedProperties) {
		return nil, false
	}
	return o.FqlDefinedProperties, true
}

// HasFqlDefinedProperties returns a boolean if a field has been set.
func (o *TransformationV1) HasFqlDefinedProperties() bool {
	if o != nil && !IsNil(o.FqlDefinedProperties) {
		return true
	}

	return false
}

// SetFqlDefinedProperties gets a reference to the given []FQLDefinedPropertyV1 and assigns it to the FqlDefinedProperties field.
func (o *TransformationV1) SetFqlDefinedProperties(v []FQLDefinedPropertyV1) {
	o.FqlDefinedProperties = v
}

// GetAllowProperties returns the AllowProperties field value if set, zero value otherwise.
func (o *TransformationV1) GetAllowProperties() []string {
	if o == nil || IsNil(o.AllowProperties) {
		var ret []string
		return ret
	}
	return o.AllowProperties
}

// GetAllowPropertiesOk returns a tuple with the AllowProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetAllowPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowProperties) {
		return nil, false
	}
	return o.AllowProperties, true
}

// HasAllowProperties returns a boolean if a field has been set.
func (o *TransformationV1) HasAllowProperties() bool {
	if o != nil && !IsNil(o.AllowProperties) {
		return true
	}

	return false
}

// SetAllowProperties gets a reference to the given []string and assigns it to the AllowProperties field.
func (o *TransformationV1) SetAllowProperties(v []string) {
	o.AllowProperties = v
}

// GetHashPropertiesConfiguration returns the HashPropertiesConfiguration field value if set, zero value otherwise.
func (o *TransformationV1) GetHashPropertiesConfiguration() HashPropertiesConfiguration {
	if o == nil || IsNil(o.HashPropertiesConfiguration) {
		var ret HashPropertiesConfiguration
		return ret
	}
	return *o.HashPropertiesConfiguration
}

// GetHashPropertiesConfigurationOk returns a tuple with the HashPropertiesConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransformationV1) GetHashPropertiesConfigurationOk() (*HashPropertiesConfiguration, bool) {
	if o == nil || IsNil(o.HashPropertiesConfiguration) {
		return nil, false
	}
	return o.HashPropertiesConfiguration, true
}

// HasHashPropertiesConfiguration returns a boolean if a field has been set.
func (o *TransformationV1) HasHashPropertiesConfiguration() bool {
	if o != nil && !IsNil(o.HashPropertiesConfiguration) {
		return true
	}

	return false
}

// SetHashPropertiesConfiguration gets a reference to the given HashPropertiesConfiguration and assigns it to the HashPropertiesConfiguration field.
func (o *TransformationV1) SetHashPropertiesConfiguration(v HashPropertiesConfiguration) {
	o.HashPropertiesConfiguration = &v
}

func (o TransformationV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["sourceId"] = o.SourceId
	if !IsNil(o.DestinationMetadataId) {
		toSerialize["destinationMetadataId"] = o.DestinationMetadataId
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["if"] = o.If
	if !IsNil(o.NewEventName) {
		toSerialize["newEventName"] = o.NewEventName
	}
	if !IsNil(o.PropertyRenames) {
		toSerialize["propertyRenames"] = o.PropertyRenames
	}
	if !IsNil(o.PropertyValueTransformations) {
		toSerialize["propertyValueTransformations"] = o.PropertyValueTransformations
	}
	if !IsNil(o.FqlDefinedProperties) {
		toSerialize["fqlDefinedProperties"] = o.FqlDefinedProperties
	}
	if !IsNil(o.AllowProperties) {
		toSerialize["allowProperties"] = o.AllowProperties
	}
	if !IsNil(o.HashPropertiesConfiguration) {
		toSerialize["hashPropertiesConfiguration"] = o.HashPropertiesConfiguration
	}
	return toSerialize, nil
}

type NullableTransformationV1 struct {
	value *TransformationV1
	isSet bool
}

func (v NullableTransformationV1) Get() *TransformationV1 {
	return v.value
}

func (v *NullableTransformationV1) Set(val *TransformationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationV1(val *TransformationV1) *NullableTransformationV1 {
	return &NullableTransformationV1{value: val, isSet: true}
}

func (v NullableTransformationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
