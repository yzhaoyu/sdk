# ParagraphStyle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamedStyleType** | Pointer to **string** |  | [optional] 
**Alignment** | Pointer to **string** |  | [optional] 
**LineSpacing** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**TextDirection** | Pointer to **string** |  | [optional] 
**SpaceAbove** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**SpaceBelow** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**BorderBetween** | Pointer to [**ParagraphBorder**](ParagraphBorder.md) |  | [optional] 
**BorderTop** | Pointer to [**ParagraphBorder**](ParagraphBorder.md) |  | [optional] 
**BorderBottom** | Pointer to [**ParagraphBorder**](ParagraphBorder.md) |  | [optional] 
**BorderLeft** | Pointer to [**ParagraphBorder**](ParagraphBorder.md) |  | [optional] 
**BorderRight** | Pointer to [**ParagraphBorder**](ParagraphBorder.md) |  | [optional] 
**IndentFirstLine** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**IndentStart** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**IndentEnd** | Pointer to [**GoogleProtobufDoubleValue**](GoogleProtobufDoubleValue.md) |  | [optional] 
**TabStops** | Pointer to [**[]TabStop**](TabStop.md) |  | [optional] 
**KeepLinesTogether** | Pointer to [**GoogleProtobufBoolValue**](GoogleProtobufBoolValue.md) |  | [optional] 
**KeepWithNext** | Pointer to [**GoogleProtobufBoolValue**](GoogleProtobufBoolValue.md) |  | [optional] 
**AvoidWidowAndOrphan** | Pointer to [**GoogleProtobufBoolValue**](GoogleProtobufBoolValue.md) |  | [optional] 
**Shading** | Pointer to [**Shading**](Shading.md) |  | [optional] 

## Methods

### NewParagraphStyle

`func NewParagraphStyle() *ParagraphStyle`

NewParagraphStyle instantiates a new ParagraphStyle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParagraphStyleWithDefaults

`func NewParagraphStyleWithDefaults() *ParagraphStyle`

NewParagraphStyleWithDefaults instantiates a new ParagraphStyle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamedStyleType

`func (o *ParagraphStyle) GetNamedStyleType() string`

GetNamedStyleType returns the NamedStyleType field if non-nil, zero value otherwise.

### GetNamedStyleTypeOk

`func (o *ParagraphStyle) GetNamedStyleTypeOk() (*string, bool)`

GetNamedStyleTypeOk returns a tuple with the NamedStyleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamedStyleType

`func (o *ParagraphStyle) SetNamedStyleType(v string)`

SetNamedStyleType sets NamedStyleType field to given value.

### HasNamedStyleType

`func (o *ParagraphStyle) HasNamedStyleType() bool`

HasNamedStyleType returns a boolean if a field has been set.

### GetAlignment

