# SpaceWarehouseSchemaOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | The collection within the Source. | 
**Enabled** | **bool** | Represents the overridden enabled state for the listed collection and / or properties. | 
**Property** | Pointer to **string** | A map that contains the properties within the collection to which the Warehouse should sync. | [optional] 

## Methods

### NewSpaceWarehouseSchemaOverride

`func NewSpaceWarehouseSchemaOverride(collection string, enabled bool, ) *SpaceWarehouseSchemaOverride`

NewSpaceWarehouseSchemaOverride instantiates a new SpaceWarehouseSchemaOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceWarehouseSchemaOverrideWithDefaults

`func NewSpaceWarehouseSchemaOverrideWithDefaults() *SpaceWarehouseSchemaOverride`

NewSpaceWarehouseSchemaOverrideWithDefaults instantiates a new SpaceWarehouseSchemaOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *SpaceWarehouseSchemaOverride) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *SpaceWarehouseSchemaOverride) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *SpaceWarehouseSchemaOverride) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetEnabled

`func (o *SpaceWarehouseSchemaOverride) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SpaceWarehouseSchemaOverride) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SpaceWarehouseSchemaOverride) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProperty

`func (o *SpaceWarehouseSchemaOverride) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SpaceWarehouseSchemaOverride) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SpaceWarehouseSchemaOverride) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SpaceWarehouseSchemaOverride) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


