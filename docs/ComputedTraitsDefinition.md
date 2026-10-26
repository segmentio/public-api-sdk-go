# ComputedTraitsDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The query language string defining the computed trait aggregation criteria. For guidance on using the query language, see the [Segment documentation site](https://segment.com/docs/api/public-api/query-language). | 
**Type** | **string** | The underlying data type being aggregated for this computed trait.  Possible values: users, accounts. | 

## Methods

### NewComputedTraitsDefinition

`func NewComputedTraitsDefinition(query string, type_ string, ) *ComputedTraitsDefinition`

NewComputedTraitsDefinition instantiates a new ComputedTraitsDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputedTraitsDefinitionWithDefaults

`func NewComputedTraitsDefinitionWithDefaults() *ComputedTraitsDefinition`

NewComputedTraitsDefinitionWithDefaults instantiates a new ComputedTraitsDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ComputedTraitsDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ComputedTraitsDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ComputedTraitsDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetType

`func (o *ComputedTraitsDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComputedTraitsDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComputedTraitsDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


