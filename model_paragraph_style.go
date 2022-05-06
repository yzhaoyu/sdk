/*
Tencent Docs SDK

Tencent docs sdk contains DocAPI, SmartsheetAPI, SheetAPI, DriveAPI and OAuthAPI

API version: 0.0.1
Contact: tencentdocs@tencent.com
*/

// Code generated by Tencent Docs (https://docs.qq.com); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ParagraphStyle struct for ParagraphStyle
type ParagraphStyle struct {
	NamedStyleType *string `json:"namedStyleType,omitempty"`
	Alignment *string `json:"alignment,omitempty"`
	LineSpacing *GoogleProtobufDoubleValue `json:"lineSpacing,omitempty"`
	TextDirection *string `json:"textDirection,omitempty"`
	SpaceAbove *GoogleProtobufDoubleValue `json:"spaceAbove,omitempty"`
	SpaceBelow *GoogleProtobufDoubleValue `json:"spaceBelow,omitempty"`
	BorderBetween *ParagraphBorder `json:"borderBetween,omitempty"`
	BorderTop *ParagraphBorder `json:"borderTop,omitempty"`
	BorderBottom *ParagraphBorder `json:"borderBottom,omitempty"`
	BorderLeft *ParagraphBorder `json:"borderLeft,omitempty"`
	BorderRight *ParagraphBorder `json:"borderRight,omitempty"`
	IndentFirstLine *GoogleProtobufDoubleValue `json:"indentFirstLine,omitempty"`
	IndentStart *GoogleProtobufDoubleValue `json:"indentStart,omitempty"`
	IndentEnd *GoogleProtobufDoubleValue `json:"indentEnd,omitempty"`
	TabStops []TabStop `json:"tabStops,omitempty"`
	KeepLinesTogether *GoogleProtobufBoolValue `json:"keepLinesTogether,omitempty"`
	KeepWithNext *GoogleProtobufBoolValue `json:"keepWithNext,omitempty"`
	AvoidWidowAndOrphan *GoogleProtobufBoolValue `json:"avoidWidowAndOrphan,omitempty"`
	Shading *Shading `json:"shading,omitempty"`
}

// NewParagraphStyle instantiates a new ParagraphStyle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParagraphStyle() *ParagraphStyle {
	this := ParagraphStyle{}
	return &this
}

// NewParagraphStyleWithDefaults instantiates a new ParagraphStyle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParagraphStyleWithDefaults() *ParagraphStyle {
	this := ParagraphStyle{}
	return &this
}

// GetNamedStyleType returns the NamedStyleType field value if set, zero value otherwise.
func (o *ParagraphStyle) GetNamedStyleType() string {
	if o == nil || o.NamedStyleType == nil {
		var ret string
		return ret
	}
	return *o.NamedStyleType
}

// GetNamedStyleTypeOk returns a tuple with the NamedStyleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetNamedStyleTypeOk() (*string, bool) {
	if o == nil || o.NamedStyleType == nil {
		return nil, false
	}
	return o.NamedStyleType, true
}

// HasNamedStyleType returns a boolean if a field has been set.
func (o *ParagraphStyle) HasNamedStyleType() bool {
	if o != nil && o.NamedStyleType != nil {
		return true
	}

	return false
}

// SetNamedStyleType gets a reference to the given string and assigns it to the NamedStyleType field.
func (o *ParagraphStyle) SetNamedStyleType(v string) {
	o.NamedStyleType = &v
}

// GetAlignment returns the Alignment field value if set, zero value otherwise.
func (o *ParagraphStyle) GetAlignment() string {
	if o == nil || o.Alignment == nil {
		var ret string
		return ret
	}
	return *o.Alignment
}

// GetAlignmentOk returns a tuple with the Alignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetAlignmentOk() (*string, bool) {
	if o == nil || o.Alignment == nil {
		return nil, false
	}
	return o.Alignment, true
}

