# EntityPropertyAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Property name. | 
**PrimaryKey** | **bool** | Whether this is a primary key. | 
**DataType** | **string** | The data type of the property. | 
**Samples** | Pointer to **[]string** | Sample values. Only present when includeSampleValues is true. | [optional] 

## Methods

### NewEntityPropertyAlpha

`func NewEntityPropertyAlpha(name string, primaryKey bool, dataType string, ) *EntityPropertyAlpha`

NewEntityPropertyAlpha instantiates a new EntityPropertyAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertyAlphaWithDefaults

`func NewEntityPropertyAlphaWithDefaults() *EntityPropertyAlpha`

NewEntityPropertyAlphaWithDefaults instantiates a new EntityPropertyAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EntityPropertyAlpha) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityPropertyAlpha) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityPropertyAlpha) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryKey

`func (o *EntityPropertyAlpha) GetPrimaryKey() bool`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *EntityPropertyAlpha) GetPrimaryKeyOk() (*bool, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *EntityPropertyAlpha) SetPrimaryKey(v bool)`

SetPrimaryKey sets PrimaryKey field to given value.


### GetDataType

`func (o *EntityPropertyAlpha) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntityPropertyAlpha) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntityPropertyAlpha) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetSamples

`func (o *EntityPropertyAlpha) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *EntityPropertyAlpha) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *EntityPropertyAlpha) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *EntityPropertyAlpha) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


