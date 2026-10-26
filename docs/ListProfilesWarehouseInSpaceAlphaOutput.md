# ListProfilesWarehouseInSpaceAlphaOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfilesWarehouses** | [**[]ProfilesWarehouseAlpha**](ProfilesWarehouseAlpha.md) | A list of Warehouses that belong to the Workspace. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListProfilesWarehouseInSpaceAlphaOutput

`func NewListProfilesWarehouseInSpaceAlphaOutput(profilesWarehouses []ProfilesWarehouseAlpha, ) *ListProfilesWarehouseInSpaceAlphaOutput`

NewListProfilesWarehouseInSpaceAlphaOutput instantiates a new ListProfilesWarehouseInSpaceAlphaOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProfilesWarehouseInSpaceAlphaOutputWithDefaults

`func NewListProfilesWarehouseInSpaceAlphaOutputWithDefaults() *ListProfilesWarehouseInSpaceAlphaOutput`

NewListProfilesWarehouseInSpaceAlphaOutputWithDefaults instantiates a new ListProfilesWarehouseInSpaceAlphaOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfilesWarehouses

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetProfilesWarehouses() []ProfilesWarehouseAlpha`

GetProfilesWarehouses returns the ProfilesWarehouses field if non-nil, zero value otherwise.

### GetProfilesWarehousesOk

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetProfilesWarehousesOk() (*[]ProfilesWarehouseAlpha, bool)`

GetProfilesWarehousesOk returns a tuple with the ProfilesWarehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilesWarehouses

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) SetProfilesWarehouses(v []ProfilesWarehouseAlpha)`

SetProfilesWarehouses sets ProfilesWarehouses field to given value.


### GetPagination

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListProfilesWarehouseInSpaceAlphaOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


