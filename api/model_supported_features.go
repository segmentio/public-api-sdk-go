/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 33.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SupportedFeatures Features that this Destination supports.  Config API note: holds `browserUnbundling` fields.
type SupportedFeatures struct {
	// This Destination's support level for cloud mode instances. The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *string `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances. Support for multiple device mode instances is currently not planned. The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *string `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
}

// NewSupportedFeatures instantiates a new SupportedFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedFeatures() *SupportedFeatures {
	this := SupportedFeatures{}
	return &this
}

// NewSupportedFeaturesWithDefaults instantiates a new SupportedFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedFeaturesWithDefaults() *SupportedFeatures {
	this := SupportedFeatures{}
	return &this
}

// GetCloudModeInstances returns the CloudModeInstances field value if set, zero value otherwise.
func (o *SupportedFeatures) GetCloudModeInstances() string {
	if o == nil || o.CloudModeInstances == nil {
		var ret string
		return ret
	}
	return *o.CloudModeInstances
}

// GetCloudModeInstancesOk returns a tuple with the CloudModeInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetCloudModeInstancesOk() (*string, bool) {
	if o == nil || o.CloudModeInstances == nil {
		return nil, false
	}
	return o.CloudModeInstances, true
}

// HasCloudModeInstances returns a boolean if a field has been set.
func (o *SupportedFeatures) HasCloudModeInstances() bool {
	if o != nil && o.CloudModeInstances != nil {
		return true
	}

	return false
}

// SetCloudModeInstances gets a reference to the given string and assigns it to the CloudModeInstances field.
func (o *SupportedFeatures) SetCloudModeInstances(v string) {
	o.CloudModeInstances = &v
}

// GetDeviceModeInstances returns the DeviceModeInstances field value if set, zero value otherwise.
func (o *SupportedFeatures) GetDeviceModeInstances() string {
	if o == nil || o.DeviceModeInstances == nil {
		var ret string
		return ret
	}
	return *o.DeviceModeInstances
}

// GetDeviceModeInstancesOk returns a tuple with the DeviceModeInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetDeviceModeInstancesOk() (*string, bool) {
	if o == nil || o.DeviceModeInstances == nil {
		return nil, false
	}
	return o.DeviceModeInstances, true
}

// HasDeviceModeInstances returns a boolean if a field has been set.
func (o *SupportedFeatures) HasDeviceModeInstances() bool {
	if o != nil && o.DeviceModeInstances != nil {
		return true
	}

	return false
}

// SetDeviceModeInstances gets a reference to the given string and assigns it to the DeviceModeInstances field.
func (o *SupportedFeatures) SetDeviceModeInstances(v string) {
	o.DeviceModeInstances = &v
}

// GetReplay returns the Replay field value if set, zero value otherwise.
func (o *SupportedFeatures) GetReplay() bool {
	if o == nil || o.Replay == nil {
		var ret bool
		return ret
	}
	return *o.Replay
}

// GetReplayOk returns a tuple with the Replay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetReplayOk() (*bool, bool) {
	if o == nil || o.Replay == nil {
		return nil, false
	}
	return o.Replay, true
}

// HasReplay returns a boolean if a field has been set.
func (o *SupportedFeatures) HasReplay() bool {
	if o != nil && o.Replay != nil {
		return true
	}

	return false
}

// SetReplay gets a reference to the given bool and assigns it to the Replay field.
func (o *SupportedFeatures) SetReplay(v bool) {
	o.Replay = &v
}

// GetBrowserUnbundling returns the BrowserUnbundling field value if set, zero value otherwise.
func (o *SupportedFeatures) GetBrowserUnbundling() bool {
	if o == nil || o.BrowserUnbundling == nil {
		var ret bool
		return ret
	}
	return *o.BrowserUnbundling
}

// GetBrowserUnbundlingOk returns a tuple with the BrowserUnbundling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetBrowserUnbundlingOk() (*bool, bool) {
	if o == nil || o.BrowserUnbundling == nil {
		return nil, false
	}
	return o.BrowserUnbundling, true
}

// HasBrowserUnbundling returns a boolean if a field has been set.
func (o *SupportedFeatures) HasBrowserUnbundling() bool {
	if o != nil && o.BrowserUnbundling != nil {
		return true
	}

	return false
}

// SetBrowserUnbundling gets a reference to the given bool and assigns it to the BrowserUnbundling field.
func (o *SupportedFeatures) SetBrowserUnbundling(v bool) {
	o.BrowserUnbundling = &v
}

// GetBrowserUnbundlingPublic returns the BrowserUnbundlingPublic field value if set, zero value otherwise.
func (o *SupportedFeatures) GetBrowserUnbundlingPublic() bool {
	if o == nil || o.BrowserUnbundlingPublic == nil {
		var ret bool
		return ret
	}
	return *o.BrowserUnbundlingPublic
}

// GetBrowserUnbundlingPublicOk returns a tuple with the BrowserUnbundlingPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedFeatures) GetBrowserUnbundlingPublicOk() (*bool, bool) {
	if o == nil || o.BrowserUnbundlingPublic == nil {
		return nil, false
	}
	return o.BrowserUnbundlingPublic, true
}

// HasBrowserUnbundlingPublic returns a boolean if a field has been set.
func (o *SupportedFeatures) HasBrowserUnbundlingPublic() bool {
	if o != nil && o.BrowserUnbundlingPublic != nil {
		return true
	}

	return false
}

// SetBrowserUnbundlingPublic gets a reference to the given bool and assigns it to the BrowserUnbundlingPublic field.
func (o *SupportedFeatures) SetBrowserUnbundlingPublic(v bool) {
	o.BrowserUnbundlingPublic = &v
}

func (o SupportedFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudModeInstances != nil {
		toSerialize["cloudModeInstances"] = o.CloudModeInstances
	}
	if o.DeviceModeInstances != nil {
		toSerialize["deviceModeInstances"] = o.DeviceModeInstances
	}
	if o.Replay != nil {
		toSerialize["replay"] = o.Replay
	}
	if o.BrowserUnbundling != nil {
		toSerialize["browserUnbundling"] = o.BrowserUnbundling
	}
	if o.BrowserUnbundlingPublic != nil {
		toSerialize["browserUnbundlingPublic"] = o.BrowserUnbundlingPublic
	}
	return json.Marshal(toSerialize)
}

type NullableSupportedFeatures struct {
	value *SupportedFeatures
	isSet bool
}

func (v NullableSupportedFeatures) Get() *SupportedFeatures {
	return v.value
}

func (v *NullableSupportedFeatures) Set(val *SupportedFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedFeatures(val *SupportedFeatures) *NullableSupportedFeatures {
	return &NullableSupportedFeatures{value: val, isSet: true}
}

func (v NullableSupportedFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


