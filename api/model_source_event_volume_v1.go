/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SourceEventVolumeV1 SourceEventVolume represents a time series of event volume for a Workspace broken down by the dimensions which the customer specifies (optional parameters).
type SourceEventVolumeV1 struct {
	Source Source `json:"source"`
	// The name of the event, if applicable.
	EventName *string `json:"eventName,omitempty"`
	// The event type, if applicable.
	EventType *string `json:"eventType,omitempty"`
	// The total count for all events in the requested time frame.
	Total float32 `json:"total"`
	// A list of the event counts broken down by the requested granularity.
	Series []SourceEventVolumeDatapointV1 `json:"series"`
}

// NewSourceEventVolumeV1 instantiates a new SourceEventVolumeV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceEventVolumeV1(
	source Source,
	total float32,
	series []SourceEventVolumeDatapointV1,
) *SourceEventVolumeV1 {
	this := SourceEventVolumeV1{}
	this.Source = source
	this.Total = total
	this.Series = series
	return &this
}

// NewSourceEventVolumeV1WithDefaults instantiates a new SourceEventVolumeV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceEventVolumeV1WithDefaults() *SourceEventVolumeV1 {
	this := SourceEventVolumeV1{}
	return &this
}

// GetSource returns the Source field value
func (o *SourceEventVolumeV1) GetSource() Source {
	if o == nil {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeV1) GetSourceOk() (*Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SourceEventVolumeV1) SetSource(v Source) {
	o.Source = v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *SourceEventVolumeV1) GetEventName() string {
	if o == nil || o.EventName == nil {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeV1) GetEventNameOk() (*string, bool) {
	if o == nil || o.EventName == nil {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *SourceEventVolumeV1) HasEventName() bool {
	if o != nil && o.EventName != nil {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *SourceEventVolumeV1) SetEventName(v string) {
	o.EventName = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *SourceEventVolumeV1) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeV1) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *SourceEventVolumeV1) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *SourceEventVolumeV1) SetEventType(v string) {
	o.EventType = &v
}

// GetTotal returns the Total field value
func (o *SourceEventVolumeV1) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeV1) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SourceEventVolumeV1) SetTotal(v float32) {
	o.Total = v
}

// GetSeries returns the Series field value
func (o *SourceEventVolumeV1) GetSeries() []SourceEventVolumeDatapointV1 {
	if o == nil {
		var ret []SourceEventVolumeDatapointV1
		return ret
	}

	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *SourceEventVolumeV1) GetSeriesOk() ([]SourceEventVolumeDatapointV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Series, true
}

// SetSeries sets field value
func (o *SourceEventVolumeV1) SetSeries(v []SourceEventVolumeDatapointV1) {
	o.Series = v
}

func (o SourceEventVolumeV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.EventName != nil {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["series"] = o.Series
	}
	return json.Marshal(toSerialize)
}

type NullableSourceEventVolumeV1 struct {
	value *SourceEventVolumeV1
	isSet bool
}

func (v NullableSourceEventVolumeV1) Get() *SourceEventVolumeV1 {
	return v.value
}

func (v *NullableSourceEventVolumeV1) Set(val *SourceEventVolumeV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceEventVolumeV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceEventVolumeV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceEventVolumeV1(val *SourceEventVolumeV1) *NullableSourceEventVolumeV1 {
	return &NullableSourceEventVolumeV1{value: val, isSet: true}
}

func (v NullableSourceEventVolumeV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceEventVolumeV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
