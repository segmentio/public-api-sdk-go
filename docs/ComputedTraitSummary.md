# ComputedTraitSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Computed trait id. | 
**SpaceId** | **string** | Space id for the computed trait. | 
**Name** | **string** | Name of the computed trait. | 
**Description** | Pointer to **string** | Description of the computed trait. | [optional] 
**Key** | **string** | Key for the computed trait. | 
**Enabled** | **bool** | Enabled/disabled status for the computed trait. | 
**Definition** | [**NullableComputedTraitsDefinition**](ComputedTraitsDefinition.md) |  | 
**Status** | Pointer to **string** | Status for the computed trait.  Possible values: Backfilling, Computing, Failed, Live, Awaiting Destinations, Disabled. | [optional] 
**CreatedBy** | **string** | User id who created the computed trait. | 
**UpdatedBy** | **string** | User id who last updated the computed trait. | 
**CreatedAt** | **string** | The timestamp of the computed trait&#39;s creation. | 
**UpdatedAt** | **string** | The timestamp of the computed trait&#39;s last change. | 
**Options** | Pointer to [**TraitOptions**](TraitOptions.md) |  | [optional] 

## Methods

### NewComputedTraitSummary

`func NewComputedTraitSummary(id string, spaceId string, name string, key string, enabled bool, definition NullableComputedTraitsDefinition, createdBy string, updatedBy string, createdAt string, updatedAt string, ) *ComputedTraitSummary`

NewComputedTraitSummary instantiates a new ComputedTraitSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputedTraitSummaryWithDefaults

`func NewComputedTraitSummaryWithDefaults() *ComputedTraitSummary`

NewComputedTraitSummaryWithDefaults instantiates a new ComputedTraitSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ComputedTraitSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComputedTraitSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComputedTraitSummary) SetId(v string)`

SetId sets Id field to given value.


### GetSpaceId

`func (o *ComputedTraitSummary) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ComputedTraitSummary) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ComputedTraitSummary) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetName

`func (o *ComputedTraitSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComputedTraitSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComputedTraitSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComputedTraitSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComputedTraitSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComputedTraitSummary) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ComputedTraitSummary) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *ComputedTraitSummary) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComputedTraitSummary) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComputedTraitSummary) SetKey(v string)`

SetKey sets Key field to given value.


### GetEnabled

`func (o *ComputedTraitSummary) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ComputedTraitSummary) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ComputedTraitSummary) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefinition

`func (o *ComputedTraitSummary) GetDefinition() ComputedTraitsDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *ComputedTraitSummary) GetDefinitionOk() (*ComputedTraitsDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *ComputedTraitSummary) SetDefinition(v ComputedTraitsDefinition)`

SetDefinition sets Definition field to given value.


### SetDefinitionNil

`func (o *ComputedTraitSummary) SetDefinitionNil(b bool)`

 SetDefinitionNil sets the value for Definition to be an explicit nil

### UnsetDefinition
`func (o *ComputedTraitSummary) UnsetDefinition()`

UnsetDefinition ensures that no value is present for Definition, not even an explicit nil
### GetStatus

`func (o *ComputedTraitSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComputedTraitSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComputedTraitSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComputedTraitSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ComputedTraitSummary) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ComputedTraitSummary) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ComputedTraitSummary) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *ComputedTraitSummary) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ComputedTraitSummary) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ComputedTraitSummary) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *ComputedTraitSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComputedTraitSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComputedTraitSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ComputedTraitSummary) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ComputedTraitSummary) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ComputedTraitSummary) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOptions

`func (o *ComputedTraitSummary) GetOptions() TraitOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ComputedTraitSummary) GetOptionsOk() (*TraitOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ComputedTraitSummary) SetOptions(v TraitOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ComputedTraitSummary) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


