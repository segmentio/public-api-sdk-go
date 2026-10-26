# SpaceWarehouseSelectiveSyncItemAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | The collection within the Source. | 
**WarehouseId** | **string** | The id of the Warehouse this sync belongs to. | 
**Enabled** | **bool** | The Enabled flag ok telling whether the Collection is enabled or not. | 
**Properties** | **map[string]interface{}** | A map that contains the properties within the collection to which the Warehouse should sync. | 

## Methods

### NewSpaceWarehouseSelectiveSyncItemAlpha

`func NewSpaceWarehouseSelectiveSyncItemAlpha(collection string, warehouseId string, enabled bool, properties map[string]interface{}, ) *SpaceWarehouseSelectiveSyncItemAlpha`

NewSpaceWarehouseSelectiveSyncItemAlpha instantiates a new SpaceWarehouseSelectiveSyncItemAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceWarehouseSelectiveSyncItemAlphaWithDefaults

`func NewSpaceWarehouseSelectiveSyncItemAlphaWithDefaults() *SpaceWarehouseSelectiveSyncItemAlpha`

NewSpaceWarehouseSelectiveSyncItemAlphaWithDefaults instantiates a new SpaceWarehouseSelectiveSyncItemAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetWarehouseId

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetEnabled

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProperties

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SpaceWarehouseSelectiveSyncItemAlpha) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