// HasAlignment returns a boolean if a field has been set.
func (o *ParagraphStyle) HasAlignment() bool {
	if o != nil && o.Alignment != nil {
		return true
	}

	return false
}

// SetAlignment gets a reference to the given string and assigns it to the Alignment field.
func (o *ParagraphStyle) SetAlignment(v string) {
	o.Alignment = &v
}

// GetLineSpacing returns the LineSpacing field value if set, zero value otherwise.
func (o *ParagraphStyle) GetLineSpacing() GoogleProtobufDoubleValue {
	if o == nil || o.LineSpacing == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.LineSpacing
}

// GetLineSpacingOk returns a tuple with the LineSpacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetLineSpacingOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.LineSpacing == nil {
		return nil, false
	}
	return o.LineSpacing, true
}

// HasLineSpacing returns a boolean if a field has been set.
func (o *ParagraphStyle) HasLineSpacing() bool {
	if o != nil && o.LineSpacing != nil {
		return true
	}

	return false
}

// SetLineSpacing gets a reference to the given GoogleProtobufDoubleValue and assigns it to the LineSpacing field.
func (o *ParagraphStyle) SetLineSpacing(v GoogleProtobufDoubleValue) {
	o.LineSpacing = &v
}

// GetTextDirection returns the TextDirection field value if set, zero value otherwise.
func (o *ParagraphStyle) GetTextDirection() string {
	if o == nil || o.TextDirection == nil {
		var ret string
		return ret
	}
	return *o.TextDirection
}

// GetTextDirectionOk returns a tuple with the TextDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetTextDirectionOk() (*string, bool) {
	if o == nil || o.TextDirection == nil {
		return nil, false
	}
	return o.TextDirection, true
}

// HasTextDirection returns a boolean if a field has been set.
func (o *ParagraphStyle) HasTextDirection() bool {
	if o != nil && o.TextDirection != nil {
		return true
	}

	return false
}

// SetTextDirection gets a reference to the given string and assigns it to the TextDirection field.
func (o *ParagraphStyle) SetTextDirection(v string) {
	o.TextDirection = &v
}

// GetSpaceAbove returns the SpaceAbove field value if set, zero value otherwise.
func (o *ParagraphStyle) GetSpaceAbove() GoogleProtobufDoubleValue {
	if o == nil || o.SpaceAbove == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.SpaceAbove
}

// GetSpaceAboveOk returns a tuple with the SpaceAbove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetSpaceAboveOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.SpaceAbove == nil {
		return nil, false
	}
	return o.SpaceAbove, true
}

// HasSpaceAbove returns a boolean if a field has been set.
func (o *ParagraphStyle) HasSpaceAbove() bool {
	if o != nil && o.SpaceAbove != nil {
		return true
	}

	return false
}

// SetSpaceAbove gets a reference to the given GoogleProtobufDoubleValue and assigns it to the SpaceAbove field.
func (o *ParagraphStyle) SetSpaceAbove(v GoogleProtobufDoubleValue) {
	o.SpaceAbove = &v
}

// GetSpaceBelow returns the SpaceBelow field value if set, zero value otherwise.
func (o *ParagraphStyle) GetSpaceBelow() GoogleProtobufDoubleValue {
	if o == nil || o.SpaceBelow == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.SpaceBelow
}

// GetSpaceBelowOk returns a tuple with the SpaceBelow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetSpaceBelowOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.SpaceBelow == nil {
		return nil, false
	}
	return o.SpaceBelow, true
}

// HasSpaceBelow returns a boolean if a field has been set.
func (o *ParagraphStyle) HasSpaceBelow() bool {
	if o != nil && o.SpaceBelow != nil {
		return true
	}

	return false
}

// SetSpaceBelow gets a reference to the given GoogleProtobufDoubleValue and assigns it to the SpaceBelow field.
func (o *ParagraphStyle) SetSpaceBelow(v GoogleProtobufDoubleValue) {
	o.SpaceBelow = &v
}

