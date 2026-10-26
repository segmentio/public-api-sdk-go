# ReadAudiencePreviewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterByExternalIds** | Pointer to **[]string** | The set of profile external identifiers being used to determine audience preview membership. Profiles will only be considered for audience preview membership if the profile has at least one external id whose key matches a value in this set. | [optional] 
**IncludeHistoricalData** | Pointer to **bool** | Determines whether data prior to the audience preview being created is included when determining audience preview membership. Note that including historical data may be needed in order to properly handle the definition specified. In these cases, Segment will automatically handle including historical data and the response will return the includeHistoricalData parameter as true. | [optional] 
**BackfillEventDataDays** | Pointer to **float32** | If specified, the value of this field indicates the number of days (specified from the date the audience preview was created) that event data will be included from when determining audience preview membership. If unspecified, defer to the value of &#x60;includeHistoricalData&#x60; to determine whether historical data is either entirely included or entirely excluded when determining audience preview membership. | [optional] 

## Methods

### NewReadAudiencePreviewOptions

`func NewReadAudiencePreviewOptions() *ReadAudiencePreviewOptions`

NewReadAudiencePreviewOptions instantiates a new ReadAudiencePreviewOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadAudiencePreviewOptionsWithDefaults

`func NewReadAudiencePreviewOptionsWithDefaults() *ReadAudiencePreviewOptions`

NewReadAudiencePreviewOptionsWithDefaults instantiates a new ReadAudiencePreviewOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterByExternalIds

`func (o *ReadAudiencePreviewOptions) GetFilterByExternalIds() []string`

GetFilterByExternalIds returns the FilterByExternalIds field if non-nil, zero value otherwise.

### GetFilterByExternalIdsOk

`func (o *ReadAudiencePreviewOptions) GetFilterByExternalIdsOk() (*[]string, bool)`

GetFilterByExternalIdsOk returns a tuple with the FilterByExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByExternalIds

`func (o *ReadAudiencePreviewOptions) SetFilterByExternalIds(v []string)`

SetFilterByExternalIds sets FilterByExternalIds field to given value.

### HasFilterByExternalIds

`func (o *ReadAudiencePreviewOptions) HasFilterByExternalIds() bool`

HasFilterByExternalIds returns a boolean if a field has been set.

### GetIncludeHistoricalData

`func (o *ReadAudiencePreviewOptions) GetIncludeHistoricalData() bool`

GetIncludeHistoricalData returns the IncludeHistoricalData field if non-nil, zero value otherwise.

### GetIncludeHistoricalDataOk

`func (o *ReadAudiencePreviewOptions) GetIncludeHistoricalDataOk() (*bool, bool)`

GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHistoricalData

`func (o *ReadAudiencePreviewOptions) SetIncludeHistoricalData(v bool)`

SetIncludeHistoricalData sets IncludeHistoricalData field to given value.

### HasIncludeHistoricalData

`func (o *ReadAudiencePreviewOptions) HasIncludeHistoricalData() bool`

HasIncludeHistoricalData returns a boolean if a field has been set.

### GetBackfillEventDataDays

`func (o *ReadAudiencePreviewOptions) GetBackfillEventDataDays() float32`

GetBackfillEventDataDays returns the BackfillEventDataDays field if non-nil, zero value otherwise.

### GetBackfillEventDataDaysOk

`func (o *ReadAudiencePreviewOptions) GetBackfillEventDataDaysOk() (*float32, bool)`

GetBackfillEventDataDaysOk returns a tuple with the BackfillEventDataDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillEventDataDays

`func (o *ReadAudiencePreviewOptions) SetBackfillEventDataDays(v float32)`

SetBackfillEventDataDays sets BackfillEventDataDays field to given value.

### HasBackfillEventDataDays

`func (o *ReadAudiencePreviewOptions) HasBackfillEventDataDays() bool`

HasBackfillEventDataDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


