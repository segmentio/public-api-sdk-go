# EntityDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The entity primary key value. | 
**IdProperty** | **string** | The entity primary key column name. | 
**RelationshipSlug** | **string** | The entity relationship slug. | 
**Properties** | Pointer to **map[string]interface{}** | Entity properties. | [optional] 
**Entities** | Pointer to **map[string]interface{}** | Related entities that are one level deeper will only be returned if those entities are referenced in the audience definition. | [optional] 

## Methods

### NewEntityDetails

`func NewEntityDetails(id string, idProperty string, relationshipSlug string, ) *EntityDetails`

NewEntityDetails instantiates a new EntityDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityDetailsWithDefaults

`func NewEntityDetailsWithDefaults() *EntityDetails`

NewEntityDetailsWithDefaults instantiates a new EntityDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityDetails) SetId(v string)`

SetId sets Id field to given value.


### GetIdProperty

`func (o *EntityDetails) GetIdProperty() string`

GetIdProperty returns the IdProperty field if non-nil, zero value otherwise.

### GetIdPropertyOk

`func (o *EntityDetails) GetIdPropertyOk() (*string, bool)`

GetIdPropertyOk returns a tuple with the IdProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdProperty

`func (o *EntityDetails) SetIdProperty(v string)`

SetIdProperty sets IdProperty field to given value.


### GetRelationshipSlug

`func (o *EntityDetails) GetRelationshipSlug() string`

GetRelationshipSlug returns the RelationshipSlug field if non-nil, zero value otherwise.

### GetRelationshipSlugOk

`func (o *EntityDetails) GetRelationshipSlugOk() (*string, bool)`

GetRelationshipSlugOk returns a tuple with the RelationshipSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipSlug

`func (o *EntityDetails) SetRelationshipSlug(v string)`

SetRelationshipSlug sets RelationshipSlug field to given value.


### GetProperties

`func (o *EntityDetails) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EntityDetails) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EntityDetails) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EntityDetails) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEntities

`func (o *EntityDetails) GetEntities() map[string]interface{}`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *EntityDetails) GetEntitiesOk() (*map[string]interface{}, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *EntityDetails) SetEntities(v map[string]interface{})`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *EntityDetails) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


