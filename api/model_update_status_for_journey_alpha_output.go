/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateStatusForJourneyAlphaOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStatusForJourneyAlphaOutput{}

// UpdateStatusForJourneyAlphaOutput Update journey status output.
type UpdateStatusForJourneyAlphaOutput struct {
	// The journey container id.
	ContainerId string `json:"containerId"`
	// The journey version.
	Version float32 `json:"version"`
	// The status of the journey.
	Status string `json:"status"`
}

// NewUpdateStatusForJourneyAlphaOutput instantiates a new UpdateStatusForJourneyAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStatusForJourneyAlphaOutput(
	containerId string,
	version float32,
	status string,
) *UpdateStatusForJourneyAlphaOutput {
	this := UpdateStatusForJourneyAlphaOutput{}
	this.ContainerId = containerId
	this.Version = version
	this.Status = status
	return &this
}

// NewUpdateStatusForJourneyAlphaOutputWithDefaults instantiates a new UpdateStatusForJourneyAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStatusForJourneyAlphaOutputWithDefaults() *UpdateStatusForJourneyAlphaOutput {
	this := UpdateStatusForJourneyAlphaOutput{}
	return &this
}

// GetContainerId returns the ContainerId field value
func (o *UpdateStatusForJourneyAlphaOutput) GetContainerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value
// and a boolean to check if the value has been set.
func (o *UpdateStatusForJourneyAlphaOutput) GetContainerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerId, true
}

// SetContainerId sets field value
func (o *UpdateStatusForJourneyAlphaOutput) SetContainerId(v string) {
	o.ContainerId = v
}

// GetVersion returns the Version field value
func (o *UpdateStatusForJourneyAlphaOutput) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateStatusForJourneyAlphaOutput) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateStatusForJourneyAlphaOutput) SetVersion(v float32) {
	o.Version = v
}

// GetStatus returns the Status field value
func (o *UpdateStatusForJourneyAlphaOutput) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateStatusForJourneyAlphaOutput) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateStatusForJourneyAlphaOutput) SetStatus(v string) {
	o.Status = v
}

func (o UpdateStatusForJourneyAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStatusForJourneyAlphaOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["containerId"] = o.ContainerId
	toSerialize["version"] = o.Version
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableUpdateStatusForJourneyAlphaOutput struct {
	value *UpdateStatusForJourneyAlphaOutput
	isSet bool
}

func (v NullableUpdateStatusForJourneyAlphaOutput) Get() *UpdateStatusForJourneyAlphaOutput {
	return v.value
}

func (v *NullableUpdateStatusForJourneyAlphaOutput) Set(val *UpdateStatusForJourneyAlphaOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStatusForJourneyAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStatusForJourneyAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStatusForJourneyAlphaOutput(
	val *UpdateStatusForJourneyAlphaOutput,
) *NullableUpdateStatusForJourneyAlphaOutput {
	return &NullableUpdateStatusForJourneyAlphaOutput{value: val, isSet: true}
}

func (v NullableUpdateStatusForJourneyAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStatusForJourneyAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
