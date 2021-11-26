// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// LogVerbosityLevel represents TL type `logVerbosityLevel#676443ea`.
type LogVerbosityLevel struct {
	// Log verbosity level
	VerbosityLevel int32
}

// LogVerbosityLevelTypeID is TL type id of LogVerbosityLevel.
const LogVerbosityLevelTypeID = 0x676443ea

// Ensuring interfaces in compile-time for LogVerbosityLevel.
var (
	_ bin.Encoder     = &LogVerbosityLevel{}
	_ bin.Decoder     = &LogVerbosityLevel{}
	_ bin.BareEncoder = &LogVerbosityLevel{}
	_ bin.BareDecoder = &LogVerbosityLevel{}
)

func (l *LogVerbosityLevel) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.VerbosityLevel == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LogVerbosityLevel) String() string {
	if l == nil {
		return "LogVerbosityLevel(nil)"
	}
	type Alias LogVerbosityLevel
	return fmt.Sprintf("LogVerbosityLevel%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LogVerbosityLevel) TypeID() uint32 {
	return LogVerbosityLevelTypeID
}

// TypeName returns name of type in TL schema.
func (*LogVerbosityLevel) TypeName() string {
	return "logVerbosityLevel"
}

// TypeInfo returns info about TL type.
func (l *LogVerbosityLevel) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "logVerbosityLevel",
		ID:   LogVerbosityLevelTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "VerbosityLevel",
			SchemaName: "verbosity_level",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LogVerbosityLevel) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logVerbosityLevel#676443ea as nil")
	}
	b.PutID(LogVerbosityLevelTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LogVerbosityLevel) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logVerbosityLevel#676443ea as nil")
	}
	b.PutInt32(l.VerbosityLevel)
	return nil
}

// Decode implements bin.Decoder.
func (l *LogVerbosityLevel) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logVerbosityLevel#676443ea to nil")
	}
	if err := b.ConsumeID(LogVerbosityLevelTypeID); err != nil {
		return fmt.Errorf("unable to decode logVerbosityLevel#676443ea: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LogVerbosityLevel) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logVerbosityLevel#676443ea to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode logVerbosityLevel#676443ea: field verbosity_level: %w", err)
		}
		l.VerbosityLevel = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LogVerbosityLevel) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode logVerbosityLevel#676443ea as nil")
	}
	b.ObjStart()
	b.PutID("logVerbosityLevel")
	b.FieldStart("verbosity_level")
	b.PutInt32(l.VerbosityLevel)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LogVerbosityLevel) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode logVerbosityLevel#676443ea to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("logVerbosityLevel"); err != nil {
				return fmt.Errorf("unable to decode logVerbosityLevel#676443ea: %w", err)
			}
		case "verbosity_level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode logVerbosityLevel#676443ea: field verbosity_level: %w", err)
			}
			l.VerbosityLevel = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVerbosityLevel returns value of VerbosityLevel field.
func (l *LogVerbosityLevel) GetVerbosityLevel() (value int32) {
	return l.VerbosityLevel
}