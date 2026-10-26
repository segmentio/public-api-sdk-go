# CancelReverseETLSyncForModelOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelId** | **string** | The id of the Model. | 
**SyncId** | **string** | The id of the Sync. | 
**ErrorCode** | Pointer to **string** | A place holder for a machine-friendly category for an error, if applicable. - \&quot;SyncAlreadyCanceled\&quot; - \&quot;SyncFinishedCannotCancel\&quot; | [optional] 
**ErrorMessage** | Pointer to **string** | A place holder for a human-readable description of the error, if applicable. - \&quot;sync already canceled\&quot; - \&quot;sync already finished\&quot;. | [optional] 
**Status** | Pointer to **string** | If no error, status will be CANCELLING, as the extract/load might take some time to cancel. | [optional] 

## Methods

### NewCancelReverseETLSyncForModelOutput

`func NewCancelReverseETLSyncForModelOutput(modelId string, syncId string, ) *CancelReverseETLSyncForModelOutput`

NewCancelReverseETLSyncForModelOutput instantiates a new CancelReverseETLSyncForModelOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelReverseETLSyncForModelOutputWithDefaults

`func NewCancelReverseETLSyncForModelOutputWithDefaults() *CancelReverseETLSyncForModelOutput`

NewCancelReverseETLSyncForModelOutputWithDefaults instantiates a new CancelReverseETLSyncForModelOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelId

`func (o *CancelReverseETLSyncForModelOutput) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *CancelReverseETLSyncForModelOutput) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *CancelReverseETLSyncForModelOutput) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetSyncId

`func (o *CancelReverseETLSyncForModelOutput) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *CancelReverseETLSyncForModelOutput) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *CancelReverseETLSyncForModelOutput) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.


### GetErrorCode

`func (o *CancelReverseETLSyncForModelOutput) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CancelReverseETLSyncForModelOutput) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CancelReverseETLSyncForModelOutput) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CancelReverseETLSyncForModelOutput) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *CancelReverseETLSyncForModelOutput) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *CancelReverseETLSyncForModelOutput) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *CancelReverseETLSyncForModelOutput) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *CancelReverseETLSyncForModelOutput) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetStatus

`func (o *CancelReverseETLSyncForModelOutput) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelReverseETLSyncForModelOutput) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelReverseETLSyncForModelOutput) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelReverseETLSyncForModelOutput) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


