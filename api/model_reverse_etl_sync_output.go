/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 46.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ReverseETLSyncOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReverseETLSyncOutput{}

// ReverseETLSyncOutput Defines the result of getting the sync status of a RETL connection.
type ReverseETLSyncOutput struct {
	// The id of the sync.
	SyncId string `json:"syncId"`
	// The Model id of the sync.
	ModelId string `json:"modelId"`
	// The Source id of the sync.
	SourceId string `json:"sourceId"`
	// The status of the sync. It currently can be IN_PROGRESS, FAIL, SUCCESS, PARTIAL_SUCCESS.
	SyncStatus string `json:"syncStatus"`
	// The duration of the sync.
	Duration string `json:"duration"`
	// When the sync started.
	StartedAt string `json:"startedAt"`
	// When the sync started.
	FinishedAt   *string           `json:"finishedAt,omitempty"`
	ExtractPhase *SyncExtractPhase `json:"extractPhase,omitempty"`
	LoadPhase    *SyncLoadPhase    `json:"loadPhase,omitempty"`
	// Error message if applicable.
	Error *string `json:"error,omitempty"`
	// Error code indicates a fatal sync error code, if applicable.
	ErrorCode *string `json:"errorCode,omitempty"`
}

// NewReverseETLSyncOutput instantiates a new ReverseETLSyncOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReverseETLSyncOutput(
	syncId string,
	modelId string,
	sourceId string,
	syncStatus string,
	duration string,
	startedAt string,
) *ReverseETLSyncOutput {
	this := ReverseETLSyncOutput{}
	this.SyncId = syncId
	this.ModelId = modelId
	this.SourceId = sourceId
	this.SyncStatus = syncStatus
	this.Duration = duration
	this.StartedAt = startedAt
	return &this
}

// NewReverseETLSyncOutputWithDefaults instantiates a new ReverseETLSyncOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReverseETLSyncOutputWithDefaults() *ReverseETLSyncOutput {
	this := ReverseETLSyncOutput{}
	return &this
}

// GetSyncId returns the SyncId field value
func (o *ReverseETLSyncOutput) GetSyncId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncId
}

// GetSyncIdOk returns a tuple with the SyncId field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetSyncIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncId, true
}

// SetSyncId sets field value
func (o *ReverseETLSyncOutput) SetSyncId(v string) {
	o.SyncId = v
}

// GetModelId returns the ModelId field value
func (o *ReverseETLSyncOutput) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *ReverseETLSyncOutput) SetModelId(v string) {
	o.ModelId = v
}

// GetSourceId returns the SourceId field value
func (o *ReverseETLSyncOutput) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *ReverseETLSyncOutput) SetSourceId(v string) {
	o.SourceId = v
}

// GetSyncStatus returns the SyncStatus field value
func (o *ReverseETLSyncOutput) GetSyncStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetSyncStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncStatus, true
}

// SetSyncStatus sets field value
func (o *ReverseETLSyncOutput) SetSyncStatus(v string) {
	o.SyncStatus = v
}

// GetDuration returns the Duration field value
func (o *ReverseETLSyncOutput) GetDuration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetDurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *ReverseETLSyncOutput) SetDuration(v string) {
	o.Duration = v
}

// GetStartedAt returns the StartedAt field value
func (o *ReverseETLSyncOutput) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *ReverseETLSyncOutput) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *ReverseETLSyncOutput) GetFinishedAt() string {
	if o == nil || IsNil(o.FinishedAt) {
		var ret string
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetFinishedAtOk() (*string, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ReverseETLSyncOutput) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given string and assigns it to the FinishedAt field.
func (o *ReverseETLSyncOutput) SetFinishedAt(v string) {
	o.FinishedAt = &v
}

// GetExtractPhase returns the ExtractPhase field value if set, zero value otherwise.
func (o *ReverseETLSyncOutput) GetExtractPhase() SyncExtractPhase {
	if o == nil || IsNil(o.ExtractPhase) {
		var ret SyncExtractPhase
		return ret
	}
	return *o.ExtractPhase
}

// GetExtractPhaseOk returns a tuple with the ExtractPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetExtractPhaseOk() (*SyncExtractPhase, bool) {
	if o == nil || IsNil(o.ExtractPhase) {
		return nil, false
	}
	return o.ExtractPhase, true
}

// HasExtractPhase returns a boolean if a field has been set.
func (o *ReverseETLSyncOutput) HasExtractPhase() bool {
	if o != nil && !IsNil(o.ExtractPhase) {
		return true
	}

	return false
}

// SetExtractPhase gets a reference to the given SyncExtractPhase and assigns it to the ExtractPhase field.
func (o *ReverseETLSyncOutput) SetExtractPhase(v SyncExtractPhase) {
	o.ExtractPhase = &v
}

// GetLoadPhase returns the LoadPhase field value if set, zero value otherwise.
func (o *ReverseETLSyncOutput) GetLoadPhase() SyncLoadPhase {
	if o == nil || IsNil(o.LoadPhase) {
		var ret SyncLoadPhase
		return ret
	}
	return *o.LoadPhase
}

// GetLoadPhaseOk returns a tuple with the LoadPhase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetLoadPhaseOk() (*SyncLoadPhase, bool) {
	if o == nil || IsNil(o.LoadPhase) {
		return nil, false
	}
	return o.LoadPhase, true
}

// HasLoadPhase returns a boolean if a field has been set.
func (o *ReverseETLSyncOutput) HasLoadPhase() bool {
	if o != nil && !IsNil(o.LoadPhase) {
		return true
	}

	return false
}

// SetLoadPhase gets a reference to the given SyncLoadPhase and assigns it to the LoadPhase field.
func (o *ReverseETLSyncOutput) SetLoadPhase(v SyncLoadPhase) {
	o.LoadPhase = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ReverseETLSyncOutput) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ReverseETLSyncOutput) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ReverseETLSyncOutput) SetError(v string) {
	o.Error = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *ReverseETLSyncOutput) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReverseETLSyncOutput) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ReverseETLSyncOutput) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ReverseETLSyncOutput) SetErrorCode(v string) {
	o.ErrorCode = &v
}

func (o ReverseETLSyncOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReverseETLSyncOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["syncId"] = o.SyncId
	toSerialize["modelId"] = o.ModelId
	toSerialize["sourceId"] = o.SourceId
	toSerialize["syncStatus"] = o.SyncStatus
	toSerialize["duration"] = o.Duration
	toSerialize["startedAt"] = o.StartedAt
	if !IsNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !IsNil(o.ExtractPhase) {
		toSerialize["extractPhase"] = o.ExtractPhase
	}
	if !IsNil(o.LoadPhase) {
		toSerialize["loadPhase"] = o.LoadPhase
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	return toSerialize, nil
}

type NullableReverseETLSyncOutput struct {
	value *ReverseETLSyncOutput
	isSet bool
}

func (v NullableReverseETLSyncOutput) Get() *ReverseETLSyncOutput {
	return v.value
}

func (v *NullableReverseETLSyncOutput) Set(val *ReverseETLSyncOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseETLSyncOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseETLSyncOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseETLSyncOutput(val *ReverseETLSyncOutput) *NullableReverseETLSyncOutput {
	return &NullableReverseETLSyncOutput{value: val, isSet: true}
}

func (v NullableReverseETLSyncOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseETLSyncOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
