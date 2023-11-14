/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 38.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateReverseEtlModelInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReverseEtlModelInput{}

// UpdateReverseEtlModelInput Defines how to update an existing Model.
type UpdateReverseEtlModelInput struct {
	// A short, human-readable description of the Model.
	Name *string `json:"name,omitempty"`
	// A longer, more descriptive explanation of the Model.
	Description *string `json:"description,omitempty"`
	// Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions.
	Enabled *bool `json:"enabled,omitempty"`
	// Determines the strategy used for triggering syncs, which will be used in conjunction with scheduleConfig.
	ScheduleStrategy *string `json:"scheduleStrategy,omitempty"`
	// Defines a configuration object used for scheduling, which can vary depending on the configured strategy, but must always be an object with at least 1 level of keys.
	ScheduleConfig map[string]interface{} `json:"scheduleConfig,omitempty"`
	// The SQL query that will be executed to extract data from the connected Source.
	Query *string `json:"query,omitempty"`
	// Indicates the column named in `query` that should be used to uniquely identify the extracted records.
	QueryIdentifierColumn *string `json:"queryIdentifierColumn,omitempty"`
}

// NewUpdateReverseEtlModelInput instantiates a new UpdateReverseEtlModelInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReverseEtlModelInput() *UpdateReverseEtlModelInput {
	this := UpdateReverseEtlModelInput{}
	return &this
}

// NewUpdateReverseEtlModelInputWithDefaults instantiates a new UpdateReverseEtlModelInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReverseEtlModelInputWithDefaults() *UpdateReverseEtlModelInput {
	this := UpdateReverseEtlModelInput{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateReverseEtlModelInput) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateReverseEtlModelInput) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateReverseEtlModelInput) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetScheduleStrategy returns the ScheduleStrategy field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetScheduleStrategy() string {
	if o == nil || IsNil(o.ScheduleStrategy) {
		var ret string
		return ret
	}
	return *o.ScheduleStrategy
}

// GetScheduleStrategyOk returns a tuple with the ScheduleStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetScheduleStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.ScheduleStrategy) {
		return nil, false
	}
	return o.ScheduleStrategy, true
}

// HasScheduleStrategy returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasScheduleStrategy() bool {
	if o != nil && !IsNil(o.ScheduleStrategy) {
		return true
	}

	return false
}

// SetScheduleStrategy gets a reference to the given string and assigns it to the ScheduleStrategy field.
func (o *UpdateReverseEtlModelInput) SetScheduleStrategy(v string) {
	o.ScheduleStrategy = &v
}

// GetScheduleConfig returns the ScheduleConfig field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetScheduleConfig() map[string]interface{} {
	if o == nil || IsNil(o.ScheduleConfig) {
		var ret map[string]interface{}
		return ret
	}
	return o.ScheduleConfig
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetScheduleConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ScheduleConfig) {
		return map[string]interface{}{}, false
	}
	return o.ScheduleConfig, true
}

// HasScheduleConfig returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasScheduleConfig() bool {
	if o != nil && !IsNil(o.ScheduleConfig) {
		return true
	}

	return false
}

// SetScheduleConfig gets a reference to the given map[string]interface{} and assigns it to the ScheduleConfig field.
func (o *UpdateReverseEtlModelInput) SetScheduleConfig(v map[string]interface{}) {
	o.ScheduleConfig = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *UpdateReverseEtlModelInput) SetQuery(v string) {
	o.Query = &v
}

// GetQueryIdentifierColumn returns the QueryIdentifierColumn field value if set, zero value otherwise.
func (o *UpdateReverseEtlModelInput) GetQueryIdentifierColumn() string {
	if o == nil || IsNil(o.QueryIdentifierColumn) {
		var ret string
		return ret
	}
	return *o.QueryIdentifierColumn
}

// GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReverseEtlModelInput) GetQueryIdentifierColumnOk() (*string, bool) {
	if o == nil || IsNil(o.QueryIdentifierColumn) {
		return nil, false
	}
	return o.QueryIdentifierColumn, true
}

// HasQueryIdentifierColumn returns a boolean if a field has been set.
func (o *UpdateReverseEtlModelInput) HasQueryIdentifierColumn() bool {
	if o != nil && !IsNil(o.QueryIdentifierColumn) {
		return true
	}

	return false
}

// SetQueryIdentifierColumn gets a reference to the given string and assigns it to the QueryIdentifierColumn field.
func (o *UpdateReverseEtlModelInput) SetQueryIdentifierColumn(v string) {
	o.QueryIdentifierColumn = &v
}

func (o UpdateReverseEtlModelInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateReverseEtlModelInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ScheduleStrategy) {
		toSerialize["scheduleStrategy"] = o.ScheduleStrategy
	}
	if !IsNil(o.ScheduleConfig) {
		toSerialize["scheduleConfig"] = o.ScheduleConfig
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.QueryIdentifierColumn) {
		toSerialize["queryIdentifierColumn"] = o.QueryIdentifierColumn
	}
	return toSerialize, nil
}

type NullableUpdateReverseEtlModelInput struct {
	value *UpdateReverseEtlModelInput
	isSet bool
}

func (v NullableUpdateReverseEtlModelInput) Get() *UpdateReverseEtlModelInput {
	return v.value
}

func (v *NullableUpdateReverseEtlModelInput) Set(val *UpdateReverseEtlModelInput) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReverseEtlModelInput) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReverseEtlModelInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReverseEtlModelInput(
	val *UpdateReverseEtlModelInput,
) *NullableUpdateReverseEtlModelInput {
	return &NullableUpdateReverseEtlModelInput{value: val, isSet: true}
}

func (v NullableUpdateReverseEtlModelInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReverseEtlModelInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
