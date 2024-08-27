/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateSchemaSettingsInSourceV1Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSchemaSettingsInSourceV1Output{}

// UpdateSchemaSettingsInSourceV1Output Output of the Source with updated settings.
type UpdateSchemaSettingsInSourceV1Output struct {
	// The id of the updated Source.  Config API note: analogous to `parent` and `name`.
	SourceId string                 `json:"sourceId"`
	Settings SourceSettingsOutputV1 `json:"settings"`
}

// NewUpdateSchemaSettingsInSourceV1Output instantiates a new UpdateSchemaSettingsInSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSchemaSettingsInSourceV1Output(
	sourceId string,
	settings SourceSettingsOutputV1,
) *UpdateSchemaSettingsInSourceV1Output {
	this := UpdateSchemaSettingsInSourceV1Output{}
	this.SourceId = sourceId
	this.Settings = settings
	return &this
}

// NewUpdateSchemaSettingsInSourceV1OutputWithDefaults instantiates a new UpdateSchemaSettingsInSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSchemaSettingsInSourceV1OutputWithDefaults() *UpdateSchemaSettingsInSourceV1Output {
	this := UpdateSchemaSettingsInSourceV1Output{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *UpdateSchemaSettingsInSourceV1Output) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Output) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *UpdateSchemaSettingsInSourceV1Output) SetSourceId(v string) {
	o.SourceId = v
}

// GetSettings returns the Settings field value
func (o *UpdateSchemaSettingsInSourceV1Output) GetSettings() SourceSettingsOutputV1 {
	if o == nil {
		var ret SourceSettingsOutputV1
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *UpdateSchemaSettingsInSourceV1Output) GetSettingsOk() (*SourceSettingsOutputV1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *UpdateSchemaSettingsInSourceV1Output) SetSettings(v SourceSettingsOutputV1) {
	o.Settings = v
}

func (o UpdateSchemaSettingsInSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSchemaSettingsInSourceV1Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["settings"] = o.Settings
	return toSerialize, nil
}

type NullableUpdateSchemaSettingsInSourceV1Output struct {
	value *UpdateSchemaSettingsInSourceV1Output
	isSet bool
}

func (v NullableUpdateSchemaSettingsInSourceV1Output) Get() *UpdateSchemaSettingsInSourceV1Output {
	return v.value
}

func (v *NullableUpdateSchemaSettingsInSourceV1Output) Set(
	val *UpdateSchemaSettingsInSourceV1Output,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSchemaSettingsInSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSchemaSettingsInSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSchemaSettingsInSourceV1Output(
	val *UpdateSchemaSettingsInSourceV1Output,
) *NullableUpdateSchemaSettingsInSourceV1Output {
	return &NullableUpdateSchemaSettingsInSourceV1Output{value: val, isSet: true}
}

func (v NullableUpdateSchemaSettingsInSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSchemaSettingsInSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
