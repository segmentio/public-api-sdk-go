# GetSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**SourceV1**](SourceV1.md) |  | 
**TrackingPlanId** | **NullableString** | The id of the Tracking Plan connected to the Source. | 

## Methods

### NewGetSourceV1Output

`func NewGetSourceV1Output(source SourceV1, trackingPlanId NullableString, ) *GetSourceV1Output`

NewGetSourceV1Output instantiates a new GetSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceV1OutputWithDefaults

`func NewGetSourceV1OutputWithDefaults() *GetSourceV1Output`

NewGetSourceV1OutputWithDefaults instantiates a new GetSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *GetSourceV1Output) GetSource() SourceV1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSourceV1Output) GetSourceOk() (*SourceV1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSourceV1Output) SetSource(v SourceV1)`

SetSource sets Source field to given value.


### GetTrackingPlanId

`func (o *GetSourceV1Output) GetTrackingPlanId() string`

GetTrackingPlanId returns the TrackingPlanId field if non-nil, zero value otherwise.

### GetTrackingPlanIdOk

`func (o *GetSourceV1Output) GetTrackingPlanIdOk() (*string, bool)`

GetTrackingPlanIdOk returns a tuple with the TrackingPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingPlanId

`func (o *GetSourceV1Output) SetTrackingPlanId(v string)`

SetTrackingPlanId sets TrackingPlanId field to given value.


### SetTrackingPlanIdNil

`func (o *GetSourceV1Output) SetTrackingPlanIdNil(b bool)`

 SetTrackingPlanIdNil sets the value for TrackingPlanId to be an explicit nil

### UnsetTrackingPlanId
`func (o *GetSourceV1Output) UnsetTrackingPlanId()`

UnsetTrackingPlanId ensures that no value is present for TrackingPlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


