/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SyncLoadPhase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncLoadPhase{}

// SyncLoadPhase Object representing the load phase + details.
type SyncLoadPhase struct {
	// Counts the subset of records successfully delivered to the downstream Destination.
	DeliverSuccessCount string `json:"deliverSuccessCount"`
	// Counts the subset of records status that failed to be delivered by either receiving a permanent error (for example, 400 Bad Request) or by exhausting all retries for temporary errors (for example, 429 Too Many Requests).
	DeliverFailureCount string `json:"deliverFailureCount"`
	// Error code indicates a fatal sync error code, if applicable.
	ErrorCode string `json:"errorCode"`
	// Time that the load phase started.
	StartedAt string `json:"startedAt"`
	// Time that the load phase finished.
	FinishedAt string `json:"finishedAt"`
}

// NewSyncLoadPhase instantiates a new SyncLoadPhase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncLoadPhase(
	deliverSuccessCount string,
	deliverFailureCount string,
	errorCode string,
	startedAt string,
	finishedAt string,
) *SyncLoadPhase {
	this := SyncLoadPhase{}
	this.DeliverSuccessCount = deliverSuccessCount
	this.DeliverFailureCount = deliverFailureCount
	this.ErrorCode = errorCode
	this.StartedAt = startedAt
	this.FinishedAt = finishedAt
	return &this
}

// NewSyncLoadPhaseWithDefaults instantiates a new SyncLoadPhase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncLoadPhaseWithDefaults() *SyncLoadPhase {
	this := SyncLoadPhase{}
	return &this
}

// GetDeliverSuccessCount returns the DeliverSuccessCount field value
func (o *SyncLoadPhase) GetDeliverSuccessCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliverSuccessCount
}

// GetDeliverSuccessCountOk returns a tuple with the DeliverSuccessCount field value
// and a boolean to check if the value has been set.
func (o *SyncLoadPhase) GetDeliverSuccessCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliverSuccessCount, true
}

// SetDeliverSuccessCount sets field value
func (o *SyncLoadPhase) SetDeliverSuccessCount(v string) {
	o.DeliverSuccessCount = v
}

// GetDeliverFailureCount returns the DeliverFailureCount field value
func (o *SyncLoadPhase) GetDeliverFailureCount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeliverFailureCount
}

// GetDeliverFailureCountOk returns a tuple with the DeliverFailureCount field value
// and a boolean to check if the value has been set.
func (o *SyncLoadPhase) GetDeliverFailureCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliverFailureCount, true
}

// SetDeliverFailureCount sets field value
func (o *SyncLoadPhase) SetDeliverFailureCount(v string) {
	o.DeliverFailureCount = v
}

// GetErrorCode returns the ErrorCode field value
func (o *SyncLoadPhase) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *SyncLoadPhase) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *SyncLoadPhase) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetStartedAt returns the StartedAt field value
func (o *SyncLoadPhase) GetStartedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *SyncLoadPhase) GetStartedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *SyncLoadPhase) SetStartedAt(v string) {
	o.StartedAt = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *SyncLoadPhase) GetFinishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *SyncLoadPhase) GetFinishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *SyncLoadPhase) SetFinishedAt(v string) {
	o.FinishedAt = v
}

func (o SyncLoadPhase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncLoadPhase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deliverSuccessCount"] = o.DeliverSuccessCount
	toSerialize["deliverFailureCount"] = o.DeliverFailureCount
	toSerialize["errorCode"] = o.ErrorCode
	toSerialize["startedAt"] = o.StartedAt
	toSerialize["finishedAt"] = o.FinishedAt
	return toSerialize, nil
}

type NullableSyncLoadPhase struct {
	value *SyncLoadPhase
	isSet bool
}

func (v NullableSyncLoadPhase) Get() *SyncLoadPhase {
	return v.value
}

func (v *NullableSyncLoadPhase) Set(val *SyncLoadPhase) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncLoadPhase) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncLoadPhase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncLoadPhase(val *SyncLoadPhase) *NullableSyncLoadPhase {
	return &NullableSyncLoadPhase{value: val, isSet: true}
}

func (v NullableSyncLoadPhase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncLoadPhase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
