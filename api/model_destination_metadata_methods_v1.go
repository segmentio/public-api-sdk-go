/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationMetadataMethodsV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationMetadataMethodsV1{}

// DestinationMetadataMethodsV1 Represents methods that a given Destination supports.
type DestinationMetadataMethodsV1 struct {
	// Identifies if the Destination supports the `pageview` method.
	Pageview *bool `json:"pageview,omitempty"`
	// Identifies if the Destination supports the `identify` method.
	Identify *bool `json:"identify,omitempty"`
	// Identifies if the Destination supports the `alias` method.
	Alias *bool `json:"alias,omitempty"`
	// Identifies if the Destination supports the `track` method.
	Track *bool `json:"track,omitempty"`
	// Identifies if the Destination supports the `group` method.
	Group *bool `json:"group,omitempty"`
}

// NewDestinationMetadataMethodsV1 instantiates a new DestinationMetadataMethodsV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataMethodsV1() *DestinationMetadataMethodsV1 {
	this := DestinationMetadataMethodsV1{}
	return &this
}

// NewDestinationMetadataMethodsV1WithDefaults instantiates a new DestinationMetadataMethodsV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataMethodsV1WithDefaults() *DestinationMetadataMethodsV1 {
	this := DestinationMetadataMethodsV1{}
	return &this
}

// GetPageview returns the Pageview field value if set, zero value otherwise.
func (o *DestinationMetadataMethodsV1) GetPageview() bool {
	if o == nil || IsNil(o.Pageview) {
		var ret bool
		return ret
	}
	return *o.Pageview
}

// GetPageviewOk returns a tuple with the Pageview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataMethodsV1) GetPageviewOk() (*bool, bool) {
	if o == nil || IsNil(o.Pageview) {
		return nil, false
	}
	return o.Pageview, true
}

// HasPageview returns a boolean if a field has been set.
func (o *DestinationMetadataMethodsV1) HasPageview() bool {
	if o != nil && !IsNil(o.Pageview) {
		return true
	}

	return false
}

// SetPageview gets a reference to the given bool and assigns it to the Pageview field.
func (o *DestinationMetadataMethodsV1) SetPageview(v bool) {
	o.Pageview = &v
}

// GetIdentify returns the Identify field value if set, zero value otherwise.
func (o *DestinationMetadataMethodsV1) GetIdentify() bool {
	if o == nil || IsNil(o.Identify) {
		var ret bool
		return ret
	}
	return *o.Identify
}

// GetIdentifyOk returns a tuple with the Identify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataMethodsV1) GetIdentifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Identify) {
		return nil, false
	}
	return o.Identify, true
}

// HasIdentify returns a boolean if a field has been set.
func (o *DestinationMetadataMethodsV1) HasIdentify() bool {
	if o != nil && !IsNil(o.Identify) {
		return true
	}

	return false
}

// SetIdentify gets a reference to the given bool and assigns it to the Identify field.
func (o *DestinationMetadataMethodsV1) SetIdentify(v bool) {
	o.Identify = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *DestinationMetadataMethodsV1) GetAlias() bool {
	if o == nil || IsNil(o.Alias) {
		var ret bool
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataMethodsV1) GetAliasOk() (*bool, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *DestinationMetadataMethodsV1) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given bool and assigns it to the Alias field.
func (o *DestinationMetadataMethodsV1) SetAlias(v bool) {
	o.Alias = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *DestinationMetadataMethodsV1) GetTrack() bool {
	if o == nil || IsNil(o.Track) {
		var ret bool
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataMethodsV1) GetTrackOk() (*bool, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *DestinationMetadataMethodsV1) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given bool and assigns it to the Track field.
func (o *DestinationMetadataMethodsV1) SetTrack(v bool) {
	o.Track = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *DestinationMetadataMethodsV1) GetGroup() bool {
	if o == nil || IsNil(o.Group) {
		var ret bool
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataMethodsV1) GetGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *DestinationMetadataMethodsV1) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given bool and assigns it to the Group field.
func (o *DestinationMetadataMethodsV1) SetGroup(v bool) {
	o.Group = &v
}

func (o DestinationMetadataMethodsV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationMetadataMethodsV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pageview) {
		toSerialize["pageview"] = o.Pageview
	}
	if !IsNil(o.Identify) {
		toSerialize["identify"] = o.Identify
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableDestinationMetadataMethodsV1 struct {
	value *DestinationMetadataMethodsV1
	isSet bool
}

func (v NullableDestinationMetadataMethodsV1) Get() *DestinationMetadataMethodsV1 {
	return v.value
}

func (v *NullableDestinationMetadataMethodsV1) Set(val *DestinationMetadataMethodsV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataMethodsV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataMethodsV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataMethodsV1(
	val *DestinationMetadataMethodsV1,
) *NullableDestinationMetadataMethodsV1 {
	return &NullableDestinationMetadataMethodsV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataMethodsV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataMethodsV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
