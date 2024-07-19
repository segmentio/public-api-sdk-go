/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateEdgeFunctionsAlphaInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEdgeFunctionsAlphaInput{}

// CreateEdgeFunctionsAlphaInput Input for CreateEdgeFunctions.
type CreateEdgeFunctionsAlphaInput struct {
	// The id of the Source associated with this Edge Function.
	UploadURL string `json:"uploadURL"`
}

// NewCreateEdgeFunctionsAlphaInput instantiates a new CreateEdgeFunctionsAlphaInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEdgeFunctionsAlphaInput(uploadURL string) *CreateEdgeFunctionsAlphaInput {
	this := CreateEdgeFunctionsAlphaInput{}
	this.UploadURL = uploadURL
	return &this
}

// NewCreateEdgeFunctionsAlphaInputWithDefaults instantiates a new CreateEdgeFunctionsAlphaInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEdgeFunctionsAlphaInputWithDefaults() *CreateEdgeFunctionsAlphaInput {
	this := CreateEdgeFunctionsAlphaInput{}
	return &this
}

// GetUploadURL returns the UploadURL field value
func (o *CreateEdgeFunctionsAlphaInput) GetUploadURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadURL
}

// GetUploadURLOk returns a tuple with the UploadURL field value
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctionsAlphaInput) GetUploadURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadURL, true
}

// SetUploadURL sets field value
func (o *CreateEdgeFunctionsAlphaInput) SetUploadURL(v string) {
	o.UploadURL = v
}

func (o CreateEdgeFunctionsAlphaInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEdgeFunctionsAlphaInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uploadURL"] = o.UploadURL
	return toSerialize, nil
}

type NullableCreateEdgeFunctionsAlphaInput struct {
	value *CreateEdgeFunctionsAlphaInput
	isSet bool
}

func (v NullableCreateEdgeFunctionsAlphaInput) Get() *CreateEdgeFunctionsAlphaInput {
	return v.value
}

func (v *NullableCreateEdgeFunctionsAlphaInput) Set(val *CreateEdgeFunctionsAlphaInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFunctionsAlphaInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFunctionsAlphaInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFunctionsAlphaInput(
	val *CreateEdgeFunctionsAlphaInput,
) *NullableCreateEdgeFunctionsAlphaInput {
	return &NullableCreateEdgeFunctionsAlphaInput{value: val, isSet: true}
}

func (v NullableCreateEdgeFunctionsAlphaInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFunctionsAlphaInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
