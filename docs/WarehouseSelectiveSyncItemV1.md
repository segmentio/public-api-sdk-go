# WarehouseSelectiveSyncItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The Source id attached to this Warehouse. | 
**Collection** | **string** | The collection within the Source. | 
**WarehouseId** | **string** | The id of the Warehouse this sync belongs to. | 
**Enabled** | **bool** | Whether this Selective Sync item is enabled. | 
**Properties** | **map[string]interface{}** | A map that contains the properties within the collection to which the Warehouse should sync. | 

## Methods

### NewWarehouseSelectiveSyncItemV1

`func NewWarehouseSelectiveSyncItemV1(sourceId string, collection string, warehouseId string, enabled bool, properties map[string]interface{}, ) *WarehouseSelectiveSyncItemV1`

NewWarehouseSelectiveSyncItemV1 instantiates a new WarehouseSelectiveSyncItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseSelectiveSyncItemV1WithDefaults

`func NewWarehouseSelectiveSyncItemV1WithDefaults() *WarehouseSelectiveSyncItemV1`

NewWarehouseSelectiveSyncItemV1WithDefaults instantiates a new WarehouseSelectiveSyncItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *WarehouseSelectiveSyncItemV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WarehouseSelectiveSyncItemV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WarehouseSelectiveSyncItemV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCollection

`func (o *WarehouseSelectiveSyncItemV1) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *WarehouseSelectiveSyncItemV1) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *WarehouseSelectiveSyncItemV1) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetWarehouseId

`func (o *WarehouseSelectiveSyncItemV1) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *WarehouseSelectiveSyncItemV1) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *WarehouseSelectiveSyncItemV1) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.


### GetEnabled

`func (o *WarehouseSelectiveSyncItemV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WarehouseSelectiveSyncItemV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WarehouseSelectiveSyncItemV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProperties

`func (o *WarehouseSelectiveSyncItemV1) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WarehouseSelectiveSyncItemV1) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WarehouseSelectiveSyncItemV1) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


