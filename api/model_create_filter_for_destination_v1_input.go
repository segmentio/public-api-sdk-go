/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 50.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the CreateFilterForDestinationV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFilterForDestinationV1Input{}

// CreateFilterForDestinationV1Input Input for CreateDestinationFilterV1.
type CreateFilterForDestinationV1Input struct {
	// The id of the Source associated with this filter.
	SourceId string `json:"sourceId"`
	// The filter's condition.
	If string `json:"if"`
	// Actions for the Destination filter.
	Actions []DestinationFilterActionV1 `json:"actions"`
	// The title of the filter.
	Title string `json:"title"`
	// The description of the filter.
	Description *string `json:"description,omitempty"`
	// When set to true, the Destination filter is active.
	Enabled bool `json:"enabled"`
}

// NewCreateFilterForDestinationV1Input instantiates a new CreateFilterForDestinationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFilterForDestinationV1Input(
	sourceId string,
	if_ string,
	actions []DestinationFilterActionV1,
	title string,
	enabled bool,
) *CreateFilterForDestinationV1Input {
	this := CreateFilterForDestinationV1Input{}
	this.SourceId = sourceId
	this.If = if_
	this.Actions = actions
	this.Title = title
	this.Enabled = enabled
	return &this
}

// NewCreateFilterForDestinationV1InputWithDefaults instantiates a new CreateFilterForDestinationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFilterForDestinationV1InputWithDefaults() *CreateFilterForDestinationV1Input {
	this := CreateFilterForDestinationV1Input{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *CreateFilterForDestinationV1Input) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *CreateFilterForDestinationV1Input) SetSourceId(v string) {
	o.SourceId = v
}

// GetIf returns the If field value
func (o *CreateFilterForDestinationV1Input) GetIf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.If
}

// GetIfOk returns a tuple with the If field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetIfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.If, true
}

// SetIf sets field value
func (o *CreateFilterForDestinationV1Input) SetIf(v string) {
	o.If = v
}

// GetActions returns the Actions field value
func (o *CreateFilterForDestinationV1Input) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		var ret []DestinationFilterActionV1
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetActionsOk() ([]DestinationFilterActionV1, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *CreateFilterForDestinationV1Input) SetActions(v []DestinationFilterActionV1) {
	o.Actions = v
}

// GetTitle returns the Title field value
func (o *CreateFilterForDestinationV1Input) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateFilterForDestinationV1Input) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateFilterForDestinationV1Input) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateFilterForDestinationV1Input) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateFilterForDestinationV1Input) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CreateFilterForDestinationV1Input) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateFilterForDestinationV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *CreateFilterForDestinationV1Input) SetEnabled(v bool) {
	o.Enabled = v
}

func (o CreateFilterForDestinationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFilterForDestinationV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourceId"] = o.SourceId
	toSerialize["if"] = o.If
	toSerialize["actions"] = o.Actions
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableCreateFilterForDestinationV1Input struct {
	value *CreateFilterForDestinationV1Input
	isSet bool
}

func (v NullableCreateFilterForDestinationV1Input) Get() *CreateFilterForDestinationV1Input {
	return v.value
}

func (v *NullableCreateFilterForDestinationV1Input) Set(val *CreateFilterForDestinationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFilterForDestinationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFilterForDestinationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFilterForDestinationV1Input(
	val *CreateFilterForDestinationV1Input,
) *NullableCreateFilterForDestinationV1Input {
	return &NullableCreateFilterForDestinationV1Input{value: val, isSet: true}
}

func (v NullableCreateFilterForDestinationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFilterForDestinationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
