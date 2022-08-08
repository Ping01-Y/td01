// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// DialogFilterClassVector is a box for Vector<DialogFilter>
type DialogFilterClassVector struct {
	// Elements of Vector<DialogFilter>
	Elems []DialogFilterClass
}

// DialogFilterClassVectorTypeID is TL type id of DialogFilterClassVector.
const DialogFilterClassVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for DialogFilterClassVector.
var (
	_ bin.Encoder     = &DialogFilterClassVector{}
	_ bin.Decoder     = &DialogFilterClassVector{}
	_ bin.BareEncoder = &DialogFilterClassVector{}
	_ bin.BareDecoder = &DialogFilterClassVector{}
)

func (vec *DialogFilterClassVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *DialogFilterClassVector) String() string {
	if vec == nil {
		return "DialogFilterClassVector(nil)"
	}
	type Alias DialogFilterClassVector
	return fmt.Sprintf("DialogFilterClassVector%+v", Alias(*vec))
}

// FillFrom fills DialogFilterClassVector from given interface.
func (vec *DialogFilterClassVector) FillFrom(from interface {
	GetElems() (value []DialogFilterClass)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogFilterClassVector) TypeID() uint32 {
	return DialogFilterClassVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogFilterClassVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *DialogFilterClassVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   DialogFilterClassVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *DialogFilterClassVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogFilter> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *DialogFilterClassVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<DialogFilter> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if v == nil {
			return fmt.Errorf("unable to encode Vector<DialogFilter>: field Elems element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<DialogFilter>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *DialogFilterClassVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogFilter> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *DialogFilterClassVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<DialogFilter> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<DialogFilter>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]DialogFilterClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDialogFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode Vector<DialogFilter>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *DialogFilterClassVector) GetElems() (value []DialogFilterClass) {
	if vec == nil {
		return
	}
	return vec.Elems
}

// MapElems returns field Elems wrapped in DialogFilterClassArray helper.
func (vec *DialogFilterClassVector) MapElems() (value DialogFilterClassArray) {
	return DialogFilterClassArray(vec.Elems)
}