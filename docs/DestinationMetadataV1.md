# DestinationMetadataV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Destination metadata.  Config API note: analogous to &#x60;name&#x60;. | 
**Name** | **string** | The user-friendly name of the Destination.  Config API note: equal to &#x60;displayName&#x60;. | 
**Description** | **string** | The description of the Destination. | 
**Slug** | **string** | The slug used to identify the Destination in the Segment app. | 
**Logos** | [**LogosBeta**](LogosBeta.md) |  | 
**Options** | [**[]IntegrationOptionBeta**](IntegrationOptionBeta.md) | Options configured for the Destination. | 
**Status** | **string** | Support status of the Destination. | 
**PreviousNames** | **[]string** | A list of names previously used by the Destination. | 
**Categories** | **[]string** | A list of categories with which the Destination is associated. | 
**Website** | **string** | A website URL for this Destination. | 
**Components** | [**[]DestinationMetadataComponentV1**](DestinationMetadataComponentV1.md) | A list of components this Destination provides. | 
**SupportedFeatures** | [**DestinationMetadataFeaturesV1**](DestinationMetadataFeaturesV1.md) |  | 
**SupportedMethods** | [**DestinationMetadataMethodsV1**](DestinationMetadataMethodsV1.md) |  | 
**SupportedPlatforms** | [**DestinationMetadataPlatformsV1**](DestinationMetadataPlatformsV1.md) |  | 
**Actions** | [**[]DestinationMetadataActionV1**](DestinationMetadataActionV1.md) | Actions available for the Destination. | 
**Presets** | [**[]DestinationMetadataSubscriptionPresetV1**](DestinationMetadataSubscriptionPresetV1.md) | Predefined Destination subscriptions that can optionally be applied when connecting a new instance of the Destination. | 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | Contact info for Integration Owners. | [optional] 
**PartnerOwned** | Pointer to **bool** | Partner Owned flag. | [optional] 
**SupportedRegions** | Pointer to **[]string** | A list of supported regions for this Destination. | [optional] 
**RegionEndpoints** | Pointer to **[]string** | The list of regional endpoints for this Destination. | [optional] 
**MultiInstanceSupportedVersion** | Pointer to **string** | This Destination&#39;s support for multiple instance types. | [optional] 

## Methods

### NewDestinationMetadataV1

`func NewDestinationMetadataV1(id string, name string, description string, slug string, logos LogosBeta, options []IntegrationOptionBeta, status string, previousNames []string, categories []string, website string, components []DestinationMetadataComponentV1, supportedFeatures DestinationMetadataFeaturesV1, supportedMethods DestinationMetadataMethodsV1, supportedPlatforms DestinationMetadataPlatformsV1, actions []DestinationMetadataActionV1, presets []DestinationMetadataSubscriptionPresetV1, ) *DestinationMetadataV1`

NewDestinationMetadataV1 instantiates a new DestinationMetadataV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataV1WithDefaults

`func NewDestinationMetadataV1WithDefaults() *DestinationMetadataV1`

NewDestinationMetadataV1WithDefaults instantiates a new DestinationMetadataV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DestinationMetadataV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationMetadataV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationMetadataV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DestinationMetadataV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationMetadataV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationMetadataV1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DestinationMetadataV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationMetadataV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationMetadataV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSlug

`func (o *DestinationMetadataV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *DestinationMetadataV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *DestinationMetadataV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetLogos

`func (o *DestinationMetadataV1) GetLogos() LogosBeta`

GetLogos returns the Logos field if non-nil, zero value otherwise.

### GetLogosOk

`func (o *DestinationMetadataV1) GetLogosOk() (*LogosBeta, bool)`

GetLogosOk returns a tuple with the Logos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogos

`func (o *DestinationMetadataV1) SetLogos(v LogosBeta)`

SetLogos sets Logos field to given value.


### GetOptions

`func (o *DestinationMetadataV1) GetOptions() []IntegrationOptionBeta`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DestinationMetadataV1) GetOptionsOk() (*[]IntegrationOptionBeta, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DestinationMetadataV1) SetOptions(v []IntegrationOptionBeta)`

SetOptions sets Options field to given value.


### GetStatus

`func (o *DestinationMetadataV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DestinationMetadataV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DestinationMetadataV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPreviousNames

`func (o *DestinationMetadataV1) GetPreviousNames() []string`

GetPreviousNames returns the PreviousNames field if non-nil, zero value otherwise.

### GetPreviousNamesOk

`func (o *DestinationMetadataV1) GetPreviousNamesOk() (*[]string, bool)`

GetPreviousNamesOk returns a tuple with the PreviousNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousNames

`func (o *DestinationMetadataV1) SetPreviousNames(v []string)`

SetPreviousNames sets PreviousNames field to given value.


### GetCategories

`func (o *DestinationMetadataV1) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *DestinationMetadataV1) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *DestinationMetadataV1) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetWebsite

`func (o *DestinationMetadataV1) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *DestinationMetadataV1) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *DestinationMetadataV1) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetComponents

`func (o *DestinationMetadataV1) GetComponents() []DestinationMetadataComponentV1`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *DestinationMetadataV1) GetComponentsOk() (*[]DestinationMetadataComponentV1, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *DestinationMetadataV1) SetComponents(v []DestinationMetadataComponentV1)`

SetComponents sets Components field to given value.


### GetSupportedFeatures

