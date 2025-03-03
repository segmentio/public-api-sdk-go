/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 57.4.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CancelReverseETLSyncForModelOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelReverseETLSyncForModelOutput{}

// CancelReverseETLSyncForModelOutput CancelReverseETLSyncForModelOutput either will return an error or a \"CANCELLING\" status.
type CancelReverseETLSyncForModelOutput struct {
	// The id of the Model.
	ModelId string `json:"modelId"`
	// The id of the Sync.
	SyncId string `json:"syncId"`
	// A place holder for a machine-friendly category for an error, if applicable. - \"SyncAlreadyCanceled\" - \"SyncFinishedCannotCancel\"
	ErrorCode *string `json:"errorCode,omitempty"`
	// A place holder for a human-readable description of the error, if applicable. - \"sync already canceled\" - \"sync already finished\".
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// If no error, status will be CANCELLING, as the extract/load might take some time to cancel.
	Status *string `json:"status,omitempty"`
}

// NewCancelReverseETLSyncForModelOutput instantiates a new CancelReverseETLSyncForModelOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelReverseETLSyncForModelOutput(
	modelId string,
	syncId string,
) *CancelReverseETLSyncForModelOutput {
	this := CancelReverseETLSyncForModelOutput{}
	this.ModelId = modelId
	this.SyncId = syncId
	return &this
}

// NewCancelReverseETLSyncForModelOutputWithDefaults instantiates a new CancelReverseETLSyncForModelOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelReverseETLSyncForModelOutputWithDefaults() *CancelReverseETLSyncForModelOutput {
	this := CancelReverseETLSyncForModelOutput{}
	return &this
}

// GetModelId returns the ModelId field value
func (o *CancelReverseETLSyncForModelOutput) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelOutput) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *CancelReverseETLSyncForModelOutput) SetModelId(v string) {
	o.ModelId = v
}

// GetSyncId returns the SyncId field value
func (o *CancelReverseETLSyncForModelOutput) GetSyncId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SyncId
}

// GetSyncIdOk returns a tuple with the SyncId field value
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelOutput) GetSyncIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SyncId, true
}

// SetSyncId sets field value
func (o *CancelReverseETLSyncForModelOutput) SetSyncId(v string) {
	o.SyncId = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CancelReverseETLSyncForModelOutput) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelOutput) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CancelReverseETLSyncForModelOutput) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CancelReverseETLSyncForModelOutput) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *CancelReverseETLSyncForModelOutput) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelOutput) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *CancelReverseETLSyncForModelOutput) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *CancelReverseETLSyncForModelOutput) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelReverseETLSyncForModelOutput) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelReverseETLSyncForModelOutput) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelReverseETLSyncForModelOutput) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CancelReverseETLSyncForModelOutput) SetStatus(v string) {
	o.Status = &v
}

func (o CancelReverseETLSyncForModelOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelReverseETLSyncForModelOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["modelId"] = o.ModelId
	toSerialize["syncId"] = o.SyncId
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCancelReverseETLSyncForModelOutput struct {
	value *CancelReverseETLSyncForModelOutput
	isSet bool
}

func (v NullableCancelReverseETLSyncForModelOutput) Get() *CancelReverseETLSyncForModelOutput {
	return v.value
}

func (v *NullableCancelReverseETLSyncForModelOutput) Set(val *CancelReverseETLSyncForModelOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelReverseETLSyncForModelOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelReverseETLSyncForModelOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelReverseETLSyncForModelOutput(
	val *CancelReverseETLSyncForModelOutput,
) *NullableCancelReverseETLSyncForModelOutput {
	return &NullableCancelReverseETLSyncForModelOutput{value: val, isSet: true}
}

func (v NullableCancelReverseETLSyncForModelOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelReverseETLSyncForModelOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
