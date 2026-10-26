# EventAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Event name. | 
**Type** | [**EventType**](EventType.md) |  | 
**LastSeenAt** | **string** | ISO timestamp of when the event was last seen. | 

## Methods

### NewEventAlpha

`func NewEventAlpha(name string, type_ EventType, lastSeenAt string, ) *EventAlpha`

NewEventAlpha instantiates a new EventAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAlphaWithDefaults

`func NewEventAlphaWithDefaults() *EventAlpha`

NewEventAlphaWithDefaults instantiates a new EventAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventAlpha) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventAlpha) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventAlpha) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EventAlpha) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventAlpha) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventAlpha) SetType(v EventType)`

SetType sets Type field to given value.


### GetLastSeenAt

`func (o *EventAlpha) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *EventAlpha) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *EventAlpha) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


