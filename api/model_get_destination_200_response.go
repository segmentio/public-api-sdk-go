/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.8
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetDestination200Response struct for GetDestination200Response
type GetDestination200Response struct {
	Data *GetDestinationV1Output `json:"data,omitempty"`
}

// NewGetDestination200Response instantiates a new GetDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDestination200Response() *GetDestination200Response {
	this := GetDestination200Response{}
	return &this
}

// NewGetDestination200ResponseWithDefaults instantiates a new GetDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDestination200ResponseWithDefaults() *GetDestination200Response {
	this := GetDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDestination200Response) GetData() GetDestinationV1Output {
	if o == nil || o.Data == nil {
		var ret GetDestinationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDestination200Response) GetDataOk() (*GetDestinationV1Output, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDestination200Response) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GetDestinationV1Output and assigns it to the Data field.
func (o *GetDestination200Response) SetData(v GetDestinationV1Output) {
	o.Data = &v
}

func (o GetDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetDestination200Response struct {
	value *GetDestination200Response
	isSet bool
}

func (v NullableGetDestination200Response) Get() *GetDestination200Response {
	return v.value
}

func (v *NullableGetDestination200Response) Set(val *GetDestination200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDestination200Response(val *GetDestination200Response) *NullableGetDestination200Response {
	return &NullableGetDestination200Response{value: val, isSet: true}
}

func (v NullableGetDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


