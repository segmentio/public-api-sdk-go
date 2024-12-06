/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.1.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Config Config contains interval duration in case of periodic or day and hours in case of specific_days. Empty if strategy is MANUAL.
type Config struct {
	ReverseEtlCronScheduleConfig         *ReverseEtlCronScheduleConfig
	ReverseEtlPeriodicScheduleConfig     *ReverseEtlPeriodicScheduleConfig
	ReverseEtlSpecificTimeScheduleConfig *ReverseEtlSpecificTimeScheduleConfig
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

	// try to unmarshal JSON data into ReverseEtlCronScheduleConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.ReverseEtlCronScheduleConfig)
	if err == nil {
		jsonReverseEtlCronScheduleConfig, _ := json.Marshal(dst.ReverseEtlCronScheduleConfig)
		if string(jsonReverseEtlCronScheduleConfig) == "{}" { // empty struct
			dst.ReverseEtlCronScheduleConfig = nil
		} else {
			return nil // data stored in dst.ReverseEtlCronScheduleConfig, return on the first match
		}
	} else {
		dst.ReverseEtlCronScheduleConfig = nil
	}

	// try to unmarshal JSON data into ReverseEtlPeriodicScheduleConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.ReverseEtlPeriodicScheduleConfig)
	if err == nil {
		jsonReverseEtlPeriodicScheduleConfig, _ := json.Marshal(
			dst.ReverseEtlPeriodicScheduleConfig,
		)
		if string(jsonReverseEtlPeriodicScheduleConfig) == "{}" { // empty struct
			dst.ReverseEtlPeriodicScheduleConfig = nil
		} else {
			return nil // data stored in dst.ReverseEtlPeriodicScheduleConfig, return on the first match
		}
	} else {
		dst.ReverseEtlPeriodicScheduleConfig = nil
	}

	// try to unmarshal JSON data into ReverseEtlSpecificTimeScheduleConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.ReverseEtlSpecificTimeScheduleConfig)
	if err == nil {
		jsonReverseEtlSpecificTimeScheduleConfig, _ := json.Marshal(
			dst.ReverseEtlSpecificTimeScheduleConfig,
		)
		if string(jsonReverseEtlSpecificTimeScheduleConfig) == "{}" { // empty struct
			dst.ReverseEtlSpecificTimeScheduleConfig = nil
		} else {
			return nil // data stored in dst.ReverseEtlSpecificTimeScheduleConfig, return on the first match
		}
	} else {
		dst.ReverseEtlSpecificTimeScheduleConfig = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Config)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Config) MarshalJSON() ([]byte, error) {
	if src.ReverseEtlCronScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlCronScheduleConfig)
	}

	if src.ReverseEtlPeriodicScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlPeriodicScheduleConfig)
	}

	if src.ReverseEtlSpecificTimeScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlSpecificTimeScheduleConfig)
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
