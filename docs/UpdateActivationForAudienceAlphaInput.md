# UpdateActivationForAudienceAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determines whether an activation is enabled. | [optional] 
**ActivationName** | Pointer to **string** | Activation name. For Warehouse Destinations, this is used as the table name. Omit this field to leave the name unchanged. An empty string is ignored rather than applied. | [optional] 
**DisplayName** | Pointer to **NullableString** | Optional human-readable label for the activation. Only supported for Warehouse Destinations. Send null to clear. When omitted, the activationName is used as the label. | [optional] 
**Personalization** | Pointer to [**PersonalizationInput**](PersonalizationInput.md) |  | [optional] 
**DestinationMapping** | Pointer to [**DestinationSubscriptionConfiguration**](DestinationSubscriptionConfiguration.md) |  | [optional] 
**PerformResync** | Pointer to **bool** | Determines whether to perform a full resync after the update. If true, the entire audience is resent to the Destination using the updated configuration. If false, the update applies only to future syncs. | [optional] 

## Methods

### NewUpdateActivationForAudienceAlphaInput

`func NewUpdateActivationForAudienceAlphaInput() *UpdateActivationForAudienceAlphaInput`

NewUpdateActivationForAudienceAlphaInput instantiates a new UpdateActivationForAudienceAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActivationForAudienceAlphaInputWithDefaults

`func NewUpdateActivationForAudienceAlphaInputWithDefaults() *UpdateActivationForAudienceAlphaInput`

NewUpdateActivationForAudienceAlphaInputWithDefaults instantiates a new UpdateActivationForAudienceAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateActivationForAudienceAlphaInput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateActivationForAudienceAlphaInput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateActivationForAudienceAlphaInput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateActivationForAudienceAlphaInput) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetActivationName

`func (o *UpdateActivationForAudienceAlphaInput) GetActivationName() string`

GetActivationName returns the ActivationName field if non-nil, zero value otherwise.

### GetActivationNameOk

`func (o *UpdateActivationForAudienceAlphaInput) GetActivationNameOk() (*string, bool)`

GetActivationNameOk returns a tuple with the ActivationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationName

`func (o *UpdateActivationForAudienceAlphaInput) SetActivationName(v string)`

SetActivationName sets ActivationName field to given value.

### HasActivationName

`func (o *UpdateActivationForAudienceAlphaInput) HasActivationName() bool`

HasActivationName returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateActivationForAudienceAlphaInput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateActivationForAudienceAlphaInput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateActivationForAudienceAlphaInput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateActivationForAudienceAlphaInput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateActivationForAudienceAlphaInput) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateActivationForAudienceAlphaInput) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPersonalization

`func (o *UpdateActivationForAudienceAlphaInput) GetPersonalization() PersonalizationInput`

GetPersonalization returns the Personalization field if non-nil, zero value otherwise.

### GetPersonalizationOk

`func (o *UpdateActivationForAudienceAlphaInput) GetPersonalizationOk() (*PersonalizationInput, bool)`

GetPersonalizationOk returns a tuple with the Personalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalization

`func (o *UpdateActivationForAudienceAlphaInput) SetPersonalization(v PersonalizationInput)`

SetPersonalization sets Personalization field to given value.

### HasPersonalization

`func (o *UpdateActivationForAudienceAlphaInput) HasPersonalization() bool`

HasPersonalization returns a boolean if a field has been set.

### GetDestinationMapping

`func (o *UpdateActivationForAudienceAlphaInput) GetDestinationMapping() DestinationSubscriptionConfiguration`

GetDestinationMapping returns the DestinationMapping field if non-nil, zero value otherwise.

### GetDestinationMappingOk

`func (o *UpdateActivationForAudienceAlphaInput) GetDestinationMappingOk() (*DestinationSubscriptionConfiguration, bool)`

GetDestinationMappingOk returns a tuple with the DestinationMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMapping

`func (o *UpdateActivationForAudienceAlphaInput) SetDestinationMapping(v DestinationSubscriptionConfiguration)`

SetDestinationMapping sets DestinationMapping field to given value.

### HasDestinationMapping

`func (o *UpdateActivationForAudienceAlphaInput) HasDestinationMapping() bool`

HasDestinationMapping returns a boolean if a field has been set.

### GetPerformResync

`func (o *UpdateActivationForAudienceAlphaInput) GetPerformResync() bool`

GetPerformResync returns the PerformResync field if non-nil, zero value otherwise.

### GetPerformResyncOk

`func (o *UpdateActivationForAudienceAlphaInput) GetPerformResyncOk() (*bool, bool)`

GetPerformResyncOk returns a tuple with the PerformResync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformResync

`func (o *UpdateActivationForAudienceAlphaInput) SetPerformResync(v bool)`

SetPerformResync sets PerformResync field to given value.

### HasPerformResync

`func (o *UpdateActivationForAudienceAlphaInput) HasPerformResync() bool`

HasPerformResync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


