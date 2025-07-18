/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.13.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Config Configuration for PERIODIC or SPECIFIC_DAYS strategies.
type Config struct {
	PeriodicConfig     *PeriodicConfig
	SpecificDaysConfig *SpecificDaysConfig
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Config) UnmarshalJSON(data []byte) error {
	var err error

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PeriodicConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.PeriodicConfig)
	if err == nil {
		jsonPeriodicConfig, _ := json.Marshal(dst.PeriodicConfig)
		if string(jsonPeriodicConfig) == "{}" { // empty struct
			dst.PeriodicConfig = nil
		} else {
			return nil // data stored in dst.PeriodicConfig, return on the first match
		}
	} else {
		dst.PeriodicConfig = nil
	}

	// try to unmarshal JSON data into SpecificDaysConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.SpecificDaysConfig)
	if err == nil {
		jsonSpecificDaysConfig, _ := json.Marshal(dst.SpecificDaysConfig)
		if string(jsonSpecificDaysConfig) == "{}" { // empty struct
			dst.SpecificDaysConfig = nil
		} else {
			return nil // data stored in dst.SpecificDaysConfig, return on the first match
		}
	} else {
		dst.SpecificDaysConfig = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Config)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Config) MarshalJSON() ([]byte, error) {
	if src.PeriodicConfig != nil {
		return json.Marshal(&src.PeriodicConfig)
	}

	if src.SpecificDaysConfig != nil {
		return json.Marshal(&src.SpecificDaysConfig)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
