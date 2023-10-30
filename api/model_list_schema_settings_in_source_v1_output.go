/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 37.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListSchemaSettingsInSourceV1Output Returns a list of settings for the Source.
type ListSchemaSettingsInSourceV1Output struct {
	// Source id.  Config API note: analogous to `parent` and `name`.
	SourceId string   `json:"sourceId"`
	Settings Settings `json:"settings"`
}

// NewListSchemaSettingsInSourceV1Output instantiates a new ListSchemaSettingsInSourceV1Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSchemaSettingsInSourceV1Output(
	sourceId string,
	settings Settings,
) *ListSchemaSettingsInSourceV1Output {
	this := ListSchemaSettingsInSourceV1Output{}
	this.SourceId = sourceId
	this.Settings = settings
	return &this
}

// NewListSchemaSettingsInSourceV1OutputWithDefaults instantiates a new ListSchemaSettingsInSourceV1Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSchemaSettingsInSourceV1OutputWithDefaults() *ListSchemaSettingsInSourceV1Output {
	this := ListSchemaSettingsInSourceV1Output{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *ListSchemaSettingsInSourceV1Output) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *ListSchemaSettingsInSourceV1Output) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *ListSchemaSettingsInSourceV1Output) SetSourceId(v string) {
	o.SourceId = v
}

// GetSettings returns the Settings field value
func (o *ListSchemaSettingsInSourceV1Output) GetSettings() Settings {
	if o == nil {
		var ret Settings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *ListSchemaSettingsInSourceV1Output) GetSettingsOk() (*Settings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *ListSchemaSettingsInSourceV1Output) SetSettings(v Settings) {
	o.Settings = v
}

func (o ListSchemaSettingsInSourceV1Output) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableListSchemaSettingsInSourceV1Output struct {
	value *ListSchemaSettingsInSourceV1Output
	isSet bool
}

func (v NullableListSchemaSettingsInSourceV1Output) Get() *ListSchemaSettingsInSourceV1Output {
	return v.value
}

func (v *NullableListSchemaSettingsInSourceV1Output) Set(val *ListSchemaSettingsInSourceV1Output) {
	v.value = val
	v.isSet = true
}

func (v NullableListSchemaSettingsInSourceV1Output) IsSet() bool {
	return v.isSet
}

func (v *NullableListSchemaSettingsInSourceV1Output) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSchemaSettingsInSourceV1Output(
	val *ListSchemaSettingsInSourceV1Output,
) *NullableListSchemaSettingsInSourceV1Output {
	return &NullableListSchemaSettingsInSourceV1Output{value: val, isSet: true}
}

func (v NullableListSchemaSettingsInSourceV1Output) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSchemaSettingsInSourceV1Output) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
