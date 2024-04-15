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

// checks if the DeliveryOverviewSourceFilterBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewSourceFilterBy{}

// DeliveryOverviewSourceFilterBy The `DeliveryOverviewSourceFilterBy` object is a map of the filterable fields and their values.
type DeliveryOverviewSourceFilterBy struct {
	// A list of strings of discard reasons.  See [Discard Record Documentation](https://segment.com/docs/connections/delivery-overview/#troubleshooting) for valid error codes.
	DiscardReason []string `json:"discardReason,omitempty"`
	// A list of strings of event names.
	EventName []string `json:"eventName,omitempty"`
	// A list of strings of event types. Valid options are: `alias`, `group`, `identify`, `page`, `screen`, and `track`.
	EventType []string `json:"eventType,omitempty"`
	// A list of strings of app versions.
	AppVersion []string `json:"appVersion,omitempty"`
}

// NewDeliveryOverviewSourceFilterBy instantiates a new DeliveryOverviewSourceFilterBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewSourceFilterBy() *DeliveryOverviewSourceFilterBy {
	this := DeliveryOverviewSourceFilterBy{}
	return &this
}

// NewDeliveryOverviewSourceFilterByWithDefaults instantiates a new DeliveryOverviewSourceFilterBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewSourceFilterByWithDefaults() *DeliveryOverviewSourceFilterBy {
	this := DeliveryOverviewSourceFilterBy{}
	return &this
}

// GetDiscardReason returns the DiscardReason field value if set, zero value otherwise.
func (o *DeliveryOverviewSourceFilterBy) GetDiscardReason() []string {
	if o == nil || IsNil(o.DiscardReason) {
		var ret []string
		return ret
	}
	return o.DiscardReason
}

// GetDiscardReasonOk returns a tuple with the DiscardReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSourceFilterBy) GetDiscardReasonOk() ([]string, bool) {
	if o == nil || IsNil(o.DiscardReason) {
		return nil, false
	}
	return o.DiscardReason, true
}

// HasDiscardReason returns a boolean if a field has been set.
func (o *DeliveryOverviewSourceFilterBy) HasDiscardReason() bool {
	if o != nil && !IsNil(o.DiscardReason) {
		return true
	}

	return false
}

// SetDiscardReason gets a reference to the given []string and assigns it to the DiscardReason field.
func (o *DeliveryOverviewSourceFilterBy) SetDiscardReason(v []string) {
	o.DiscardReason = v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DeliveryOverviewSourceFilterBy) GetEventName() []string {
	if o == nil || IsNil(o.EventName) {
		var ret []string
		return ret
	}
	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSourceFilterBy) GetEventNameOk() ([]string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DeliveryOverviewSourceFilterBy) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given []string and assigns it to the EventName field.
func (o *DeliveryOverviewSourceFilterBy) SetEventName(v []string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DeliveryOverviewSourceFilterBy) GetEventType() []string {
	if o == nil || IsNil(o.EventType) {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSourceFilterBy) GetEventTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DeliveryOverviewSourceFilterBy) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *DeliveryOverviewSourceFilterBy) SetEventType(v []string) {
	o.EventType = v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *DeliveryOverviewSourceFilterBy) GetAppVersion() []string {
	if o == nil || IsNil(o.AppVersion) {
		var ret []string
		return ret
	}
	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSourceFilterBy) GetAppVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeliveryOverviewSourceFilterBy) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given []string and assigns it to the AppVersion field.
func (o *DeliveryOverviewSourceFilterBy) SetAppVersion(v []string) {
	o.AppVersion = v
}

func (o DeliveryOverviewSourceFilterBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewSourceFilterBy) ToMap() (map[string]interface{}, error) {
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

type NullableDeliveryOverviewSourceFilterBy struct {
	value *DeliveryOverviewSourceFilterBy
	isSet bool
}

func (v NullableDeliveryOverviewSourceFilterBy) Get() *DeliveryOverviewSourceFilterBy {
	return v.value
}

func (v *NullableDeliveryOverviewSourceFilterBy) Set(val *DeliveryOverviewSourceFilterBy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewSourceFilterBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewSourceFilterBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewSourceFilterBy(
	val *DeliveryOverviewSourceFilterBy,
) *NullableDeliveryOverviewSourceFilterBy {
	return &NullableDeliveryOverviewSourceFilterBy{value: val, isSet: true}
}

func (v NullableDeliveryOverviewSourceFilterBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewSourceFilterBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
