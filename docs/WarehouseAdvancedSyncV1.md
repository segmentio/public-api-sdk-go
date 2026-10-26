# WarehouseAdvancedSyncV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HourOfDay** | **float32** | The hour of day for which to enable/disable a sync, between 0 and 23. | 
**Enabled** | **bool** | Enable to the sync at the specified hour. | 

## Methods

### NewWarehouseAdvancedSyncV1

`func NewWarehouseAdvancedSyncV1(hourOfDay float32, enabled bool, ) *WarehouseAdvancedSyncV1`

NewWarehouseAdvancedSyncV1 instantiates a new WarehouseAdvancedSyncV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseAdvancedSyncV1WithDefaults

`func NewWarehouseAdvancedSyncV1WithDefaults() *WarehouseAdvancedSyncV1`

NewWarehouseAdvancedSyncV1WithDefaults instantiates a new WarehouseAdvancedSyncV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHourOfDay

`func (o *WarehouseAdvancedSyncV1) GetHourOfDay() float32`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *WarehouseAdvancedSyncV1) GetHourOfDayOk() (*float32, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *WarehouseAdvancedSyncV1) SetHourOfDay(v float32)`

SetHourOfDay sets HourOfDay field to given value.


### GetEnabled

`func (o *WarehouseAdvancedSyncV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WarehouseAdvancedSyncV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WarehouseAdvancedSyncV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


