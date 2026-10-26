# UpdateFilterByIdInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | **string** | The Integration id of the resource. | 
**Enabled** | Pointer to **bool** | Whether the filter is enabled. | [optional] 
**Name** | Pointer to **string** | The name of the filter. | [optional] 
**Description** | Pointer to **string** | The description of the filter. | [optional] 
**If** | Pointer to **string** | The \&quot;if\&quot; statement for a filter. | [optional] 
**Drop** | Pointer to **bool** | Whether the event is dropped. | [optional] 

## Methods

### NewUpdateFilterByIdInput

`func NewUpdateFilterByIdInput(integrationId string, ) *UpdateFilterByIdInput`

NewUpdateFilterByIdInput instantiates a new UpdateFilterByIdInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFilterByIdInputWithDefaults

`func NewUpdateFilterByIdInputWithDefaults() *UpdateFilterByIdInput`

NewUpdateFilterByIdInputWithDefaults instantiates a new UpdateFilterByIdInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *UpdateFilterByIdInput) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UpdateFilterByIdInput) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UpdateFilterByIdInput) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetEnabled

`func (o *UpdateFilterByIdInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateFilterByIdInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateFilterByIdInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateFilterByIdInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *UpdateFilterByIdInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateFilterByIdInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateFilterByIdInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateFilterByIdInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateFilterByIdInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateFilterByIdInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateFilterByIdInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateFilterByIdInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIf

`func (o *UpdateFilterByIdInput) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *UpdateFilterByIdInput) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *UpdateFilterByIdInput) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *UpdateFilterByIdInput) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetDrop

`func (o *UpdateFilterByIdInput) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *UpdateFilterByIdInput) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *UpdateFilterByIdInput) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *UpdateFilterByIdInput) HasDrop() bool`

HasDrop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


