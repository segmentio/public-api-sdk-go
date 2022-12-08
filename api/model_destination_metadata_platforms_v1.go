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

// DestinationMetadataPlatformsV1 Represents platforms that a given Destination supports.
type DestinationMetadataPlatformsV1 struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
}

// NewDestinationMetadataPlatformsV1 instantiates a new DestinationMetadataPlatformsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataPlatformsV1() *DestinationMetadataPlatformsV1 {
	this := DestinationMetadataPlatformsV1{}
	return &this
}

// NewDestinationMetadataPlatformsV1WithDefaults instantiates a new DestinationMetadataPlatformsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataPlatformsV1WithDefaults() *DestinationMetadataPlatformsV1 {
	this := DestinationMetadataPlatformsV1{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *DestinationMetadataPlatformsV1) GetBrowser() bool {
	if o == nil || o.Browser == nil {
		var ret bool
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataPlatformsV1) GetBrowserOk() (*bool, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *DestinationMetadataPlatformsV1) HasBrowser() bool {
	if o != nil && o.Browser != nil {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given bool and assigns it to the Browser field.
func (o *DestinationMetadataPlatformsV1) SetBrowser(v bool) {
	o.Browser = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *DestinationMetadataPlatformsV1) GetServer() bool {
	if o == nil || o.Server == nil {
		var ret bool
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataPlatformsV1) GetServerOk() (*bool, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *DestinationMetadataPlatformsV1) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given bool and assigns it to the Server field.
func (o *DestinationMetadataPlatformsV1) SetServer(v bool) {
	o.Server = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *DestinationMetadataPlatformsV1) GetMobile() bool {
	if o == nil || o.Mobile == nil {
		var ret bool
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataPlatformsV1) GetMobileOk() (*bool, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *DestinationMetadataPlatformsV1) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given bool and assigns it to the Mobile field.
func (o *DestinationMetadataPlatformsV1) SetMobile(v bool) {
	o.Mobile = &v
}

func (o DestinationMetadataPlatformsV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}
	if o.Server != nil {
		toSerialize["server"] = o.Server
	}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationMetadataPlatformsV1 struct {
	value *DestinationMetadataPlatformsV1
	isSet bool
}

func (v NullableDestinationMetadataPlatformsV1) Get() *DestinationMetadataPlatformsV1 {
	return v.value
}

func (v *NullableDestinationMetadataPlatformsV1) Set(val *DestinationMetadataPlatformsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataPlatformsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataPlatformsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataPlatformsV1(
	val *DestinationMetadataPlatformsV1,
) *NullableDestinationMetadataPlatformsV1 {
	return &NullableDestinationMetadataPlatformsV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataPlatformsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataPlatformsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
