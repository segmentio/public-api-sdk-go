/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DestinationV1 Business tools or apps that you can connect to the data flowing through Segment.  This is equal to the Destination object in Config API, with the following fields omitted: - catalogId - createTime - updateTime - connectionMode.
type DestinationV1 struct {
	// The unique identifier of this instance of a Destination.  Config API note: analogous to `name`.
	Id string `json:"id"`
	// The name of this instance of a Destination.  Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// Whether this instance of a Destination receives data.
	Enabled bool `json:"enabled"`
	Metadata Metadata `json:"metadata"`
	// The id of a Source connected to this instance of a Destination.  Config API note: analogous to `parent`.
	SourceId string `json:"sourceId"`
	// The collection of settings associated with a Destination.  Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
}

// NewDestinationV1 instantiates a new DestinationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationV1(id string, enabled bool, metadata Metadata, sourceId string, settings map[string]interface{}) *DestinationV1 {
	this := DestinationV1{}
	this.Id = id
	this.Enabled = enabled
	this.Metadata = metadata
	this.SourceId = sourceId
	this.Settings = settings
	return &this
}

// NewDestinationV1WithDefaults instantiates a new DestinationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationV1WithDefaults() *DestinationV1 {
	this := DestinationV1{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationV1) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DestinationV1) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DestinationV1) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DestinationV1) SetName(v string) {
	o.Name = &v
}

// GetEnabled returns the Enabled field value
func (o *DestinationV1) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DestinationV1) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMetadata returns the Metadata field value
func (o *DestinationV1) GetMetadata() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *DestinationV1) SetMetadata(v Metadata) {
	o.Metadata = v
}

// GetSourceId returns the SourceId field value
func (o *DestinationV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *DestinationV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetSettings returns the Settings field value
func (o *DestinationV1) GetSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *DestinationV1) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings, true
}

// SetSettings sets field value
func (o *DestinationV1) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

func (o DestinationV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationV1 struct {
	value *DestinationV1
	isSet bool
}

func (v NullableDestinationV1) Get() *DestinationV1 {
	return v.value
}

func (v *NullableDestinationV1) Set(val *DestinationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationV1(val *DestinationV1) *NullableDestinationV1 {
	return &NullableDestinationV1{value: val, isSet: true}
}

func (v NullableDestinationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