// GetBorderBetween returns the BorderBetween field value if set, zero value otherwise.
func (o *ParagraphStyle) GetBorderBetween() ParagraphBorder {
	if o == nil || o.BorderBetween == nil {
		var ret ParagraphBorder
		return ret
	}
	return *o.BorderBetween
}

// GetBorderBetweenOk returns a tuple with the BorderBetween field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetBorderBetweenOk() (*ParagraphBorder, bool) {
	if o == nil || o.BorderBetween == nil {
		return nil, false
	}
	return o.BorderBetween, true
}

// HasBorderBetween returns a boolean if a field has been set.
func (o *ParagraphStyle) HasBorderBetween() bool {
	if o != nil && o.BorderBetween != nil {
		return true
	}

	return false
}

// SetBorderBetween gets a reference to the given ParagraphBorder and assigns it to the BorderBetween field.
func (o *ParagraphStyle) SetBorderBetween(v ParagraphBorder) {
	o.BorderBetween = &v
}

// GetBorderTop returns the BorderTop field value if set, zero value otherwise.
func (o *ParagraphStyle) GetBorderTop() ParagraphBorder {
	if o == nil || o.BorderTop == nil {
		var ret ParagraphBorder
		return ret
	}
	return *o.BorderTop
}

// GetBorderTopOk returns a tuple with the BorderTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetBorderTopOk() (*ParagraphBorder, bool) {
	if o == nil || o.BorderTop == nil {
		return nil, false
	}
	return o.BorderTop, true
}

// HasBorderTop returns a boolean if a field has been set.
func (o *ParagraphStyle) HasBorderTop() bool {
	if o != nil && o.BorderTop != nil {
		return true
	}

	return false
}

// SetBorderTop gets a reference to the given ParagraphBorder and assigns it to the BorderTop field.
func (o *ParagraphStyle) SetBorderTop(v ParagraphBorder) {
	o.BorderTop = &v
}

// GetBorderBottom returns the BorderBottom field value if set, zero value otherwise.
func (o *ParagraphStyle) GetBorderBottom() ParagraphBorder {
	if o == nil || o.BorderBottom == nil {
		var ret ParagraphBorder
		return ret
	}
	return *o.BorderBottom
}

// GetBorderBottomOk returns a tuple with the BorderBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetBorderBottomOk() (*ParagraphBorder, bool) {
	if o == nil || o.BorderBottom == nil {
		return nil, false
	}
	return o.BorderBottom, true
}

// HasBorderBottom returns a boolean if a field has been set.
func (o *ParagraphStyle) HasBorderBottom() bool {
	if o != nil && o.BorderBottom != nil {
		return true
	}

	return false
}

// SetBorderBottom gets a reference to the given ParagraphBorder and assigns it to the BorderBottom field.
func (o *ParagraphStyle) SetBorderBottom(v ParagraphBorder) {
	o.BorderBottom = &v
}

// GetBorderLeft returns the BorderLeft field value if set, zero value otherwise.
func (o *ParagraphStyle) GetBorderLeft() ParagraphBorder {
	if o == nil || o.BorderLeft == nil {
		var ret ParagraphBorder
		return ret
	}
	return *o.BorderLeft
}

// GetBorderLeftOk returns a tuple with the BorderLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetBorderLeftOk() (*ParagraphBorder, bool) {
	if o == nil || o.BorderLeft == nil {
		return nil, false
	}
	return o.BorderLeft, true
}

// HasBorderLeft returns a boolean if a field has been set.
func (o *ParagraphStyle) HasBorderLeft() bool {
	if o != nil && o.BorderLeft != nil {
		return true
	}

	return false
}

// SetBorderLeft gets a reference to the given ParagraphBorder and assigns it to the BorderLeft field.
func (o *ParagraphStyle) SetBorderLeft(v ParagraphBorder) {
	o.BorderLeft = &v
}

