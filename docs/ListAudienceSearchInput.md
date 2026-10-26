# ListAudienceSearchInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Field to filter by. | 
**Query** | **string** | Text to match against the selected field (max 255 characters). When searching by KEY, a prefix match is used. For all other options, a full wildcard search will be used. | 

## Methods

### NewListAudienceSearchInput

`func NewListAudienceSearchInput(type_ string, query string, ) *ListAudienceSearchInput`

NewListAudienceSearchInput instantiates a new ListAudienceSearchInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAudienceSearchInputWithDefaults

`func NewListAudienceSearchInputWithDefaults() *ListAudienceSearchInput`

NewListAudienceSearchInputWithDefaults instantiates a new ListAudienceSearchInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListAudienceSearchInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListAudienceSearchInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListAudienceSearchInput) SetType(v string)`

SetType sets Type field to given value.


### GetQuery

`func (o *ListAudienceSearchInput) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ListAudienceSearchInput) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ListAudienceSearchInput) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


