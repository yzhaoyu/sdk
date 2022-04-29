# GetViewsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Next** | Pointer to **int32** |  | [optional] 
**Views** | Pointer to [**[]ViewResource**](ViewResource.md) |  | [optional] 

## Methods

### NewGetViewsResponse

`func NewGetViewsResponse() *GetViewsResponse`

NewGetViewsResponse instantiates a new GetViewsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetViewsResponseWithDefaults

`func NewGetViewsResponseWithDefaults() *GetViewsResponse`

NewGetViewsResponseWithDefaults instantiates a new GetViewsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetViewsResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetViewsResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetViewsResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetViewsResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHasMore

`func (o *GetViewsResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetViewsResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetViewsResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetViewsResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNext

`func (o *GetViewsResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetViewsResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetViewsResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetViewsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetViews

`func (o *GetViewsResponse) GetViews() []ViewResource`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetViewsResponse) GetViewsOk() (*[]ViewResource, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetViewsResponse) SetViews(v []ViewResource)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetViewsResponse) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


