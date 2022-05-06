# CreateFileReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Type** | **string** |  | 
**TemplateID** | Pointer to **string** |  | [optional] 
**FolderID** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateFileReq

`func NewCreateFileReq(title string, type_ string, ) *CreateFileReq`

NewCreateFileReq instantiates a new CreateFileReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileReqWithDefaults

`func NewCreateFileReqWithDefaults() *CreateFileReq`

NewCreateFileReqWithDefaults instantiates a new CreateFileReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateFileReq) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateFileReq) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateFileReq) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *CreateFileReq) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateFileReq) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateFileReq) SetType(v string)`

SetType sets Type field to given value.


### GetTemplateID

`func (o *CreateFileReq) GetTemplateID() string`

GetTemplateID returns the TemplateID field if non-nil, zero value otherwise.

### GetTemplateIDOk

`func (o *CreateFileReq) GetTemplateIDOk() (*string, bool)`

GetTemplateIDOk returns a tuple with the TemplateID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateID

`func (o *CreateFileReq) SetTemplateID(v string)`

SetTemplateID sets TemplateID field to given value.

### HasTemplateID

`func (o *CreateFileReq) HasTemplateID() bool`

HasTemplateID returns a boolean if a field has been set.

### GetFolderID

`func (o *CreateFileReq) GetFolderID() string`

GetFolderID returns the FolderID field if non-nil, zero value otherwise.

### GetFolderIDOk

`func (o *CreateFileReq) GetFolderIDOk() (*string, bool)`

GetFolderIDOk returns a tuple with the FolderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderID

`func (o *CreateFileReq) SetFolderID(v string)`

SetFolderID sets FolderID field to given value.

### HasFolderID

`func (o *CreateFileReq) HasFolderID() bool`

HasFolderID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


