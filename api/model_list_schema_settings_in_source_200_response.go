/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListSchemaSettingsInSource200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSchemaSettingsInSource200Response{}

// ListSchemaSettingsInSource200Response struct for ListSchemaSettingsInSource200Response
type ListSchemaSettingsInSource200Response struct {
	Data *ListSchemaSettingsInSourceV1Output `json:"data,omitempty"`
}

// NewListSchemaSettingsInSource200Response instantiates a new ListSchemaSettingsInSource200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSchemaSettingsInSource200Response() *ListSchemaSettingsInSource200Response {
	this := ListSchemaSettingsInSource200Response{}
	return &this
}

// NewListSchemaSettingsInSource200ResponseWithDefaults instantiates a new ListSchemaSettingsInSource200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSchemaSettingsInSource200ResponseWithDefaults() *ListSchemaSettingsInSource200Response {
	this := ListSchemaSettingsInSource200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListSchemaSettingsInSource200Response) GetData() ListSchemaSettingsInSourceV1Output {
	if o == nil || IsNil(o.Data) {
		var ret ListSchemaSettingsInSourceV1Output
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSchemaSettingsInSource200Response) GetDataOk() (*ListSchemaSettingsInSourceV1Output, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListSchemaSettingsInSource200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ListSchemaSettingsInSourceV1Output and assigns it to the Data field.
func (o *ListSchemaSettingsInSource200Response) SetData(v ListSchemaSettingsInSourceV1Output) {
	o.Data = &v
}

func (o ListSchemaSettingsInSource200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSchemaSettingsInSource200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListSchemaSettingsInSource200Response struct {
	value *ListSchemaSettingsInSource200Response
	isSet bool
}

func (v NullableListSchemaSettingsInSource200Response) Get() *ListSchemaSettingsInSource200Response {
	return v.value
}

func (v *NullableListSchemaSettingsInSource200Response) Set(
	val *ListSchemaSettingsInSource200Response,
) {
	v.value = val
	v.isSet = true
}

func (v NullableListSchemaSettingsInSource200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListSchemaSettingsInSource200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSchemaSettingsInSource200Response(
	val *ListSchemaSettingsInSource200Response,
) *NullableListSchemaSettingsInSource200Response {
	return &NullableListSchemaSettingsInSource200Response{value: val, isSet: true}
}

func (v NullableListSchemaSettingsInSource200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSchemaSettingsInSource200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
