# PersonalizationInputEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **[]string** | The entity properties to include in the event sent to the Destination. | 
**RelationshipSlug** | **string** | The relationship slug for the entity. | 

## Methods

### NewPersonalizationInputEntity

`func NewPersonalizationInputEntity(properties []string, relationshipSlug string, ) *PersonalizationInputEntity`

NewPersonalizationInputEntity instantiates a new PersonalizationInputEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonalizationInputEntityWithDefaults

`func NewPersonalizationInputEntityWithDefaults() *PersonalizationInputEntity`

NewPersonalizationInputEntityWithDefaults instantiates a new PersonalizationInputEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PersonalizationInputEntity) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PersonalizationInputEntity) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PersonalizationInputEntity) SetProperties(v []string)`

SetProperties sets Properties field to given value.


### GetRelationshipSlug

`func (o *PersonalizationInputEntity) GetRelationshipSlug() string`

GetRelationshipSlug returns the RelationshipSlug field if non-nil, zero value otherwise.

### GetRelationshipSlugOk

`func (o *PersonalizationInputEntity) GetRelationshipSlugOk() (*string, bool)`

GetRelationshipSlugOk returns a tuple with the RelationshipSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipSlug

`func (o *PersonalizationInputEntity) SetRelationshipSlug(v string)`

SetRelationshipSlug sets RelationshipSlug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


