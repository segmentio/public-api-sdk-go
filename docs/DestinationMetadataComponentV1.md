# DestinationMetadataComponentV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The component type. | 
**Code** | **string** | Link to the repository hosting the code for this component. | 
**Owner** | Pointer to **string** | The owner of this component. Either &#39;SEGMENT&#39; or &#39;PARTNER&#39;. | [optional] 

## Methods

### NewDestinationMetadataComponentV1

`func NewDestinationMetadataComponentV1(type_ string, code string, ) *DestinationMetadataComponentV1`

NewDestinationMetadataComponentV1 instantiates a new DestinationMetadataComponentV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationMetadataComponentV1WithDefaults

`func NewDestinationMetadataComponentV1WithDefaults() *DestinationMetadataComponentV1`

NewDestinationMetadataComponentV1WithDefaults instantiates a new DestinationMetadataComponentV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationMetadataComponentV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationMetadataComponentV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationMetadataComponentV1) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *DestinationMetadataComponentV1) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DestinationMetadataComponentV1) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DestinationMetadataComponentV1) SetCode(v string)`

SetCode sets Code field to given value.


### GetOwner

`func (o *DestinationMetadataComponentV1) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DestinationMetadataComponentV1) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DestinationMetadataComponentV1) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DestinationMetadataComponentV1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


