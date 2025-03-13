/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubscriptionRequest{}

// GetSubscriptionRequest struct for GetSubscriptionRequest
type GetSubscriptionRequest struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used.
	Type string `json:"type"`
}

// NewGetSubscriptionRequest instantiates a new GetSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubscriptionRequest(key string, type_ string) *GetSubscriptionRequest {
	this := GetSubscriptionRequest{}
	this.Key = key
	this.Type = type_
	return &this
}

// NewGetSubscriptionRequestWithDefaults instantiates a new GetSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubscriptionRequestWithDefaults() *GetSubscriptionRequest {
	this := GetSubscriptionRequest{}
	return &this
}

// GetKey returns the Key field value
func (o *GetSubscriptionRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *GetSubscriptionRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *GetSubscriptionRequest) SetKey(v string) {
	o.Key = v
}

// GetType returns the Type field value
func (o *GetSubscriptionRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetSubscriptionRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetSubscriptionRequest) SetType(v string) {
	o.Type = v
}

func (o GetSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableGetSubscriptionRequest struct {
	value *GetSubscriptionRequest
	isSet bool
}

func (v NullableGetSubscriptionRequest) Get() *GetSubscriptionRequest {
	return v.value
}

func (v *NullableGetSubscriptionRequest) Set(val *GetSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubscriptionRequest(
	val *GetSubscriptionRequest,
) *NullableGetSubscriptionRequest {
	return &NullableGetSubscriptionRequest{value: val, isSet: true}
}

func (v NullableGetSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
