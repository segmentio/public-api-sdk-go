# CancelReverseETLSyncForModelInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReasonForCanceling** | Pointer to **float32** | The reason for canceling the sync. - IncorrectModel &#x3D; 0 - IncorrectDest &#x3D; 1 - IncorrectKeys &#x3D; 2 - IncorrectMapping &#x3D; 3 - Other &#x3D; 4 | [optional] 

## Methods

### NewCancelReverseETLSyncForModelInput

`func NewCancelReverseETLSyncForModelInput() *CancelReverseETLSyncForModelInput`

NewCancelReverseETLSyncForModelInput instantiates a new CancelReverseETLSyncForModelInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelReverseETLSyncForModelInputWithDefaults

`func NewCancelReverseETLSyncForModelInputWithDefaults() *CancelReverseETLSyncForModelInput`

NewCancelReverseETLSyncForModelInputWithDefaults instantiates a new CancelReverseETLSyncForModelInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReasonForCanceling

`func (o *CancelReverseETLSyncForModelInput) GetReasonForCanceling() float32`

GetReasonForCanceling returns the ReasonForCanceling field if non-nil, zero value otherwise.

### GetReasonForCancelingOk

`func (o *CancelReverseETLSyncForModelInput) GetReasonForCancelingOk() (*float32, bool)`

GetReasonForCancelingOk returns a tuple with the ReasonForCanceling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForCanceling

`func (o *CancelReverseETLSyncForModelInput) SetReasonForCanceling(v float32)`

SetReasonForCanceling sets ReasonForCanceling field to given value.

### HasReasonForCanceling

`func (o *CancelReverseETLSyncForModelInput) HasReasonForCanceling() bool`

HasReasonForCanceling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


