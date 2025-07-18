/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationMetadataFeaturesV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationMetadataFeaturesV1{}

// DestinationMetadataFeaturesV1 Represents features that a given Destination supports.
type DestinationMetadataFeaturesV1 struct {
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

// NewDestinationMetadataFeaturesV1 instantiates a new DestinationMetadataFeaturesV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataFeaturesV1() *DestinationMetadataFeaturesV1 {
	this := DestinationMetadataFeaturesV1{}
	return &this
}

// NewDestinationMetadataFeaturesV1WithDefaults instantiates a new DestinationMetadataFeaturesV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataFeaturesV1WithDefaults() *DestinationMetadataFeaturesV1 {
	this := DestinationMetadataFeaturesV1{}
	return &this
}

// GetCloudModeInstances returns the CloudModeInstances field value if set, zero value otherwise.
func (o *DestinationMetadataFeaturesV1) GetCloudModeInstances() string {
	if o == nil || IsNil(o.CloudModeInstances) {
		var ret string
		return ret
	}
	return *o.CloudModeInstances
}

// GetCloudModeInstancesOk returns a tuple with the CloudModeInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataFeaturesV1) GetCloudModeInstancesOk() (*string, bool) {
	if o == nil || IsNil(o.CloudModeInstances) {
		return nil, false
	}
	return o.CloudModeInstances, true
}

// HasCloudModeInstances returns a boolean if a field has been set.
func (o *DestinationMetadataFeaturesV1) HasCloudModeInstances() bool {
	if o != nil && !IsNil(o.CloudModeInstances) {
		return true
	}

	return false
}

// SetCloudModeInstances gets a reference to the given string and assigns it to the CloudModeInstances field.
func (o *DestinationMetadataFeaturesV1) SetCloudModeInstances(v string) {
	o.CloudModeInstances = &v
}

// GetDeviceModeInstances returns the DeviceModeInstances field value if set, zero value otherwise.
func (o *DestinationMetadataFeaturesV1) GetDeviceModeInstances() string {
	if o == nil || IsNil(o.DeviceModeInstances) {
		var ret string
		return ret
	}
	return *o.DeviceModeInstances
}

// GetDeviceModeInstancesOk returns a tuple with the DeviceModeInstances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataFeaturesV1) GetDeviceModeInstancesOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModeInstances) {
		return nil, false
	}
	return o.DeviceModeInstances, true
}

// HasDeviceModeInstances returns a boolean if a field has been set.
func (o *DestinationMetadataFeaturesV1) HasDeviceModeInstances() bool {
	if o != nil && !IsNil(o.DeviceModeInstances) {
		return true
	}

	return false
}

// SetDeviceModeInstances gets a reference to the given string and assigns it to the DeviceModeInstances field.
func (o *DestinationMetadataFeaturesV1) SetDeviceModeInstances(v string) {
	o.DeviceModeInstances = &v
}

// GetReplay returns the Replay field value if set, zero value otherwise.
func (o *DestinationMetadataFeaturesV1) GetReplay() bool {
	if o == nil || IsNil(o.Replay) {
		var ret bool
		return ret
	}
	return *o.Replay
}

// GetReplayOk returns a tuple with the Replay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataFeaturesV1) GetReplayOk() (*bool, bool) {
	if o == nil || IsNil(o.Replay) {
		return nil, false
	}
	return o.Replay, true
}

// HasReplay returns a boolean if a field has been set.
func (o *DestinationMetadataFeaturesV1) HasReplay() bool {
	if o != nil && !IsNil(o.Replay) {
		return true
	}

	return false
}

// SetReplay gets a reference to the given bool and assigns it to the Replay field.
func (o *DestinationMetadataFeaturesV1) SetReplay(v bool) {
	o.Replay = &v
}

// GetBrowserUnbundling returns the BrowserUnbundling field value if set, zero value otherwise.
func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundling() bool {
	if o == nil || IsNil(o.BrowserUnbundling) {
		var ret bool
		return ret
	}
	return *o.BrowserUnbundling
}

// GetBrowserUnbundlingOk returns a tuple with the BrowserUnbundling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingOk() (*bool, bool) {
	if o == nil || IsNil(o.BrowserUnbundling) {
		return nil, false
	}
	return o.BrowserUnbundling, true
}

// HasBrowserUnbundling returns a boolean if a field has been set.
func (o *DestinationMetadataFeaturesV1) HasBrowserUnbundling() bool {
	if o != nil && !IsNil(o.BrowserUnbundling) {
		return true
	}

	return false
}

// SetBrowserUnbundling gets a reference to the given bool and assigns it to the BrowserUnbundling field.
func (o *DestinationMetadataFeaturesV1) SetBrowserUnbundling(v bool) {
	o.BrowserUnbundling = &v
}

// GetBrowserUnbundlingPublic returns the BrowserUnbundlingPublic field value if set, zero value otherwise.
func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublic() bool {
	if o == nil || IsNil(o.BrowserUnbundlingPublic) {
		var ret bool
		return ret
	}
	return *o.BrowserUnbundlingPublic
}

// GetBrowserUnbundlingPublicOk returns a tuple with the BrowserUnbundlingPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataFeaturesV1) GetBrowserUnbundlingPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.BrowserUnbundlingPublic) {
		return nil, false
	}
	return o.BrowserUnbundlingPublic, true
}

// HasBrowserUnbundlingPublic returns a boolean if a field has been set.
func (o *DestinationMetadataFeaturesV1) HasBrowserUnbundlingPublic() bool {
	if o != nil && !IsNil(o.BrowserUnbundlingPublic) {
		return true
	}

	return false
}

// SetBrowserUnbundlingPublic gets a reference to the given bool and assigns it to the BrowserUnbundlingPublic field.
func (o *DestinationMetadataFeaturesV1) SetBrowserUnbundlingPublic(v bool) {
	o.BrowserUnbundlingPublic = &v
}

func (o DestinationMetadataFeaturesV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationMetadataFeaturesV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudModeInstances) {
		toSerialize["cloudModeInstances"] = o.CloudModeInstances
	}
	if !IsNil(o.DeviceModeInstances) {
		toSerialize["deviceModeInstances"] = o.DeviceModeInstances
	}
	if !IsNil(o.Replay) {
		toSerialize["replay"] = o.Replay
	}
	if !IsNil(o.BrowserUnbundling) {
		toSerialize["browserUnbundling"] = o.BrowserUnbundling
	}
	if !IsNil(o.BrowserUnbundlingPublic) {
		toSerialize["browserUnbundlingPublic"] = o.BrowserUnbundlingPublic
	}
	return toSerialize, nil
}

type NullableDestinationMetadataFeaturesV1 struct {
	value *DestinationMetadataFeaturesV1
	isSet bool
}

func (v NullableDestinationMetadataFeaturesV1) Get() *DestinationMetadataFeaturesV1 {
	return v.value
}

func (v *NullableDestinationMetadataFeaturesV1) Set(val *DestinationMetadataFeaturesV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataFeaturesV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataFeaturesV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataFeaturesV1(
	val *DestinationMetadataFeaturesV1,
) *NullableDestinationMetadataFeaturesV1 {
	return &NullableDestinationMetadataFeaturesV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataFeaturesV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataFeaturesV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