// GetBorderRight returns the BorderRight field value if set, zero value otherwise.
func (o *ParagraphStyle) GetBorderRight() ParagraphBorder {
	if o == nil || o.BorderRight == nil {
		var ret ParagraphBorder
		return ret
	}
	return *o.BorderRight
}

// GetBorderRightOk returns a tuple with the BorderRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetBorderRightOk() (*ParagraphBorder, bool) {
	if o == nil || o.BorderRight == nil {
		return nil, false
	}
	return o.BorderRight, true
}

// HasBorderRight returns a boolean if a field has been set.
func (o *ParagraphStyle) HasBorderRight() bool {
	if o != nil && o.BorderRight != nil {
		return true
	}

	return false
}

// SetBorderRight gets a reference to the given ParagraphBorder and assigns it to the BorderRight field.
func (o *ParagraphStyle) SetBorderRight(v ParagraphBorder) {
	o.BorderRight = &v
}

// GetIndentFirstLine returns the IndentFirstLine field value if set, zero value otherwise.
func (o *ParagraphStyle) GetIndentFirstLine() GoogleProtobufDoubleValue {
	if o == nil || o.IndentFirstLine == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.IndentFirstLine
}

// GetIndentFirstLineOk returns a tuple with the IndentFirstLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetIndentFirstLineOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.IndentFirstLine == nil {
		return nil, false
	}
	return o.IndentFirstLine, true
}

// HasIndentFirstLine returns a boolean if a field has been set.
func (o *ParagraphStyle) HasIndentFirstLine() bool {
	if o != nil && o.IndentFirstLine != nil {
		return true
	}

	return false
}

// SetIndentFirstLine gets a reference to the given GoogleProtobufDoubleValue and assigns it to the IndentFirstLine field.
func (o *ParagraphStyle) SetIndentFirstLine(v GoogleProtobufDoubleValue) {
	o.IndentFirstLine = &v
}

// GetIndentStart returns the IndentStart field value if set, zero value otherwise.
func (o *ParagraphStyle) GetIndentStart() GoogleProtobufDoubleValue {
	if o == nil || o.IndentStart == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.IndentStart
}

// GetIndentStartOk returns a tuple with the IndentStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetIndentStartOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.IndentStart == nil {
		return nil, false
	}
	return o.IndentStart, true
}

// HasIndentStart returns a boolean if a field has been set.
func (o *ParagraphStyle) HasIndentStart() bool {
	if o != nil && o.IndentStart != nil {
		return true
	}

	return false
}

// SetIndentStart gets a reference to the given GoogleProtobufDoubleValue and assigns it to the IndentStart field.
func (o *ParagraphStyle) SetIndentStart(v GoogleProtobufDoubleValue) {
	o.IndentStart = &v
}

// GetIndentEnd returns the IndentEnd field value if set, zero value otherwise.
func (o *ParagraphStyle) GetIndentEnd() GoogleProtobufDoubleValue {
	if o == nil || o.IndentEnd == nil {
		var ret GoogleProtobufDoubleValue
		return ret
	}
	return *o.IndentEnd
}

// GetIndentEndOk returns a tuple with the IndentEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetIndentEndOk() (*GoogleProtobufDoubleValue, bool) {
	if o == nil || o.IndentEnd == nil {
		return nil, false
	}
	return o.IndentEnd, true
}

// HasIndentEnd returns a boolean if a field has been set.
func (o *ParagraphStyle) HasIndentEnd() bool {
	if o != nil && o.IndentEnd != nil {
		return true
	}

	return false
}

// SetIndentEnd gets a reference to the given GoogleProtobufDoubleValue and assigns it to the IndentEnd field.
func (o *ParagraphStyle) SetIndentEnd(v GoogleProtobufDoubleValue) {
	o.IndentEnd = &v
}

