# ListSubscriptionsFromDestinationAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | [**[]DestinationSubscription**](DestinationSubscription.md) | A list of Destination subscriptions. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListSubscriptionsFromDestinationAlphaOutput

`func NewListSubscriptionsFromDestinationAlphaOutput(subscriptions []DestinationSubscription, ) *ListSubscriptionsFromDestinationAlphaOutput`

NewListSubscriptionsFromDestinationAlphaOutput instantiates a new ListSubscriptionsFromDestinationAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionsFromDestinationAlphaOutputWithDefaults

`func NewListSubscriptionsFromDestinationAlphaOutputWithDefaults() *ListSubscriptionsFromDestinationAlphaOutput`

NewListSubscriptionsFromDestinationAlphaOutputWithDefaults instantiates a new ListSubscriptionsFromDestinationAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *ListSubscriptionsFromDestinationAlphaOutput) GetSubscriptions() []DestinationSubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ListSubscriptionsFromDestinationAlphaOutput) GetSubscriptionsOk() (*[]DestinationSubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ListSubscriptionsFromDestinationAlphaOutput) SetSubscriptions(v []DestinationSubscription)`

SetSubscriptions sets Subscriptions field to given value.


### GetPagination

`func (o *ListSubscriptionsFromDestinationAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSubscriptionsFromDestinationAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSubscriptionsFromDestinationAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSubscriptionsFromDestinationAlphaOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


