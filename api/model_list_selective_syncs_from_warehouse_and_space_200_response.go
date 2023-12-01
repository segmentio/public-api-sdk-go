/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.2.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSelectiveSyncsFromWarehouseAndSpace200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSelectiveSyncsFromWarehouseAndSpace200Response{}

// ListSelectiveSyncsFromWarehouseAndSpace200Response struct for ListSelectiveSyncsFromWarehouseAndSpace200Response
type ListSelectiveSyncsFromWarehouseAndSpace200Response struct {
	Data *ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput `json:"data,omitempty"`
}

// NewListSelectiveSyncsFromWarehouseAndSpace200Response instantiates a new ListSelectiveSyncsFromWarehouseAndSpace200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSelectiveSyncsFromWarehouseAndSpace200Response() *ListSelectiveSyncsFromWarehouseAndSpace200Response {
	this := ListSelectiveSyncsFromWarehouseAndSpace200Response{}
	return &this
}

// NewListSelectiveSyncsFromWarehouseAndSpace200ResponseWithDefaults instantiates a new ListSelectiveSyncsFromWarehouseAndSpace200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSelectiveSyncsFromWarehouseAndSpace200ResponseWithDefaults() *ListSelectiveSyncsFromWarehouseAndSpace200Response {
	this := ListSelectiveSyncsFromWarehouseAndSpace200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSelectiveSyncsFromWarehouseAndSpace200Response) GetData() ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput {
	if o == nil || IsNil(o.Data) {
		var ret ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpace200Response) GetDataOk() (*ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSelectiveSyncsFromWarehouseAndSpace200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput and assigns it to the Data field.
func (o *ListSelectiveSyncsFromWarehouseAndSpace200Response) SetData(
	v ListSelectiveSyncsFromWarehouseAndSpaceAlphaOutput,
) {
	o.Data = &v
}

func (o ListSelectiveSyncsFromWarehouseAndSpace200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSelectiveSyncsFromWarehouseAndSpace200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListSelectiveSyncsFromWarehouseAndSpace200Response struct {
	value *ListSelectiveSyncsFromWarehouseAndSpace200Response
	isSet bool
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpace200Response) Get() *ListSelectiveSyncsFromWarehouseAndSpace200Response {
	return v.value
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpace200Response) Set(
	val *ListSelectiveSyncsFromWarehouseAndSpace200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpace200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpace200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSelectiveSyncsFromWarehouseAndSpace200Response(
	val *ListSelectiveSyncsFromWarehouseAndSpace200Response,
) *NullableListSelectiveSyncsFromWarehouseAndSpace200Response {
	return &NullableListSelectiveSyncsFromWarehouseAndSpace200Response{value: val, isSet: true}
}

func (v NullableListSelectiveSyncsFromWarehouseAndSpace200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSelectiveSyncsFromWarehouseAndSpace200Response) UnmarshalJSON(
	src []byte,
) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
