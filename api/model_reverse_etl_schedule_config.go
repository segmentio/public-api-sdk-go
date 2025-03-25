/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 58.1.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ReverseEtlScheduleConfig struct for ReverseEtlScheduleConfig
type ReverseEtlScheduleConfig struct {
	ReverseEtlCronScheduleConfig         *ReverseEtlCronScheduleConfig
	ReverseEtlDbtCloudScheduleConfig     *ReverseEtlDbtCloudScheduleConfig
	ReverseEtlPeriodicScheduleConfig     *ReverseEtlPeriodicScheduleConfig
	ReverseEtlSpecificTimeScheduleConfig *ReverseEtlSpecificTimeScheduleConfig
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReverseEtlScheduleConfig) UnmarshalJSON(data []byte) error {
	var err error

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

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

	// try to unmarshal JSON data into ReverseEtlDbtCloudScheduleConfig
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.ReverseEtlDbtCloudScheduleConfig)
	if err == nil {
		jsonReverseEtlDbtCloudScheduleConfig, _ := json.Marshal(
			dst.ReverseEtlDbtCloudScheduleConfig,
		)
		if string(jsonReverseEtlDbtCloudScheduleConfig) == "{}" { // empty struct
			dst.ReverseEtlDbtCloudScheduleConfig = nil
		} else {
			return nil // data stored in dst.ReverseEtlDbtCloudScheduleConfig, return on the first match
		}
	} else {
		dst.ReverseEtlDbtCloudScheduleConfig = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReverseEtlScheduleConfig)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReverseEtlScheduleConfig) MarshalJSON() ([]byte, error) {
	if src.ReverseEtlCronScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlCronScheduleConfig)
	}

	if src.ReverseEtlDbtCloudScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlDbtCloudScheduleConfig)
	}

	if src.ReverseEtlPeriodicScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlPeriodicScheduleConfig)
	}

	if src.ReverseEtlSpecificTimeScheduleConfig != nil {
		return json.Marshal(&src.ReverseEtlSpecificTimeScheduleConfig)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReverseEtlScheduleConfig struct {
	value *ReverseEtlScheduleConfig
	isSet bool
}

func (v NullableReverseEtlScheduleConfig) Get() *ReverseEtlScheduleConfig {
	return v.value
}

func (v *NullableReverseEtlScheduleConfig) Set(val *ReverseEtlScheduleConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReverseEtlScheduleConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReverseEtlScheduleConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReverseEtlScheduleConfig(
	val *ReverseEtlScheduleConfig,
) *NullableReverseEtlScheduleConfig {
	return &NullableReverseEtlScheduleConfig{value: val, isSet: true}
}

func (v NullableReverseEtlScheduleConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReverseEtlScheduleConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
