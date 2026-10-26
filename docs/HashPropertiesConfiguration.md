# HashPropertiesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | Which algorithm to use to hash to properties. | 
**Key** | Pointer to **string** | Optional key to hash with. | [optional] 
**Encoding** | Pointer to **string** | Optional encoding to use for the hashing. | [optional] 
**Paths** | **[]string** | The paths to the properties to be hashed. | 

## Methods

### NewHashPropertiesConfiguration

`func NewHashPropertiesConfiguration(algorithm string, paths []string, ) *HashPropertiesConfiguration`

NewHashPropertiesConfiguration instantiates a new HashPropertiesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashPropertiesConfigurationWithDefaults

`func NewHashPropertiesConfigurationWithDefaults() *HashPropertiesConfiguration`

NewHashPropertiesConfigurationWithDefaults instantiates a new HashPropertiesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *HashPropertiesConfiguration) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *HashPropertiesConfiguration) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *HashPropertiesConfiguration) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetKey

`func (o *HashPropertiesConfiguration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HashPropertiesConfiguration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HashPropertiesConfiguration) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *HashPropertiesConfiguration) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetEncoding

`func (o *HashPropertiesConfiguration) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *HashPropertiesConfiguration) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *HashPropertiesConfiguration) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *HashPropertiesConfiguration) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetPaths

`func (o *HashPropertiesConfiguration) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *HashPropertiesConfiguration) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *HashPropertiesConfiguration) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


