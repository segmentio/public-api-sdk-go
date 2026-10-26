# TrackSourceSettingsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUnplannedEvents** | Pointer to **bool** | Enable to allow unplanned track events.  Config API note: equal to &#x60;allowUnplannedTrackEvents&#x60;. | [optional] 
**AllowUnplannedEventProperties** | Pointer to **bool** | Enable to allow unplanned track event properties.  Config API note: equal to &#x60;allowUnplannedTrackEventProperties&#x60;. | [optional] 
**AllowEventOnViolations** | Pointer to **bool** | Allow track event on violations.  Config API note: equal to &#x60;allowTrackEventOnViolations&#x60;. | [optional] 
**AllowPropertiesOnViolations** | Pointer to **bool** | Enable to allow track properties on violations.  Config API note: equal to &#x60;allowTrackEventPropertiesOnViolations&#x60;. | [optional] 
**CommonEventOnViolations** | Pointer to **string** | The common track event on violations.  Config API note: equal to &#x60;commonTrackEventOnViolations&#x60;. | [optional] 

## Methods

### NewTrackSourceSettingsV1

`func NewTrackSourceSettingsV1() *TrackSourceSettingsV1`

NewTrackSourceSettingsV1 instantiates a new TrackSourceSettingsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackSourceSettingsV1WithDefaults

`func NewTrackSourceSettingsV1WithDefaults() *TrackSourceSettingsV1`

NewTrackSourceSettingsV1WithDefaults instantiates a new TrackSourceSettingsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUnplannedEvents

`func (o *TrackSourceSettingsV1) GetAllowUnplannedEvents() bool`

GetAllowUnplannedEvents returns the AllowUnplannedEvents field if non-nil, zero value otherwise.

### GetAllowUnplannedEventsOk

`func (o *TrackSourceSettingsV1) GetAllowUnplannedEventsOk() (*bool, bool)`

GetAllowUnplannedEventsOk returns a tuple with the AllowUnplannedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnplannedEvents

`func (o *TrackSourceSettingsV1) SetAllowUnplannedEvents(v bool)`

SetAllowUnplannedEvents sets AllowUnplannedEvents field to given value.

### HasAllowUnplannedEvents

`func (o *TrackSourceSettingsV1) HasAllowUnplannedEvents() bool`

HasAllowUnplannedEvents returns a boolean if a field has been set.

### GetAllowUnplannedEventProperties

`func (o *TrackSourceSettingsV1) GetAllowUnplannedEventProperties() bool`

GetAllowUnplannedEventProperties returns the AllowUnplannedEventProperties field if non-nil, zero value otherwise.

### GetAllowUnplannedEventPropertiesOk

`func (o *TrackSourceSettingsV1) GetAllowUnplannedEventPropertiesOk() (*bool, bool)`

GetAllowUnplannedEventPropertiesOk returns a tuple with the AllowUnplannedEventProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnplannedEventProperties

`func (o *TrackSourceSettingsV1) SetAllowUnplannedEventProperties(v bool)`

SetAllowUnplannedEventProperties sets AllowUnplannedEventProperties field to given value.

### HasAllowUnplannedEventProperties

`func (o *TrackSourceSettingsV1) HasAllowUnplannedEventProperties() bool`

HasAllowUnplannedEventProperties returns a boolean if a field has been set.

### GetAllowEventOnViolations

`func (o *TrackSourceSettingsV1) GetAllowEventOnViolations() bool`

GetAllowEventOnViolations returns the AllowEventOnViolations field if non-nil, zero value otherwise.

### GetAllowEventOnViolationsOk

`func (o *TrackSourceSettingsV1) GetAllowEventOnViolationsOk() (*bool, bool)`

GetAllowEventOnViolationsOk returns a tuple with the AllowEventOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEventOnViolations

`func (o *TrackSourceSettingsV1) SetAllowEventOnViolations(v bool)`

SetAllowEventOnViolations sets AllowEventOnViolations field to given value.

### HasAllowEventOnViolations

`func (o *TrackSourceSettingsV1) HasAllowEventOnViolations() bool`

HasAllowEventOnViolations returns a boolean if a field has been set.

### GetAllowPropertiesOnViolations

`func (o *TrackSourceSettingsV1) GetAllowPropertiesOnViolations() bool`

GetAllowPropertiesOnViolations returns the AllowPropertiesOnViolations field if non-nil, zero value otherwise.

### GetAllowPropertiesOnViolationsOk

`func (o *TrackSourceSettingsV1) GetAllowPropertiesOnViolationsOk() (*bool, bool)`

GetAllowPropertiesOnViolationsOk returns a tuple with the AllowPropertiesOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPropertiesOnViolations

`func (o *TrackSourceSettingsV1) SetAllowPropertiesOnViolations(v bool)`

SetAllowPropertiesOnViolations sets AllowPropertiesOnViolations field to given value.

### HasAllowPropertiesOnViolations

`func (o *TrackSourceSettingsV1) HasAllowPropertiesOnViolations() bool`

HasAllowPropertiesOnViolations returns a boolean if a field has been set.

### GetCommonEventOnViolations

`func (o *TrackSourceSettingsV1) GetCommonEventOnViolations() string`

GetCommonEventOnViolations returns the CommonEventOnViolations field if non-nil, zero value otherwise.

### GetCommonEventOnViolationsOk

`func (o *TrackSourceSettingsV1) GetCommonEventOnViolationsOk() (*string, bool)`

GetCommonEventOnViolationsOk returns a tuple with the CommonEventOnViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonEventOnViolations

`func (o *TrackSourceSettingsV1) SetCommonEventOnViolations(v string)`

SetCommonEventOnViolations sets CommonEventOnViolations field to given value.

### HasCommonEventOnViolations

`func (o *TrackSourceSettingsV1) HasCommonEventOnViolations() bool`

HasCommonEventOnViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


