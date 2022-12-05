/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Regulation The regulate request.
type Regulation struct {
	// The id of the regulate request.
	Id string `json:"id"`
	// The id of the Workspace that the regulate request belongs to.
	WorkspaceId string `json:"workspaceId"`
	// The current status of the regulate request.
	OverallStatus string `json:"overallStatus"`
	// The timestamp of when the request finished.
	FinishedAt string `json:"finishedAt"`
	// The timestamp of the creation of the request.
	CreatedAt string `json:"createdAt"`
	// The status of each stream including all the Destinations that correspond to the stream.
	StreamStatus []StreamStatusV1 `json:"streamStatus"`
}

// NewRegulation instantiates a new Regulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegulation(
	id string,
	workspaceId string,
	overallStatus string,
	finishedAt string,
	createdAt string,
	streamStatus []StreamStatusV1,
) *Regulation {
	this := Regulation{}
	this.Id = id
	this.WorkspaceId = workspaceId
	this.OverallStatus = overallStatus
	this.FinishedAt = finishedAt
	this.CreatedAt = createdAt
	this.StreamStatus = streamStatus
	return &this
}

// NewRegulationWithDefaults instantiates a new Regulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegulationWithDefaults() *Regulation {
	this := Regulation{}
	return &this
}

// GetId returns the Id field value
func (o *Regulation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Regulation) SetId(v string) {
	o.Id = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *Regulation) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *Regulation) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetOverallStatus returns the OverallStatus field value
func (o *Regulation) GetOverallStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OverallStatus
}

// GetOverallStatusOk returns a tuple with the OverallStatus field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetOverallStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverallStatus, true
}

// SetOverallStatus sets field value
func (o *Regulation) SetOverallStatus(v string) {
	o.OverallStatus = v
}

// GetFinishedAt returns the FinishedAt field value
func (o *Regulation) GetFinishedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetFinishedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishedAt, true
}

// SetFinishedAt sets field value
func (o *Regulation) SetFinishedAt(v string) {
	o.FinishedAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Regulation) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Regulation) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetStreamStatus returns the StreamStatus field value
func (o *Regulation) GetStreamStatus() []StreamStatusV1 {
	if o == nil {
		var ret []StreamStatusV1
		return ret
	}

	return o.StreamStatus
}

// GetStreamStatusOk returns a tuple with the StreamStatus field value
// and a boolean to check if the value has been set.
func (o *Regulation) GetStreamStatusOk() ([]StreamStatusV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreamStatus, true
}

// SetStreamStatus sets field value
func (o *Regulation) SetStreamStatus(v []StreamStatusV1) {
	o.StreamStatus = v
}

func (o Regulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if true {
		toSerialize["overallStatus"] = o.OverallStatus
	}
	if true {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["streamStatus"] = o.StreamStatus
	}
	return json.Marshal(toSerialize)
}

type NullableRegulation struct {
	value *Regulation
	isSet bool
}

func (v NullableRegulation) Get() *Regulation {
	return v.value
}

func (v *NullableRegulation) Set(val *Regulation) {
	v.value = val
	v.isSet = true
}

func (v NullableRegulation) IsSet() bool {
	return v.isSet
}

func (v *NullableRegulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegulation(val *Regulation) *NullableRegulation {
	return &NullableRegulation{value: val, isSet: true}
}

func (v NullableRegulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
