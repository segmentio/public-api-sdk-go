/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 34.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateEdgeFunctions200Response struct for CreateEdgeFunctions200Response
type CreateEdgeFunctions200Response struct {
	Data *CreateEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

// NewCreateEdgeFunctions200Response instantiates a new CreateEdgeFunctions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEdgeFunctions200Response() *CreateEdgeFunctions200Response {
	this := CreateEdgeFunctions200Response{}
	return &this
}

// NewCreateEdgeFunctions200ResponseWithDefaults instantiates a new CreateEdgeFunctions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEdgeFunctions200ResponseWithDefaults() *CreateEdgeFunctions200Response {
	this := CreateEdgeFunctions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateEdgeFunctions200Response) GetData() CreateEdgeFunctionsAlphaOutput {
	if o == nil || o.Data == nil {
		var ret CreateEdgeFunctionsAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEdgeFunctions200Response) GetDataOk() (*CreateEdgeFunctionsAlphaOutput, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateEdgeFunctions200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CreateEdgeFunctionsAlphaOutput and assigns it to the Data field.
func (o *CreateEdgeFunctions200Response) SetData(v CreateEdgeFunctionsAlphaOutput) {
	o.Data = &v
}

func (o CreateEdgeFunctions200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateEdgeFunctions200Response struct {
	value *CreateEdgeFunctions200Response
	isSet bool
}

func (v NullableCreateEdgeFunctions200Response) Get() *CreateEdgeFunctions200Response {
	return v.value
}

func (v *NullableCreateEdgeFunctions200Response) Set(val *CreateEdgeFunctions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEdgeFunctions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEdgeFunctions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEdgeFunctions200Response(
	val *CreateEdgeFunctions200Response,
) *NullableCreateEdgeFunctions200Response {
	return &NullableCreateEdgeFunctions200Response{value: val, isSet: true}
}

func (v NullableCreateEdgeFunctions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEdgeFunctions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
