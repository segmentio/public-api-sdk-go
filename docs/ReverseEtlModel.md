# ReverseEtlModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the Model. | 
**SourceId** | **string** | The id for the attached Source. | 
**Name** | **string** | A short, human-readable description of the Model. | 
**Description** | **string** | A longer, more descriptive explanation of the Model. | 
**Enabled** | **bool** | Indicates whether the Model should have syncs enabled. When disabled, no syncs will be triggered, regardless of the enabled status of the attached destinations/subscriptions. | 
**Query** | **string** | The SQL query that will be executed to extract data from the connected Source. | 
**QueryIdentifierColumn** | **string** | Indicates the column named in &#x60;query&#x60; that should be used to uniquely identify the extracted records. | 

## Methods

### NewReverseEtlModel

`func NewReverseEtlModel(id string, sourceId string, name string, description string, enabled bool, query string, queryIdentifierColumn string, ) *ReverseEtlModel`

NewReverseEtlModel instantiates a new ReverseEtlModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseEtlModelWithDefaults

`func NewReverseEtlModelWithDefaults() *ReverseEtlModel`

NewReverseEtlModelWithDefaults instantiates a new ReverseEtlModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReverseEtlModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReverseEtlModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReverseEtlModel) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *ReverseEtlModel) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ReverseEtlModel) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ReverseEtlModel) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetName

`func (o *ReverseEtlModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReverseEtlModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReverseEtlModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReverseEtlModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReverseEtlModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReverseEtlModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *ReverseEtlModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReverseEtlModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReverseEtlModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetQuery

`func (o *ReverseEtlModel) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ReverseEtlModel) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ReverseEtlModel) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQueryIdentifierColumn

`func (o *ReverseEtlModel) GetQueryIdentifierColumn() string`

GetQueryIdentifierColumn returns the QueryIdentifierColumn field if non-nil, zero value otherwise.

### GetQueryIdentifierColumnOk

`func (o *ReverseEtlModel) GetQueryIdentifierColumnOk() (*string, bool)`

GetQueryIdentifierColumnOk returns a tuple with the QueryIdentifierColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIdentifierColumn

`func (o *ReverseEtlModel) SetQueryIdentifierColumn(v string)`

SetQueryIdentifierColumn sets QueryIdentifierColumn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


