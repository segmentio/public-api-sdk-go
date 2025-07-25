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

// checks if the SyncV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncV1{}

// SyncV1 Represents a sync between a Source and Warehouse.  A sync occurs when data from a Source is loaded into a Warehouse.
type SyncV1 struct {
	// The id of the Source loaded in the sync.
	SourceId string `json:"sourceId"`
	// The start time of the sync.
	Start string `json:"start"`
	// The time the sync completed. Returns null if unfinished.
	End NullableString `json:"end"`
	// The status of the sync.
	Status string `json:"status"`
	// The duration of the sync in seconds. Returns the partial duration if the sync has not finished yet.
	Duration float32 `json:"duration"`
	// The human-readable counterpart of `duration`.
	HumanDuration string `json:"humanDuration"`
	// The number of rows synced into the Warehouse.
	Count float32 `json:"count"`
	// Notices that contain the events that occurred during the sync.
	Notices []SyncNoticeV1 `json:"notices"`
}

// NewSyncV1 instantiates a new SyncV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncV1(
	sourceId string,
	start string,
	end NullableString,
	status string,
	duration float32,
	humanDuration string,
	count float32,
	notices []SyncNoticeV1,
) *SyncV1 {
	this := SyncV1{}
	this.SourceId = sourceId
	this.Start = start
	this.End = end
	this.Status = status
	this.Duration = duration
	this.HumanDuration = humanDuration
	this.Count = count
	this.Notices = notices
	return &this
}

// NewSyncV1WithDefaults instantiates a new SyncV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncV1WithDefaults() *SyncV1 {
	this := SyncV1{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *SyncV1) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *SyncV1) SetSourceId(v string) {
	o.SourceId = v
}

// GetStart returns the Start field value
func (o *SyncV1) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *SyncV1) SetStart(v string) {
	o.Start = v
}

// GetEnd returns the End field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SyncV1) GetEnd() string {
	if o == nil || o.End.Get() == nil {
		var ret string
		return ret
	}

	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SyncV1) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// SetEnd sets field value
func (o *SyncV1) SetEnd(v string) {
	o.End.Set(&v)
}

// GetStatus returns the Status field value
func (o *SyncV1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SyncV1) SetStatus(v string) {
	o.Status = v
}

// GetDuration returns the Duration field value
func (o *SyncV1) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SyncV1) SetDuration(v float32) {
	o.Duration = v
}

// GetHumanDuration returns the HumanDuration field value
func (o *SyncV1) GetHumanDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HumanDuration
}

// GetHumanDurationOk returns a tuple with the HumanDuration field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetHumanDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HumanDuration, true
}

// SetHumanDuration sets field value
func (o *SyncV1) SetHumanDuration(v string) {
	o.HumanDuration = v
}

// GetCount returns the Count field value
func (o *SyncV1) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *SyncV1) SetCount(v float32) {
	o.Count = v
}

// GetNotices returns the Notices field value
func (o *SyncV1) GetNotices() []SyncNoticeV1 {
	if o == nil {
		var ret []SyncNoticeV1
		return ret
	}

	return o.Notices
}

// GetNoticesOk returns a tuple with the Notices field value
// and a boolean to check if the value has been set.
func (o *SyncV1) GetNoticesOk() ([]SyncNoticeV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notices, true
}

// SetNotices sets field value
func (o *SyncV1) SetNotices(v []SyncNoticeV1) {
	o.Notices = v
}

func (o SyncV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["start"] = o.Start
	toSerialize["end"] = o.End.Get()
	toSerialize["status"] = o.Status
	toSerialize["duration"] = o.Duration
	toSerialize["humanDuration"] = o.HumanDuration
	toSerialize["count"] = o.Count
	toSerialize["notices"] = o.Notices
	return toSerialize, nil
}

type NullableSyncV1 struct {
	value *SyncV1
	isSet bool
}

func (v NullableSyncV1) Get() *SyncV1 {
	return v.value
}

func (v *NullableSyncV1) Set(val *SyncV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncV1(val *SyncV1) *NullableSyncV1 {
	return &NullableSyncV1{value: val, isSet: true}
}

func (v NullableSyncV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
