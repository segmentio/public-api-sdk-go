# DestinationMetadataActionFieldV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The primary key of the field. | 
**SortOrder** | **float32** | The order this particular field is (used in the UI for displaying the fields in a specified order). | 
**FieldKey** | **string** | A unique machine-readable key for the field. Should ideally match the expected key in the action\\&#39;s API request. | 
**Label** | **string** | A human-readable label for this value. | 
**Type** | **string** | The data type for this value. | 
**Description** | **string** | A human-readable description of this value. You can use Markdown. | 
**Placeholder** | Pointer to **string** | An example value displayed but not saved. | [optional] 
**DefaultValue** | Pointer to **interface{}** | A default value that is saved the first time an action is created. | [optional] 
**Required** | **bool** | Whether this field is required. | 
**Multiple** | **bool** | Whether a user can provide multiples of this field. | 
**Choices** | Pointer to **interface{}** | A list of machine-readable value/label pairs to populate a static dropdown. | [optional] 
**Dynamic** | **bool** | Whether this field should execute a dynamic request to fetch choices to populate a dropdown. When true, &#x60;choices&#x60; is ignored. | 
**AllowNull** | **bool** | Whether this field allows null values. | 
**Hidden** | Pointer to **bool** | Whether the action field should be hidden or not. | [optional] 

## Methods

### NewDestinationMetadataActionFieldV1

`func NewDestinationMetadataActionFieldV1(id string, sortOrder float32, fieldKey string, label string, type_ string, description string, required bool, multiple bool, dynamic bool, allowNull bool, ) *DestinationMetadataActionFieldV1`

NewDestinationMetadataActionFieldV1 instantiates a new DestinationMetadataActionFieldV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataActionFieldV1WithDefaults

`func NewDestinationMetadataActionFieldV1WithDefaults() *DestinationMetadataActionFieldV1`

NewDestinationMetadataActionFieldV1WithDefaults instantiates a new DestinationMetadataActionFieldV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationMetadataActionFieldV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationMetadataActionFieldV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationMetadataActionFieldV1) SetId(v string)`

SetId sets Id field to given value.


### GetSortOrder

`func (o *DestinationMetadataActionFieldV1) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *DestinationMetadataActionFieldV1) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *DestinationMetadataActionFieldV1) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetFieldKey

`func (o *DestinationMetadataActionFieldV1) GetFieldKey() string`

GetFieldKey returns the FieldKey field if non-nil, zero value otherwise.

### GetFieldKeyOk

`func (o *DestinationMetadataActionFieldV1) GetFieldKeyOk() (*string, bool)`

GetFieldKeyOk returns a tuple with the FieldKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldKey

`func (o *DestinationMetadataActionFieldV1) SetFieldKey(v string)`

SetFieldKey sets FieldKey field to given value.


### GetLabel

`func (o *DestinationMetadataActionFieldV1) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DestinationMetadataActionFieldV1) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DestinationMetadataActionFieldV1) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *DestinationMetadataActionFieldV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationMetadataActionFieldV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationMetadataActionFieldV1) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *DestinationMetadataActionFieldV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationMetadataActionFieldV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationMetadataActionFieldV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlaceholder

`func (o *DestinationMetadataActionFieldV1) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *DestinationMetadataActionFieldV1) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *DestinationMetadataActionFieldV1) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *DestinationMetadataActionFieldV1) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetDefaultValue

`func (o *DestinationMetadataActionFieldV1) GetDefaultValue() interface{}`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DestinationMetadataActionFieldV1) GetDefaultValueOk() (*interface{}, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DestinationMetadataActionFieldV1) SetDefaultValue(v interface{})`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DestinationMetadataActionFieldV1) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *DestinationMetadataActionFieldV1) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *DestinationMetadataActionFieldV1) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetRequired

`func (o *DestinationMetadataActionFieldV1) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DestinationMetadataActionFieldV1) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DestinationMetadataActionFieldV1) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetMultiple

`func (o *DestinationMetadataActionFieldV1) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *DestinationMetadataActionFieldV1) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *DestinationMetadataActionFieldV1) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.


### GetChoices

`func (o *DestinationMetadataActionFieldV1) GetChoices() interface{}`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *DestinationMetadataActionFieldV1) GetChoicesOk() (*interface{}, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *DestinationMetadataActionFieldV1) SetChoices(v interface{})`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *DestinationMetadataActionFieldV1) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### SetChoicesNil

`func (o *DestinationMetadataActionFieldV1) SetChoicesNil(b bool)`

 SetChoicesNil sets the value for Choices to be an explicit nil

### UnsetChoices
`func (o *DestinationMetadataActionFieldV1) UnsetChoices()`

UnsetChoices ensures that no value is present for Choices, not even an explicit nil
### GetDynamic

`func (o *DestinationMetadataActionFieldV1) GetDynamic() bool`

GetDynamic returns the Dynamic field if non-nil, zero value otherwise.

### GetDynamicOk

`func (o *DestinationMetadataActionFieldV1) GetDynamicOk() (*bool, bool)`

GetDynamicOk returns a tuple with the Dynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic

`func (o *DestinationMetadataActionFieldV1) SetDynamic(v bool)`

SetDynamic sets Dynamic field to given value.


### GetAllowNull

`func (o *DestinationMetadataActionFieldV1) GetAllowNull() bool`

GetAllowNull returns the AllowNull field if non-nil, zero value otherwise.

### GetAllowNullOk

`func (o *DestinationMetadataActionFieldV1) GetAllowNullOk() (*bool, bool)`

GetAllowNullOk returns a tuple with the AllowNull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNull

`func (o *DestinationMetadataActionFieldV1) SetAllowNull(v bool)`

SetAllowNull sets AllowNull field to given value.


### GetHidden

`func (o *DestinationMetadataActionFieldV1) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DestinationMetadataActionFieldV1) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DestinationMetadataActionFieldV1) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *DestinationMetadataActionFieldV1) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


