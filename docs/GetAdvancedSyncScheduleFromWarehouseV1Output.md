# GetAdvancedSyncScheduleFromWarehouseV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if an advanced sync schedule is enabled for this Warehouse. | 
**Schedule** | Pointer to [**AdvancedWarehouseSyncScheduleV1Output**](AdvancedWarehouseSyncScheduleV1Output.md) |  | [optional] 

## Methods

### NewGetAdvancedSyncScheduleFromWarehouseV1Output

`func NewGetAdvancedSyncScheduleFromWarehouseV1Output(enabled bool, ) *GetAdvancedSyncScheduleFromWarehouseV1Output`

NewGetAdvancedSyncScheduleFromWarehouseV1Output instantiates a new GetAdvancedSyncScheduleFromWarehouseV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdvancedSyncScheduleFromWarehouseV1OutputWithDefaults

`func NewGetAdvancedSyncScheduleFromWarehouseV1OutputWithDefaults() *GetAdvancedSyncScheduleFromWarehouseV1Output`

NewGetAdvancedSyncScheduleFromWarehouseV1OutputWithDefaults instantiates a new GetAdvancedSyncScheduleFromWarehouseV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchedule

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetSchedule() AdvancedWarehouseSyncScheduleV1Output`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) GetScheduleOk() (*AdvancedWarehouseSyncScheduleV1Output, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) SetSchedule(v AdvancedWarehouseSyncScheduleV1Output)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetAdvancedSyncScheduleFromWarehouseV1Output) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


