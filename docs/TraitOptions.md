# TraitOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeHistoricalData** | Pointer to **bool** | Determines whether data prior to the computed trait being created is included when determining the computed trait value. Note that including historical data may be needed in order to properly handle the definition specified. In these cases, Segment will automatically handle including historical data and the response will return the includeHistoricalData parameter as true. | [optional] 
**IncludeAnonymousUsers** | Pointer to **bool** | Determines whether anonymous users should be included when determining the computed trait value. | [optional] 
**BackfillDurationDays** | Pointer to **float32** | If specified, the value of this field indicates the number of days, specified from the date the audience was created, that event data will be included from when determining audience membership. If unspecified, defer to the value of &#x60;includeHistoricalData&#x60; to determine whether historical data is either entirely included or entirely excluded when determining audience membership. | [optional] 

## Methods

### NewTraitOptions

`func NewTraitOptions() *TraitOptions`

NewTraitOptions instantiates a new TraitOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraitOptionsWithDefaults

`func NewTraitOptionsWithDefaults() *TraitOptions`

NewTraitOptionsWithDefaults instantiates a new TraitOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeHistoricalData

`func (o *TraitOptions) GetIncludeHistoricalData() bool`

GetIncludeHistoricalData returns the IncludeHistoricalData field if non-nil, zero value otherwise.

### GetIncludeHistoricalDataOk

`func (o *TraitOptions) GetIncludeHistoricalDataOk() (*bool, bool)`

GetIncludeHistoricalDataOk returns a tuple with the IncludeHistoricalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHistoricalData

`func (o *TraitOptions) SetIncludeHistoricalData(v bool)`

SetIncludeHistoricalData sets IncludeHistoricalData field to given value.

### HasIncludeHistoricalData

`func (o *TraitOptions) HasIncludeHistoricalData() bool`

HasIncludeHistoricalData returns a boolean if a field has been set.

### GetIncludeAnonymousUsers

`func (o *TraitOptions) GetIncludeAnonymousUsers() bool`

GetIncludeAnonymousUsers returns the IncludeAnonymousUsers field if non-nil, zero value otherwise.

### GetIncludeAnonymousUsersOk

`func (o *TraitOptions) GetIncludeAnonymousUsersOk() (*bool, bool)`

GetIncludeAnonymousUsersOk returns a tuple with the IncludeAnonymousUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnonymousUsers

`func (o *TraitOptions) SetIncludeAnonymousUsers(v bool)`

SetIncludeAnonymousUsers sets IncludeAnonymousUsers field to given value.

### HasIncludeAnonymousUsers

`func (o *TraitOptions) HasIncludeAnonymousUsers() bool`

HasIncludeAnonymousUsers returns a boolean if a field has been set.

### GetBackfillDurationDays

`func (o *TraitOptions) GetBackfillDurationDays() float32`

GetBackfillDurationDays returns the BackfillDurationDays field if non-nil, zero value otherwise.

### GetBackfillDurationDaysOk

`func (o *TraitOptions) GetBackfillDurationDaysOk() (*float32, bool)`

GetBackfillDurationDaysOk returns a tuple with the BackfillDurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackfillDurationDays

`func (o *TraitOptions) SetBackfillDurationDays(v float32)`

SetBackfillDurationDays sets BackfillDurationDays field to given value.

### HasBackfillDurationDays

`func (o *TraitOptions) HasBackfillDurationDays() bool`

HasBackfillDurationDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


