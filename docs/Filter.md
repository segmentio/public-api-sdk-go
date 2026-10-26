# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The newly created filter id. | 
**WorkspaceId** | **string** | The Workspace id to create the filter. | 
**IntegrationId** | **string** | The Integration id of the resource. | 
**Enabled** | Pointer to **bool** | Whether the filter is enabled. | [optional] 
**Name** | Pointer to **string** | The name of the filter. | [optional] 
**Description** | Pointer to **string** | The description of the filter. | [optional] 
**If** | Pointer to **string** | The \&quot;if\&quot; statement for a filter. | [optional] 
**Drop** | Pointer to **bool** | Whether the event is dropped. | [optional] 
**CreatedAt** | **string** | The timestamp of this filter&#39;s creation. | 
**UpdatedAt** | **string** | The timestamp of this filter&#39;s last change. | 

## Methods

### NewFilter

`func NewFilter(id string, workspaceId string, integrationId string, createdAt string, updatedAt string, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Filter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Filter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Filter) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *Filter) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Filter) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Filter) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetIntegrationId

`func (o *Filter) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *Filter) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *Filter) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetEnabled

`func (o *Filter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Filter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Filter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Filter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *Filter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Filter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Filter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Filter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Filter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIf

`func (o *Filter) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *Filter) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *Filter) SetIf(v string)`

SetIf sets If field to given value.

### HasIf

`func (o *Filter) HasIf() bool`

HasIf returns a boolean if a field has been set.

### GetDrop

`func (o *Filter) GetDrop() bool`

GetDrop returns the Drop field if non-nil, zero value otherwise.

### GetDropOk

`func (o *Filter) GetDropOk() (*bool, bool)`

GetDropOk returns a tuple with the Drop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrop

`func (o *Filter) SetDrop(v bool)`

SetDrop sets Drop field to given value.

### HasDrop

`func (o *Filter) HasDrop() bool`

HasDrop returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Filter) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Filter) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Filter) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Filter) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Filter) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Filter) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


