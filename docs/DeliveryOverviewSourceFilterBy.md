# DeliveryOverviewSourceFilterBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscardReason** | Pointer to **[]string** | A list of strings of discard reasons.  See [Discard Record Documentation](https://segment.com/docs/connections/delivery-overview/#troubleshooting) for valid error codes. | [optional] 
**EventName** | Pointer to **[]string** | A list of strings of event names. | [optional] 
**EventType** | Pointer to **[]string** | A list of strings of event types. Valid options are: &#x60;alias&#x60;, &#x60;group&#x60;, &#x60;identify&#x60;, &#x60;page&#x60;, &#x60;screen&#x60;, and &#x60;track&#x60;. | [optional] 
**AppVersion** | Pointer to **[]string** | A list of strings of app versions. | [optional] 

## Methods

### NewDeliveryOverviewSourceFilterBy

`func NewDeliveryOverviewSourceFilterBy() *DeliveryOverviewSourceFilterBy`

NewDeliveryOverviewSourceFilterBy instantiates a new DeliveryOverviewSourceFilterBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryOverviewSourceFilterByWithDefaults

`func NewDeliveryOverviewSourceFilterByWithDefaults() *DeliveryOverviewSourceFilterBy`

NewDeliveryOverviewSourceFilterByWithDefaults instantiates a new DeliveryOverviewSourceFilterBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscardReason

`func (o *DeliveryOverviewSourceFilterBy) GetDiscardReason() []string`

GetDiscardReason returns the DiscardReason field if non-nil, zero value otherwise.

### GetDiscardReasonOk

`func (o *DeliveryOverviewSourceFilterBy) GetDiscardReasonOk() (*[]string, bool)`

GetDiscardReasonOk returns a tuple with the DiscardReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardReason

`func (o *DeliveryOverviewSourceFilterBy) SetDiscardReason(v []string)`

SetDiscardReason sets DiscardReason field to given value.

### HasDiscardReason

`func (o *DeliveryOverviewSourceFilterBy) HasDiscardReason() bool`

HasDiscardReason returns a boolean if a field has been set.

### GetEventName

`func (o *DeliveryOverviewSourceFilterBy) GetEventName() []string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DeliveryOverviewSourceFilterBy) GetEventNameOk() (*[]string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DeliveryOverviewSourceFilterBy) SetEventName(v []string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DeliveryOverviewSourceFilterBy) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetEventType

`func (o *DeliveryOverviewSourceFilterBy) GetEventType() []string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *DeliveryOverviewSourceFilterBy) GetEventTypeOk() (*[]string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *DeliveryOverviewSourceFilterBy) SetEventType(v []string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *DeliveryOverviewSourceFilterBy) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAppVersion

`func (o *DeliveryOverviewSourceFilterBy) GetAppVersion() []string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeliveryOverviewSourceFilterBy) GetAppVersionOk() (*[]string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeliveryOverviewSourceFilterBy) SetAppVersion(v []string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeliveryOverviewSourceFilterBy) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


