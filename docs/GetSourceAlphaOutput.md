# GetSourceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**SourceAlpha**](SourceAlpha.md) |  | 
**TrackingPlanId** | **NullableString** | The id of the Tracking Plan connected to the Source. | 

## Methods

### NewGetSourceAlphaOutput

`func NewGetSourceAlphaOutput(source SourceAlpha, trackingPlanId NullableString, ) *GetSourceAlphaOutput`

NewGetSourceAlphaOutput instantiates a new GetSourceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSourceAlphaOutputWithDefaults

`func NewGetSourceAlphaOutputWithDefaults() *GetSourceAlphaOutput`

NewGetSourceAlphaOutputWithDefaults instantiates a new GetSourceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *GetSourceAlphaOutput) GetSource() SourceAlpha`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetSourceAlphaOutput) GetSourceOk() (*SourceAlpha, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetSourceAlphaOutput) SetSource(v SourceAlpha)`

SetSource sets Source field to given value.


### GetTrackingPlanId

`func (o *GetSourceAlphaOutput) GetTrackingPlanId() string`

GetTrackingPlanId returns the TrackingPlanId field if non-nil, zero value otherwise.

### GetTrackingPlanIdOk

`func (o *GetSourceAlphaOutput) GetTrackingPlanIdOk() (*string, bool)`

GetTrackingPlanIdOk returns a tuple with the TrackingPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingPlanId

`func (o *GetSourceAlphaOutput) SetTrackingPlanId(v string)`

SetTrackingPlanId sets TrackingPlanId field to given value.


### SetTrackingPlanIdNil

`func (o *GetSourceAlphaOutput) SetTrackingPlanIdNil(b bool)`

 SetTrackingPlanIdNil sets the value for TrackingPlanId to be an explicit nil

### UnsetTrackingPlanId
`func (o *GetSourceAlphaOutput) UnsetTrackingPlanId()`

UnsetTrackingPlanId ensures that no value is present for TrackingPlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


