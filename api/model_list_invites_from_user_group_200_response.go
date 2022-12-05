/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 33.0.3
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ListInvitesFromUserGroup200Response struct for ListInvitesFromUserGroup200Response
type ListInvitesFromUserGroup200Response struct {
	Data *ListInvitesFromUserGroupV1Output `json:"data,omitempty"`
}

// NewListInvitesFromUserGroup200Response instantiates a new ListInvitesFromUserGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInvitesFromUserGroup200Response() *ListInvitesFromUserGroup200Response {
	this := ListInvitesFromUserGroup200Response{}
	return &this
}

// NewListInvitesFromUserGroup200ResponseWithDefaults instantiates a new ListInvitesFromUserGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInvitesFromUserGroup200ResponseWithDefaults() *ListInvitesFromUserGroup200Response {
	this := ListInvitesFromUserGroup200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListInvitesFromUserGroup200Response) GetData() ListInvitesFromUserGroupV1Output {
	if o == nil || o.Data == nil {
		var ret ListInvitesFromUserGroupV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInvitesFromUserGroup200Response) GetDataOk() (*ListInvitesFromUserGroupV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListInvitesFromUserGroup200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ListInvitesFromUserGroupV1Output and assigns it to the Data field.
func (o *ListInvitesFromUserGroup200Response) SetData(v ListInvitesFromUserGroupV1Output) {
	o.Data = &v
}

func (o ListInvitesFromUserGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListInvitesFromUserGroup200Response struct {
	value *ListInvitesFromUserGroup200Response
	isSet bool
}

func (v NullableListInvitesFromUserGroup200Response) Get() *ListInvitesFromUserGroup200Response {
	return v.value
}

func (v *NullableListInvitesFromUserGroup200Response) Set(
	val *ListInvitesFromUserGroup200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListInvitesFromUserGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListInvitesFromUserGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInvitesFromUserGroup200Response(
	val *ListInvitesFromUserGroup200Response,
) *NullableListInvitesFromUserGroup200Response {
	return &NullableListInvitesFromUserGroup200Response{value: val, isSet: true}
}

func (v NullableListInvitesFromUserGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInvitesFromUserGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
