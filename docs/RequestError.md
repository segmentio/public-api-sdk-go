# RequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type for this error (validation, server, unknown, etc). | 
**Message** | Pointer to **string** | An error message attached to this error. | [optional] 
**Field** | Pointer to **string** | The name of an input field from the request that triggered this error. | [optional] 
**Data** | Pointer to **interface{}** | Any extra data associated with this error. | [optional] 
**Status** | Pointer to **float32** | Http status code. | [optional] 

## Methods

### NewRequestError

`func NewRequestError(type_ string, ) *RequestError`

NewRequestError instantiates a new RequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestErrorWithDefaults

`func NewRequestErrorWithDefaults() *RequestError`

NewRequestErrorWithDefaults instantiates a new RequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestError) SetType(v string)`

SetType sets Type field to given value.


### GetMessage

`func (o *RequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RequestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetField

`func (o *RequestError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RequestError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RequestError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *RequestError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetData

`func (o *RequestError) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RequestError) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RequestError) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *RequestError) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *RequestError) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *RequestError) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetStatus

`func (o *RequestError) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestError) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestError) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


