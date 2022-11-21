/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.2
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SupportedPlatforms Platforms from which the Destination receives events.  Config API note: equal to `platforms`.
type SupportedPlatforms struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
}

// NewSupportedPlatforms instantiates a new SupportedPlatforms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedPlatforms() *SupportedPlatforms {
	this := SupportedPlatforms{}
	return &this
}

// NewSupportedPlatformsWithDefaults instantiates a new SupportedPlatforms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedPlatformsWithDefaults() *SupportedPlatforms {
	this := SupportedPlatforms{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *SupportedPlatforms) GetBrowser() bool {
	if o == nil || o.Browser == nil {
		var ret bool
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPlatforms) GetBrowserOk() (*bool, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *SupportedPlatforms) HasBrowser() bool {
	if o != nil && o.Browser != nil {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given bool and assigns it to the Browser field.
func (o *SupportedPlatforms) SetBrowser(v bool) {
	o.Browser = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *SupportedPlatforms) GetServer() bool {
	if o == nil || o.Server == nil {
		var ret bool
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPlatforms) GetServerOk() (*bool, bool) {
	if o == nil || o.Server == nil {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *SupportedPlatforms) HasServer() bool {
	if o != nil && o.Server != nil {
		return true
	}

	return false
}

// SetServer gets a reference to the given bool and assigns it to the Server field.
func (o *SupportedPlatforms) SetServer(v bool) {
	o.Server = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *SupportedPlatforms) GetMobile() bool {
	if o == nil || o.Mobile == nil {
		var ret bool
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedPlatforms) GetMobileOk() (*bool, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *SupportedPlatforms) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given bool and assigns it to the Mobile field.
func (o *SupportedPlatforms) SetMobile(v bool) {
	o.Mobile = &v
}

func (o SupportedPlatforms) MarshalJSON() ([]byte, error) {
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

type NullableSupportedPlatforms struct {
	value *SupportedPlatforms
	isSet bool
}

func (v NullableSupportedPlatforms) Get() *SupportedPlatforms {
	return v.value
}

func (v *NullableSupportedPlatforms) Set(val *SupportedPlatforms) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedPlatforms) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedPlatforms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedPlatforms(val *SupportedPlatforms) *NullableSupportedPlatforms {
	return &NullableSupportedPlatforms{value: val, isSet: true}
}

func (v NullableSupportedPlatforms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedPlatforms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


