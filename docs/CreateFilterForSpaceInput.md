# CreateFilterForSpaceInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The Space id to filter on. | 
**Enabled** | Pointer to **bool** | Whether the filter is enabled. | [optional] 
**Name** | **string** | The name of the filter. | 
**Description** | Pointer to **string** | The description of the filter. | [optional] 
**If** | **string** | The \&quot;if\&quot; statement for a filter. | 
**Drop** | Pointer to **bool** | Whether the event is dropped. | [optional] 

## Methods

### NewCreateFilterForSpaceInput

`func NewCreateFilterForSpaceInput(integrationId string, name string, if_ string, ) *CreateFilterForSpaceInput`

NewCreateFilterForSpaceInput instantiates a new CreateFilterForSpaceInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFilterForSpaceInputWithDefaults

`func NewCreateFilterForSpaceInputWithDefaults() *CreateFilterForSpaceInput`

NewCreateFilterForSpaceInputWithDefaults instantiates a new CreateFilterForSpaceInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *CreateFilterForSpaceInput) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CreateFilterForSpaceInput) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CreateFilterForSpaceInput) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetEnabled

`func (o *CreateFilterForSpaceInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateFilterForSpaceInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateFilterForSpaceInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateFilterForSpaceInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *CreateFilterForSpaceInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFilterForSpaceInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFilterForSpaceInput) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateFilterForSpaceInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFilterForSpaceInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFilterForSpaceInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFilterForSpaceInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIf

`func (o *CreateFilterForSpaceInput) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *CreateFilterForSpaceInput) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *CreateFilterForSpaceInput) SetIf(v string)`

SetIf sets If field to given value.


### GetDrop

`func (o *CreateFilterForSpaceInput) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *CreateFilterForSpaceInput) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *CreateFilterForSpaceInput) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *CreateFilterForSpaceInput) HasDrop() bool`

HasDrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


