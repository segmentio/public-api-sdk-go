/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SyncExtractPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncExtractPhase{}

// SyncExtractPhase Object representing the extract phase + details.
type SyncExtractPhase struct {
	// Counts the subset of records with status=new, which indicates records that were created/inserted/added.
	AddedCount string `json:"addedCount"`
	// Counts the subset of records with status=updated, which indicates records that were modified/updated.
	UpdatedCount string `json:"updatedCount"`
	// Counts the subset of records with status=deleted, which indicates records that were deleted/removed.
	DeletedCount string `json:"deletedCount"`
	// Counts the total number of records/rows handled by extract, across all statuses.
	ExtractCount string `json:"extractCount"`
	// Error code indicates a fatal sync error code, if applicable.
	ErrorCode string `json:"errorCode"`
	// Time that the extract phase started.
	StartedAt string `json:"startedAt"`
	// Time that the extract phase finished.
	FinishedAt string `json:"finishedAt"`
}

// NewSyncExtractPhase instantiates a new SyncExtractPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncExtractPhase(
	addedCount string,
	updatedCount string,
	deletedCount string,
	extractCount string,
	errorCode string,
	startedAt string,
	finishedAt string,
) *SyncExtractPhase {
	this := SyncExtractPhase{}
	this.AddedCount = addedCount
	this.UpdatedCount = updatedCount
	this.DeletedCount = deletedCount
	this.ExtractCount = extractCount
	this.ErrorCode = errorCode
	this.StartedAt = startedAt
	this.FinishedAt = finishedAt
	return &this
}

// NewSyncExtractPhaseWithDefaults instantiates a new SyncExtractPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncExtractPhaseWithDefaults() *SyncExtractPhase {
	this := SyncExtractPhase{}
	return &this
}

// GetAddedCount returns the AddedCount field value
func (o *SyncExtractPhase) GetAddedCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddedCount
}

// GetAddedCountOk returns a tuple with the AddedCount field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetAddedCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddedCount, true
}

// SetAddedCount sets field value
func (o *SyncExtractPhase) SetAddedCount(v string) {
	o.AddedCount = v
}

// GetUpdatedCount returns the UpdatedCount field value
func (o *SyncExtractPhase) GetUpdatedCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedCount
}

// GetUpdatedCountOk returns a tuple with the UpdatedCount field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetUpdatedCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedCount, true
}

// SetUpdatedCount sets field value
func (o *SyncExtractPhase) SetUpdatedCount(v string) {
	o.UpdatedCount = v
}

// GetDeletedCount returns the DeletedCount field value
func (o *SyncExtractPhase) GetDeletedCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedCount
}

// GetDeletedCountOk returns a tuple with the DeletedCount field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetDeletedCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedCount, true
}

// SetDeletedCount sets field value
func (o *SyncExtractPhase) SetDeletedCount(v string) {
	o.DeletedCount = v
}

// GetExtractCount returns the ExtractCount field value
func (o *SyncExtractPhase) GetExtractCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtractCount
}

// GetExtractCountOk returns a tuple with the ExtractCount field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetExtractCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtractCount, true
}

// SetExtractCount sets field value
func (o *SyncExtractPhase) SetExtractCount(v string) {
	o.ExtractCount = v
}

// GetErrorCode returns the ErrorCode field value
func (o *SyncExtractPhase) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *SyncExtractPhase) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetStartedAt returns the StartedAt field value
func (o *SyncExtractPhase) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *SyncExtractPhase) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *SyncExtractPhase) GetFinishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *SyncExtractPhase) GetFinishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *SyncExtractPhase) SetFinishedAt(v string) {
	o.FinishedAt = v
}

func (o SyncExtractPhase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncExtractPhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addedCount"] = o.AddedCount
	toSerialize["updatedCount"] = o.UpdatedCount
	toSerialize["deletedCount"] = o.DeletedCount
	toSerialize["extractCount"] = o.ExtractCount
	toSerialize["errorCode"] = o.ErrorCode
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["finishedAt"] = o.FinishedAt
	return toSerialize, nil
}

type NullableSyncExtractPhase struct {
	value *SyncExtractPhase
	isSet bool
}

func (v NullableSyncExtractPhase) Get() *SyncExtractPhase {
	return v.value
}

func (v *NullableSyncExtractPhase) Set(val *SyncExtractPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncExtractPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncExtractPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncExtractPhase(val *SyncExtractPhase) *NullableSyncExtractPhase {
	return &NullableSyncExtractPhase{value: val, isSet: true}
}

func (v NullableSyncExtractPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncExtractPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
