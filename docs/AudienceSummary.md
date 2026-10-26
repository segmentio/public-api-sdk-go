# AudienceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Audience id. | 
**SpaceId** | **string** | Space id for the audience. | 
**Name** | **string** | Name of the audience. | 
**Description** | Pointer to **string** | Description of the audience. | [optional] 
**Key** | **string** | Key for the audience. | 
**Enabled** | **bool** | Enabled/disabled status for the audience. | 
**Definition** | [**NullableAudienceDefinition**](AudienceDefinition.md) |  | 
**Status** | Pointer to **string** | Status for the audience.  Possible values: Backfilling, Computing, Disabled, Error, Failed, Live, Locked, Not Computing, Preparing, Awaiting Destinations. | [optional] 
**CreatedBy** | **string** | User id who created the audience. | 
**UpdatedBy** | **string** | User id who last updated the audience. | 
**CreatedAt** | **string** | Date the audience was created. | 
**UpdatedAt** | **string** | Date the audience was last updated. | 
**AudienceType** | **string** | Denotes the type of audience product. | 
**Options** | Pointer to [**AudienceOptions**](AudienceOptions.md) |  | [optional] 

## Methods

### NewAudienceSummary

`func NewAudienceSummary(id string, spaceId string, name string, key string, enabled bool, definition NullableAudienceDefinition, createdBy string, updatedBy string, createdAt string, updatedAt string, audienceType string, ) *AudienceSummary`

NewAudienceSummary instantiates a new AudienceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSummaryWithDefaults

`func NewAudienceSummaryWithDefaults() *AudienceSummary`

NewAudienceSummaryWithDefaults instantiates a new AudienceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudienceSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceSummary) SetId(v string)`

SetId sets Id field to given value.


### GetSpaceId

`func (o *AudienceSummary) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *AudienceSummary) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *AudienceSummary) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetName

`func (o *AudienceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AudienceSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AudienceSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AudienceSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AudienceSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *AudienceSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AudienceSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AudienceSummary) SetKey(v string)`

SetKey sets Key field to given value.


### GetEnabled

`func (o *AudienceSummary) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AudienceSummary) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AudienceSummary) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefinition

`func (o *AudienceSummary) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *AudienceSummary) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *AudienceSummary) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### SetDefinitionNil

`func (o *AudienceSummary) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *AudienceSummary) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetStatus

`func (o *AudienceSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudienceSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudienceSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AudienceSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AudienceSummary) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AudienceSummary) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AudienceSummary) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *AudienceSummary) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AudienceSummary) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AudienceSummary) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *AudienceSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AudienceSummary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AudienceSummary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AudienceSummary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAudienceType

`func (o *AudienceSummary) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *AudienceSummary) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *AudienceSummary) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.


### GetOptions

`func (o *AudienceSummary) GetOptions() AudienceOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AudienceSummary) GetOptionsOk() (*AudienceOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AudienceSummary) SetOptions(v AudienceOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AudienceSummary) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


