# TrackingPlanV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Tracking Plan&#39;s identifier.  Config API note: analogous to &#x60;name&#x60;. | 
**Name** | Pointer to **string** | The Tracking Plan&#39;s name.  Config API note: equal to &#x60;displayName&#x60;. | [optional] 
**Slug** | Pointer to **string** | URL-friendly slug of this Tracking Plan.  Config API note: equal to &#x60;name&#x60;. | [optional] 
**Description** | Pointer to **string** | The Tracking Plan&#39;s description. | [optional] 
**Type** | **string** | The Tracking Plan&#39;s type. | 
**UpdatedAt** | Pointer to **string** | The timestamp of the last change to the Tracking Plan.  Config API note: equal to &#x60;updateTime&#x60;. | [optional] 
**CreatedAt** | Pointer to **string** | The timestamp of this Tracking Plan&#39;s creation.  Config API note: equal to &#x60;createTime&#x60;. | [optional] 

## Methods

### NewTrackingPlanV1

`func NewTrackingPlanV1(id string, type_ string, ) *TrackingPlanV1`

NewTrackingPlanV1 instantiates a new TrackingPlanV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingPlanV1WithDefaults

`func NewTrackingPlanV1WithDefaults() *TrackingPlanV1`

NewTrackingPlanV1WithDefaults instantiates a new TrackingPlanV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrackingPlanV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrackingPlanV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrackingPlanV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TrackingPlanV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrackingPlanV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrackingPlanV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrackingPlanV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *TrackingPlanV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TrackingPlanV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TrackingPlanV1) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TrackingPlanV1) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetDescription

`func (o *TrackingPlanV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackingPlanV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackingPlanV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrackingPlanV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *TrackingPlanV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TrackingPlanV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TrackingPlanV1) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *TrackingPlanV1) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TrackingPlanV1) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TrackingPlanV1) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TrackingPlanV1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TrackingPlanV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TrackingPlanV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TrackingPlanV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TrackingPlanV1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


