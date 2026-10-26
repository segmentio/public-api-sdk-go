# RequestErrorEnvelope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]RequestError**](RequestError.md) |  | 

## Methods

### NewRequestErrorEnvelope

`func NewRequestErrorEnvelope(errors []RequestError, ) *RequestErrorEnvelope`

NewRequestErrorEnvelope instantiates a new RequestErrorEnvelope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestErrorEnvelopeWithDefaults

`func NewRequestErrorEnvelopeWithDefaults() *RequestErrorEnvelope`

NewRequestErrorEnvelopeWithDefaults instantiates a new RequestErrorEnvelope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *RequestErrorEnvelope) GetErrors() []RequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RequestErrorEnvelope) GetErrorsOk() (*[]RequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RequestErrorEnvelope) SetErrors(v []RequestError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


