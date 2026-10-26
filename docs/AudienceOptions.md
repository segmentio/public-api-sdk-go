# AudienceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeHistoricalData** | Pointer to **bool** | Determines whether data prior to the audience being created is included when determining audience membership. Note that including historical data may be needed in order to properly handle the definition specified. In these cases, Segment will automatically handle including historical data and the response will return the includeHistoricalData parameter as true. | [optional] 
**FilterByExternalIds** | Pointer to **[]string** | The set of profile external identifiers being used to determine audience membership. Profiles will only be considered for audience membership if the profile has at least one external id whose key matches a value in this set. If unspecified, a default set of external identifiers will be used: [&#39;user_id&#39;, &#39;email&#39;, &#39;android.idfa&#39;, &#39;ios.idfa&#39;]. | [optional] 
**BackfillEventDataDays** | Pointer to **float32** | If specified, the value of this field indicates the number of days (specified from the date the audience was created) that event data will be included from when determining audience membership. If not specified, the default is set to the maximum event window size, or 7 days if no window condition is defined. Note that this is applied on a best-effort basis and may not always be applicable. In such cases, the response will not return a backfillEventDataDays value, and all available data will be taken into account. Note that includeHistoricalData must be set to true. | [optional] 

## Methods

### NewAudienceOptions

`func NewAudienceOptions() *AudienceOptions`

NewAudienceOptions instantiates a new AudienceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudienceOptionsWithDefaults

`func NewAudienceOptionsWithDefaults() *AudienceOptions`

NewAudienceOptionsWithDefaults instantiates a new AudienceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeHistoricalData

`func (o *AudienceOptions) GetIncludeHistoricalData() bool`

GetIncludeHistoricalData returns the IncludeHistoricalData field if non-nil, zero value otherwise.

### GetIncludeHistoricalDataOk

`func (o *AudienceOptions) GetIncludeHistoricalDataOk() (*bool, bool)`

GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHistoricalData

`func (o *AudienceOptions) SetIncludeHistoricalData(v bool)`

SetIncludeHistoricalData sets IncludeHistoricalData field to given value.

### HasIncludeHistoricalData

`func (o *AudienceOptions) HasIncludeHistoricalData() bool`

HasIncludeHistoricalData returns a boolean if a field has been set.

### GetFilterByExternalIds

`func (o *AudienceOptions) GetFilterByExternalIds() []string`

GetFilterByExternalIds returns the FilterByExternalIds field if non-nil, zero value otherwise.

### GetFilterByExternalIdsOk

`func (o *AudienceOptions) GetFilterByExternalIdsOk() (*[]string, bool)`

GetFilterByExternalIdsOk returns a tuple with the FilterByExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByExternalIds

`func (o *AudienceOptions) SetFilterByExternalIds(v []string)`

SetFilterByExternalIds sets FilterByExternalIds field to given value.

### HasFilterByExternalIds

`func (o *AudienceOptions) HasFilterByExternalIds() bool`

HasFilterByExternalIds returns a boolean if a field has been set.

### GetBackfillEventDataDays

`func (o *AudienceOptions) GetBackfillEventDataDays() float32`

GetBackfillEventDataDays returns the BackfillEventDataDays field if non-nil, zero value otherwise.

### GetBackfillEventDataDaysOk

`func (o *AudienceOptions) GetBackfillEventDataDaysOk() (*float32, bool)`

GetBackfillEventDataDaysOk returns a tuple with the BackfillEventDataDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillEventDataDays

`func (o *AudienceOptions) SetBackfillEventDataDays(v float32)`

SetBackfillEventDataDays sets BackfillEventDataDays field to given value.

### HasBackfillEventDataDays

`func (o *AudienceOptions) HasBackfillEventDataDays() bool`

HasBackfillEventDataDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


