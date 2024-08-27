/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 53.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GetFilterInDestination200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFilterInDestination200Response{}

// GetFilterInDestination200Response struct for GetFilterInDestination200Response
type GetFilterInDestination200Response struct {
	Data *GetFilterInDestinationV1Output `json:"data,omitempty"`
}

// NewGetFilterInDestination200Response instantiates a new GetFilterInDestination200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFilterInDestination200Response() *GetFilterInDestination200Response {
	this := GetFilterInDestination200Response{}
	return &this
}

// NewGetFilterInDestination200ResponseWithDefaults instantiates a new GetFilterInDestination200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFilterInDestination200ResponseWithDefaults() *GetFilterInDestination200Response {
	this := GetFilterInDestination200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFilterInDestination200Response) GetData() GetFilterInDestinationV1Output {
	if o == nil || IsNil(o.Data) {
		var ret GetFilterInDestinationV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilterInDestination200Response) GetDataOk() (*GetFilterInDestinationV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFilterInDestination200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetFilterInDestinationV1Output and assigns it to the Data field.
func (o *GetFilterInDestination200Response) SetData(v GetFilterInDestinationV1Output) {
	o.Data = &v
}

func (o GetFilterInDestination200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFilterInDestination200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetFilterInDestination200Response struct {
	value *GetFilterInDestination200Response
	isSet bool
}

func (v NullableGetFilterInDestination200Response) Get() *GetFilterInDestination200Response {
	return v.value
}

func (v *NullableGetFilterInDestination200Response) Set(val *GetFilterInDestination200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFilterInDestination200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFilterInDestination200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFilterInDestination200Response(
	val *GetFilterInDestination200Response,
) *NullableGetFilterInDestination200Response {
	return &NullableGetFilterInDestination200Response{value: val, isSet: true}
}

func (v NullableGetFilterInDestination200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFilterInDestination200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
