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

// checks if the UpdateGroupSubscriptionStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateGroupSubscriptionStatusResponse{}

// UpdateGroupSubscriptionStatusResponse struct for UpdateGroupSubscriptionStatusResponse
type UpdateGroupSubscriptionStatusResponse struct {
	// Name of the group.
	Name string `json:"name"`
	// The user subscribed, unsubscribed, or on initial status.
	Status string `json:"status"`
	// The group id.
	Id string `json:"id"`
}

// NewUpdateGroupSubscriptionStatusResponse instantiates a new UpdateGroupSubscriptionStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGroupSubscriptionStatusResponse(
	name string,
	status string,
	id string,
) *UpdateGroupSubscriptionStatusResponse {
	this := UpdateGroupSubscriptionStatusResponse{}
	this.Name = name
	this.Status = status
	this.Id = id
	return &this
}

// NewUpdateGroupSubscriptionStatusResponseWithDefaults instantiates a new UpdateGroupSubscriptionStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGroupSubscriptionStatusResponseWithDefaults() *UpdateGroupSubscriptionStatusResponse {
	this := UpdateGroupSubscriptionStatusResponse{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateGroupSubscriptionStatusResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupSubscriptionStatusResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateGroupSubscriptionStatusResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *UpdateGroupSubscriptionStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupSubscriptionStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UpdateGroupSubscriptionStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value
func (o *UpdateGroupSubscriptionStatusResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateGroupSubscriptionStatusResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateGroupSubscriptionStatusResponse) SetId(v string) {
	o.Id = v
}

func (o UpdateGroupSubscriptionStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateGroupSubscriptionStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableUpdateGroupSubscriptionStatusResponse struct {
	value *UpdateGroupSubscriptionStatusResponse
	isSet bool
}

func (v NullableUpdateGroupSubscriptionStatusResponse) Get() *UpdateGroupSubscriptionStatusResponse {
	return v.value
}

func (v *NullableUpdateGroupSubscriptionStatusResponse) Set(
	val *UpdateGroupSubscriptionStatusResponse,
) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGroupSubscriptionStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGroupSubscriptionStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGroupSubscriptionStatusResponse(
	val *UpdateGroupSubscriptionStatusResponse,
) *NullableUpdateGroupSubscriptionStatusResponse {
	return &NullableUpdateGroupSubscriptionStatusResponse{value: val, isSet: true}
}

func (v NullableUpdateGroupSubscriptionStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGroupSubscriptionStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
