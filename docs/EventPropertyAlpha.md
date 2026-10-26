# EventPropertyAlpha

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Property name. | 
**Type** | [**EventPropertyType**](EventPropertyType.md) |  | 
**EventName** | **string** | The event this property belongs to. | 
**LastSeenAt** | **string** | ISO timestamp of when the property was last seen. | 
**Samples** | Pointer to **[]string** | Sample values for this property. Only present when includeSampleValues is true. | [optional] 

## Methods

### NewEventPropertyAlpha

`func NewEventPropertyAlpha(name string, type_ EventPropertyType, eventName string, lastSeenAt string, ) *EventPropertyAlpha`

NewEventPropertyAlpha instantiates a new EventPropertyAlpha object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPropertyAlphaWithDefaults

`func NewEventPropertyAlphaWithDefaults() *EventPropertyAlpha`

NewEventPropertyAlphaWithDefaults instantiates a new EventPropertyAlpha object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventPropertyAlpha) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventPropertyAlpha) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventPropertyAlpha) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EventPropertyAlpha) GetType() EventPropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventPropertyAlpha) GetTypeOk() (*EventPropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventPropertyAlpha) SetType(v EventPropertyType)`

SetType sets Type field to given value.


### GetEventName

`func (o *EventPropertyAlpha) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *EventPropertyAlpha) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *EventPropertyAlpha) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetLastSeenAt

`func (o *EventPropertyAlpha) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *EventPropertyAlpha) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *EventPropertyAlpha) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetSamples

`func (o *EventPropertyAlpha) GetSamples() []string`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *EventPropertyAlpha) GetSamplesOk() (*[]string, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *EventPropertyAlpha) SetSamples(v []string)`

SetSamples sets Samples field to given value.

### HasSamples

`func (o *EventPropertyAlpha) HasSamples() bool`

HasSamples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