`func (o *DestinationMetadataV1) GetSupportedFeatures() DestinationMetadataFeaturesV1`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DestinationMetadataV1) GetSupportedFeaturesOk() (*DestinationMetadataFeaturesV1, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DestinationMetadataV1) SetSupportedFeatures(v DestinationMetadataFeaturesV1)`

SetSupportedFeatures sets SupportedFeatures field to given value.


### GetSupportedMethods

`func (o *DestinationMetadataV1) GetSupportedMethods() DestinationMetadataMethodsV1`

GetSupportedMethods returns the SupportedMethods field if non-nil, zero value otherwise.

### GetSupportedMethodsOk

`func (o *DestinationMetadataV1) GetSupportedMethodsOk() (*DestinationMetadataMethodsV1, bool)`

GetSupportedMethodsOk returns a tuple with the SupportedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMethods

`func (o *DestinationMetadataV1) SetSupportedMethods(v DestinationMetadataMethodsV1)`

SetSupportedMethods sets SupportedMethods field to given value.


### GetSupportedPlatforms

`func (o *DestinationMetadataV1) GetSupportedPlatforms() DestinationMetadataPlatformsV1`

GetSupportedPlatforms returns the SupportedPlatforms field if non-nil, zero value otherwise.

### GetSupportedPlatformsOk

`func (o *DestinationMetadataV1) GetSupportedPlatformsOk() (*DestinationMetadataPlatformsV1, bool)`

GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedPlatforms

`func (o *DestinationMetadataV1) SetSupportedPlatforms(v DestinationMetadataPlatformsV1)`

SetSupportedPlatforms sets SupportedPlatforms field to given value.


### GetActions

`func (o *DestinationMetadataV1) GetActions() []DestinationMetadataActionV1`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DestinationMetadataV1) GetActionsOk() (*[]DestinationMetadataActionV1, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DestinationMetadataV1) SetActions(v []DestinationMetadataActionV1)`

SetActions sets Actions field to given value.


### GetPresets

`func (o *DestinationMetadataV1) GetPresets() []DestinationMetadataSubscriptionPresetV1`

GetPresets returns the Presets field if non-nil, zero value otherwise.

### GetPresetsOk

`func (o *DestinationMetadataV1) GetPresetsOk() (*[]DestinationMetadataSubscriptionPresetV1, bool)`

GetPresetsOk returns a tuple with the Presets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresets

`func (o *DestinationMetadataV1) SetPresets(v []DestinationMetadataSubscriptionPresetV1)`

SetPresets sets Presets field to given value.


### GetContacts

`func (o *DestinationMetadataV1) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *DestinationMetadataV1) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *DestinationMetadataV1) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *DestinationMetadataV1) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetPartnerOwned

`func (o *DestinationMetadataV1) GetPartnerOwned() bool`

GetPartnerOwned returns the PartnerOwned field if non-nil, zero value otherwise.

### GetPartnerOwnedOk

`func (o *DestinationMetadataV1) GetPartnerOwnedOk() (*bool, bool)`

GetPartnerOwnedOk returns a tuple with the PartnerOwned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerOwned

`func (o *DestinationMetadataV1) SetPartnerOwned(v bool)`

SetPartnerOwned sets PartnerOwned field to given value.

### HasPartnerOwned

`func (o *DestinationMetadataV1) HasPartnerOwned() bool`

HasPartnerOwned returns a boolean if a field has been set.

### GetSupportedRegions

`func (o *DestinationMetadataV1) GetSupportedRegions() []string`

GetSupportedRegions returns the SupportedRegions field if non-nil, zero value otherwise.

### GetSupportedRegionsOk

`func (o *DestinationMetadataV1) GetSupportedRegionsOk() (*[]string, bool)`

GetSupportedRegionsOk returns a tuple with the SupportedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRegions

`func (o *DestinationMetadataV1) SetSupportedRegions(v []string)`

SetSupportedRegions sets SupportedRegions field to given value.

### HasSupportedRegions

`func (o *DestinationMetadataV1) HasSupportedRegions() bool`

HasSupportedRegions returns a boolean if a field has been set.

### GetRegionEndpoints

`func (o *DestinationMetadataV1) GetRegionEndpoints() []string`

GetRegionEndpoints returns the RegionEndpoints field if non-nil, zero value otherwise.

### GetRegionEndpointsOk

`func (o *DestinationMetadataV1) GetRegionEndpointsOk() (*[]string, bool)`

GetRegionEndpointsOk returns a tuple with the RegionEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionEndpoints

`func (o *DestinationMetadataV1) SetRegionEndpoints(v []string)`

SetRegionEndpoints sets RegionEndpoints field to given value.

### HasRegionEndpoints

`func (o *DestinationMetadataV1) HasRegionEndpoints() bool`

HasRegionEndpoints returns a boolean if a field has been set.

### GetMultiInstanceSupportedVersion

`func (o *DestinationMetadataV1) GetMultiInstanceSupportedVersion() string`

GetMultiInstanceSupportedVersion returns the MultiInstanceSupportedVersion field if non-nil, zero value otherwise.

### GetMultiInstanceSupportedVersionOk

`func (o *DestinationMetadataV1) GetMultiInstanceSupportedVersionOk() (*string, bool)`

GetMultiInstanceSupportedVersionOk returns a tuple with the MultiInstanceSupportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiInstanceSupportedVersion

`func (o *DestinationMetadataV1) SetMultiInstanceSupportedVersion(v string)`

SetMultiInstanceSupportedVersion sets MultiInstanceSupportedVersion field to given value.

### HasMultiInstanceSupportedVersion

`func (o *DestinationMetadataV1) HasMultiInstanceSupportedVersion() bool`

HasMultiInstanceSupportedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


