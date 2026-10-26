# SourceEventVolumeV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to [**EventSourceV1**](EventSourceV1.md) |  | [optional] 
**EventName** | Pointer to **string** | The name of the event, if applicable. | [optional] 
**EventType** | Pointer to **string** | The event type, if applicable. | [optional] 
**Total** | **float32** | The total count for all events in the requested time frame. | 
**Series** | [**[]SourceEventVolumeDatapointV1**](SourceEventVolumeDatapointV1.md) | A list of the event counts broken down by the requested granularity. | 

## Methods

### NewSourceEventVolumeV1

`func NewSourceEventVolumeV1(total float32, series []SourceEventVolumeDatapointV1, ) *SourceEventVolumeV1`

NewSourceEventVolumeV1 instantiates a new SourceEventVolumeV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceEventVolumeV1WithDefaults

`func NewSourceEventVolumeV1WithDefaults() *SourceEventVolumeV1`

NewSourceEventVolumeV1WithDefaults instantiates a new SourceEventVolumeV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *SourceEventVolumeV1) GetSource() EventSourceV1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SourceEventVolumeV1) GetSourceOk() (*EventSourceV1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SourceEventVolumeV1) SetSource(v EventSourceV1)`

SetSource sets Source field to given value.

### HasSource

`func (o *SourceEventVolumeV1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEventName

`func (o *SourceEventVolumeV1) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *SourceEventVolumeV1) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *SourceEventVolumeV1) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *SourceEventVolumeV1) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *SourceEventVolumeV1) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SourceEventVolumeV1) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SourceEventVolumeV1) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SourceEventVolumeV1) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTotal

`func (o *SourceEventVolumeV1) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SourceEventVolumeV1) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SourceEventVolumeV1) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetSeries

`func (o *SourceEventVolumeV1) GetSeries() []SourceEventVolumeDatapointV1`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SourceEventVolumeV1) GetSeriesOk() (*[]SourceEventVolumeDatapointV1, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SourceEventVolumeV1) SetSeries(v []SourceEventVolumeDatapointV1)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


