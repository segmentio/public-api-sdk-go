# CreateDownloadAlphaInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionId** | **string** | The collection&#39;s unique id. | 
**Hour** | **string** | The ISO8601 formatted timestamp corresponding to a specific hour and day to retrieve data for. E.g.: 2025-05-07T23:00:00Z Objects are bucketed by hour and a month of data is retained. | 

## Methods

### NewCreateDownloadAlphaInput

`func NewCreateDownloadAlphaInput(collectionId string, hour string, ) *CreateDownloadAlphaInput`

NewCreateDownloadAlphaInput instantiates a new CreateDownloadAlphaInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDownloadAlphaInputWithDefaults

`func NewCreateDownloadAlphaInputWithDefaults() *CreateDownloadAlphaInput`

NewCreateDownloadAlphaInputWithDefaults instantiates a new CreateDownloadAlphaInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionId

`func (o *CreateDownloadAlphaInput) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CreateDownloadAlphaInput) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CreateDownloadAlphaInput) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetHour

`func (o *CreateDownloadAlphaInput) GetHour() string`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *CreateDownloadAlphaInput) GetHourOk() (*string, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *CreateDownloadAlphaInput) SetHour(v string)`

SetHour sets Hour field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


