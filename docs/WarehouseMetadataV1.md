# WarehouseMetadataV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of this object. | 
**Name** | **string** | The name of this object. | 
**Slug** | **string** | A human-readable, unique identifier for object. | 
**Description** | **string** | A description, in English, of this object. | 
**Logos** | [**LogosBeta**](LogosBeta.md) |  | 
**Options** | [**[]IntegrationOptionBeta**](IntegrationOptionBeta.md) | The Integration options for this object. | 

## Methods

### NewWarehouseMetadataV1

`func NewWarehouseMetadataV1(id string, name string, slug string, description string, logos LogosBeta, options []IntegrationOptionBeta, ) *WarehouseMetadataV1`

NewWarehouseMetadataV1 instantiates a new WarehouseMetadataV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseMetadataV1WithDefaults

`func NewWarehouseMetadataV1WithDefaults() *WarehouseMetadataV1`

NewWarehouseMetadataV1WithDefaults instantiates a new WarehouseMetadataV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseMetadataV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseMetadataV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseMetadataV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WarehouseMetadataV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseMetadataV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseMetadataV1) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WarehouseMetadataV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WarehouseMetadataV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WarehouseMetadataV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *WarehouseMetadataV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WarehouseMetadataV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WarehouseMetadataV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogos

`func (o *WarehouseMetadataV1) GetLogos() LogosBeta`

GetLogos returns the Logos field if non-nil, zero value otherwise.

### GetLogosOk

`func (o *WarehouseMetadataV1) GetLogosOk() (*LogosBeta, bool)`

GetLogosOk returns a tuple with the Logos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogos

`func (o *WarehouseMetadataV1) SetLogos(v LogosBeta)`

SetLogos sets Logos field to given value.


### GetOptions

`func (o *WarehouseMetadataV1) GetOptions() []IntegrationOptionBeta`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WarehouseMetadataV1) GetOptionsOk() (*[]IntegrationOptionBeta, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WarehouseMetadataV1) SetOptions(v []IntegrationOptionBeta)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


