# FunctionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An identifier for this Function. | [optional] 
**ResourceType** | Pointer to **string** | The Function type.  Config API note: equal to &#x60;type&#x60;. | [optional] 
**CreatedAt** | Pointer to **string** | The time this Function was created. | [optional] 
**CreatedBy** | Pointer to **string** | The id of the user who created this Function. | [optional] 
**Code** | Pointer to **string** | The Function code. | [optional] 
**DeployedAt** | Pointer to **NullableString** | The time of this Function&#39;s last deployment. | [optional] 
**Settings** | Pointer to [**[]FunctionSettingV1**](FunctionSettingV1.md) | The list of settings for this Function. | [optional] 
**DisplayName** | Pointer to **string** | A display name for this Function. | [optional] 
**Description** | Pointer to **string** | A description for this Function. | [optional] 
**LogoUrl** | Pointer to **string** | The URL of the logo for this Function. | [optional] 
**PreviewWebhookUrl** | Pointer to **string** | The preview webhook URL for this Function. | [optional] 
**BatchMaxCount** | Pointer to **float32** | The max count of the batch for this Function. | [optional] 
**CatalogId** | Pointer to **string** | The catalog id of this Function. | [optional] 
**IsLatestVersion** | Pointer to **bool** | Whether the deployment of this Function is the latest version. | [optional] 

## Methods

### NewFunctionV1

`func NewFunctionV1() *FunctionV1`

NewFunctionV1 instantiates a new FunctionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionV1WithDefaults

`func NewFunctionV1WithDefaults() *FunctionV1`

NewFunctionV1WithDefaults instantiates a new FunctionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FunctionV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionV1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FunctionV1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceType

`func (o *FunctionV1) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *FunctionV1) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *FunctionV1) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *FunctionV1) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionV1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *FunctionV1) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FunctionV1) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FunctionV1) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FunctionV1) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCode

`func (o *FunctionV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FunctionV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FunctionV1) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FunctionV1) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDeployedAt

`func (o *FunctionV1) GetDeployedAt() string`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *FunctionV1) GetDeployedAtOk() (*string, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *FunctionV1) SetDeployedAt(v string)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *FunctionV1) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### SetDeployedAtNil

`func (o *FunctionV1) SetDeployedAtNil(b bool)`

 SetDeployedAtNil sets the value for DeployedAt to be an explicit nil

### UnsetDeployedAt
`func (o *FunctionV1) UnsetDeployedAt()`

UnsetDeployedAt ensures that no value is present for DeployedAt, not even an explicit nil
### GetSettings

`func (o *FunctionV1) GetSettings() []FunctionSettingV1`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *FunctionV1) GetSettingsOk() (*[]FunctionSettingV1, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *FunctionV1) SetSettings(v []FunctionSettingV1)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *FunctionV1) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetDisplayName

`func (o *FunctionV1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionV1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionV1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionV1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *FunctionV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogoUrl

`func (o *FunctionV1) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *FunctionV1) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *FunctionV1) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *FunctionV1) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetPreviewWebhookUrl

`func (o *FunctionV1) GetPreviewWebhookUrl() string`

GetPreviewWebhookUrl returns the PreviewWebhookUrl field if non-nil, zero value otherwise.

### GetPreviewWebhookUrlOk

`func (o *FunctionV1) GetPreviewWebhookUrlOk() (*string, bool)`

GetPreviewWebhookUrlOk returns a tuple with the PreviewWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewWebhookUrl

`func (o *FunctionV1) SetPreviewWebhookUrl(v string)`

SetPreviewWebhookUrl sets PreviewWebhookUrl field to given value.

### HasPreviewWebhookUrl

`func (o *FunctionV1) HasPreviewWebhookUrl() bool`

HasPreviewWebhookUrl returns a boolean if a field has been set.

### GetBatchMaxCount

`func (o *FunctionV1) GetBatchMaxCount() float32`

GetBatchMaxCount returns the BatchMaxCount field if non-nil, zero value otherwise.

### GetBatchMaxCountOk

`func (o *FunctionV1) GetBatchMaxCountOk() (*float32, bool)`

GetBatchMaxCountOk returns a tuple with the BatchMaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchMaxCount

`func (o *FunctionV1) SetBatchMaxCount(v float32)`

SetBatchMaxCount sets BatchMaxCount field to given value.

### HasBatchMaxCount

`func (o *FunctionV1) HasBatchMaxCount() bool`

HasBatchMaxCount returns a boolean if a field has been set.

### GetCatalogId

`func (o *FunctionV1) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *FunctionV1) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *FunctionV1) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *FunctionV1) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetIsLatestVersion

`func (o *FunctionV1) GetIsLatestVersion() bool`

GetIsLatestVersion returns the IsLatestVersion field if non-nil, zero value otherwise.

### GetIsLatestVersionOk

`func (o *FunctionV1) GetIsLatestVersionOk() (*bool, bool)`

GetIsLatestVersionOk returns a tuple with the IsLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestVersion

`func (o *FunctionV1) SetIsLatestVersion(v bool)`

SetIsLatestVersion sets IsLatestVersion field to given value.

### HasIsLatestVersion

`func (o *FunctionV1) HasIsLatestVersion() bool`

HasIsLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


