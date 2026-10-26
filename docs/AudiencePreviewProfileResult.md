# AudiencePreviewProfileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Segment id. | 
**Entities** | Pointer to [**map[string][]EntityDetails**](array.md) | Associated entities. | [optional] 

## Methods

### NewAudiencePreviewProfileResult

`func NewAudiencePreviewProfileResult(id string, ) *AudiencePreviewProfileResult`

NewAudiencePreviewProfileResult instantiates a new AudiencePreviewProfileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudiencePreviewProfileResultWithDefaults

`func NewAudiencePreviewProfileResultWithDefaults() *AudiencePreviewProfileResult`

NewAudiencePreviewProfileResultWithDefaults instantiates a new AudiencePreviewProfileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AudiencePreviewProfileResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AudiencePreviewProfileResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AudiencePreviewProfileResult) SetId(v string)`

SetId sets Id field to given value.


### GetEntities

`func (o *AudiencePreviewProfileResult) GetEntities() map[string][]EntityDetails`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *AudiencePreviewProfileResult) GetEntitiesOk() (*map[string][]EntityDetails, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *AudiencePreviewProfileResult) SetEntities(v map[string][]EntityDetails)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *AudiencePreviewProfileResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


