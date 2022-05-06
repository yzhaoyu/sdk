# SmartsheetentityGetViewsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Views** | Pointer to [**[]SmartsheetentityView**](SmartsheetentityView.md) |  | [optional] 

## Methods

### NewSmartsheetentityGetViewsResponse

`func NewSmartsheetentityGetViewsResponse() *SmartsheetentityGetViewsResponse`

NewSmartsheetentityGetViewsResponse instantiates a new SmartsheetentityGetViewsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartsheetentityGetViewsResponseWithDefaults

`func NewSmartsheetentityGetViewsResponseWithDefaults() *SmartsheetentityGetViewsResponse`

NewSmartsheetentityGetViewsResponseWithDefaults instantiates a new SmartsheetentityGetViewsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SmartsheetentityGetViewsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SmartsheetentityGetViewsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SmartsheetentityGetViewsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SmartsheetentityGetViewsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *SmartsheetentityGetViewsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *SmartsheetentityGetViewsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *SmartsheetentityGetViewsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *SmartsheetentityGetViewsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *SmartsheetentityGetViewsResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SmartsheetentityGetViewsResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SmartsheetentityGetViewsResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *SmartsheetentityGetViewsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetViews

`func (o *SmartsheetentityGetViewsResponse) GetViews() []SmartsheetentityView`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *SmartsheetentityGetViewsResponse) GetViewsOk() (*[]SmartsheetentityView, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *SmartsheetentityGetViewsResponse) SetViews(v []SmartsheetentityView)`

SetViews sets Views field to given value.

### HasViews

`func (o *SmartsheetentityGetViewsResponse) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


