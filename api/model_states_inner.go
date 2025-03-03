/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// StatesInner struct for StatesInner
type StatesInner struct {
	AudienceExitRule *AudienceExitRule
	DestinationState *DestinationState
	EventExitRule    *EventExitRule
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StatesInner) UnmarshalJSON(data []byte) error {
	var err error

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	// try to unmarshal JSON data into AudienceExitRule
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.AudienceExitRule)
	if err == nil {
		jsonAudienceExitRule, _ := json.Marshal(dst.AudienceExitRule)
		if string(jsonAudienceExitRule) == "{}" { // empty struct
			dst.AudienceExitRule = nil
		} else {
			return nil // data stored in dst.AudienceExitRule, return on the first match
		}
	} else {
		dst.AudienceExitRule = nil
	}

	// try to unmarshal JSON data into DestinationState
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.DestinationState)
	if err == nil {
		jsonDestinationState, _ := json.Marshal(dst.DestinationState)
		if string(jsonDestinationState) == "{}" { // empty struct
			dst.DestinationState = nil
		} else {
			return nil // data stored in dst.DestinationState, return on the first match
		}
	} else {
		dst.DestinationState = nil
	}

	// try to unmarshal JSON data into EventExitRule
	decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&dst.EventExitRule)
	if err == nil {
		jsonEventExitRule, _ := json.Marshal(dst.EventExitRule)
		if string(jsonEventExitRule) == "{}" { // empty struct
			dst.EventExitRule = nil
		} else {
			return nil // data stored in dst.EventExitRule, return on the first match
		}
	} else {
		dst.EventExitRule = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(StatesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *StatesInner) MarshalJSON() ([]byte, error) {
	if src.AudienceExitRule != nil {
		return json.Marshal(&src.AudienceExitRule)
	}

	if src.DestinationState != nil {
		return json.Marshal(&src.DestinationState)
	}

	if src.EventExitRule != nil {
		return json.Marshal(&src.EventExitRule)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableStatesInner struct {
	value *StatesInner
	isSet bool
}

func (v NullableStatesInner) Get() *StatesInner {
	return v.value
}

func (v *NullableStatesInner) Set(val *StatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatesInner(val *StatesInner) *NullableStatesInner {
	return &NullableStatesInner{value: val, isSet: true}
}

func (v NullableStatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
