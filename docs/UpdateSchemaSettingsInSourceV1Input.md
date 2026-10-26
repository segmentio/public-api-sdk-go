# UpdateSchemaSettingsInSourceV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Track** | Pointer to [**TrackSourceSettingsV1**](TrackSourceSettingsV1.md) |  | [optional] 
**Identify** | Pointer to [**IdentifySourceSettingsV1**](IdentifySourceSettingsV1.md) |  | [optional] 
**Group** | Pointer to [**GroupSourceSettingsV1**](GroupSourceSettingsV1.md) |  | [optional] 
**ForwardingViolationsTo** | Pointer to **string** | Source id to forward violations to. | [optional] 
**ForwardingBlockedEventsTo** | Pointer to **string** | Source id to forward blocked events to. | [optional] 

## Methods

### NewUpdateSchemaSettingsInSourceV1Input

`func NewUpdateSchemaSettingsInSourceV1Input() *UpdateSchemaSettingsInSourceV1Input`

NewUpdateSchemaSettingsInSourceV1Input instantiates a new UpdateSchemaSettingsInSourceV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSchemaSettingsInSourceV1InputWithDefaults

`func NewUpdateSchemaSettingsInSourceV1InputWithDefaults() *UpdateSchemaSettingsInSourceV1Input`

NewUpdateSchemaSettingsInSourceV1InputWithDefaults instantiates a new UpdateSchemaSettingsInSourceV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrack

`func (o *UpdateSchemaSettingsInSourceV1Input) GetTrack() TrackSourceSettingsV1`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *UpdateSchemaSettingsInSourceV1Input) GetTrackOk() (*TrackSourceSettingsV1, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *UpdateSchemaSettingsInSourceV1Input) SetTrack(v TrackSourceSettingsV1)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *UpdateSchemaSettingsInSourceV1Input) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetIdentify

`func (o *UpdateSchemaSettingsInSourceV1Input) GetIdentify() IdentifySourceSettingsV1`

GetIdentify returns the Identify field if non-nil, zero value otherwise.

### GetIdentifyOk

`func (o *UpdateSchemaSettingsInSourceV1Input) GetIdentifyOk() (*IdentifySourceSettingsV1, bool)`

GetIdentifyOk returns a tuple with the Identify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentify

`func (o *UpdateSchemaSettingsInSourceV1Input) SetIdentify(v IdentifySourceSettingsV1)`

SetIdentify sets Identify field to given value.

### HasIdentify

`func (o *UpdateSchemaSettingsInSourceV1Input) HasIdentify() bool`

HasIdentify returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateSchemaSettingsInSourceV1Input) GetGroup() GroupSourceSettingsV1`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateSchemaSettingsInSourceV1Input) GetGroupOk() (*GroupSourceSettingsV1, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateSchemaSettingsInSourceV1Input) SetGroup(v GroupSourceSettingsV1)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateSchemaSettingsInSourceV1Input) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetForwardingViolationsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingViolationsTo() string`

GetForwardingViolationsTo returns the ForwardingViolationsTo field if non-nil, zero value otherwise.

### GetForwardingViolationsToOk

`func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingViolationsToOk() (*string, bool)`

GetForwardingViolationsToOk returns a tuple with the ForwardingViolationsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingViolationsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) SetForwardingViolationsTo(v string)`

SetForwardingViolationsTo sets ForwardingViolationsTo field to given value.

### HasForwardingViolationsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) HasForwardingViolationsTo() bool`

HasForwardingViolationsTo returns a boolean if a field has been set.

### GetForwardingBlockedEventsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingBlockedEventsTo() string`

GetForwardingBlockedEventsTo returns the ForwardingBlockedEventsTo field if non-nil, zero value otherwise.

### GetForwardingBlockedEventsToOk

`func (o *UpdateSchemaSettingsInSourceV1Input) GetForwardingBlockedEventsToOk() (*string, bool)`

GetForwardingBlockedEventsToOk returns a tuple with the ForwardingBlockedEventsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardingBlockedEventsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) SetForwardingBlockedEventsTo(v string)`

SetForwardingBlockedEventsTo sets ForwardingBlockedEventsTo field to given value.

### HasForwardingBlockedEventsTo

`func (o *UpdateSchemaSettingsInSourceV1Input) HasForwardingBlockedEventsTo() bool`

HasForwardingBlockedEventsTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


