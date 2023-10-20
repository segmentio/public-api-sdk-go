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

// Transformation4 The updated Transformation.
type Transformation4 struct {
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
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:   Track -\\> \"event='\\<eventName\\>'\"   Identify -\\> \"type='identify'\"   Group -\\> \"type='group'\"
	If string `json:"if"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// Optional array for defining new properties in FQL. Limited to 1 property right now.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
}

// NewTransformation4 instantiates a new Transformation4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformation4(
	id string,
	name string,
	sourceId string,
	enabled bool,
	if_ string,
) *Transformation4 {
	this := Transformation4{}
	this.Id = id
	this.Name = name
	this.SourceId = sourceId
	this.Enabled = enabled
	this.If = if_
	return &this
}

// NewTransformation4WithDefaults instantiates a new Transformation4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformation4WithDefaults() *Transformation4 {
	this := Transformation4{}
	return &this
}

// GetId returns the Id field value
func (o *Transformation4) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Transformation4) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Transformation4) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Transformation4) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Transformation4) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Transformation4) SetName(v string) {
	o.Name = v
}

// GetSourceId returns the SourceId field value
func (o *Transformation4) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *Transformation4) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *Transformation4) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value if set, zero value otherwise.
func (o *Transformation4) GetDestinationMetadataId() string {
	if o == nil || o.DestinationMetadataId == nil {
		var ret string
		return ret
	}
	return *o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transformation4) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil || o.DestinationMetadataId == nil {
		return nil, false
	}
	return o.DestinationMetadataId, true
}

// HasDestinationMetadataId returns a boolean if a field has been set.
func (o *Transformation4) HasDestinationMetadataId() bool {
	if o != nil && o.DestinationMetadataId != nil {
		return true
	}

	return false
}

// SetDestinationMetadataId gets a reference to the given string and assigns it to the DestinationMetadataId field.
func (o *Transformation4) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = &v
}

// GetEnabled returns the Enabled field value
func (o *Transformation4) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Transformation4) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Transformation4) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIf returns the If field value
func (o *Transformation4) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *Transformation4) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *Transformation4) SetIf(v string) {
	o.If = v
}

// GetNewEventName returns the NewEventName field value if set, zero value otherwise.
func (o *Transformation4) GetNewEventName() string {
	if o == nil || o.NewEventName == nil {
		var ret string
		return ret
	}
	return *o.NewEventName
}

// GetNewEventNameOk returns a tuple with the NewEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transformation4) GetNewEventNameOk() (*string, bool) {
	if o == nil || o.NewEventName == nil {
		return nil, false
	}
	return o.NewEventName, true
}

// HasNewEventName returns a boolean if a field has been set.
func (o *Transformation4) HasNewEventName() bool {
	if o != nil && o.NewEventName != nil {
		return true
	}

	return false
}

// SetNewEventName gets a reference to the given string and assigns it to the NewEventName field.
func (o *Transformation4) SetNewEventName(v string) {
	o.NewEventName = &v
}

// GetPropertyRenames returns the PropertyRenames field value if set, zero value otherwise.
func (o *Transformation4) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil || o.PropertyRenames == nil {
		var ret []PropertyRenameV1
		return ret
	}
	return o.PropertyRenames
}

// GetPropertyRenamesOk returns a tuple with the PropertyRenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transformation4) GetPropertyRenamesOk() ([]PropertyRenameV1, bool) {
	if o == nil || o.PropertyRenames == nil {
		return nil, false
	}
	return o.PropertyRenames, true
}

// HasPropertyRenames returns a boolean if a field has been set.
func (o *Transformation4) HasPropertyRenames() bool {
	if o != nil && o.PropertyRenames != nil {
		return true
	}

	return false
}

// SetPropertyRenames gets a reference to the given []PropertyRenameV1 and assigns it to the PropertyRenames field.
func (o *Transformation4) SetPropertyRenames(v []PropertyRenameV1) {
	o.PropertyRenames = v
}

// GetPropertyValueTransformations returns the PropertyValueTransformations field value if set, zero value otherwise.
func (o *Transformation4) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil || o.PropertyValueTransformations == nil {
		var ret []PropertyValueTransformationV1
		return ret
	}
	return o.PropertyValueTransformations
}

// GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transformation4) GetPropertyValueTransformationsOk() ([]PropertyValueTransformationV1, bool) {
	if o == nil || o.PropertyValueTransformations == nil {
		return nil, false
	}
	return o.PropertyValueTransformations, true
}

// HasPropertyValueTransformations returns a boolean if a field has been set.
func (o *Transformation4) HasPropertyValueTransformations() bool {
	if o != nil && o.PropertyValueTransformations != nil {
		return true
	}

	return false
}

// SetPropertyValueTransformations gets a reference to the given []PropertyValueTransformationV1 and assigns it to the PropertyValueTransformations field.
func (o *Transformation4) SetPropertyValueTransformations(v []PropertyValueTransformationV1) {
	o.PropertyValueTransformations = v
}

// GetFqlDefinedProperties returns the FqlDefinedProperties field value if set, zero value otherwise.
func (o *Transformation4) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil || o.FqlDefinedProperties == nil {
		var ret []FQLDefinedPropertyV1
		return ret
	}
	return o.FqlDefinedProperties
}

// GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transformation4) GetFqlDefinedPropertiesOk() ([]FQLDefinedPropertyV1, bool) {
	if o == nil || o.FqlDefinedProperties == nil {
		return nil, false
	}
	return o.FqlDefinedProperties, true
}

// HasFqlDefinedProperties returns a boolean if a field has been set.
func (o *Transformation4) HasFqlDefinedProperties() bool {
	if o != nil && o.FqlDefinedProperties != nil {
		return true
	}

	return false
}

// SetFqlDefinedProperties gets a reference to the given []FQLDefinedPropertyV1 and assigns it to the FqlDefinedProperties field.
func (o *Transformation4) SetFqlDefinedProperties(v []FQLDefinedPropertyV1) {
	o.FqlDefinedProperties = v
}

func (o Transformation4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.DestinationMetadataId != nil {
		toSerialize["destinationMetadataId"] = o.DestinationMetadataId
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["if"] = o.If
	}
	if o.NewEventName != nil {
		toSerialize["newEventName"] = o.NewEventName
	}
	if o.PropertyRenames != nil {
		toSerialize["propertyRenames"] = o.PropertyRenames
	}
	if o.PropertyValueTransformations != nil {
		toSerialize["propertyValueTransformations"] = o.PropertyValueTransformations
	}
	if o.FqlDefinedProperties != nil {
		toSerialize["fqlDefinedProperties"] = o.FqlDefinedProperties
	}
	return json.Marshal(toSerialize)
}

type NullableTransformation4 struct {
	value *Transformation4
	isSet bool
}

func (v NullableTransformation4) Get() *Transformation4 {
	return v.value
}

func (v *NullableTransformation4) Set(val *Transformation4) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformation4) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformation4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformation4(val *Transformation4) *NullableTransformation4 {
	return &NullableTransformation4{value: val, isSet: true}
}

func (v NullableTransformation4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformation4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
