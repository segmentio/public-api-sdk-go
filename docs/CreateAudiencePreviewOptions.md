# CreateAudiencePreviewOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterByExternalIds** | Pointer to **[]string** | The set of profile external identifiers being used to determine audience preview membership. Profiles will only be considered for audience preview membership if the profile has at least one external id whose key matches a value in this set. If unspecified, a default set of external identifiers will be used: &#x60;[&#39;user_id&#39;, &#39;email&#39;, &#39;android.idfa&#39;, &#39;ios.idfa&#39;]&#x60;. | [optional] 
**BackfillEventDataDays** | Pointer to **float32** | If specified, the value of this field indicates the number of days (specified from the date the audience preview was created) that event data will be included from when determining audience preview membership. If not specified, the default is set to the maximum event window size, or 7 days if no window condition is defined. Note that this is applied on a best-effort basis and may not always be applicable. In such cases, the response will not return a backfillEventDataDays value, and all available data will be taken into account. | [optional] 

## Methods

### NewCreateAudiencePreviewOptions

`func NewCreateAudiencePreviewOptions() *CreateAudiencePreviewOptions`

NewCreateAudiencePreviewOptions instantiates a new CreateAudiencePreviewOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAudiencePreviewOptionsWithDefaults

`func NewCreateAudiencePreviewOptionsWithDefaults() *CreateAudiencePreviewOptions`

NewCreateAudiencePreviewOptionsWithDefaults instantiates a new CreateAudiencePreviewOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterByExternalIds

`func (o *CreateAudiencePreviewOptions) GetFilterByExternalIds() []string`

GetFilterByExternalIds returns the FilterByExternalIds field if non-nil, zero value otherwise.

### GetFilterByExternalIdsOk

`func (o *CreateAudiencePreviewOptions) GetFilterByExternalIdsOk() (*[]string, bool)`

GetFilterByExternalIdsOk returns a tuple with the FilterByExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterByExternalIds

`func (o *CreateAudiencePreviewOptions) SetFilterByExternalIds(v []string)`

SetFilterByExternalIds sets FilterByExternalIds field to given value.

### HasFilterByExternalIds

`func (o *CreateAudiencePreviewOptions) HasFilterByExternalIds() bool`

HasFilterByExternalIds returns a boolean if a field has been set.

### GetBackfillEventDataDays

`func (o *CreateAudiencePreviewOptions) GetBackfillEventDataDays() float32`

GetBackfillEventDataDays returns the BackfillEventDataDays field if non-nil, zero value otherwise.

### GetBackfillEventDataDaysOk

`func (o *CreateAudiencePreviewOptions) GetBackfillEventDataDaysOk() (*float32, bool)`

GetBackfillEventDataDaysOk returns a tuple with the BackfillEventDataDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillEventDataDays

`func (o *CreateAudiencePreviewOptions) SetBackfillEventDataDays(v float32)`

SetBackfillEventDataDays sets BackfillEventDataDays field to given value.

### HasBackfillEventDataDays

`func (o *CreateAudiencePreviewOptions) HasBackfillEventDataDays() bool`

HasBackfillEventDataDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


