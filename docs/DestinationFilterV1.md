# DestinationFilterV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id of this filter. | 
**SourceId** | **string** | The id of the Source associated with this filter. | 
**DestinationId** | **string** | The id of the Destination associated with this filter. | 
**If** | **string** | A condition that defines whether to apply this filter to a payload. | 
**Actions** | [**[]DestinationFilterActionV1**](DestinationFilterActionV1.md) | A list of actions this filter performs. | 
**Title** | **string** | A title for this filter. | 
**Description** | Pointer to **string** | A description for this filter. | [optional] 
**Enabled** | **bool** | When set to true, this filter is active. | 
**CreatedAt** | **string** | The timestamp of this filter&#39;s creation. | 
**UpdatedAt** | **string** | The timestamp of this filter&#39;s last change. | 

## Methods

### NewDestinationFilterV1

`func NewDestinationFilterV1(id string, sourceId string, destinationId string, if_ string, actions []DestinationFilterActionV1, title string, enabled bool, createdAt string, updatedAt string, ) *DestinationFilterV1`

NewDestinationFilterV1 instantiates a new DestinationFilterV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationFilterV1WithDefaults

`func NewDestinationFilterV1WithDefaults() *DestinationFilterV1`

NewDestinationFilterV1WithDefaults instantiates a new DestinationFilterV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationFilterV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationFilterV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationFilterV1) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *DestinationFilterV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DestinationFilterV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DestinationFilterV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *DestinationFilterV1) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationFilterV1) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationFilterV1) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetIf

`func (o *DestinationFilterV1) GetIf() string`

GetIf returns the If field if non-nil, zero value otherwise.

### GetIfOk

`func (o *DestinationFilterV1) GetIfOk() (*string, bool)`

GetIfOk returns a tuple with the If field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIf

`func (o *DestinationFilterV1) SetIf(v string)`

SetIf sets If field to given value.


### GetActions

`func (o *DestinationFilterV1) GetActions() []DestinationFilterActionV1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DestinationFilterV1) GetActionsOk() (*[]DestinationFilterActionV1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DestinationFilterV1) SetActions(v []DestinationFilterActionV1)`

SetActions sets Actions field to given value.


### GetTitle

`func (o *DestinationFilterV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DestinationFilterV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DestinationFilterV1) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DestinationFilterV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationFilterV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationFilterV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationFilterV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationFilterV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationFilterV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationFilterV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *DestinationFilterV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationFilterV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationFilterV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DestinationFilterV1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationFilterV1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationFilterV1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


