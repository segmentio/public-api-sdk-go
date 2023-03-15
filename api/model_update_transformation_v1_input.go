/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// UpdateTransformationV1Input The input to update a Transformation.
type UpdateTransformationV1Input struct {
	// The name of the Transformation.
	Name *string `json:"name,omitempty"`
	// The optional Source to be associated with the Transformation.
	SourceId *string `json:"sourceId,omitempty"`
	// The optional Destination metadata to be associated with the Transformation.
	DestinationMetadataId *string `json:"destinationMetadataId,omitempty"`
	// If the Transformation should be enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// If statement ([FQL](https://segment.com/docs/config-api/fql/)) to match events.  For standard event matchers, use the following:   Track -\\> \"event='\\<eventName\\>'\"   Identify -\\> \"type='identify'\"   Group -\\> \"type='group'\"
	If *string `json:"if,omitempty"`
	// Optional new event name for renaming events. Works only for 'track' event type.
	NewEventName *string `json:"newEventName,omitempty"`
	// Optional array for renaming properties collected by your events.
	PropertyRenames []PropertyRenameV1 `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationV1 `json:"propertyValueTransformations,omitempty"`
	// Optional array for updating properties defined in [FQL](https://segment.com/docs/config-api/fql/). Currently limited to 1 property.
	FqlDefinedProperties []FQLDefinedPropertyV1 `json:"fqlDefinedProperties,omitempty"`
}

// NewUpdateTransformationV1Input instantiates a new UpdateTransformationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransformationV1Input() *UpdateTransformationV1Input {
	this := UpdateTransformationV1Input{}
	return &this
}

// NewUpdateTransformationV1InputWithDefaults instantiates a new UpdateTransformationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransformationV1InputWithDefaults() *UpdateTransformationV1Input {
	this := UpdateTransformationV1Input{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTransformationV1Input) SetName(v string) {
	o.Name = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *UpdateTransformationV1Input) SetSourceId(v string) {
	o.SourceId = &v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetDestinationMetadataId() string {
	if o == nil || o.DestinationMetadataId == nil {
		var ret string
		return ret
	}
	return *o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil || o.DestinationMetadataId == nil {
		return nil, false
	}
	return o.DestinationMetadataId, true
}

// HasDestinationMetadataId returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasDestinationMetadataId() bool {
	if o != nil && o.DestinationMetadataId != nil {
		return true
	}

	return false
}

// SetDestinationMetadataId gets a reference to the given string and assigns it to the DestinationMetadataId field.
func (o *UpdateTransformationV1Input) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateTransformationV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIf returns the If field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetIf() string {
	if o == nil || o.If == nil {
		var ret string
		return ret
	}
	return *o.If
}

// GetIfOk returns a tuple with the If field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetIfOk() (*string, bool) {
	if o == nil || o.If == nil {
		return nil, false
	}
	return o.If, true
}

// HasIf returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasIf() bool {
	if o != nil && o.If != nil {
		return true
	}

	return false
}

// SetIf gets a reference to the given string and assigns it to the If field.
func (o *UpdateTransformationV1Input) SetIf(v string) {
	o.If = &v
}

// GetNewEventName returns the NewEventName field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetNewEventName() string {
	if o == nil || o.NewEventName == nil {
		var ret string
		return ret
	}
	return *o.NewEventName
}

// GetNewEventNameOk returns a tuple with the NewEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetNewEventNameOk() (*string, bool) {
	if o == nil || o.NewEventName == nil {
		return nil, false
	}
	return o.NewEventName, true
}

// HasNewEventName returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasNewEventName() bool {
	if o != nil && o.NewEventName != nil {
		return true
	}

	return false
}

// SetNewEventName gets a reference to the given string and assigns it to the NewEventName field.
func (o *UpdateTransformationV1Input) SetNewEventName(v string) {
	o.NewEventName = &v
}

// GetPropertyRenames returns the PropertyRenames field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetPropertyRenames() []PropertyRenameV1 {
	if o == nil || o.PropertyRenames == nil {
		var ret []PropertyRenameV1
		return ret
	}
	return o.PropertyRenames
}

// GetPropertyRenamesOk returns a tuple with the PropertyRenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetPropertyRenamesOk() ([]PropertyRenameV1, bool) {
	if o == nil || o.PropertyRenames == nil {
		return nil, false
	}
	return o.PropertyRenames, true
}

// HasPropertyRenames returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasPropertyRenames() bool {
	if o != nil && o.PropertyRenames != nil {
		return true
	}

	return false
}

// SetPropertyRenames gets a reference to the given []PropertyRenameV1 and assigns it to the PropertyRenames field.
func (o *UpdateTransformationV1Input) SetPropertyRenames(v []PropertyRenameV1) {
	o.PropertyRenames = v
}

// GetPropertyValueTransformations returns the PropertyValueTransformations field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetPropertyValueTransformations() []PropertyValueTransformationV1 {
	if o == nil || o.PropertyValueTransformations == nil {
		var ret []PropertyValueTransformationV1
		return ret
	}
	return o.PropertyValueTransformations
}

// GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetPropertyValueTransformationsOk() ([]PropertyValueTransformationV1, bool) {
	if o == nil || o.PropertyValueTransformations == nil {
		return nil, false
	}
	return o.PropertyValueTransformations, true
}

// HasPropertyValueTransformations returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasPropertyValueTransformations() bool {
	if o != nil && o.PropertyValueTransformations != nil {
		return true
	}

	return false
}

// SetPropertyValueTransformations gets a reference to the given []PropertyValueTransformationV1 and assigns it to the PropertyValueTransformations field.
func (o *UpdateTransformationV1Input) SetPropertyValueTransformations(
	v []PropertyValueTransformationV1,
) {
	o.PropertyValueTransformations = v
}

// GetFqlDefinedProperties returns the FqlDefinedProperties field value if set, zero value otherwise.
func (o *UpdateTransformationV1Input) GetFqlDefinedProperties() []FQLDefinedPropertyV1 {
	if o == nil || o.FqlDefinedProperties == nil {
		var ret []FQLDefinedPropertyV1
		return ret
	}
	return o.FqlDefinedProperties
}

// GetFqlDefinedPropertiesOk returns a tuple with the FqlDefinedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationV1Input) GetFqlDefinedPropertiesOk() ([]FQLDefinedPropertyV1, bool) {
	if o == nil || o.FqlDefinedProperties == nil {
		return nil, false
	}
	return o.FqlDefinedProperties, true
}

// HasFqlDefinedProperties returns a boolean if a field has been set.
func (o *UpdateTransformationV1Input) HasFqlDefinedProperties() bool {
	if o != nil && o.FqlDefinedProperties != nil {
		return true
	}

	return false
}

// SetFqlDefinedProperties gets a reference to the given []FQLDefinedPropertyV1 and assigns it to the FqlDefinedProperties field.
func (o *UpdateTransformationV1Input) SetFqlDefinedProperties(v []FQLDefinedPropertyV1) {
	o.FqlDefinedProperties = v
}

func (o UpdateTransformationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.DestinationMetadataId != nil {
		toSerialize["destinationMetadataId"] = o.DestinationMetadataId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.If != nil {
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

type NullableUpdateTransformationV1Input struct {
	value *UpdateTransformationV1Input
	isSet bool
}

func (v NullableUpdateTransformationV1Input) Get() *UpdateTransformationV1Input {
	return v.value
}

func (v *NullableUpdateTransformationV1Input) Set(val *UpdateTransformationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransformationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransformationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransformationV1Input(
	val *UpdateTransformationV1Input,
) *NullableUpdateTransformationV1Input {
	return &NullableUpdateTransformationV1Input{value: val, isSet: true}
}

func (v NullableUpdateTransformationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransformationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
