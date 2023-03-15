/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateTransformationV1Input The input to create a Transformation.
type CreateTransformationV1Input struct {
	// The name of the Transformation.
	Name string `json:"name"`
	// The Source to be associated with the Transformation.
	SourceId string `json:"sourceId"`
	// The optional Destination metadata id to be associated with the Transformation.
	DestinationMetadataId *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation should be enabled.
	Enabled bool `json:"enabled"`
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:   Track -\\> \"event='\\<eventName\\>'\"   Identify -\\> \"type='identify'\"   Group -\\> \"type='group'\"
	If string `json:"if"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// Optional array for defining new properties in [FQL](https://segment.com/docs/config-api/fql/). Currently limited to 1 property.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
}

// NewCreateTransformationV1Input instantiates a new CreateTransformationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransformationV1Input(
	name string,
	sourceId string,
	enabled bool,
	if_ string,
) *CreateTransformationV1Input {
	this := CreateTransformationV1Input{}
	this.Name = name
	this.SourceId = sourceId
	this.Enabled = enabled
	this.If = if_
	return &this
}

// NewCreateTransformationV1InputWithDefaults instantiates a new CreateTransformationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransformationV1InputWithDefaults() *CreateTransformationV1Input {
	this := CreateTransformationV1Input{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTransformationV1Input) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTransformationV1Input) SetName(v string) {
	o.Name = v
}

// GetSourceId returns the SourceId field value
func (o *CreateTransformationV1Input) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CreateTransformationV1Input) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value if set, zero value otherwise.
func (o *CreateTransformationV1Input) GetDestinationMetadataId() string {
	if o == nil || o.DestinationMetadataId == nil {
		var ret string
		return ret
	}
	return *o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil || o.DestinationMetadataId == nil {
		return nil, false
	}
	return o.DestinationMetadataId, true
}

// HasDestinationMetadataId returns a boolean if a field has been set.
func (o *CreateTransformationV1Input) HasDestinationMetadataId() bool {
	if o != nil && o.DestinationMetadataId != nil {
		return true
	}

	return false
}

// SetDestinationMetadataId gets a reference to the given string and assigns it to the DestinationMetadataId field.
func (o *CreateTransformationV1Input) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateTransformationV1Input) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateTransformationV1Input) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIf returns the If field value
func (o *CreateTransformationV1Input) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *CreateTransformationV1Input) SetIf(v string) {
	o.If = v
}

// GetNewEventName returns the NewEventName field value if set, zero value otherwise.
func (o *CreateTransformationV1Input) GetNewEventName() string {
	if o == nil || o.NewEventName == nil {
		var ret string
		return ret
	}
	return *o.NewEventName
}

// GetNewEventNameOk returns a tuple with the NewEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetNewEventNameOk() (*string, bool) {
	if o == nil || o.NewEventName == nil {
		return nil, false
	}
	return o.NewEventName, true
}

// HasNewEventName returns a boolean if a field has been set.
func (o *CreateTransformationV1Input) HasNewEventName() bool {
	if o != nil && o.NewEventName != nil {
		return true
	}

	return false
}

// SetNewEventName gets a reference to the given string and assigns it to the NewEventName field.
func (o *CreateTransformationV1Input) SetNewEventName(v string) {
	o.NewEventName = &v
}

// GetPropertyRenames returns the PropertyRenames field value if set, zero value otherwise.
func (o *CreateTransformationV1Input) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil || o.PropertyRenames == nil {
		var ret []PropertyRenameV1
		return ret
	}
	return o.PropertyRenames
}

// GetPropertyRenamesOk returns a tuple with the PropertyRenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetPropertyRenamesOk() ([]PropertyRenameV1, bool) {
	if o == nil || o.PropertyRenames == nil {
		return nil, false
	}
	return o.PropertyRenames, true
}

// HasPropertyRenames returns a boolean if a field has been set.
func (o *CreateTransformationV1Input) HasPropertyRenames() bool {
	if o != nil && o.PropertyRenames != nil {
		return true
	}

	return false
}

// SetPropertyRenames gets a reference to the given []PropertyRenameV1 and assigns it to the PropertyRenames field.
func (o *CreateTransformationV1Input) SetPropertyRenames(v []PropertyRenameV1) {
	o.PropertyRenames = v
}

// GetPropertyValueTransformations returns the PropertyValueTransformations field value if set, zero value otherwise.
func (o *CreateTransformationV1Input) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil || o.PropertyValueTransformations == nil {
		var ret []PropertyValueTransformationV1
		return ret
	}
	return o.PropertyValueTransformations
}

// GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetPropertyValueTransformationsOk() ([]PropertyValueTransformationV1, bool) {
	if o == nil || o.PropertyValueTransformations == nil {
		return nil, false
	}
	return o.PropertyValueTransformations, true
}

// HasPropertyValueTransformations returns a boolean if a field has been set.
func (o *CreateTransformationV1Input) HasPropertyValueTransformations() bool {
	if o != nil && o.PropertyValueTransformations != nil {
		return true
	}

	return false
}

// SetPropertyValueTransformations gets a reference to the given []PropertyValueTransformationV1 and assigns it to the PropertyValueTransformations field.
func (o *CreateTransformationV1Input) SetPropertyValueTransformations(
	v []PropertyValueTransformationV1,
) {
	o.PropertyValueTransformations = v
}

// GetFqlDefinedProperties returns the FqlDefinedProperties field value if set, zero value otherwise.
func (o *CreateTransformationV1Input) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil || o.FqlDefinedProperties == nil {
		var ret []FQLDefinedPropertyV1
		return ret
	}
	return o.FqlDefinedProperties
}

// GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformationV1Input) GetFqlDefinedPropertiesOk() ([]FQLDefinedPropertyV1, bool) {
	if o == nil || o.FqlDefinedProperties == nil {
		return nil, false
	}
	return o.FqlDefinedProperties, true
}

// HasFqlDefinedProperties returns a boolean if a field has been set.
func (o *CreateTransformationV1Input) HasFqlDefinedProperties() bool {
	if o != nil && o.FqlDefinedProperties != nil {
		return true
	}

	return false
}

// SetFqlDefinedProperties gets a reference to the given []FQLDefinedPropertyV1 and assigns it to the FqlDefinedProperties field.
func (o *CreateTransformationV1Input) SetFqlDefinedProperties(v []FQLDefinedPropertyV1) {
	o.FqlDefinedProperties = v
}

func (o CreateTransformationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCreateTransformationV1Input struct {
	value *CreateTransformationV1Input
	isSet bool
}

func (v NullableCreateTransformationV1Input) Get() *CreateTransformationV1Input {
	return v.value
}

func (v *NullableCreateTransformationV1Input) Set(val *CreateTransformationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransformationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransformationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransformationV1Input(
	val *CreateTransformationV1Input,
) *NullableCreateTransformationV1Input {
	return &NullableCreateTransformationV1Input{value: val, isSet: true}
}

func (v NullableCreateTransformationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransformationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
