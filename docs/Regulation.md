# Regulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the regulate request. | 
**WorkspaceId** | **string** | The id of the Workspace that the regulate request belongs to. | 
**OverallStatus** | **string** | The current status of the regulate request. | 
**FinishedAt** | Pointer to **string** | The timestamp of when the request finished. | [optional] 
**CreatedAt** | **string** | The timestamp of the creation of the request. | 
**StreamStatus** | [**[]StreamStatusV1**](StreamStatusV1.md) | The status of each stream including all the Destinations that correspond to the stream. | 

## Methods

### NewRegulation

`func NewRegulation(id string, workspaceId string, overallStatus string, createdAt string, streamStatus []StreamStatusV1, ) *Regulation`

NewRegulation instantiates a new Regulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegulationWithDefaults

`func NewRegulationWithDefaults() *Regulation`

NewRegulationWithDefaults instantiates a new Regulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Regulation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Regulation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Regulation) SetId(v string)`

SetId sets Id field to given value.


### GetWorkspaceId

`func (o *Regulation) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Regulation) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Regulation) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetOverallStatus

`func (o *Regulation) GetOverallStatus() string`

GetOverallStatus returns the OverallStatus field if non-nil, zero value otherwise.

### GetOverallStatusOk

`func (o *Regulation) GetOverallStatusOk() (*string, bool)`

GetOverallStatusOk returns a tuple with the OverallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallStatus

`func (o *Regulation) SetOverallStatus(v string)`

SetOverallStatus sets OverallStatus field to given value.


### GetFinishedAt

`func (o *Regulation) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Regulation) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Regulation) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Regulation) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Regulation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Regulation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Regulation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetStreamStatus

`func (o *Regulation) GetStreamStatus() []StreamStatusV1`

GetStreamStatus returns the StreamStatus field if non-nil, zero value otherwise.

### GetStreamStatusOk

`func (o *Regulation) GetStreamStatusOk() (*[]StreamStatusV1, bool)`

GetStreamStatusOk returns a tuple with the StreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamStatus

`func (o *Regulation) SetStreamStatus(v []StreamStatusV1)`

SetStreamStatus sets StreamStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


