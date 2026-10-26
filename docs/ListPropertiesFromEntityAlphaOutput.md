# ListPropertiesFromEntityAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityProperties** | [**[]EntityPropertyAlpha**](EntityPropertyAlpha.md) | List of properties for the entity. | 
**Pagination** | [**ForwardOnlyPaginationOutput**](ForwardOnlyPaginationOutput.md) |  | 

## Methods

### NewListPropertiesFromEntityAlphaOutput

`func NewListPropertiesFromEntityAlphaOutput(entityProperties []EntityPropertyAlpha, pagination ForwardOnlyPaginationOutput, ) *ListPropertiesFromEntityAlphaOutput`

NewListPropertiesFromEntityAlphaOutput instantiates a new ListPropertiesFromEntityAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPropertiesFromEntityAlphaOutputWithDefaults

`func NewListPropertiesFromEntityAlphaOutputWithDefaults() *ListPropertiesFromEntityAlphaOutput`

NewListPropertiesFromEntityAlphaOutputWithDefaults instantiates a new ListPropertiesFromEntityAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityProperties

`func (o *ListPropertiesFromEntityAlphaOutput) GetEntityProperties() []EntityPropertyAlpha`

GetEntityProperties returns the EntityProperties field if non-nil, zero value otherwise.

### GetEntityPropertiesOk

`func (o *ListPropertiesFromEntityAlphaOutput) GetEntityPropertiesOk() (*[]EntityPropertyAlpha, bool)`

GetEntityPropertiesOk returns a tuple with the EntityProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityProperties

`func (o *ListPropertiesFromEntityAlphaOutput) SetEntityProperties(v []EntityPropertyAlpha)`

SetEntityProperties sets EntityProperties field to given value.


### GetPagination

`func (o *ListPropertiesFromEntityAlphaOutput) GetPagination() ForwardOnlyPaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPropertiesFromEntityAlphaOutput) GetPaginationOk() (*ForwardOnlyPaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPropertiesFromEntityAlphaOutput) SetPagination(v ForwardOnlyPaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


