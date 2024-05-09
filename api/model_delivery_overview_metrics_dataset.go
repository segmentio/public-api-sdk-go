/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 49.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeliveryOverviewMetricsDataset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewMetricsDataset{}

// DeliveryOverviewMetricsDataset Dataset within GetDeliveryOverviewMetricsBetaOutput.
type DeliveryOverviewMetricsDataset struct {
	// The name of the event if groupBy[] included 'eventName' in the request.
	EventName *string `json:"eventName,omitempty"`
	// The version of the app if groupBy[] included 'appVersion' in the request.
	AppVersion *string `json:"appVersion,omitempty"`
	// The event type if groupBy[] included 'eventType' in the request.
	EventType *string `json:"eventType,omitempty"`
	// The discard reason for dropped events if groupBy[] included 'discardReason' in the request.
	DiscardReason *string `json:"discardReason,omitempty"`
	// Holds the count of all event counts over the timeframe of the series.
	Total float32 `json:"total"`
	// A list of the event counts broken down by the requested granularity, timeframe, and groupBy options.
	Series []DeliveryOverviewMetricsDatapoint `json:"series"`
	// The number of events successfully delivered upon retry.
	TotalRetryCount *float32 `json:"totalRetryCount,omitempty"`
}

// NewDeliveryOverviewMetricsDataset instantiates a new DeliveryOverviewMetricsDataset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewMetricsDataset(
	total float32,
	series []DeliveryOverviewMetricsDatapoint,
) *DeliveryOverviewMetricsDataset {
	this := DeliveryOverviewMetricsDataset{}
	this.Total = total
	this.Series = series
	return &this
}

// NewDeliveryOverviewMetricsDatasetWithDefaults instantiates a new DeliveryOverviewMetricsDataset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewMetricsDatasetWithDefaults() *DeliveryOverviewMetricsDataset {
	this := DeliveryOverviewMetricsDataset{}
	return &this
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDataset) GetEventName() string {
	if o == nil || IsNil(o.EventName) {
		var ret string
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetEventNameOk() (*string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDataset) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given string and assigns it to the EventName field.
func (o *DeliveryOverviewMetricsDataset) SetEventName(v string) {
	o.EventName = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDataset) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDataset) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *DeliveryOverviewMetricsDataset) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDataset) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDataset) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *DeliveryOverviewMetricsDataset) SetEventType(v string) {
	o.EventType = &v
}

// GetDiscardReason returns the DiscardReason field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDataset) GetDiscardReason() string {
	if o == nil || IsNil(o.DiscardReason) {
		var ret string
		return ret
	}
	return *o.DiscardReason
}

// GetDiscardReasonOk returns a tuple with the DiscardReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetDiscardReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DiscardReason) {
		return nil, false
	}
	return o.DiscardReason, true
}

// HasDiscardReason returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDataset) HasDiscardReason() bool {
	if o != nil && !IsNil(o.DiscardReason) {
		return true
	}

	return false
}

// SetDiscardReason gets a reference to the given string and assigns it to the DiscardReason field.
func (o *DeliveryOverviewMetricsDataset) SetDiscardReason(v string) {
	o.DiscardReason = &v
}

// GetTotal returns the Total field value
func (o *DeliveryOverviewMetricsDataset) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *DeliveryOverviewMetricsDataset) SetTotal(v float32) {
	o.Total = v
}

// GetSeries returns the Series field value
func (o *DeliveryOverviewMetricsDataset) GetSeries() []DeliveryOverviewMetricsDatapoint {
	if o == nil {
		var ret []DeliveryOverviewMetricsDatapoint
		return ret
	}

	return o.Series
}

// GetSeriesOk returns a tuple with the Series field value
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetSeriesOk() ([]DeliveryOverviewMetricsDatapoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Series, true
}

// SetSeries sets field value
func (o *DeliveryOverviewMetricsDataset) SetSeries(v []DeliveryOverviewMetricsDatapoint) {
	o.Series = v
}

// GetTotalRetryCount returns the TotalRetryCount field value if set, zero value otherwise.
func (o *DeliveryOverviewMetricsDataset) GetTotalRetryCount() float32 {
	if o == nil || IsNil(o.TotalRetryCount) {
		var ret float32
		return ret
	}
	return *o.TotalRetryCount
}

// GetTotalRetryCountOk returns a tuple with the TotalRetryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewMetricsDataset) GetTotalRetryCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalRetryCount) {
		return nil, false
	}
	return o.TotalRetryCount, true
}

// HasTotalRetryCount returns a boolean if a field has been set.
func (o *DeliveryOverviewMetricsDataset) HasTotalRetryCount() bool {
	if o != nil && !IsNil(o.TotalRetryCount) {
		return true
	}

	return false
}

// SetTotalRetryCount gets a reference to the given float32 and assigns it to the TotalRetryCount field.
func (o *DeliveryOverviewMetricsDataset) SetTotalRetryCount(v float32) {
	o.TotalRetryCount = &v
}

func (o DeliveryOverviewMetricsDataset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewMetricsDataset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.AppVersion) {
		toSerialize["appVersion"] = o.AppVersion
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.DiscardReason) {
		toSerialize["discardReason"] = o.DiscardReason
	}
	toSerialize["total"] = o.Total
	toSerialize["series"] = o.Series
	if !IsNil(o.TotalRetryCount) {
		toSerialize["totalRetryCount"] = o.TotalRetryCount
	}
	return toSerialize, nil
}

type NullableDeliveryOverviewMetricsDataset struct {
	value *DeliveryOverviewMetricsDataset
	isSet bool
}

func (v NullableDeliveryOverviewMetricsDataset) Get() *DeliveryOverviewMetricsDataset {
	return v.value
}

func (v *NullableDeliveryOverviewMetricsDataset) Set(val *DeliveryOverviewMetricsDataset) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewMetricsDataset) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewMetricsDataset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewMetricsDataset(
	val *DeliveryOverviewMetricsDataset,
) *NullableDeliveryOverviewMetricsDataset {
	return &NullableDeliveryOverviewMetricsDataset{value: val, isSet: true}
}

func (v NullableDeliveryOverviewMetricsDataset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewMetricsDataset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
