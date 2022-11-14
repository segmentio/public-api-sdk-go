/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API. 

API version: 32.0.4
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import "encoding/json"

type ModelMap struct {
	value map[string]interface{}
	isSet bool
}

type NullableModelMap struct {
	value *ModelMap
	isSet bool
}

func (v NullableModelMap) Get() *ModelMap {
	return v.value
}

func (v *NullableModelMap) Set(val *ModelMap) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMap) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMap(val *ModelMap) *NullableModelMap {
	return &NullableModelMap{value: val, isSet: true}
}

func (v NullableModelMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
