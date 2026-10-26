# AudiencePreviewEntitiesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Entity id. | 
**IdProperty** | **string** | The entity primary key column name. | 
**RelationshipSlug** | **string** | The entity relationship slug. | 
**Properties** | Pointer to **map[string]interface{}** | Entity properties. | [optional] 
**Entities** | Pointer to **map[string]interface{}** | Related entities that are one level deeper will only be returned if those entities are referenced in the audience definition. | [optional] 
**Profiles** | Pointer to [**[]Profile**](Profile.md) | Related list of profiles. | [optional] 
**ProfilesTruncated** | **bool** | Indicates if only a subset of the profiles associated with the entity were returned. | 

## Methods

### NewAudiencePreviewEntitiesResult

`func NewAudiencePreviewEntitiesResult(id string, idProperty string, relationshipSlug string, profilesTruncated bool, ) *AudiencePreviewEntitiesResult`

NewAudiencePreviewEntitiesResult instantiates a new AudiencePreviewEntitiesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencePreviewEntitiesResultWithDefaults

`func NewAudiencePreviewEntitiesResultWithDefaults() *AudiencePreviewEntitiesResult`

NewAudiencePreviewEntitiesResultWithDefaults instantiates a new AudiencePreviewEntitiesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudiencePreviewEntitiesResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudiencePreviewEntitiesResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudiencePreviewEntitiesResult) SetId(v string)`

SetId sets Id field to given value.


### GetIdProperty

`func (o *AudiencePreviewEntitiesResult) GetIdProperty() string`

GetIdProperty returns the IdProperty field if non-nil, zero value otherwise.

### GetIdPropertyOk

`func (o *AudiencePreviewEntitiesResult) GetIdPropertyOk() (*string, bool)`

GetIdPropertyOk returns a tuple with the IdProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProperty

`func (o *AudiencePreviewEntitiesResult) SetIdProperty(v string)`

SetIdProperty sets IdProperty field to given value.


### GetRelationshipSlug

`func (o *AudiencePreviewEntitiesResult) GetRelationshipSlug() string`

GetRelationshipSlug returns the RelationshipSlug field if non-nil, zero value otherwise.

### GetRelationshipSlugOk

`func (o *AudiencePreviewEntitiesResult) GetRelationshipSlugOk() (*string, bool)`

GetRelationshipSlugOk returns a tuple with the RelationshipSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipSlug

`func (o *AudiencePreviewEntitiesResult) SetRelationshipSlug(v string)`

SetRelationshipSlug sets RelationshipSlug field to given value.


### GetProperties

`func (o *AudiencePreviewEntitiesResult) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AudiencePreviewEntitiesResult) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AudiencePreviewEntitiesResult) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AudiencePreviewEntitiesResult) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEntities

`func (o *AudiencePreviewEntitiesResult) GetEntities() map[string]interface{}`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AudiencePreviewEntitiesResult) GetEntitiesOk() (*map[string]interface{}, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AudiencePreviewEntitiesResult) SetEntities(v map[string]interface{})`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AudiencePreviewEntitiesResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetProfiles

`func (o *AudiencePreviewEntitiesResult) GetProfiles() []Profile`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *AudiencePreviewEntitiesResult) GetProfilesOk() (*[]Profile, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *AudiencePreviewEntitiesResult) SetProfiles(v []Profile)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *AudiencePreviewEntitiesResult) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetProfilesTruncated

`func (o *AudiencePreviewEntitiesResult) GetProfilesTruncated() bool`

GetProfilesTruncated returns the ProfilesTruncated field if non-nil, zero value otherwise.

### GetProfilesTruncatedOk

`func (o *AudiencePreviewEntitiesResult) GetProfilesTruncatedOk() (*bool, bool)`

GetProfilesTruncatedOk returns a tuple with the ProfilesTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesTruncated

`func (o *AudiencePreviewEntitiesResult) SetProfilesTruncated(v bool)`

SetProfilesTruncated sets ProfilesTruncated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


