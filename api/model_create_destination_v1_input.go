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

// CreateDestinationV1Input Creates a new Destination.
type CreateDestinationV1Input struct {
	// The id of the Source to connect to this Destination instance.  Config API note: analogous to `parent`.
	SourceId string `json:"sourceId"`
	// The id of the metadata to link to the new Destination.
	MetadataId string `json:"metadataId"`
	// Whether this Destination should receive data.
	Enabled *bool `json:"enabled,omitempty"`
	// Defines the display name of the Destination.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// An object that contains settings for the Destination based on the \"required\" and \"advanced\" settings present in the Destination metadata.  Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
}

// NewCreateDestinationV1Input instantiates a new CreateDestinationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDestinationV1Input(
	sourceId string,
	metadataId string,
	settings map[string]interface{},
) *CreateDestinationV1Input {
	this := CreateDestinationV1Input{}
	this.SourceId = sourceId
	this.MetadataId = metadataId
	this.Settings = settings
	return &this
}

// NewCreateDestinationV1InputWithDefaults instantiates a new CreateDestinationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDestinationV1InputWithDefaults() *CreateDestinationV1Input {
	this := CreateDestinationV1Input{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *CreateDestinationV1Input) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CreateDestinationV1Input) SetSourceId(v string) {
	o.SourceId = v
}

// GetMetadataId returns the MetadataId field value
func (o *CreateDestinationV1Input) GetMetadataId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Input) GetMetadataIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataId, true
}

// SetMetadataId sets field value
func (o *CreateDestinationV1Input) SetMetadataId(v string) {
	o.MetadataId = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateDestinationV1Input) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateDestinationV1Input) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateDestinationV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateDestinationV1Input) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Input) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateDestinationV1Input) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateDestinationV1Input) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value
func (o *CreateDestinationV1Input) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CreateDestinationV1Input) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *CreateDestinationV1Input) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o CreateDestinationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["metadataId"] = o.MetadataId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDestinationV1Input struct {
	value *CreateDestinationV1Input
	isSet bool
}

func (v NullableCreateDestinationV1Input) Get() *CreateDestinationV1Input {
	return v.value
}

func (v *NullableCreateDestinationV1Input) Set(val *CreateDestinationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDestinationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDestinationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDestinationV1Input(
	val *CreateDestinationV1Input,
) *NullableCreateDestinationV1Input {
	return &NullableCreateDestinationV1Input{value: val, isSet: true}
}

func (v NullableCreateDestinationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDestinationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
