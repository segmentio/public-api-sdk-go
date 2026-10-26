# EntityPathAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the entity path. | 
**EntityName** | **string** | The display name of the entity. | 
**PathName** | **string** | The display name of the path (root → node). | 
**EntitySlug** | **string** | The slug of the entity. | 
**RelationshipSlugs** | **[]string** | Ordered slugs for all relationships from root to this node. | 
**ChildRelationshipSlugs** | **[]string** | Slugs of child relationships from this node. | 
**ParentRelationshipSlugPath** | **string** | Full slug path from root to this node, colon-delimited (for example \&quot;rel-a:rel-b\&quot;). Empty string for root nodes. | 

## Methods

### NewEntityPathAlpha

`func NewEntityPathAlpha(id string, entityName string, pathName string, entitySlug string, relationshipSlugs []string, childRelationshipSlugs []string, parentRelationshipSlugPath string, ) *EntityPathAlpha`

NewEntityPathAlpha instantiates a new EntityPathAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPathAlphaWithDefaults

`func NewEntityPathAlphaWithDefaults() *EntityPathAlpha`

NewEntityPathAlphaWithDefaults instantiates a new EntityPathAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityPathAlpha) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityPathAlpha) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityPathAlpha) SetId(v string)`

SetId sets Id field to given value.


### GetEntityName

`func (o *EntityPathAlpha) GetEntityName() string`

GetEntityName returns the EntityName field if non-nil, zero value otherwise.

### GetEntityNameOk

`func (o *EntityPathAlpha) GetEntityNameOk() (*string, bool)`

GetEntityNameOk returns a tuple with the EntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityName

`func (o *EntityPathAlpha) SetEntityName(v string)`

SetEntityName sets EntityName field to given value.


### GetPathName

`func (o *EntityPathAlpha) GetPathName() string`

GetPathName returns the PathName field if non-nil, zero value otherwise.

### GetPathNameOk

`func (o *EntityPathAlpha) GetPathNameOk() (*string, bool)`

GetPathNameOk returns a tuple with the PathName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathName

`func (o *EntityPathAlpha) SetPathName(v string)`

SetPathName sets PathName field to given value.


### GetEntitySlug

`func (o *EntityPathAlpha) GetEntitySlug() string`

GetEntitySlug returns the EntitySlug field if non-nil, zero value otherwise.

### GetEntitySlugOk

`func (o *EntityPathAlpha) GetEntitySlugOk() (*string, bool)`

GetEntitySlugOk returns a tuple with the EntitySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySlug

`func (o *EntityPathAlpha) SetEntitySlug(v string)`

SetEntitySlug sets EntitySlug field to given value.


### GetRelationshipSlugs

`func (o *EntityPathAlpha) GetRelationshipSlugs() []string`

GetRelationshipSlugs returns the RelationshipSlugs field if non-nil, zero value otherwise.

### GetRelationshipSlugsOk

`func (o *EntityPathAlpha) GetRelationshipSlugsOk() (*[]string, bool)`

GetRelationshipSlugsOk returns a tuple with the RelationshipSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipSlugs

`func (o *EntityPathAlpha) SetRelationshipSlugs(v []string)`

SetRelationshipSlugs sets RelationshipSlugs field to given value.


### GetChildRelationshipSlugs

`func (o *EntityPathAlpha) GetChildRelationshipSlugs() []string`

GetChildRelationshipSlugs returns the ChildRelationshipSlugs field if non-nil, zero value otherwise.

### GetChildRelationshipSlugsOk

`func (o *EntityPathAlpha) GetChildRelationshipSlugsOk() (*[]string, bool)`

GetChildRelationshipSlugsOk returns a tuple with the ChildRelationshipSlugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildRelationshipSlugs

`func (o *EntityPathAlpha) SetChildRelationshipSlugs(v []string)`

SetChildRelationshipSlugs sets ChildRelationshipSlugs field to given value.


### GetParentRelationshipSlugPath

`func (o *EntityPathAlpha) GetParentRelationshipSlugPath() string`

GetParentRelationshipSlugPath returns the ParentRelationshipSlugPath field if non-nil, zero value otherwise.

### GetParentRelationshipSlugPathOk

`func (o *EntityPathAlpha) GetParentRelationshipSlugPathOk() (*string, bool)`

GetParentRelationshipSlugPathOk returns a tuple with the ParentRelationshipSlugPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRelationshipSlugPath

`func (o *EntityPathAlpha) SetParentRelationshipSlugPath(v string)`

SetParentRelationshipSlugPath sets ParentRelationshipSlugPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


