# ActivationOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The activation id. | 
**Enabled** | **bool** | Activation Enabled Status. | 
**WorkspaceId** | **string** | The WORKSPACE id. | 
**SpaceId** | **string** | The space id. | 
**AudienceId** | **string** | The audience id. | 
**ConnectionId** | **string** | The connection id. | 
**ActivationType** | **string** | Determines when an event is sent to the Destination.  Possible values: - **Audience Entered**: Sends an event when a profile or entity enters the audience. - **Audience Exited**: Sends an event when a profile or entity exits the audience. - **Audience Membership Changed**: Sends an event for both entries and exits. This does not apply to entities. - **Entity Added**: Sends an event when an entity is added to the audience. - **Entity Removed**: Sends an event when an entity is removed from the audience. - **Entity Values Changed**: Sends an event when an entity&#39;s property values change. Restricted to freshness-enabled Destinations. - **All Events**: Sends events for all trigger conditions (enters, exits, and entity value changes for freshness Destinations; enters and exits only for non-freshness Destinations).  Note that events are sent for the profile, unless the audience is a Linked Audience. In that case, events are sent for the target entity defined for that audience. | 
**ActivationName** | **string** | Activation name. For Warehouse Destinations, this is the table name. | 
**DisplayName** | Pointer to **NullableString** | Human-readable label for the activation. Only present for Warehouse Destinations that have a display name configured. When null, the activationName serves as the label. | [optional] 
**Personalization** | [**PersonalizationOutput**](PersonalizationOutput.md) |  | 
**DestinationMapping** | Pointer to [**DestinationSubscriptionConfiguration**](DestinationSubscriptionConfiguration.md) |  | [optional] 
**PerformResync** | Pointer to **bool** | Indicates if a full resync is currently pending or in progress. | [optional] 

## Methods

### NewActivationOutput

`func NewActivationOutput(id string, enabled bool, workspaceId string, spaceId string, audienceId string, connectionId string, activationType string, activationName string, personalization PersonalizationOutput, ) *ActivationOutput`

NewActivationOutput instantiates a new ActivationOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationOutputWithDefaults

`func NewActivationOutputWithDefaults() *ActivationOutput`

NewActivationOutputWithDefaults instantiates a new ActivationOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivationOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivationOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivationOutput) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *ActivationOutput) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActivationOutput) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActivationOutput) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetWorkspaceId

`func (o *ActivationOutput) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ActivationOutput) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ActivationOutput) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetSpaceId

`func (o *ActivationOutput) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *ActivationOutput) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *ActivationOutput) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetAudienceId

`func (o *ActivationOutput) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *ActivationOutput) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *ActivationOutput) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.


### GetConnectionId

`func (o *ActivationOutput) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ActivationOutput) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ActivationOutput) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetActivationType

`func (o *ActivationOutput) GetActivationType() string`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *ActivationOutput) GetActivationTypeOk() (*string, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *ActivationOutput) SetActivationType(v string)`

SetActivationType sets ActivationType field to given value.


### GetActivationName

`func (o *ActivationOutput) GetActivationName() string`

GetActivationName returns the ActivationName field if non-nil, zero value otherwise.

### GetActivationNameOk

`func (o *ActivationOutput) GetActivationNameOk() (*string, bool)`

GetActivationNameOk returns a tuple with the ActivationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationName

`func (o *ActivationOutput) SetActivationName(v string)`

SetActivationName sets ActivationName field to given value.


### GetDisplayName

`func (o *ActivationOutput) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ActivationOutput) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ActivationOutput) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ActivationOutput) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ActivationOutput) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ActivationOutput) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetPersonalization

`func (o *ActivationOutput) GetPersonalization() PersonalizationOutput`

GetPersonalization returns the Personalization field if non-nil, zero value otherwise.

### GetPersonalizationOk

`func (o *ActivationOutput) GetPersonalizationOk() (*PersonalizationOutput, bool)`

GetPersonalizationOk returns a tuple with the Personalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalization

`func (o *ActivationOutput) SetPersonalization(v PersonalizationOutput)`

SetPersonalization sets Personalization field to given value.


### GetDestinationMapping

`func (o *ActivationOutput) GetDestinationMapping() DestinationSubscriptionConfiguration`

GetDestinationMapping returns the DestinationMapping field if non-nil, zero value otherwise.

### GetDestinationMappingOk

`func (o *ActivationOutput) GetDestinationMappingOk() (*DestinationSubscriptionConfiguration, bool)`

GetDestinationMappingOk returns a tuple with the DestinationMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationMapping

`func (o *ActivationOutput) SetDestinationMapping(v DestinationSubscriptionConfiguration)`

SetDestinationMapping sets DestinationMapping field to given value.

### HasDestinationMapping

`func (o *ActivationOutput) HasDestinationMapping() bool`

HasDestinationMapping returns a boolean if a field has been set.

### GetPerformResync

`func (o *ActivationOutput) GetPerformResync() bool`

GetPerformResync returns the PerformResync field if non-nil, zero value otherwise.

### GetPerformResyncOk

`func (o *ActivationOutput) GetPerformResyncOk() (*bool, bool)`

GetPerformResyncOk returns a tuple with the PerformResync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformResync

`func (o *ActivationOutput) SetPerformResync(v bool)`

SetPerformResync sets PerformResync field to given value.

### HasPerformResync

`func (o *ActivationOutput) HasPerformResync() bool`

HasPerformResync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