`func (o *ParagraphStyle) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *ParagraphStyle) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *ParagraphStyle) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *ParagraphStyle) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetLineSpacing

`func (o *ParagraphStyle) GetLineSpacing() GoogleProtobufDoubleValue`

GetLineSpacing returns the LineSpacing field if non-nil, zero value otherwise.

### GetLineSpacingOk

`func (o *ParagraphStyle) GetLineSpacingOk() (*GoogleProtobufDoubleValue, bool)`

GetLineSpacingOk returns a tuple with the LineSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineSpacing

`func (o *ParagraphStyle) SetLineSpacing(v GoogleProtobufDoubleValue)`

SetLineSpacing sets LineSpacing field to given value.

### HasLineSpacing

`func (o *ParagraphStyle) HasLineSpacing() bool`

HasLineSpacing returns a boolean if a field has been set.

### GetTextDirection

`func (o *ParagraphStyle) GetTextDirection() string`

GetTextDirection returns the TextDirection field if non-nil, zero value otherwise.

### GetTextDirectionOk

`func (o *ParagraphStyle) GetTextDirectionOk() (*string, bool)`

GetTextDirectionOk returns a tuple with the TextDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDirection

`func (o *ParagraphStyle) SetTextDirection(v string)`

SetTextDirection sets TextDirection field to given value.

### HasTextDirection

`func (o *ParagraphStyle) HasTextDirection() bool`

HasTextDirection returns a boolean if a field has been set.

### GetSpaceAbove

`func (o *ParagraphStyle) GetSpaceAbove() GoogleProtobufDoubleValue`

GetSpaceAbove returns the SpaceAbove field if non-nil, zero value otherwise.

### GetSpaceAboveOk

`func (o *ParagraphStyle) GetSpaceAboveOk() (*GoogleProtobufDoubleValue, bool)`

GetSpaceAboveOk returns a tuple with the SpaceAbove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAbove

`func (o *ParagraphStyle) SetSpaceAbove(v GoogleProtobufDoubleValue)`

SetSpaceAbove sets SpaceAbove field to given value.

### HasSpaceAbove

`func (o *ParagraphStyle) HasSpaceAbove() bool`

HasSpaceAbove returns a boolean if a field has been set.

### GetSpaceBelow

`func (o *ParagraphStyle) GetSpaceBelow() GoogleProtobufDoubleValue`

GetSpaceBelow returns the SpaceBelow field if non-nil, zero value otherwise.

### GetSpaceBelowOk

`func (o *ParagraphStyle) GetSpaceBelowOk() (*GoogleProtobufDoubleValue, bool)`

GetSpaceBelowOk returns a tuple with the SpaceBelow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceBelow

`func (o *ParagraphStyle) SetSpaceBelow(v GoogleProtobufDoubleValue)`

SetSpaceBelow sets SpaceBelow field to given value.

### HasSpaceBelow

`func (o *ParagraphStyle) HasSpaceBelow() bool`

HasSpaceBelow returns a boolean if a field has been set.

### GetBorderBetween

`func (o *ParagraphStyle) GetBorderBetween() ParagraphBorder`

GetBorderBetween returns the BorderBetween field if non-nil, zero value otherwise.

### GetBorderBetweenOk

`func (o *ParagraphStyle) GetBorderBetweenOk() (*ParagraphBorder, bool)`

GetBorderBetweenOk returns a tuple with the BorderBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderBetween

`func (o *ParagraphStyle) SetBorderBetween(v ParagraphBorder)`

SetBorderBetween sets BorderBetween field to given value.

### HasBorderBetween

`func (o *ParagraphStyle) HasBorderBetween() bool`

HasBorderBetween returns a boolean if a field has been set.

### GetBorderTop

`func (o *ParagraphStyle) GetBorderTop() ParagraphBorder`

GetBorderTop returns the BorderTop field if non-nil, zero value otherwise.

### GetBorderTopOk

`func (o *ParagraphStyle) GetBorderTopOk() (*ParagraphBorder, bool)`

GetBorderTopOk returns a tuple with the BorderTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderTop

`func (o *ParagraphStyle) SetBorderTop(v ParagraphBorder)`

SetBorderTop sets BorderTop field to given value.

### HasBorderTop

`func (o *ParagraphStyle) HasBorderTop() bool`

HasBorderTop returns a boolean if a field has been set.

### GetBorderBottom

`func (o *ParagraphStyle) GetBorderBottom() ParagraphBorder`

GetBorderBottom returns the BorderBottom field if non-nil, zero value otherwise.

### GetBorderBottomOk

`func (o *ParagraphStyle) GetBorderBottomOk() (*ParagraphBorder, bool)`

GetBorderBottomOk returns a tuple with the BorderBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderBottom

`func (o *ParagraphStyle) SetBorderBottom(v ParagraphBorder)`

SetBorderBottom sets BorderBottom field to given value.

### HasBorderBottom

`func (o *ParagraphStyle) HasBorderBottom() bool`

HasBorderBottom returns a boolean if a field has been set.

### GetBorderLeft

`func (o *ParagraphStyle) GetBorderLeft() ParagraphBorder`

GetBorderLeft returns the BorderLeft field if non-nil, zero value otherwise.

### GetBorderLeftOk

`func (o *ParagraphStyle) GetBorderLeftOk() (*ParagraphBorder, bool)`

GetBorderLeftOk returns a tuple with the BorderLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderLeft

`func (o *ParagraphStyle) SetBorderLeft(v ParagraphBorder)`

SetBorderLeft sets BorderLeft field to given value.

### HasBorderLeft

`func (o *ParagraphStyle) HasBorderLeft() bool`

HasBorderLeft returns a boolean if a field has been set.

### GetBorderRight

`func (o *ParagraphStyle) GetBorderRight() ParagraphBorder`

GetBorderRight returns the BorderRight field if non-nil, zero value otherwise.

### GetBorderRightOk

`func (o *ParagraphStyle) GetBorderRightOk() (*ParagraphBorder, bool)`

GetBorderRightOk returns a tuple with the BorderRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderRight

`func (o *ParagraphStyle) SetBorderRight(v ParagraphBorder)`

SetBorderRight sets BorderRight field to given value.

### HasBorderRight

`func (o *ParagraphStyle) HasBorderRight() bool`

HasBorderRight returns a boolean if a field has been set.

### GetIndentFirstLine

`func (o *ParagraphStyle) GetIndentFirstLine() GoogleProtobufDoubleValue`

GetIndentFirstLine returns the IndentFirstLine field if non-nil, zero value otherwise.

### GetIndentFirstLineOk

`func (o *ParagraphStyle) GetIndentFirstLineOk() (*GoogleProtobufDoubleValue, bool)`

GetIndentFirstLineOk returns a tuple with the IndentFirstLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentFirstLine

`func (o *ParagraphStyle) SetIndentFirstLine(v GoogleProtobufDoubleValue)`

SetIndentFirstLine sets IndentFirstLine field to given value.

### HasIndentFirstLine

`func (o *ParagraphStyle) HasIndentFirstLine() bool`

HasIndentFirstLine returns a boolean if a field has been set.

### GetIndentStart

`func (o *ParagraphStyle) GetIndentStart() GoogleProtobufDoubleValue`

GetIndentStart returns the IndentStart field if non-nil, zero value otherwise.

### GetIndentStartOk

`func (o *ParagraphStyle) GetIndentStartOk() (*GoogleProtobufDoubleValue, bool)`

GetIndentStartOk returns a tuple with the IndentStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentStart

`func (o *ParagraphStyle) SetIndentStart(v GoogleProtobufDoubleValue)`

SetIndentStart sets IndentStart field to given value.

### HasIndentStart

`func (o *ParagraphStyle) HasIndentStart() bool`

HasIndentStart returns a boolean if a field has been set.

### GetIndentEnd

`func (o *ParagraphStyle) GetIndentEnd() GoogleProtobufDoubleValue`

GetIndentEnd returns the IndentEnd field if non-nil, zero value otherwise.

### GetIndentEndOk

`func (o *ParagraphStyle) GetIndentEndOk() (*GoogleProtobufDoubleValue, bool)`

GetIndentEndOk returns a tuple with the IndentEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndentEnd

`func (o *ParagraphStyle) SetIndentEnd(v GoogleProtobufDoubleValue)`

SetIndentEnd sets IndentEnd field to given value.

### HasIndentEnd

`func (o *ParagraphStyle) HasIndentEnd() bool`

HasIndentEnd returns a boolean if a field has been set.

### GetTabStops

`func (o *ParagraphStyle) GetTabStops() []TabStop`

GetTabStops returns the TabStops field if non-nil, zero value otherwise.

### GetTabStopsOk

`func (o *ParagraphStyle) GetTabStopsOk() (*[]TabStop, bool)`

GetTabStopsOk returns a tuple with the TabStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabStops

`func (o *ParagraphStyle) SetTabStops(v []TabStop)`

SetTabStops sets TabStops field to given value.

### HasTabStops

`func (o *ParagraphStyle) HasTabStops() bool`

HasTabStops returns a boolean if a field has been set.

### GetKeepLinesTogether

`func (o *ParagraphStyle) GetKeepLinesTogether() GoogleProtobufBoolValue`

GetKeepLinesTogether returns the KeepLinesTogether field if non-nil, zero value otherwise.

### GetKeepLinesTogetherOk

`func (o *ParagraphStyle) GetKeepLinesTogetherOk() (*GoogleProtobufBoolValue, bool)`

GetKeepLinesTogetherOk returns a tuple with the KeepLinesTogether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLinesTogether

`func (o *ParagraphStyle) SetKeepLinesTogether(v GoogleProtobufBoolValue)`

SetKeepLinesTogether sets KeepLinesTogether field to given value.

### HasKeepLinesTogether

`func (o *ParagraphStyle) HasKeepLinesTogether() bool`

HasKeepLinesTogether returns a boolean if a field has been set.

### GetKeepWithNext

`func (o *ParagraphStyle) GetKeepWithNext() GoogleProtobufBoolValue`

GetKeepWithNext returns the KeepWithNext field if non-nil, zero value otherwise.

### GetKeepWithNextOk

`func (o *ParagraphStyle) GetKeepWithNextOk() (*GoogleProtobufBoolValue, bool)`

GetKeepWithNextOk returns a tuple with the KeepWithNext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepWithNext

`func (o *ParagraphStyle) SetKeepWithNext(v GoogleProtobufBoolValue)`

SetKeepWithNext sets KeepWithNext field to given value.

### HasKeepWithNext

`func (o *ParagraphStyle) HasKeepWithNext() bool`

HasKeepWithNext returns a boolean if a field has been set.

### GetAvoidWidowAndOrphan

`func (o *ParagraphStyle) GetAvoidWidowAndOrphan() GoogleProtobufBoolValue`

GetAvoidWidowAndOrphan returns the AvoidWidowAndOrphan field if non-nil, zero value otherwise.

### GetAvoidWidowAndOrphanOk

`func (o *ParagraphStyle) GetAvoidWidowAndOrphanOk() (*GoogleProtobufBoolValue, bool)`

GetAvoidWidowAndOrphanOk returns a tuple with the AvoidWidowAndOrphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidWidowAndOrphan

`func (o *ParagraphStyle) SetAvoidWidowAndOrphan(v GoogleProtobufBoolValue)`

SetAvoidWidowAndOrphan sets AvoidWidowAndOrphan field to given value.

### HasAvoidWidowAndOrphan

`func (o *ParagraphStyle) HasAvoidWidowAndOrphan() bool`

HasAvoidWidowAndOrphan returns a boolean if a field has been set.

### GetShading

`func (o *ParagraphStyle) GetShading() Shading`

GetShading returns the Shading field if non-nil, zero value otherwise.

### GetShadingOk

`func (o *ParagraphStyle) GetShadingOk() (*Shading, bool)`

GetShadingOk returns a tuple with the Shading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShading

`func (o *ParagraphStyle) SetShading(v Shading)`

SetShading sets Shading field to given value.

### HasShading

`func (o *ParagraphStyle) HasShading() bool`

HasShading returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


