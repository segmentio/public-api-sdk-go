# SourceEventVolumeDatapointV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **string** | The timestamp that corresponds to the beginning of the window given by the requested granularity. | 
**Count** | **float32** | The number of valid events Segment received in the given window, after the events completed the validation process. | 

## Methods

### NewSourceEventVolumeDatapointV1

`func NewSourceEventVolumeDatapointV1(time string, count float32, ) *SourceEventVolumeDatapointV1`

NewSourceEventVolumeDatapointV1 instantiates a new SourceEventVolumeDatapointV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceEventVolumeDatapointV1WithDefaults

`func NewSourceEventVolumeDatapointV1WithDefaults() *SourceEventVolumeDatapointV1`

NewSourceEventVolumeDatapointV1WithDefaults instantiates a new SourceEventVolumeDatapointV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *SourceEventVolumeDatapointV1) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SourceEventVolumeDatapointV1) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SourceEventVolumeDatapointV1) SetTime(v string)`

SetTime sets Time field to given value.


### GetCount

`func (o *SourceEventVolumeDatapointV1) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SourceEventVolumeDatapointV1) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SourceEventVolumeDatapointV1) SetCount(v float32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


