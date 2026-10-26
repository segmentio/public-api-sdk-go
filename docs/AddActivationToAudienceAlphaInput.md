# AddActivationToAudienceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines whether an activation is enabled. | [optional] 
**PerformResync** | **bool** | Determines whether to perform a full resync upon creation. If true, the entire audience is resent to the Destination from scratch. If false, only future changes will be synced. | 
**ActivationType** | **string** | Determines when an event is sent to the Destination.  Possible values:  - **Audience Entered**: Sends an event when a profile or entity enters the audience. - **Audience Exited**: Sends an event when a profile or entity exits the audience. - **Audience Membership Changed**: Sends an event for both entries and exits. This does not apply to entities. - **Entity Added**: Sends an event when an entity is added to the audience. - **Entity Removed**: Sends an event when an entity is removed from the audience. - **Entity Values Changed**: Sends an event when an entity&#39;s property values change. Restricted to freshness-enabled Destinations. - **All Events**: Sends events for all trigger conditions (enters, exits, and entity value changes for freshness Destinations; enters and exits only for non-freshness Destinations).  Note that events are sent for the profile, unless the audience is a Linked Audience. In that case, events are sent for the target entity defined for that audience.  Activations on Warehouse Destinations support only track-based activation types. **Audience Membership Changed** is identify-based, so requests using it for a Warehouse Destination are rejected. Use **Audience Entered** or **Audience Exited** instead. | 
**ActivationName** | **string** | Activation name. For Warehouse Destinations, this is used as the table name. | 
**DisplayName** | Pointer to **string** | Optional human-readable label for the activation. Only supported for Warehouse Destinations. When omitted, the activationName is used as the label. | [optional] 
**Personalization** | Pointer to [**PersonalizationInput**](PersonalizationInput.md) |  | [optional] 
**DestinationMapping** | Pointer to [**DestinationSubscriptionConfiguration**](DestinationSubscriptionConfiguration.md) |  | [optional] 

## Methods

### NewAddActivationToAudienceAlphaInput

`func NewAddActivationToAudienceAlphaInput(performResync bool, activationType string, activationName string, ) *AddActivationToAudienceAlphaInput`

NewAddActivationToAudienceAlphaInput instantiates a new AddActivationToAudienceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddActivationToAudienceAlphaInputWithDefaults

`func NewAddActivationToAudienceAlphaInputWithDefaults() *AddActivationToAudienceAlphaInput`

NewAddActivationToAudienceAlphaInputWithDefaults instantiates a new AddActivationToAudienceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AddActivationToAudienceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddActivationToAudienceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddActivationToAudienceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddActivationToAudienceAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPerformResync

`func (o *AddActivationToAudienceAlphaInput) GetPerformResync() bool`

GetPerformResync returns the PerformResync field if non-nil, zero value otherwise.

### GetPerformResyncOk

`func (o *AddActivationToAudienceAlphaInput) GetPerformResyncOk() (*bool, bool)`

GetPerformResyncOk returns a tuple with the PerformResync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformResync

`func (o *AddActivationToAudienceAlphaInput) SetPerformResync(v bool)`

SetPerformResync sets PerformResync field to given value.


### GetActivationType

`func (o *AddActivationToAudienceAlphaInput) GetActivationType() string`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *AddActivationToAudienceAlphaInput) GetActivationTypeOk() (*string, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *AddActivationToAudienceAlphaInput) SetActivationType(v string)`

SetActivationType sets ActivationType field to given value.


### GetActivationName

`func (o *AddActivationToAudienceAlphaInput) GetActivationName() string`

GetActivationName returns the ActivationName field if non-nil, zero value otherwise.

### GetActivationNameOk

`func (o *AddActivationToAudienceAlphaInput) GetActivationNameOk() (*string, bool)`

GetActivationNameOk returns a tuple with the ActivationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationName

`func (o *AddActivationToAudienceAlphaInput) SetActivationName(v string)`

SetActivationName sets ActivationName field to given value.


### GetDisplayName

`func (o *AddActivationToAudienceAlphaInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddActivationToAudienceAlphaInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddActivationToAudienceAlphaInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AddActivationToAudienceAlphaInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPersonalization

`func (o *AddActivationToAudienceAlphaInput) GetPersonalization() PersonalizationInput`

GetPersonalization returns the Personalization field if non-nil, zero value otherwise.

### GetPersonalizationOk

`func (o *AddActivationToAudienceAlphaInput) GetPersonalizationOk() (*PersonalizationInput, bool)`

GetPersonalizationOk returns a tuple with the Personalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalization

`func (o *AddActivationToAudienceAlphaInput) SetPersonalization(v PersonalizationInput)`

SetPersonalization sets Personalization field to given value.

### HasPersonalization

`func (o *AddActivationToAudienceAlphaInput) HasPersonalization() bool`

HasPersonalization returns a boolean if a field has been set.

### GetDestinationMapping

`func (o *AddActivationToAudienceAlphaInput) GetDestinationMapping() DestinationSubscriptionConfiguration`

GetDestinationMapping returns the DestinationMapping field if non-nil, zero value otherwise.

### GetDestinationMappingOk

`func (o *AddActivationToAudienceAlphaInput) GetDestinationMappingOk() (*DestinationSubscriptionConfiguration, bool)`

GetDestinationMappingOk returns a tuple with the DestinationMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMapping

`func (o *AddActivationToAudienceAlphaInput) SetDestinationMapping(v DestinationSubscriptionConfiguration)`

SetDestinationMapping sets DestinationMapping field to given value.

### HasDestinationMapping

`func (o *AddActivationToAudienceAlphaInput) HasDestinationMapping() bool`

HasDestinationMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


