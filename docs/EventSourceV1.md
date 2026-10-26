# EventSourceV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Source where the events came from. | 
**Name** | Pointer to **string** | The name of the Source, if applicable. | [optional] 
**Slug** | Pointer to **string** | The slug of the Source, if applicable. | [optional] 

## Methods

### NewEventSourceV1

`func NewEventSourceV1(id string, ) *EventSourceV1`

NewEventSourceV1 instantiates a new EventSourceV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSourceV1WithDefaults

`func NewEventSourceV1WithDefaults() *EventSourceV1`

NewEventSourceV1WithDefaults instantiates a new EventSourceV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventSourceV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventSourceV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventSourceV1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EventSourceV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventSourceV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventSourceV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventSourceV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSlug

`func (o *EventSourceV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EventSourceV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EventSourceV1) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EventSourceV1) HasSlug() bool`

HasSlug returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