// GetTabStops returns the TabStops field value if set, zero value otherwise.
func (o *ParagraphStyle) GetTabStops() []TabStop {
	if o == nil || o.TabStops == nil {
		var ret []TabStop
		return ret
	}
	return o.TabStops
}

// GetTabStopsOk returns a tuple with the TabStops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetTabStopsOk() ([]TabStop, bool) {
	if o == nil || o.TabStops == nil {
		return nil, false
	}
	return o.TabStops, true
}

// HasTabStops returns a boolean if a field has been set.
func (o *ParagraphStyle) HasTabStops() bool {
	if o != nil && o.TabStops != nil {
		return true
	}

	return false
}

// SetTabStops gets a reference to the given []TabStop and assigns it to the TabStops field.
func (o *ParagraphStyle) SetTabStops(v []TabStop) {
	o.TabStops = v
}

// GetKeepLinesTogether returns the KeepLinesTogether field value if set, zero value otherwise.
func (o *ParagraphStyle) GetKeepLinesTogether() GoogleProtobufBoolValue {
	if o == nil || o.KeepLinesTogether == nil {
		var ret GoogleProtobufBoolValue
		return ret
	}
	return *o.KeepLinesTogether
}

// GetKeepLinesTogetherOk returns a tuple with the KeepLinesTogether field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetKeepLinesTogetherOk() (*GoogleProtobufBoolValue, bool) {
	if o == nil || o.KeepLinesTogether == nil {
		return nil, false
	}
	return o.KeepLinesTogether, true
}

// HasKeepLinesTogether returns a boolean if a field has been set.
func (o *ParagraphStyle) HasKeepLinesTogether() bool {
	if o != nil && o.KeepLinesTogether != nil {
		return true
	}

	return false
}

// SetKeepLinesTogether gets a reference to the given GoogleProtobufBoolValue and assigns it to the KeepLinesTogether field.
func (o *ParagraphStyle) SetKeepLinesTogether(v GoogleProtobufBoolValue) {
	o.KeepLinesTogether = &v
}

// GetKeepWithNext returns the KeepWithNext field value if set, zero value otherwise.
func (o *ParagraphStyle) GetKeepWithNext() GoogleProtobufBoolValue {
	if o == nil || o.KeepWithNext == nil {
		var ret GoogleProtobufBoolValue
		return ret
	}
	return *o.KeepWithNext
}

// GetKeepWithNextOk returns a tuple with the KeepWithNext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetKeepWithNextOk() (*GoogleProtobufBoolValue, bool) {
	if o == nil || o.KeepWithNext == nil {
		return nil, false
	}
	return o.KeepWithNext, true
}

// HasKeepWithNext returns a boolean if a field has been set.
func (o *ParagraphStyle) HasKeepWithNext() bool {
	if o != nil && o.KeepWithNext != nil {
		return true
	}

	return false
}

// SetKeepWithNext gets a reference to the given GoogleProtobufBoolValue and assigns it to the KeepWithNext field.
func (o *ParagraphStyle) SetKeepWithNext(v GoogleProtobufBoolValue) {
	o.KeepWithNext = &v
}

// GetAvoidWidowAndOrphan returns the AvoidWidowAndOrphan field value if set, zero value otherwise.
func (o *ParagraphStyle) GetAvoidWidowAndOrphan() GoogleProtobufBoolValue {
	if o == nil || o.AvoidWidowAndOrphan == nil {
		var ret GoogleProtobufBoolValue
		return ret
	}
	return *o.AvoidWidowAndOrphan
}

// GetAvoidWidowAndOrphanOk returns a tuple with the AvoidWidowAndOrphan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetAvoidWidowAndOrphanOk() (*GoogleProtobufBoolValue, bool) {
	if o == nil || o.AvoidWidowAndOrphan == nil {
		return nil, false
	}
	return o.AvoidWidowAndOrphan, true
}

// HasAvoidWidowAndOrphan returns a boolean if a field has been set.
func (o *ParagraphStyle) HasAvoidWidowAndOrphan() bool {
	if o != nil && o.AvoidWidowAndOrphan != nil {
		return true
	}

	return false
}

