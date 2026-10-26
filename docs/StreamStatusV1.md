# StreamStatusV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DestinationStatus** | [**[]DestinationStatusV1**](DestinationStatusV1.md) |  | 

## Methods

### NewStreamStatusV1

`func NewStreamStatusV1(id string, destinationStatus []DestinationStatusV1, ) *StreamStatusV1`

NewStreamStatusV1 instantiates a new StreamStatusV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamStatusV1WithDefaults

`func NewStreamStatusV1WithDefaults() *StreamStatusV1`

NewStreamStatusV1WithDefaults instantiates a new StreamStatusV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamStatusV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamStatusV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamStatusV1) SetId(v string)`

SetId sets Id field to given value.


### GetDestinationStatus

`func (o *StreamStatusV1) GetDestinationStatus() []DestinationStatusV1`

GetDestinationStatus returns the DestinationStatus field if non-nil, zero value otherwise.

### GetDestinationStatusOk

`func (o *StreamStatusV1) GetDestinationStatusOk() (*[]DestinationStatusV1, bool)`

GetDestinationStatusOk returns a tuple with the DestinationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationStatus

`func (o *StreamStatusV1) SetDestinationStatus(v []DestinationStatusV1)`

SetDestinationStatus sets DestinationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


