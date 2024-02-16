/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 43.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateFilterForDestinationV1Input type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilterForDestinationV1Input{}

// UpdateFilterForDestinationV1Input Input for UpdateDestinationFilterV1.
type UpdateFilterForDestinationV1Input struct {
	// The FQL if condition to update.
	If *string `json:"if,omitempty"`
	// Actions for this Destination filter.
	Actions []DestinationFilterActionV1 `json:"actions,omitempty"`
	// The title to update.
	Title *string `json:"title,omitempty"`
	// The description of this filter.
	Description NullableString `json:"description,omitempty"`
	// When set to true, this Destination filter is active.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateFilterForDestinationV1Input instantiates a new UpdateFilterForDestinationV1Input object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterForDestinationV1Input() *UpdateFilterForDestinationV1Input {
	this := UpdateFilterForDestinationV1Input{}
	return &this
}

// NewUpdateFilterForDestinationV1InputWithDefaults instantiates a new UpdateFilterForDestinationV1Input object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterForDestinationV1InputWithDefaults() *UpdateFilterForDestinationV1Input {
	this := UpdateFilterForDestinationV1Input{}
	return &this
}

// GetIf returns the If field value if set, zero value otherwise.
func (o *UpdateFilterForDestinationV1Input) GetIf() string {
	if o == nil || IsNil(o.If) {
		var ret string
		return ret
	}
	return *o.If
}

// GetIfOk returns a tuple with the If field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestinationV1Input) GetIfOk() (*string, bool) {
	if o == nil || IsNil(o.If) {
		return nil, false
	}
	return o.If, true
}

// HasIf returns a boolean if a field has been set.
func (o *UpdateFilterForDestinationV1Input) HasIf() bool {
	if o != nil && !IsNil(o.If) {
		return true
	}

	return false
}

// SetIf gets a reference to the given string and assigns it to the If field.
func (o *UpdateFilterForDestinationV1Input) SetIf(v string) {
	o.If = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *UpdateFilterForDestinationV1Input) GetActions() []DestinationFilterActionV1 {
	if o == nil || IsNil(o.Actions) {
		var ret []DestinationFilterActionV1
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestinationV1Input) GetActionsOk() ([]DestinationFilterActionV1, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *UpdateFilterForDestinationV1Input) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []DestinationFilterActionV1 and assigns it to the Actions field.
func (o *UpdateFilterForDestinationV1Input) SetActions(v []DestinationFilterActionV1) {
	o.Actions = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateFilterForDestinationV1Input) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestinationV1Input) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateFilterForDestinationV1Input) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateFilterForDestinationV1Input) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateFilterForDestinationV1Input) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateFilterForDestinationV1Input) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateFilterForDestinationV1Input) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *UpdateFilterForDestinationV1Input) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *UpdateFilterForDestinationV1Input) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *UpdateFilterForDestinationV1Input) UnsetDescription() {
	o.Description.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateFilterForDestinationV1Input) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterForDestinationV1Input) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateFilterForDestinationV1Input) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateFilterForDestinationV1Input) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateFilterForDestinationV1Input) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilterForDestinationV1Input) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.If) {
		toSerialize["if"] = o.If
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateFilterForDestinationV1Input struct {
	value *UpdateFilterForDestinationV1Input
	isSet bool
}

func (v NullableUpdateFilterForDestinationV1Input) Get() *UpdateFilterForDestinationV1Input {
	return v.value
}

func (v *NullableUpdateFilterForDestinationV1Input) Set(val *UpdateFilterForDestinationV1Input) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterForDestinationV1Input) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterForDestinationV1Input) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterForDestinationV1Input(
	val *UpdateFilterForDestinationV1Input,
) *NullableUpdateFilterForDestinationV1Input {
	return &NullableUpdateFilterForDestinationV1Input{value: val, isSet: true}
}

func (v NullableUpdateFilterForDestinationV1Input) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterForDestinationV1Input) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