// SetAvoidWidowAndOrphan gets a reference to the given GoogleProtobufBoolValue and assigns it to the AvoidWidowAndOrphan field.
func (o *ParagraphStyle) SetAvoidWidowAndOrphan(v GoogleProtobufBoolValue) {
	o.AvoidWidowAndOrphan = &v
}

// GetShading returns the Shading field value if set, zero value otherwise.
func (o *ParagraphStyle) GetShading() Shading {
	if o == nil || o.Shading == nil {
		var ret Shading
		return ret
	}
	return *o.Shading
}

// GetShadingOk returns a tuple with the Shading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParagraphStyle) GetShadingOk() (*Shading, bool) {
	if o == nil || o.Shading == nil {
		return nil, false
	}
	return o.Shading, true
}

// HasShading returns a boolean if a field has been set.
func (o *ParagraphStyle) HasShading() bool {
	if o != nil && o.Shading != nil {
		return true
	}

	return false
}

// SetShading gets a reference to the given Shading and assigns it to the Shading field.
func (o *ParagraphStyle) SetShading(v Shading) {
	o.Shading = &v
}

func (o ParagraphStyle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NamedStyleType != nil {
		toSerialize["namedStyleType"] = o.NamedStyleType
	}
	if o.Alignment != nil {
		toSerialize["alignment"] = o.Alignment
	}
	if o.LineSpacing != nil {
		toSerialize["lineSpacing"] = o.LineSpacing
	}
	if o.TextDirection != nil {
		toSerialize["textDirection"] = o.TextDirection
	}
	if o.SpaceAbove != nil {
		toSerialize["spaceAbove"] = o.SpaceAbove
	}
	if o.SpaceBelow != nil {
		toSerialize["spaceBelow"] = o.SpaceBelow
	}
	if o.BorderBetween != nil {
		toSerialize["borderBetween"] = o.BorderBetween
	}
	if o.BorderTop != nil {
		toSerialize["borderTop"] = o.BorderTop
	}
	if o.BorderBottom != nil {
		toSerialize["borderBottom"] = o.BorderBottom
	}
	if o.BorderLeft != nil {
		toSerialize["borderLeft"] = o.BorderLeft
	}
	if o.BorderRight != nil {
		toSerialize["borderRight"] = o.BorderRight
	}
	if o.IndentFirstLine != nil {
		toSerialize["indentFirstLine"] = o.IndentFirstLine
	}
	if o.IndentStart != nil {
		toSerialize["indentStart"] = o.IndentStart
	}
	if o.IndentEnd != nil {
		toSerialize["indentEnd"] = o.IndentEnd
	}
	if o.TabStops != nil {
		toSerialize["tabStops"] = o.TabStops
	}
	if o.KeepLinesTogether != nil {
		toSerialize["keepLinesTogether"] = o.KeepLinesTogether
	}
	if o.KeepWithNext != nil {
		toSerialize["keepWithNext"] = o.KeepWithNext
	}
	if o.AvoidWidowAndOrphan != nil {
		toSerialize["avoidWidowAndOrphan"] = o.AvoidWidowAndOrphan
	}
	if o.Shading != nil {
		toSerialize["shading"] = o.Shading
	}
	return json.Marshal(toSerialize)
}

type NullableParagraphStyle struct {
	value *ParagraphStyle
	isSet bool
}

func (v NullableParagraphStyle) Get() *ParagraphStyle {
	return v.value
}

func (v *NullableParagraphStyle) Set(val *ParagraphStyle) {
	v.value = val
	v.isSet = true
}

func (v NullableParagraphStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableParagraphStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParagraphStyle(val *ParagraphStyle) *NullableParagraphStyle {
	return &NullableParagraphStyle{value: val, isSet: true}
}

func (v NullableParagraphStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParagraphStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


