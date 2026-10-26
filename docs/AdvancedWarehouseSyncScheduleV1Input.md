# AdvancedWarehouseSyncScheduleV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Times** | [**[]WarehouseAdvancedSyncV1**](WarehouseAdvancedSyncV1.md) | A list that contains the times when syncs should occur. | 
**Timezone** | **string** | A TZ-database timezone for this sync schedule. | 

## Methods

### NewAdvancedWarehouseSyncScheduleV1Input

`func NewAdvancedWarehouseSyncScheduleV1Input(times []WarehouseAdvancedSyncV1, timezone string, ) *AdvancedWarehouseSyncScheduleV1Input`

NewAdvancedWarehouseSyncScheduleV1Input instantiates a new AdvancedWarehouseSyncScheduleV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedWarehouseSyncScheduleV1InputWithDefaults

`func NewAdvancedWarehouseSyncScheduleV1InputWithDefaults() *AdvancedWarehouseSyncScheduleV1Input`

NewAdvancedWarehouseSyncScheduleV1InputWithDefaults instantiates a new AdvancedWarehouseSyncScheduleV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimes

`func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimes() []WarehouseAdvancedSyncV1`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimesOk() (*[]WarehouseAdvancedSyncV1, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *AdvancedWarehouseSyncScheduleV1Input) SetTimes(v []WarehouseAdvancedSyncV1)`

SetTimes sets Times field to given value.


### GetTimezone

`func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AdvancedWarehouseSyncScheduleV1Input) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AdvancedWarehouseSyncScheduleV1Input) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


