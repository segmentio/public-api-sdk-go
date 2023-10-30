/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 36.3.1
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateReverseEtlModelInput Defines how to create a new Model.
type CreateReverseEtlModelInput struct {
	// Indicates which Source to attach this model to.
	SourceId string `json:"sourceId"`
	// A short, human-readable description of the Model.
	Name string `json:"name"`
	// A longer, more descriptive explanation of the Model.
	Description string `json:"description"`
	// Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions.
	Enabled bool `json:"enabled"`
	// Determines the strategy used for triggering syncs, which will be used in conjunction with scheduleConfig.
	ScheduleStrategy string `json:"scheduleStrategy"`
	// Depending on the chosen strategy, configures the schedule for this model.
	ScheduleConfig NullableModelMap `json:"scheduleConfig"`
	// The SQL query that will be executed to extract data from the connected Source.
	Query string `json:"query"`
	// Indicates the column named in `query` that should be used to uniquely identify the extracted records.
	QueryIdentifierColumn string `json:"queryIdentifierColumn"`
}

// NewCreateReverseEtlModelInput instantiates a new CreateReverseEtlModelInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReverseEtlModelInput(
	sourceId string,
	name string,
	description string,
	enabled bool,
	scheduleStrategy string,
	scheduleConfig NullableModelMap,
	query string,
	queryIdentifierColumn string,
) *CreateReverseEtlModelInput {
	this := CreateReverseEtlModelInput{}
	this.SourceId = sourceId
	this.Name = name
	this.Description = description
	this.Enabled = enabled
	this.ScheduleStrategy = scheduleStrategy
	this.ScheduleConfig = scheduleConfig
	this.Query = query
	this.QueryIdentifierColumn = queryIdentifierColumn
	return &this
}

// NewCreateReverseEtlModelInputWithDefaults instantiates a new CreateReverseEtlModelInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReverseEtlModelInputWithDefaults() *CreateReverseEtlModelInput {
	this := CreateReverseEtlModelInput{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *CreateReverseEtlModelInput) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CreateReverseEtlModelInput) SetSourceId(v string) {
	o.SourceId = v
}

// GetName returns the Name field value
func (o *CreateReverseEtlModelInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateReverseEtlModelInput) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *CreateReverseEtlModelInput) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateReverseEtlModelInput) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value
func (o *CreateReverseEtlModelInput) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateReverseEtlModelInput) SetEnabled(v bool) {
	o.Enabled = v
}

// GetScheduleStrategy returns the ScheduleStrategy field value
func (o *CreateReverseEtlModelInput) GetScheduleStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleStrategy
}

// GetScheduleStrategyOk returns a tuple with the ScheduleStrategy field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetScheduleStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleStrategy, true
}

// SetScheduleStrategy sets field value
func (o *CreateReverseEtlModelInput) SetScheduleStrategy(v string) {
	o.ScheduleStrategy = v
}

// GetScheduleConfig returns the ScheduleConfig field value
// If the value is explicit nil, the zero value for ModelMap will be returned
func (o *CreateReverseEtlModelInput) GetScheduleConfig() ModelMap {
	if o == nil || o.ScheduleConfig.Get() == nil {
		var ret ModelMap
		return ret
	}

	return *o.ScheduleConfig.Get()
}

// GetScheduleConfigOk returns a tuple with the ScheduleConfig field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateReverseEtlModelInput) GetScheduleConfigOk() (*ModelMap, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleConfig.Get(), o.ScheduleConfig.IsSet()
}

// SetScheduleConfig sets field value
func (o *CreateReverseEtlModelInput) SetScheduleConfig(v ModelMap) {
	o.ScheduleConfig.Set(&v)
}

// GetQuery returns the Query field value
func (o *CreateReverseEtlModelInput) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *CreateReverseEtlModelInput) SetQuery(v string) {
	o.Query = v
}

// GetQueryIdentifierColumn returns the QueryIdentifierColumn field value
func (o *CreateReverseEtlModelInput) GetQueryIdentifierColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryIdentifierColumn
}

// GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field value
// and a boolean to check if the value has been set.
func (o *CreateReverseEtlModelInput) GetQueryIdentifierColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryIdentifierColumn, true
}

// SetQueryIdentifierColumn sets field value
func (o *CreateReverseEtlModelInput) SetQueryIdentifierColumn(v string) {
	o.QueryIdentifierColumn = v
}

func (o CreateReverseEtlModelInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["scheduleStrategy"] = o.ScheduleStrategy
	}
	if true {
		toSerialize["scheduleConfig"] = o.ScheduleConfig.Get()
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if true {
		toSerialize["queryIdentifierColumn"] = o.QueryIdentifierColumn
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReverseEtlModelInput struct {
	value *CreateReverseEtlModelInput
	isSet bool
}

func (v NullableCreateReverseEtlModelInput) Get() *CreateReverseEtlModelInput {
	return v.value
}

func (v *NullableCreateReverseEtlModelInput) Set(val *CreateReverseEtlModelInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReverseEtlModelInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReverseEtlModelInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReverseEtlModelInput(
	val *CreateReverseEtlModelInput,
) *NullableCreateReverseEtlModelInput {
	return &NullableCreateReverseEtlModelInput{value: val, isSet: true}
}

func (v NullableCreateReverseEtlModelInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReverseEtlModelInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
