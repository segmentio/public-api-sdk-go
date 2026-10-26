# AudienceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEntity** | Pointer to **NullableString** | The target entity relationship slug, only applicable for linked audiences. Use &#x60;&#x60;&#x60;profile&#x60;&#x60;&#x60; when the targeting the profile.  Note that the value defaults to profile if not specified. Also note, that the value will be returned as null if the target entity is removed from the data graph. | [optional] 
**Query** | **string** | The query language string defining the audience segmentation criteria.  For guidance on using the query language, see the [Segment documentation site](https://segment.com/docs/api/public-api/query-language). | 

## Methods

### NewAudienceDefinition

`func NewAudienceDefinition(query string, ) *AudienceDefinition`

NewAudienceDefinition instantiates a new AudienceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceDefinitionWithDefaults

`func NewAudienceDefinitionWithDefaults() *AudienceDefinition`

NewAudienceDefinitionWithDefaults instantiates a new AudienceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEntity

`func (o *AudienceDefinition) GetTargetEntity() string`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *AudienceDefinition) GetTargetEntityOk() (*string, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *AudienceDefinition) SetTargetEntity(v string)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *AudienceDefinition) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### SetTargetEntityNil

`func (o *AudienceDefinition) SetTargetEntityNil(b bool)`

 SetTargetEntityNil sets the value for TargetEntity to be an explicit nil

### UnsetTargetEntity
`func (o *AudienceDefinition) UnsetTargetEntity()`

UnsetTargetEntity ensures that no value is present for TargetEntity, not even an explicit nil
### GetQuery

`func (o *AudienceDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AudienceDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AudienceDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


