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

// checks if the DeliveryOverviewMetricsDatapoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewMetricsDatapoint{}

// DeliveryOverviewMetricsDatapoint Series within DeliveryOverviewMetricsDataset.
type DeliveryOverviewMetricsDatapoint struct {
	// The timestamp corresponding to the beginning of the window given by the requested granularity.
	Time string `json:"time"`
	// Holds the number of events within the specified granularity and groupBy options.
	Count float32 `json:"count"`
	// The number of retried events that were successfully delivered.
	RetryCount *float32 `json:"retryCount,omitempty"`
}

// NewDeliveryOverviewMetricsDatapoint instantiates a new DeliveryOverviewMetricsDatapoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewMetricsDatapoint(
	time string,
	count float32,
) *DeliveryOverviewMetricsDatapoint {
	this := DeliveryOverviewMetricsDatapoint{}
	this.Time = time
	this.Count = count
	return &this
}

// NewDeliveryOverviewMetricsDatapointWithDefaults instantiates a new DeliveryOverviewMetricsDatapoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewMetricsDatapointWithDefaults() *DeliveryOverviewMetricsDatapoint {
	this := DeliveryOverviewMetricsDatapoint{}
	return &this
}

// GetTime returns the Time field value
func (o *DeliveryOverviewMetricsDatapoint) GetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDatapoint) GetTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *DeliveryOverviewMetricsDatapoint) SetTime(v string) {
	o.Time = v
}

// GetCount returns the Count field value
func (o *DeliveryOverviewMetricsDatapoint) GetCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDatapoint) GetCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *DeliveryOverviewMetricsDatapoint) SetCount(v float32) {
	o.Count = v
}

// GetRetryCount returns the RetryCount field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDatapoint) GetRetryCount() float32 {
	if o == nil || IsNil(o.RetryCount) {
		var ret float32
		return ret
	}
	return *o.RetryCount
}

// GetRetryCountOk returns a tuple with the RetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDatapoint) GetRetryCountOk() (*float32, bool) {
	if o == nil || IsNil(o.RetryCount) {
		return nil, false
	}
	return o.RetryCount, true
}

// HasRetryCount returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDatapoint) HasRetryCount() bool {
	if o != nil && !IsNil(o.RetryCount) {
		return true
	}

	return false
}

// SetRetryCount gets a reference to the given float32 and assigns it to the RetryCount field.
func (o *DeliveryOverviewMetricsDatapoint) SetRetryCount(v float32) {
	o.RetryCount = &v
}

func (o DeliveryOverviewMetricsDatapoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewMetricsDatapoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	toSerialize["count"] = o.Count
	if !IsNil(o.RetryCount) {
		toSerialize["retryCount"] = o.RetryCount
	}
	return toSerialize, nil
}

type NullableDeliveryOverviewMetricsDatapoint struct {
	value *DeliveryOverviewMetricsDatapoint
	isSet bool
}

func (v NullableDeliveryOverviewMetricsDatapoint) Get() *DeliveryOverviewMetricsDatapoint {
	return v.value
}

func (v *NullableDeliveryOverviewMetricsDatapoint) Set(val *DeliveryOverviewMetricsDatapoint) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewMetricsDatapoint) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewMetricsDatapoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewMetricsDatapoint(
	val *DeliveryOverviewMetricsDatapoint,
) *NullableDeliveryOverviewMetricsDatapoint {
	return &NullableDeliveryOverviewMetricsDatapoint{value: val, isSet: true}
}

func (v NullableDeliveryOverviewMetricsDatapoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewMetricsDatapoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
