# DestinationMetadataActionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The primary key of the action. | 
**Slug** | **string** | A machine-readable key unique to the action definition. | 
**Name** | **string** | A human-readable name for the action. | 
**Description** | **string** | A human-readable description of the action. May include Markdown. | 
**Platform** | **string** | The platform on which this action runs. | 
**Hidden** | **bool** | Whether the action should be hidden. | 
**DefaultTrigger** | **NullableString** | The default value used as the trigger when connecting this action. | 
**Fields** | [**[]DestinationMetadataActionFieldV1**](DestinationMetadataActionFieldV1.md) | The fields expected in order to perform the action. | 

## Methods

### NewDestinationMetadataActionV1

`func NewDestinationMetadataActionV1(id string, slug string, name string, description string, platform string, hidden bool, defaultTrigger NullableString, fields []DestinationMetadataActionFieldV1, ) *DestinationMetadataActionV1`

NewDestinationMetadataActionV1 instantiates a new DestinationMetadataActionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataActionV1WithDefaults

`func NewDestinationMetadataActionV1WithDefaults() *DestinationMetadataActionV1`

NewDestinationMetadataActionV1WithDefaults instantiates a new DestinationMetadataActionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationMetadataActionV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationMetadataActionV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationMetadataActionV1) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *DestinationMetadataActionV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DestinationMetadataActionV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DestinationMetadataActionV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *DestinationMetadataActionV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationMetadataActionV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationMetadataActionV1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DestinationMetadataActionV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationMetadataActionV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationMetadataActionV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPlatform

`func (o *DestinationMetadataActionV1) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DestinationMetadataActionV1) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DestinationMetadataActionV1) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetHidden

`func (o *DestinationMetadataActionV1) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *DestinationMetadataActionV1) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *DestinationMetadataActionV1) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetDefaultTrigger

`func (o *DestinationMetadataActionV1) GetDefaultTrigger() string`

GetDefaultTrigger returns the DefaultTrigger field if non-nil, zero value otherwise.

### GetDefaultTriggerOk

`func (o *DestinationMetadataActionV1) GetDefaultTriggerOk() (*string, bool)`

GetDefaultTriggerOk returns a tuple with the DefaultTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTrigger

`func (o *DestinationMetadataActionV1) SetDefaultTrigger(v string)`

SetDefaultTrigger sets DefaultTrigger field to given value.


### SetDefaultTriggerNil

`func (o *DestinationMetadataActionV1) SetDefaultTriggerNil(b bool)`

 SetDefaultTriggerNil sets the value for DefaultTrigger to be an explicit nil

### UnsetDefaultTrigger
`func (o *DestinationMetadataActionV1) UnsetDefaultTrigger()`

UnsetDefaultTrigger ensures that no value is present for DefaultTrigger, not even an explicit nil
### GetFields

`func (o *DestinationMetadataActionV1) GetFields() []DestinationMetadataActionFieldV1`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *DestinationMetadataActionV1) GetFieldsOk() (*[]DestinationMetadataActionFieldV1, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *DestinationMetadataActionV1) SetFields(v []DestinationMetadataActionFieldV1)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


