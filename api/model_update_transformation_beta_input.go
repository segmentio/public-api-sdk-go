/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateTransformationBetaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateTransformationBetaInput{}

// UpdateTransformationBetaInput The input to update a Transformation.
type UpdateTransformationBetaInput struct {
	// ID of the Transformation to update.
	TransformationId string `json:"transformationId"`
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
	PropertyRenames []PropertyRenameBeta `json:"propertyRenames,omitempty"`
	// Optional array for transforming properties and values collected by your events. Limited to 10 properties.
	PropertyValueTransformations []PropertyValueTransformationBeta `json:"propertyValueTransformations,omitempty"`
}

// NewUpdateTransformationBetaInput instantiates a new UpdateTransformationBetaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTransformationBetaInput(transformationId string) *UpdateTransformationBetaInput {
	this := UpdateTransformationBetaInput{}
	this.TransformationId = transformationId
	return &this
}

// NewUpdateTransformationBetaInputWithDefaults instantiates a new UpdateTransformationBetaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTransformationBetaInputWithDefaults() *UpdateTransformationBetaInput {
	this := UpdateTransformationBetaInput{}
	return &this
}

// GetTransformationId returns the TransformationId field value
func (o *UpdateTransformationBetaInput) GetTransformationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransformationId
}

// GetTransformationIdOk returns a tuple with the TransformationId field value
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetTransformationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformationId, true
}

// SetTransformationId sets field value
func (o *UpdateTransformationBetaInput) SetTransformationId(v string) {
	o.TransformationId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateTransformationBetaInput) SetName(v string) {
	o.Name = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *UpdateTransformationBetaInput) SetSourceId(v string) {
	o.SourceId = &v
}

// GetDestinationMetadataId returns the DestinationMetadataId field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetDestinationMetadataId() string {
	if o == nil || IsNil(o.DestinationMetadataId) {
		var ret string
		return ret
	}
	return *o.DestinationMetadataId
}

// GetDestinationMetadataIdOk returns a tuple with the DestinationMetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetDestinationMetadataIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationMetadataId) {
		return nil, false
	}
	return o.DestinationMetadataId, true
}

// HasDestinationMetadataId returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasDestinationMetadataId() bool {
	if o != nil && !IsNil(o.DestinationMetadataId) {
		return true
	}

	return false
}

// SetDestinationMetadataId gets a reference to the given string and assigns it to the DestinationMetadataId field.
func (o *UpdateTransformationBetaInput) SetDestinationMetadataId(v string) {
	o.DestinationMetadataId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateTransformationBetaInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIf returns the If field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetIf() string {
	if o == nil || IsNil(o.If) {
		var ret string
		return ret
	}
	return *o.If
}

// GetIfOk returns a tuple with the If field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetIfOk() (*string, bool) {
	if o == nil || IsNil(o.If) {
		return nil, false
	}
	return o.If, true
}

// HasIf returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasIf() bool {
	if o != nil && !IsNil(o.If) {
		return true
	}

	return false
}

// SetIf gets a reference to the given string and assigns it to the If field.
func (o *UpdateTransformationBetaInput) SetIf(v string) {
	o.If = &v
}

// GetNewEventName returns the NewEventName field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetNewEventName() string {
	if o == nil || IsNil(o.NewEventName) {
		var ret string
		return ret
	}
	return *o.NewEventName
}

// GetNewEventNameOk returns a tuple with the NewEventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetNewEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.NewEventName) {
		return nil, false
	}
	return o.NewEventName, true
}

// HasNewEventName returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasNewEventName() bool {
	if o != nil && !IsNil(o.NewEventName) {
		return true
	}

	return false
}

// SetNewEventName gets a reference to the given string and assigns it to the NewEventName field.
func (o *UpdateTransformationBetaInput) SetNewEventName(v string) {
	o.NewEventName = &v
}

// GetPropertyRenames returns the PropertyRenames field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetPropertyRenames() []PropertyRenameBeta {
	if o == nil || IsNil(o.PropertyRenames) {
		var ret []PropertyRenameBeta
		return ret
	}
	return o.PropertyRenames
}

// GetPropertyRenamesOk returns a tuple with the PropertyRenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetPropertyRenamesOk() ([]PropertyRenameBeta, bool) {
	if o == nil || IsNil(o.PropertyRenames) {
		return nil, false
	}
	return o.PropertyRenames, true
}

// HasPropertyRenames returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasPropertyRenames() bool {
	if o != nil && !IsNil(o.PropertyRenames) {
		return true
	}

	return false
}

// SetPropertyRenames gets a reference to the given []PropertyRenameBeta and assigns it to the PropertyRenames field.
func (o *UpdateTransformationBetaInput) SetPropertyRenames(v []PropertyRenameBeta) {
	o.PropertyRenames = v
}

// GetPropertyValueTransformations returns the PropertyValueTransformations field value if set, zero value otherwise.
func (o *UpdateTransformationBetaInput) GetPropertyValueTransformations() []PropertyValueTransformationBeta {
	if o == nil || IsNil(o.PropertyValueTransformations) {
		var ret []PropertyValueTransformationBeta
		return ret
	}
	return o.PropertyValueTransformations
}

// GetPropertyValueTransformationsOk returns a tuple with the PropertyValueTransformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTransformationBetaInput) GetPropertyValueTransformationsOk() ([]PropertyValueTransformationBeta, bool) {
	if o == nil || IsNil(o.PropertyValueTransformations) {
		return nil, false
	}
	return o.PropertyValueTransformations, true
}

// HasPropertyValueTransformations returns a boolean if a field has been set.
func (o *UpdateTransformationBetaInput) HasPropertyValueTransformations() bool {
	if o != nil && !IsNil(o.PropertyValueTransformations) {
		return true
	}

	return false
}

// SetPropertyValueTransformations gets a reference to the given []PropertyValueTransformationBeta and assigns it to the PropertyValueTransformations field.
func (o *UpdateTransformationBetaInput) SetPropertyValueTransformations(
	v []PropertyValueTransformationBeta,
) {
	o.PropertyValueTransformations = v
}

func (o UpdateTransformationBetaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateTransformationBetaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transformationId"] = o.TransformationId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	if !IsNil(o.DestinationMetadataId) {
		toSerialize["destinationMetadataId"] = o.DestinationMetadataId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.If) {
		toSerialize["if"] = o.If
	}
	if !IsNil(o.NewEventName) {
		toSerialize["newEventName"] = o.NewEventName
	}
	if !IsNil(o.PropertyRenames) {
		toSerialize["propertyRenames"] = o.PropertyRenames
	}
	if !IsNil(o.PropertyValueTransformations) {
		toSerialize["propertyValueTransformations"] = o.PropertyValueTransformations
	}
	return toSerialize, nil
}

type NullableUpdateTransformationBetaInput struct {
	value *UpdateTransformationBetaInput
	isSet bool
}

func (v NullableUpdateTransformationBetaInput) Get() *UpdateTransformationBetaInput {
	return v.value
}

func (v *NullableUpdateTransformationBetaInput) Set(val *UpdateTransformationBetaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTransformationBetaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTransformationBetaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTransformationBetaInput(
	val *UpdateTransformationBetaInput,
) *NullableUpdateTransformationBetaInput {
	return &NullableUpdateTransformationBetaInput{value: val, isSet: true}
}

func (v NullableUpdateTransformationBetaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTransformationBetaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
