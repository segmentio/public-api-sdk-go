/*
Segment Public API

The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.  All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.  See the next sections for more information on how to use the Segment Public API.

API version: 39.0.0
Contact: friends@segment.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DestinationMetadataActionFieldV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationMetadataActionFieldV1{}

// DestinationMetadataActionFieldV1 Represents a field used in configuring an action.
type DestinationMetadataActionFieldV1 struct {
	// The primary key of the field.
	Id string `json:"id"`
	// The order this particular field is (used in the UI for displaying the fields in a specified order).
	SortOrder float32 `json:"sortOrder"`
	// A unique machine-readable key for the field. Should ideally match the expected key in the action\\'s API request.
	FieldKey string `json:"fieldKey"`
	// A human-readable label for this value.
	Label string `json:"label"`
	// The data type for this value.
	Type string `json:"type"`
	// A human-readable description of this value. You can use Markdown.
	Description string `json:"description"`
	// An example value displayed but not saved.
	Placeholder *string `json:"placeholder,omitempty"`
	// A default value that is saved the first time an action is created.
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	// Whether this field is required.
	Required bool `json:"required"`
	// Whether a user can provide multiples of this field.
	Multiple bool `json:"multiple"`
	// A list of machine-readable value/label pairs to populate a static dropdown.
	Choices interface{} `json:"choices,omitempty"`
	// Whether this field should execute a dynamic request to fetch choices to populate a dropdown. When true, `choices` is ignored.
	Dynamic bool `json:"dynamic"`
	// Whether this field allows null values.
	AllowNull bool `json:"allowNull"`
}

// NewDestinationMetadataActionFieldV1 instantiates a new DestinationMetadataActionFieldV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationMetadataActionFieldV1(
	id string,
	sortOrder float32,
	fieldKey string,
	label string,
	type_ string,
	description string,
	required bool,
	multiple bool,
	dynamic bool,
	allowNull bool,
) *DestinationMetadataActionFieldV1 {
	this := DestinationMetadataActionFieldV1{}
	this.Id = id
	this.SortOrder = sortOrder
	this.FieldKey = fieldKey
	this.Label = label
	this.Type = type_
	this.Description = description
	this.Required = required
	this.Multiple = multiple
	this.Dynamic = dynamic
	this.AllowNull = allowNull
	return &this
}

// NewDestinationMetadataActionFieldV1WithDefaults instantiates a new DestinationMetadataActionFieldV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationMetadataActionFieldV1WithDefaults() *DestinationMetadataActionFieldV1 {
	this := DestinationMetadataActionFieldV1{}
	return &this
}

// GetId returns the Id field value
func (o *DestinationMetadataActionFieldV1) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DestinationMetadataActionFieldV1) SetId(v string) {
	o.Id = v
}

// GetSortOrder returns the SortOrder field value
func (o *DestinationMetadataActionFieldV1) GetSortOrder() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetSortOrderOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrder, true
}

// SetSortOrder sets field value
func (o *DestinationMetadataActionFieldV1) SetSortOrder(v float32) {
	o.SortOrder = v
}

// GetFieldKey returns the FieldKey field value
func (o *DestinationMetadataActionFieldV1) GetFieldKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetFieldKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldKey, true
}

// SetFieldKey sets field value
func (o *DestinationMetadataActionFieldV1) SetFieldKey(v string) {
	o.FieldKey = v
}

// GetLabel returns the Label field value
func (o *DestinationMetadataActionFieldV1) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *DestinationMetadataActionFieldV1) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *DestinationMetadataActionFieldV1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationMetadataActionFieldV1) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *DestinationMetadataActionFieldV1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DestinationMetadataActionFieldV1) SetDescription(v string) {
	o.Description = v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *DestinationMetadataActionFieldV1) GetPlaceholder() string {
	if o == nil || IsNil(o.Placeholder) {
		var ret string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetPlaceholderOk() (*string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *DestinationMetadataActionFieldV1) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given string and assigns it to the Placeholder field.
func (o *DestinationMetadataActionFieldV1) SetPlaceholder(v string) {
	o.Placeholder = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DestinationMetadataActionFieldV1) GetDefaultValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestinationMetadataActionFieldV1) GetDefaultValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return &o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *DestinationMetadataActionFieldV1) HasDefaultValue() bool {
	if o != nil && IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given interface{} and assigns it to the DefaultValue field.
func (o *DestinationMetadataActionFieldV1) SetDefaultValue(v interface{}) {
	o.DefaultValue = v
}

// GetRequired returns the Required field value
func (o *DestinationMetadataActionFieldV1) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *DestinationMetadataActionFieldV1) SetRequired(v bool) {
	o.Required = v
}

// GetMultiple returns the Multiple field value
func (o *DestinationMetadataActionFieldV1) GetMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Multiple
}

// GetMultipleOk returns a tuple with the Multiple field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Multiple, true
}

// SetMultiple sets field value
func (o *DestinationMetadataActionFieldV1) SetMultiple(v bool) {
	o.Multiple = v
}

// GetChoices returns the Choices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DestinationMetadataActionFieldV1) GetChoices() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Choices
}

// GetChoicesOk returns a tuple with the Choices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DestinationMetadataActionFieldV1) GetChoicesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Choices) {
		return nil, false
	}
	return &o.Choices, true
}

// HasChoices returns a boolean if a field has been set.
func (o *DestinationMetadataActionFieldV1) HasChoices() bool {
	if o != nil && IsNil(o.Choices) {
		return true
	}

	return false
}

// SetChoices gets a reference to the given interface{} and assigns it to the Choices field.
func (o *DestinationMetadataActionFieldV1) SetChoices(v interface{}) {
	o.Choices = v
}

// GetDynamic returns the Dynamic field value
func (o *DestinationMetadataActionFieldV1) GetDynamic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetDynamicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dynamic, true
}

// SetDynamic sets field value
func (o *DestinationMetadataActionFieldV1) SetDynamic(v bool) {
	o.Dynamic = v
}

// GetAllowNull returns the AllowNull field value
func (o *DestinationMetadataActionFieldV1) GetAllowNull() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowNull
}

// GetAllowNullOk returns a tuple with the AllowNull field value
// and a boolean to check if the value has been set.
func (o *DestinationMetadataActionFieldV1) GetAllowNullOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowNull, true
}

// SetAllowNull sets field value
func (o *DestinationMetadataActionFieldV1) SetAllowNull(v bool) {
	o.AllowNull = v
}

func (o DestinationMetadataActionFieldV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationMetadataActionFieldV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["sortOrder"] = o.SortOrder
	toSerialize["fieldKey"] = o.FieldKey
	toSerialize["label"] = o.Label
	toSerialize["type"] = o.Type
	toSerialize["description"] = o.Description
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	toSerialize["required"] = o.Required
	toSerialize["multiple"] = o.Multiple
	if o.Choices != nil {
		toSerialize["choices"] = o.Choices
	}
	toSerialize["dynamic"] = o.Dynamic
	toSerialize["allowNull"] = o.AllowNull
	return toSerialize, nil
}

type NullableDestinationMetadataActionFieldV1 struct {
	value *DestinationMetadataActionFieldV1
	isSet bool
}

func (v NullableDestinationMetadataActionFieldV1) Get() *DestinationMetadataActionFieldV1 {
	return v.value
}

func (v *NullableDestinationMetadataActionFieldV1) Set(val *DestinationMetadataActionFieldV1) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationMetadataActionFieldV1) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationMetadataActionFieldV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationMetadataActionFieldV1(
	val *DestinationMetadataActionFieldV1,
) *NullableDestinationMetadataActionFieldV1 {
	return &NullableDestinationMetadataActionFieldV1{value: val, isSet: true}
}

func (v NullableDestinationMetadataActionFieldV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationMetadataActionFieldV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
