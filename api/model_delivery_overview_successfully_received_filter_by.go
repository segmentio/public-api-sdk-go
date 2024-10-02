/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 55.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeliveryOverviewSuccessfullyReceivedFilterBy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryOverviewSuccessfullyReceivedFilterBy{}

// DeliveryOverviewSuccessfullyReceivedFilterBy The `DeliveryOverviewSuccessfullyReceivedFilterBy` object is a map of the filterable fields and their values for the Successfully Received pipeline step.
type DeliveryOverviewSuccessfullyReceivedFilterBy struct {
	// A list of strings of event names.
	EventName []string `json:"eventName,omitempty"`
	// A list of strings of event types. Valid options are: `alias`, `group`, `identify`, `page`, `screen`, and `track`.
	EventType []string `json:"eventType,omitempty"`
	// A list of strings of app versions.
	AppVersion []string `json:"appVersion,omitempty"`
}

// NewDeliveryOverviewSuccessfullyReceivedFilterBy instantiates a new DeliveryOverviewSuccessfullyReceivedFilterBy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryOverviewSuccessfullyReceivedFilterBy() *DeliveryOverviewSuccessfullyReceivedFilterBy {
	this := DeliveryOverviewSuccessfullyReceivedFilterBy{}
	return &this
}

// NewDeliveryOverviewSuccessfullyReceivedFilterByWithDefaults instantiates a new DeliveryOverviewSuccessfullyReceivedFilterBy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryOverviewSuccessfullyReceivedFilterByWithDefaults() *DeliveryOverviewSuccessfullyReceivedFilterBy {
	this := DeliveryOverviewSuccessfullyReceivedFilterBy{}
	return &this
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetEventName() []string {
	if o == nil || IsNil(o.EventName) {
		var ret []string
		return ret
	}
	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetEventNameOk() ([]string, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given []string and assigns it to the EventName field.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) SetEventName(v []string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetEventType() []string {
	if o == nil || IsNil(o.EventType) {
		var ret []string
		return ret
	}
	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetEventTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given []string and assigns it to the EventType field.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) SetEventType(v []string) {
	o.EventType = v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetAppVersion() []string {
	if o == nil || IsNil(o.AppVersion) {
		var ret []string
		return ret
	}
	return o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) GetAppVersionOk() ([]string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given []string and assigns it to the AppVersion field.
func (o *DeliveryOverviewSuccessfullyReceivedFilterBy) SetAppVersion(v []string) {
	o.AppVersion = v
}

func (o DeliveryOverviewSuccessfullyReceivedFilterBy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliveryOverviewSuccessfullyReceivedFilterBy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableDeliveryOverviewSuccessfullyReceivedFilterBy struct {
	value *DeliveryOverviewSuccessfullyReceivedFilterBy
	isSet bool
}

func (v NullableDeliveryOverviewSuccessfullyReceivedFilterBy) Get() *DeliveryOverviewSuccessfullyReceivedFilterBy {
	return v.value
}

func (v *NullableDeliveryOverviewSuccessfullyReceivedFilterBy) Set(
	val *DeliveryOverviewSuccessfullyReceivedFilterBy,
) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryOverviewSuccessfullyReceivedFilterBy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryOverviewSuccessfullyReceivedFilterBy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryOverviewSuccessfullyReceivedFilterBy(
	val *DeliveryOverviewSuccessfullyReceivedFilterBy,
) *NullableDeliveryOverviewSuccessfullyReceivedFilterBy {
	return &NullableDeliveryOverviewSuccessfullyReceivedFilterBy{value: val, isSet: true}
}

func (v NullableDeliveryOverviewSuccessfullyReceivedFilterBy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryOverviewSuccessfullyReceivedFilterBy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
