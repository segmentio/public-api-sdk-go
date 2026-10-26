# AdvancedWarehouseSyncScheduleV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Times** | Pointer to [**[]WarehouseAdvancedSyncV1**](WarehouseAdvancedSyncV1.md) | A list that contains the times when syncs should occur. | [optional] 
**Timezone** | Pointer to **string** | A TZ-database timezone for this sync schedule. | [optional] 

## Methods

### NewAdvancedWarehouseSyncScheduleV1Output

`func NewAdvancedWarehouseSyncScheduleV1Output() *AdvancedWarehouseSyncScheduleV1Output`

NewAdvancedWarehouseSyncScheduleV1Output instantiates a new AdvancedWarehouseSyncScheduleV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedWarehouseSyncScheduleV1OutputWithDefaults

`func NewAdvancedWarehouseSyncScheduleV1OutputWithDefaults() *AdvancedWarehouseSyncScheduleV1Output`

NewAdvancedWarehouseSyncScheduleV1OutputWithDefaults instantiates a new AdvancedWarehouseSyncScheduleV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimes

`func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimes() []WarehouseAdvancedSyncV1`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimesOk() (*[]WarehouseAdvancedSyncV1, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *AdvancedWarehouseSyncScheduleV1Output) SetTimes(v []WarehouseAdvancedSyncV1)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *AdvancedWarehouseSyncScheduleV1Output) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTimezone

`func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AdvancedWarehouseSyncScheduleV1Output) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AdvancedWarehouseSyncScheduleV1Output) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AdvancedWarehouseSyncScheduleV1Output) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


