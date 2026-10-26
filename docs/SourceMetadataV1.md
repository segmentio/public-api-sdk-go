# SourceMetadataV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id for this Source metadata in the Segment catalog.  Config API note: analogous to &#x60;name&#x60;. | 
**Name** | **string** | The user-friendly name of this Source.  Config API note: equal to &#x60;displayName&#x60;. | 
**Slug** | **string** | The slug that identifies this Source in the Segment app.  Config API note: equal to &#x60;name&#x60;. | 
**Description** | **string** | The description of this Source. | 
**Logos** | [**LogosBeta**](LogosBeta.md) |  | 
**Options** | [**[]IntegrationOptionBeta**](IntegrationOptionBeta.md) | Options for this Source. | 
**Categories** | **[]string** | A list of categories this Source belongs to. | 
**IsCloudEventSource** | **bool** | True if this is a Cloud Event Source. | 
**Status** | **string** | Support status of the Source. | 
**PartnerOwned** | Pointer to **bool** | Partner Owned flag. | [optional] 

## Methods

### NewSourceMetadataV1

`func NewSourceMetadataV1(id string, name string, slug string, description string, logos LogosBeta, options []IntegrationOptionBeta, categories []string, isCloudEventSource bool, status string, ) *SourceMetadataV1`

NewSourceMetadataV1 instantiates a new SourceMetadataV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceMetadataV1WithDefaults

`func NewSourceMetadataV1WithDefaults() *SourceMetadataV1`

NewSourceMetadataV1WithDefaults instantiates a new SourceMetadataV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceMetadataV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceMetadataV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceMetadataV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SourceMetadataV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceMetadataV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceMetadataV1) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *SourceMetadataV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SourceMetadataV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SourceMetadataV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *SourceMetadataV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceMetadataV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceMetadataV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLogos

`func (o *SourceMetadataV1) GetLogos() LogosBeta`

GetLogos returns the Logos field if non-nil, zero value otherwise.

### GetLogosOk

`func (o *SourceMetadataV1) GetLogosOk() (*LogosBeta, bool)`

GetLogosOk returns a tuple with the Logos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogos

`func (o *SourceMetadataV1) SetLogos(v LogosBeta)`

SetLogos sets Logos field to given value.


### GetOptions

`func (o *SourceMetadataV1) GetOptions() []IntegrationOptionBeta`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SourceMetadataV1) GetOptionsOk() (*[]IntegrationOptionBeta, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SourceMetadataV1) SetOptions(v []IntegrationOptionBeta)`

SetOptions sets Options field to given value.


### GetCategories

`func (o *SourceMetadataV1) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *SourceMetadataV1) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *SourceMetadataV1) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetIsCloudEventSource

`func (o *SourceMetadataV1) GetIsCloudEventSource() bool`

GetIsCloudEventSource returns the IsCloudEventSource field if non-nil, zero value otherwise.

### GetIsCloudEventSourceOk

`func (o *SourceMetadataV1) GetIsCloudEventSourceOk() (*bool, bool)`

GetIsCloudEventSourceOk returns a tuple with the IsCloudEventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudEventSource

`func (o *SourceMetadataV1) SetIsCloudEventSource(v bool)`

SetIsCloudEventSource sets IsCloudEventSource field to given value.


### GetStatus

`func (o *SourceMetadataV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceMetadataV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceMetadataV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPartnerOwned

`func (o *SourceMetadataV1) GetPartnerOwned() bool`

GetPartnerOwned returns the PartnerOwned field if non-nil, zero value otherwise.

### GetPartnerOwnedOk

`func (o *SourceMetadataV1) GetPartnerOwnedOk() (*bool, bool)`

GetPartnerOwnedOk returns a tuple with the PartnerOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerOwned

`func (o *SourceMetadataV1) SetPartnerOwned(v bool)`

SetPartnerOwned sets PartnerOwned field to given value.

### HasPartnerOwned

`func (o *SourceMetadataV1) HasPartnerOwned() bool`

HasPartnerOwned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


