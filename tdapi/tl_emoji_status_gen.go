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

// EmojiStatus represents TL type `emojiStatus#4118a266`.
type EmojiStatus struct {
	// Identifier of the custom emoji in stickerFormatTgs format. If the custom emoji belongs
	// to the sticker set GetOption("themed_emoji_statuses_sticker_set_id"), then it's color
	// must be changed to the color of the Telegram Premium badge
	CustomEmojiID int64
}

// EmojiStatusTypeID is TL type id of EmojiStatus.
const EmojiStatusTypeID = 0x4118a266

// Ensuring interfaces in compile-time for EmojiStatus.
var (
	_ bin.Encoder     = &EmojiStatus{}
	_ bin.Decoder     = &EmojiStatus{}
	_ bin.BareEncoder = &EmojiStatus{}
	_ bin.BareDecoder = &EmojiStatus{}
)

func (e *EmojiStatus) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.CustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiStatus) String() string {
	if e == nil {
		return "EmojiStatus(nil)"
	}
	type Alias EmojiStatus
	return fmt.Sprintf("EmojiStatus%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiStatus) TypeID() uint32 {
	return EmojiStatusTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiStatus) TypeName() string {
	return "emojiStatus"
}

// TypeInfo returns info about TL type.
func (e *EmojiStatus) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiStatus",
		ID:   EmojiStatusTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomEmojiID",
			SchemaName: "custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiStatus) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatus#4118a266 as nil")
	}
	b.PutID(EmojiStatusTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiStatus) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatus#4118a266 as nil")
	}
	b.PutLong(e.CustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiStatus) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatus#4118a266 to nil")
	}
	if err := b.ConsumeID(EmojiStatusTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiStatus#4118a266: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiStatus) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatus#4118a266 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode emojiStatus#4118a266: field custom_emoji_id: %w", err)
		}
		e.CustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmojiStatus) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiStatus#4118a266 as nil")
	}
	b.ObjStart()
	b.PutID("emojiStatus")
	b.Comma()
	b.FieldStart("custom_emoji_id")
	b.PutLong(e.CustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmojiStatus) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiStatus#4118a266 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emojiStatus"); err != nil {
				return fmt.Errorf("unable to decode emojiStatus#4118a266: %w", err)
			}
		case "custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode emojiStatus#4118a266: field custom_emoji_id: %w", err)
			}
			e.CustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomEmojiID returns value of CustomEmojiID field.
func (e *EmojiStatus) GetCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.CustomEmojiID
}
