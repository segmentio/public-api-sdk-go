# CreateValidationInWarehouseV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetadataId** | **string** | The id of the Warehouse metadata type. | 
**Settings** | **map[string]interface{}** | A key-value object that contains instance-specific Warehouse settings. | 

## Methods

### NewCreateValidationInWarehouseV1Input

`func NewCreateValidationInWarehouseV1Input(metadataId string, settings map[string]interface{}, ) *CreateValidationInWarehouseV1Input`

NewCreateValidationInWarehouseV1Input instantiates a new CreateValidationInWarehouseV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateValidationInWarehouseV1InputWithDefaults

`func NewCreateValidationInWarehouseV1InputWithDefaults() *CreateValidationInWarehouseV1Input`

NewCreateValidationInWarehouseV1InputWithDefaults instantiates a new CreateValidationInWarehouseV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadataId

`func (o *CreateValidationInWarehouseV1Input) GetMetadataId() string`

GetMetadataId returns the MetadataId field if non-nil, zero value otherwise.

### GetMetadataIdOk

`func (o *CreateValidationInWarehouseV1Input) GetMetadataIdOk() (*string, bool)`

GetMetadataIdOk returns a tuple with the MetadataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataId

`func (o *CreateValidationInWarehouseV1Input) SetMetadataId(v string)`

SetMetadataId sets MetadataId field to given value.


### GetSettings

`func (o *CreateValidationInWarehouseV1Input) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateValidationInWarehouseV1Input) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateValidationInWarehouseV1Input) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


