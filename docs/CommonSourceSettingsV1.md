# CommonSourceSettingsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Track** | Pointer to [**TrackSourceSettingsV1**](TrackSourceSettingsV1.md) |  | [optional] 
**Identify** | Pointer to [**IdentifySourceSettingsV1**](IdentifySourceSettingsV1.md) |  | [optional] 
**Group** | Pointer to [**GroupSourceSettingsV1**](GroupSourceSettingsV1.md) |  | [optional] 
**ForwardingViolationsTo** | Pointer to **string** | SourceId to forward violations to. | [optional] 
**ForwardingBlockedEventsTo** | Pointer to **string** | SourceId to forward blocked events to. | [optional] 

## Methods

### NewCommonSourceSettingsV1

`func NewCommonSourceSettingsV1() *CommonSourceSettingsV1`

NewCommonSourceSettingsV1 instantiates a new CommonSourceSettingsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonSourceSettingsV1WithDefaults

`func NewCommonSourceSettingsV1WithDefaults() *CommonSourceSettingsV1`

NewCommonSourceSettingsV1WithDefaults instantiates a new CommonSourceSettingsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrack

`func (o *CommonSourceSettingsV1) GetTrack() TrackSourceSettingsV1`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *CommonSourceSettingsV1) GetTrackOk() (*TrackSourceSettingsV1, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *CommonSourceSettingsV1) SetTrack(v TrackSourceSettingsV1)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *CommonSourceSettingsV1) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetIdentify

`func (o *CommonSourceSettingsV1) GetIdentify() IdentifySourceSettingsV1`

GetIdentify returns the Identify field if non-nil, zero value otherwise.

### GetIdentifyOk

`func (o *CommonSourceSettingsV1) GetIdentifyOk() (*IdentifySourceSettingsV1, bool)`

GetIdentifyOk returns a tuple with the Identify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentify

`func (o *CommonSourceSettingsV1) SetIdentify(v IdentifySourceSettingsV1)`

SetIdentify sets Identify field to given value.

### HasIdentify

`func (o *CommonSourceSettingsV1) HasIdentify() bool`

HasIdentify returns a boolean if a field has been set.

### GetGroup

`func (o *CommonSourceSettingsV1) GetGroup() GroupSourceSettingsV1`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CommonSourceSettingsV1) GetGroupOk() (*GroupSourceSettingsV1, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CommonSourceSettingsV1) SetGroup(v GroupSourceSettingsV1)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CommonSourceSettingsV1) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetForwardingViolationsTo

`func (o *CommonSourceSettingsV1) GetForwardingViolationsTo() string`

GetForwardingViolationsTo returns the ForwardingViolationsTo field if non-nil, zero value otherwise.

### GetForwardingViolationsToOk

`func (o *CommonSourceSettingsV1) GetForwardingViolationsToOk() (*string, bool)`

GetForwardingViolationsToOk returns a tuple with the ForwardingViolationsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingViolationsTo

`func (o *CommonSourceSettingsV1) SetForwardingViolationsTo(v string)`

SetForwardingViolationsTo sets ForwardingViolationsTo field to given value.

### HasForwardingViolationsTo

`func (o *CommonSourceSettingsV1) HasForwardingViolationsTo() bool`

HasForwardingViolationsTo returns a boolean if a field has been set.

### GetForwardingBlockedEventsTo

`func (o *CommonSourceSettingsV1) GetForwardingBlockedEventsTo() string`

GetForwardingBlockedEventsTo returns the ForwardingBlockedEventsTo field if non-nil, zero value otherwise.

### GetForwardingBlockedEventsToOk

`func (o *CommonSourceSettingsV1) GetForwardingBlockedEventsToOk() (*string, bool)`

GetForwardingBlockedEventsToOk returns a tuple with the ForwardingBlockedEventsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingBlockedEventsTo

`func (o *CommonSourceSettingsV1) SetForwardingBlockedEventsTo(v string)`

SetForwardingBlockedEventsTo sets ForwardingBlockedEventsTo field to given value.

### HasForwardingBlockedEventsTo

`func (o *CommonSourceSettingsV1) HasForwardingBlockedEventsTo() bool`

HasForwardingBlockedEventsTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


