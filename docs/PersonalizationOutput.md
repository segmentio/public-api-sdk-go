# PersonalizationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | [**Profile1**](Profile1.md) |  | 
**Entities** | Pointer to [**[]PersonalizationInputEntity**](PersonalizationInputEntity.md) | The entities, and the properties of each entity, included in the event sent to the Destination. Only applicable to Linked Audiences. | [optional] 
**SyncEntityPropertyChanges** | Pointer to **bool** | Sync entity property changes to the Destination. Only applicable if activationType is \&quot;Audience Membership Changed\&quot;. | [optional] 

## Methods

### NewPersonalizationOutput

`func NewPersonalizationOutput(profile Profile1, ) *PersonalizationOutput`

NewPersonalizationOutput instantiates a new PersonalizationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalizationOutputWithDefaults

`func NewPersonalizationOutputWithDefaults() *PersonalizationOutput`

NewPersonalizationOutputWithDefaults instantiates a new PersonalizationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *PersonalizationOutput) GetProfile() Profile1`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PersonalizationOutput) GetProfileOk() (*Profile1, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PersonalizationOutput) SetProfile(v Profile1)`

SetProfile sets Profile field to given value.


### GetEntities

`func (o *PersonalizationOutput) GetEntities() []PersonalizationInputEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PersonalizationOutput) GetEntitiesOk() (*[]PersonalizationInputEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PersonalizationOutput) SetEntities(v []PersonalizationInputEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PersonalizationOutput) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetSyncEntityPropertyChanges

`func (o *PersonalizationOutput) GetSyncEntityPropertyChanges() bool`

GetSyncEntityPropertyChanges returns the SyncEntityPropertyChanges field if non-nil, zero value otherwise.

### GetSyncEntityPropertyChangesOk

`func (o *PersonalizationOutput) GetSyncEntityPropertyChangesOk() (*bool, bool)`

GetSyncEntityPropertyChangesOk returns a tuple with the SyncEntityPropertyChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEntityPropertyChanges

`func (o *PersonalizationOutput) SetSyncEntityPropertyChanges(v bool)`

SetSyncEntityPropertyChanges sets SyncEntityPropertyChanges field to given value.

### HasSyncEntityPropertyChanges

`func (o *PersonalizationOutput) HasSyncEntityPropertyChanges() bool`

HasSyncEntityPropertyChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


