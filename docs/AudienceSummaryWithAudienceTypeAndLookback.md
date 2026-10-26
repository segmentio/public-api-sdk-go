# AudienceSummaryWithAudienceTypeAndLookback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputeCadence** | [**AudienceComputeCadence**](AudienceComputeCadence.md) |  | 
**Size** | Pointer to [**AudienceSize**](AudienceSize.md) |  | [optional] 
**Options** | Pointer to [**AudienceOptionsWithLookback**](AudienceOptionsWithLookback.md) |  | [optional] 
**Schedules** | Pointer to [**[]AudienceSchedule**](AudienceSchedule.md) | List of schedules for the audience. | [optional] 
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

## Methods

### NewAudienceSummaryWithAudienceTypeAndLookback

`func NewAudienceSummaryWithAudienceTypeAndLookback(computeCadence AudienceComputeCadence, id string, spaceId string, name string, key string, enabled bool, definition NullableAudienceDefinition, createdBy string, updatedBy string, createdAt string, updatedAt string, audienceType string, ) *AudienceSummaryWithAudienceTypeAndLookback`

NewAudienceSummaryWithAudienceTypeAndLookback instantiates a new AudienceSummaryWithAudienceTypeAndLookback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceSummaryWithAudienceTypeAndLookbackWithDefaults

`func NewAudienceSummaryWithAudienceTypeAndLookbackWithDefaults() *AudienceSummaryWithAudienceTypeAndLookback`

NewAudienceSummaryWithAudienceTypeAndLookbackWithDefaults instantiates a new AudienceSummaryWithAudienceTypeAndLookback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputeCadence

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetComputeCadence() AudienceComputeCadence`

GetComputeCadence returns the ComputeCadence field if non-nil, zero value otherwise.

### GetComputeCadenceOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetComputeCadenceOk() (*AudienceComputeCadence, bool)`

GetComputeCadenceOk returns a tuple with the ComputeCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeCadence

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetComputeCadence(v AudienceComputeCadence)`

SetComputeCadence sets ComputeCadence field to given value.


### GetSize

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSize() AudienceSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSizeOk() (*AudienceSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSize(v AudienceSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *AudienceSummaryWithAudienceTypeAndLookback) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetOptions

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetOptions() AudienceOptionsWithLookback`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetOptionsOk() (*AudienceOptionsWithLookback, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetOptions(v AudienceOptionsWithLookback)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AudienceSummaryWithAudienceTypeAndLookback) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSchedules

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSchedules() []AudienceSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSchedulesOk() (*[]AudienceSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSchedules(v []AudienceSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *AudienceSummaryWithAudienceTypeAndLookback) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetId

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetId(v string)`

SetId sets Id field to given value.


### GetSpaceId

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetName

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AudienceSummaryWithAudienceTypeAndLookback) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetKey(v string)`

SetKey sets Key field to given value.


### GetEnabled

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefinition

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDefinition() AudienceDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetDefinitionOk() (*AudienceDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetDefinition(v AudienceDefinition)`

SetDefinition sets Definition field to given value.


### SetDefinitionNil

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *AudienceSummaryWithAudienceTypeAndLookback) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetStatus

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AudienceSummaryWithAudienceTypeAndLookback) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetAudienceType

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetAudienceType() string`

GetAudienceType returns the AudienceType field if non-nil, zero value otherwise.

### GetAudienceTypeOk

`func (o *AudienceSummaryWithAudienceTypeAndLookback) GetAudienceTypeOk() (*string, bool)`

GetAudienceTypeOk returns a tuple with the AudienceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceType

`func (o *AudienceSummaryWithAudienceTypeAndLookback) SetAudienceType(v string)`

SetAudienceType sets AudienceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


