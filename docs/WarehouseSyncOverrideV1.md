# WarehouseSyncOverrideV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The id of the Source this schema item applies to. | 
**Collection** | Pointer to **string** | The collection the schema item override applies to. | [optional] 
**Enabled** | **bool** | Enable to apply the override. | 
**Property** | Pointer to **string** | A property within the collection to which the override applies. | [optional] 

## Methods

### NewWarehouseSyncOverrideV1

`func NewWarehouseSyncOverrideV1(sourceId string, enabled bool, ) *WarehouseSyncOverrideV1`

NewWarehouseSyncOverrideV1 instantiates a new WarehouseSyncOverrideV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseSyncOverrideV1WithDefaults

`func NewWarehouseSyncOverrideV1WithDefaults() *WarehouseSyncOverrideV1`

NewWarehouseSyncOverrideV1WithDefaults instantiates a new WarehouseSyncOverrideV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *WarehouseSyncOverrideV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WarehouseSyncOverrideV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WarehouseSyncOverrideV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCollection

`func (o *WarehouseSyncOverrideV1) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *WarehouseSyncOverrideV1) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *WarehouseSyncOverrideV1) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *WarehouseSyncOverrideV1) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetEnabled

`func (o *WarehouseSyncOverrideV1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WarehouseSyncOverrideV1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WarehouseSyncOverrideV1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProperty

`func (o *WarehouseSyncOverrideV1) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *WarehouseSyncOverrideV1) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *WarehouseSyncOverrideV1) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *WarehouseSyncOverrideV1) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


