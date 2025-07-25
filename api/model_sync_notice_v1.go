/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SyncNoticeV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncNoticeV1{}

// SyncNoticeV1 Represents a notice within a sync for a Source and Warehouse pair.
type SyncNoticeV1 struct {
	// The severity of the notice.
	Level string `json:"level"`
	// The human-readable message that describes the notice.
	Message string `json:"message"`
	// The timestamp of this sync notice's creation.
	CreatedAt string `json:"createdAt"`
}

// NewSyncNoticeV1 instantiates a new SyncNoticeV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncNoticeV1(level string, message string, createdAt string) *SyncNoticeV1 {
	this := SyncNoticeV1{}
	this.Level = level
	this.Message = message
	this.CreatedAt = createdAt
	return &this
}

// NewSyncNoticeV1WithDefaults instantiates a new SyncNoticeV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncNoticeV1WithDefaults() *SyncNoticeV1 {
	this := SyncNoticeV1{}
	return &this
}

// GetLevel returns the Level field value
func (o *SyncNoticeV1) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *SyncNoticeV1) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *SyncNoticeV1) SetLevel(v string) {
	o.Level = v
}

// GetMessage returns the Message field value
func (o *SyncNoticeV1) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SyncNoticeV1) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SyncNoticeV1) SetMessage(v string) {
	o.Message = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SyncNoticeV1) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SyncNoticeV1) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SyncNoticeV1) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o SyncNoticeV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncNoticeV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["level"] = o.Level
	toSerialize["message"] = o.Message
	toSerialize["createdAt"] = o.CreatedAt
	return toSerialize, nil
}

type NullableSyncNoticeV1 struct {
	value *SyncNoticeV1
	isSet bool
}

func (v NullableSyncNoticeV1) Get() *SyncNoticeV1 {
	return v.value
}

func (v *NullableSyncNoticeV1) Set(val *SyncNoticeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncNoticeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncNoticeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncNoticeV1(val *SyncNoticeV1) *NullableSyncNoticeV1 {
	return &NullableSyncNoticeV1{value: val, isSet: true}
}

func (v NullableSyncNoticeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncNoticeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
