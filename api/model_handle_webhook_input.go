/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// HandleWebhookInput Function webhook input.
type HandleWebhookInput struct {
	// The Workspace id.
	W string `json:"w"`
	// The webhook nonce.
	N string `json:"n"`
	// The webhook timestamp.
	T string `json:"t"`
	// The webhook signature.
	S string `json:"s"`
}

// NewHandleWebhookInput instantiates a new HandleWebhookInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandleWebhookInput(w string, n string, t string, s string) *HandleWebhookInput {
	this := HandleWebhookInput{}
	this.W = w
	this.N = n
	this.T = t
	this.S = s
	return &this
}

// NewHandleWebhookInputWithDefaults instantiates a new HandleWebhookInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandleWebhookInputWithDefaults() *HandleWebhookInput {
	this := HandleWebhookInput{}
	return &this
}

// GetW returns the W field value
func (o *HandleWebhookInput) GetW() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.W
}

// GetWOk returns a tuple with the W field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookInput) GetWOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.W, true
}

// SetW sets field value
func (o *HandleWebhookInput) SetW(v string) {
	o.W = v
}

// GetN returns the N field value
func (o *HandleWebhookInput) GetN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.N
}

// GetNOk returns a tuple with the N field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookInput) GetNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N, true
}

// SetN sets field value
func (o *HandleWebhookInput) SetN(v string) {
	o.N = v
}

// GetT returns the T field value
func (o *HandleWebhookInput) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookInput) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *HandleWebhookInput) SetT(v string) {
	o.T = v
}

// GetS returns the S field value
func (o *HandleWebhookInput) GetS() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.S
}

// GetSOk returns a tuple with the S field value
// and a boolean to check if the value has been set.
func (o *HandleWebhookInput) GetSOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.S, true
}

// SetS sets field value
func (o *HandleWebhookInput) SetS(v string) {
	o.S = v
}

func (o HandleWebhookInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["w"] = o.W
	}
	if true {
		toSerialize["n"] = o.N
	}
	if true {
		toSerialize["t"] = o.T
	}
	if true {
		toSerialize["s"] = o.S
	}
	return json.Marshal(toSerialize)
}

type NullableHandleWebhookInput struct {
	value *HandleWebhookInput
	isSet bool
}

func (v NullableHandleWebhookInput) Get() *HandleWebhookInput {
	return v.value
}

func (v *NullableHandleWebhookInput) Set(val *HandleWebhookInput) {
	v.value = val
	v.isSet = true
}

func (v NullableHandleWebhookInput) IsSet() bool {
	return v.isSet
}

func (v *NullableHandleWebhookInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandleWebhookInput(val *HandleWebhookInput) *NullableHandleWebhookInput {
	return &NullableHandleWebhookInput{value: val, isSet: true}
}

func (v NullableHandleWebhookInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandleWebhookInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
