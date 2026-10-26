# EchoAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The HTTP method used for this round-trip.  Currently, this endpoint supports only &#x60;get&#x60; and &#x60;post&#x60; methods. | 
**Message** | **string** | The string passed in the &#x60;message&#x60; input field. | 
**Headers** | **map[string]interface{}** | The request&#39;s HTTP headers. | 

## Methods

### NewEchoAlphaOutput

`func NewEchoAlphaOutput(method string, message string, headers map[string]interface{}, ) *EchoAlphaOutput`

NewEchoAlphaOutput instantiates a new EchoAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEchoAlphaOutputWithDefaults

`func NewEchoAlphaOutputWithDefaults() *EchoAlphaOutput`

NewEchoAlphaOutputWithDefaults instantiates a new EchoAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *EchoAlphaOutput) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *EchoAlphaOutput) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *EchoAlphaOutput) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetMessage

`func (o *EchoAlphaOutput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EchoAlphaOutput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EchoAlphaOutput) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetHeaders

`func (o *EchoAlphaOutput) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EchoAlphaOutput) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EchoAlphaOutput) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


