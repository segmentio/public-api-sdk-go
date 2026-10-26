# PersonalizationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 
**Entities** | Pointer to [**[]PersonalizationInputEntity**](PersonalizationInputEntity.md) | The entities, and the properties of each entity, to include in the event sent to the Destination. Only applicable to Linked Audiences. Providing entities for a Classic audience returns a 400 error, as Classic audiences support profile properties only. | [optional] 
**SyncEntityPropertyChanges** | Pointer to **bool** | Sync entity property changes to the Destination. Only applicable if activationType is \&quot;Audience Membership Changed\&quot;. | [optional] 

## Methods

### NewPersonalizationInput

`func NewPersonalizationInput() *PersonalizationInput`

NewPersonalizationInput instantiates a new PersonalizationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalizationInputWithDefaults

`func NewPersonalizationInputWithDefaults() *PersonalizationInput`

NewPersonalizationInputWithDefaults instantiates a new PersonalizationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *PersonalizationInput) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PersonalizationInput) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PersonalizationInput) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PersonalizationInput) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetEntities

`func (o *PersonalizationInput) GetEntities() []PersonalizationInputEntity`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *PersonalizationInput) GetEntitiesOk() (*[]PersonalizationInputEntity, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *PersonalizationInput) SetEntities(v []PersonalizationInputEntity)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *PersonalizationInput) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetSyncEntityPropertyChanges

`func (o *PersonalizationInput) GetSyncEntityPropertyChanges() bool`

GetSyncEntityPropertyChanges returns the SyncEntityPropertyChanges field if non-nil, zero value otherwise.

### GetSyncEntityPropertyChangesOk

`func (o *PersonalizationInput) GetSyncEntityPropertyChangesOk() (*bool, bool)`

GetSyncEntityPropertyChangesOk returns a tuple with the SyncEntityPropertyChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncEntityPropertyChanges

`func (o *PersonalizationInput) SetSyncEntityPropertyChanges(v bool)`

SetSyncEntityPropertyChanges sets SyncEntityPropertyChanges field to given value.

### HasSyncEntityPropertyChanges

`func (o *PersonalizationInput) HasSyncEntityPropertyChanges() bool`

HasSyncEntityPropertyChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


