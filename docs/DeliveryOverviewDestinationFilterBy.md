# DeliveryOverviewDestinationFilterBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscardReason** | Pointer to **[]string** | A list of strings of discard reasons.  See [Discard Record Documentation](https://segment.com/docs/connections/delivery-overview/#troubleshooting) for valid error codes. | [optional] 
**EventName** | Pointer to **[]string** | A list of strings of event names. | [optional] 
**EventType** | Pointer to **[]string** | A list of strings of event types. Valid options are: &#x60;alias&#x60;, &#x60;group&#x60;, &#x60;identify&#x60;, &#x60;page&#x60;, &#x60;screen&#x60;, and &#x60;track&#x60;. | [optional] 
**AppVersion** | Pointer to **[]string** | A list of strings of app versions. | [optional] 
**SubscriptionId** | Pointer to **[]string** | A list of strings of subscription IDs for Actions Destinations. | [optional] 
**ActivationId** | Pointer to **[]string** | A list of strings of event context IDs from a Linked Audience mapping/activation. | [optional] 
**AudienceId** | Pointer to **[]string** | A list of strings of audience IDs for a Linked Audience. | [optional] 
**SpaceId** | Pointer to **[]string** | A list of strings of space IDs for a Linked Audience. | [optional] 

## Methods

### NewDeliveryOverviewDestinationFilterBy

`func NewDeliveryOverviewDestinationFilterBy() *DeliveryOverviewDestinationFilterBy`

NewDeliveryOverviewDestinationFilterBy instantiates a new DeliveryOverviewDestinationFilterBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOverviewDestinationFilterByWithDefaults

`func NewDeliveryOverviewDestinationFilterByWithDefaults() *DeliveryOverviewDestinationFilterBy`

NewDeliveryOverviewDestinationFilterByWithDefaults instantiates a new DeliveryOverviewDestinationFilterBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscardReason

`func (o *DeliveryOverviewDestinationFilterBy) GetDiscardReason() []string`

GetDiscardReason returns the DiscardReason field if non-nil, zero value otherwise.

### GetDiscardReasonOk

`func (o *DeliveryOverviewDestinationFilterBy) GetDiscardReasonOk() (*[]string, bool)`

GetDiscardReasonOk returns a tuple with the DiscardReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardReason

`func (o *DeliveryOverviewDestinationFilterBy) SetDiscardReason(v []string)`

SetDiscardReason sets DiscardReason field to given value.

### HasDiscardReason

`func (o *DeliveryOverviewDestinationFilterBy) HasDiscardReason() bool`

HasDiscardReason returns a boolean if a field has been set.

### GetEventName

`func (o *DeliveryOverviewDestinationFilterBy) GetEventName() []string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DeliveryOverviewDestinationFilterBy) GetEventNameOk() (*[]string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DeliveryOverviewDestinationFilterBy) SetEventName(v []string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DeliveryOverviewDestinationFilterBy) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *DeliveryOverviewDestinationFilterBy) GetEventType() []string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DeliveryOverviewDestinationFilterBy) GetEventTypeOk() (*[]string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DeliveryOverviewDestinationFilterBy) SetEventType(v []string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *DeliveryOverviewDestinationFilterBy) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAppVersion

`func (o *DeliveryOverviewDestinationFilterBy) GetAppVersion() []string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeliveryOverviewDestinationFilterBy) GetAppVersionOk() (*[]string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeliveryOverviewDestinationFilterBy) SetAppVersion(v []string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeliveryOverviewDestinationFilterBy) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DeliveryOverviewDestinationFilterBy) GetSubscriptionId() []string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DeliveryOverviewDestinationFilterBy) GetSubscriptionIdOk() (*[]string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DeliveryOverviewDestinationFilterBy) SetSubscriptionId(v []string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DeliveryOverviewDestinationFilterBy) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetActivationId

`func (o *DeliveryOverviewDestinationFilterBy) GetActivationId() []string`

GetActivationId returns the ActivationId field if non-nil, zero value otherwise.

### GetActivationIdOk

`func (o *DeliveryOverviewDestinationFilterBy) GetActivationIdOk() (*[]string, bool)`

GetActivationIdOk returns a tuple with the ActivationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationId

`func (o *DeliveryOverviewDestinationFilterBy) SetActivationId(v []string)`

SetActivationId sets ActivationId field to given value.

### HasActivationId

`func (o *DeliveryOverviewDestinationFilterBy) HasActivationId() bool`

HasActivationId returns a boolean if a field has been set.

### GetAudienceId

`func (o *DeliveryOverviewDestinationFilterBy) GetAudienceId() []string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *DeliveryOverviewDestinationFilterBy) GetAudienceIdOk() (*[]string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *DeliveryOverviewDestinationFilterBy) SetAudienceId(v []string)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *DeliveryOverviewDestinationFilterBy) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetSpaceId

`func (o *DeliveryOverviewDestinationFilterBy) GetSpaceId() []string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DeliveryOverviewDestinationFilterBy) GetSpaceIdOk() (*[]string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DeliveryOverviewDestinationFilterBy) SetSpaceId(v []string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DeliveryOverviewDestinationFilterBy) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


