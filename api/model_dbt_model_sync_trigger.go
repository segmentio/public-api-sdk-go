/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DbtModelSyncTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbtModelSyncTrigger{}

// DbtModelSyncTrigger Defines the DBT Model Sync Trigger.
type DbtModelSyncTrigger struct {
	// The id of the DBT Model Sync.
	Id string `json:"id"`
	// The Source id that was triggered.
	SourceId *string `json:"sourceId,omitempty"`
	// The status of the trigger.
	Status string `json:"status"`
}

// NewDbtModelSyncTrigger instantiates a new DbtModelSyncTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbtModelSyncTrigger(id string, status string) *DbtModelSyncTrigger {
	this := DbtModelSyncTrigger{}
	this.Id = id
	this.Status = status
	return &this
}

// NewDbtModelSyncTriggerWithDefaults instantiates a new DbtModelSyncTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbtModelSyncTriggerWithDefaults() *DbtModelSyncTrigger {
	this := DbtModelSyncTrigger{}
	return &this
}

// GetId returns the Id field value
func (o *DbtModelSyncTrigger) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DbtModelSyncTrigger) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DbtModelSyncTrigger) SetId(v string) {
	o.Id = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *DbtModelSyncTrigger) GetSourceId() string {
	if o == nil || IsNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbtModelSyncTrigger) GetSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *DbtModelSyncTrigger) HasSourceId() bool {
	if o != nil && !IsNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *DbtModelSyncTrigger) SetSourceId(v string) {
	o.SourceId = &v
}

// GetStatus returns the Status field value
func (o *DbtModelSyncTrigger) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DbtModelSyncTrigger) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DbtModelSyncTrigger) SetStatus(v string) {
	o.Status = v
}

func (o DbtModelSyncTrigger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbtModelSyncTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.SourceId) {
		toSerialize["sourceId"] = o.SourceId
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableDbtModelSyncTrigger struct {
	value *DbtModelSyncTrigger
	isSet bool
}

func (v NullableDbtModelSyncTrigger) Get() *DbtModelSyncTrigger {
	return v.value
}

func (v *NullableDbtModelSyncTrigger) Set(val *DbtModelSyncTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableDbtModelSyncTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableDbtModelSyncTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbtModelSyncTrigger(val *DbtModelSyncTrigger) *NullableDbtModelSyncTrigger {
	return &NullableDbtModelSyncTrigger{value: val, isSet: true}
}

func (v NullableDbtModelSyncTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbtModelSyncTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
