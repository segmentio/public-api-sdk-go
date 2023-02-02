/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GenerateUploadURLForEdgeFunctionsAlphaOutput Output for GenerateSignedUrl.
type GenerateUploadURLForEdgeFunctionsAlphaOutput struct {
	// A temporary URL that can be used to upload your Edge Functions bundle. Expires in 15 minutes.
	UploadURL string `json:"uploadURL"`
}

// NewGenerateUploadURLForEdgeFunctionsAlphaOutput instantiates a new GenerateUploadURLForEdgeFunctionsAlphaOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateUploadURLForEdgeFunctionsAlphaOutput(
	uploadURL string,
) *GenerateUploadURLForEdgeFunctionsAlphaOutput {
	this := GenerateUploadURLForEdgeFunctionsAlphaOutput{}
	this.UploadURL = uploadURL
	return &this
}

// NewGenerateUploadURLForEdgeFunctionsAlphaOutputWithDefaults instantiates a new GenerateUploadURLForEdgeFunctionsAlphaOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateUploadURLForEdgeFunctionsAlphaOutputWithDefaults() *GenerateUploadURLForEdgeFunctionsAlphaOutput {
	this := GenerateUploadURLForEdgeFunctionsAlphaOutput{}
	return &this
}

// GetUploadURL returns the UploadURL field value
func (o *GenerateUploadURLForEdgeFunctionsAlphaOutput) GetUploadURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadURL
}

// GetUploadURLOk returns a tuple with the UploadURL field value
// and a boolean to check if the value has been set.
func (o *GenerateUploadURLForEdgeFunctionsAlphaOutput) GetUploadURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadURL, true
}

// SetUploadURL sets field value
func (o *GenerateUploadURLForEdgeFunctionsAlphaOutput) SetUploadURL(v string) {
	o.UploadURL = v
}

func (o GenerateUploadURLForEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["uploadURL"] = o.UploadURL
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateUploadURLForEdgeFunctionsAlphaOutput struct {
	value *GenerateUploadURLForEdgeFunctionsAlphaOutput
	isSet bool
}

func (v NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) Get() *GenerateUploadURLForEdgeFunctionsAlphaOutput {
	return v.value
}

func (v *NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) Set(
	val *GenerateUploadURLForEdgeFunctionsAlphaOutput,
) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateUploadURLForEdgeFunctionsAlphaOutput(
	val *GenerateUploadURLForEdgeFunctionsAlphaOutput,
) *NullableGenerateUploadURLForEdgeFunctionsAlphaOutput {
	return &NullableGenerateUploadURLForEdgeFunctionsAlphaOutput{value: val, isSet: true}
}

func (v NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateUploadURLForEdgeFunctionsAlphaOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
