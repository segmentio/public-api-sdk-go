/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 48.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeliveryOverviewFilterBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewFilterBy{}

// DeliveryOverviewFilterBy The `DeliveryOverviewFilterBy` object is a map of the filterable fields and their values.
type DeliveryOverviewFilterBy struct {
	// A list of strings of discard reasons. Can be used to filter all Source and Destination steps, except for Successfully Received.  See [Discard Record Documentation](https://segment.com/docs/connections/delivery-overview/#troubleshooting) for valid error codes.
	DiscardReason []string `json:"discardReason,omitempty"`
	// A list of strings of event names.
	EventName []string `json:"eventName,omitempty"`
	// A list of strings of event types. Valid options are: `alias`, `group`, `identify`, `page`, `screen`, and `track`.
	EventType []string `json:"eventType,omitempty"`
	// A list of strings of app versions.
	AppVersion []string `json:"appVersion,omitempty"`
}

// NewDeliveryOverviewFilterBy instantiates a new DeliveryOverviewFilterBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewFilterBy() *DeliveryOverviewFilterBy {
	this := DeliveryOverviewFilterBy{}
	return &this
}

// NewDeliveryOverviewFilterByWithDefaults instantiates a new DeliveryOverviewFilterBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewFilterByWithDefaults() *DeliveryOverviewFilterBy {
	this := DeliveryOverviewFilterBy{}
	return &this
}

// GetDiscardReason returns the DiscardReason field value if set, zero value otherwise.
func (o *DeliveryOverviewFilterBy) GetDiscardReason() []string {
	if o == nil || IsNil(o.DiscardReason) {
		var ret []string
		return ret
	}
	return o.DiscardReason
}

// GetDiscardReasonOk returns a tuple with the DiscardReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewFilterBy) GetDiscardReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.DiscardReason) {
		return nil, false
	}
	return o.DiscardReason, true
}

// HasDiscardReason returns a boolean if a field has been set.
func (o *DeliveryOverviewFilterBy) HasDiscardReason() bool {
	if o != nil && !IsNil(o.DiscardReason) {
		return true
	}

	return false
}

// SetDiscardReason gets a reference to the given []string and assigns it to the DiscardReason field.
func (o *DeliveryOverviewFilterBy) SetDiscardReason(v []string) {
	o.DiscardReason = v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DeliveryOverviewFilterBy) GetEventName() []string {
	if o == nil || IsNil(o.EventName) {
		var ret []string
		return ret
	}
	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewFilterBy) GetEventNameOk() ([]string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DeliveryOverviewFilterBy) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given []string and assigns it to the EventName field.
func (o *DeliveryOverviewFilterBy) SetEventName(v []string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DeliveryOverviewFilterBy) GetEventType() []string {
	if o == nil || IsNil(o.EventType) {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewFilterBy) GetEventTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DeliveryOverviewFilterBy) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *DeliveryOverviewFilterBy) SetEventType(v []string) {
	o.EventType = v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *DeliveryOverviewFilterBy) GetAppVersion() []string {
	if o == nil || IsNil(o.AppVersion) {
		var ret []string
		return ret
	}
	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewFilterBy) GetAppVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeliveryOverviewFilterBy) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given []string and assigns it to the AppVersion field.
func (o *DeliveryOverviewFilterBy) SetAppVersion(v []string) {
	o.AppVersion = v
}

func (o DeliveryOverviewFilterBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewFilterBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscardReason) {
		toSerialize["discardReason"] = o.DiscardReason
	}
	if !IsNil(o.EventName) {
		toSerialize["eventName"] = o.EventName
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.AppVersion) {
		toSerialize["appVersion"] = o.AppVersion
	}
	return toSerialize, nil
}

type NullableDeliveryOverviewFilterBy struct {
	value *DeliveryOverviewFilterBy
	isSet bool
}

func (v NullableDeliveryOverviewFilterBy) Get() *DeliveryOverviewFilterBy {
	return v.value
}

func (v *NullableDeliveryOverviewFilterBy) Set(val *DeliveryOverviewFilterBy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewFilterBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewFilterBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewFilterBy(
	val *DeliveryOverviewFilterBy,
) *NullableDeliveryOverviewFilterBy {
	return &NullableDeliveryOverviewFilterBy{value: val, isSet: true}
}

func (v NullableDeliveryOverviewFilterBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewFilterBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
