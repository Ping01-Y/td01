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
)

// BotsGetBotCommandsRequest represents TL type `bots.getBotCommands#e34c0dd6`.
//
// See https://core.telegram.org/method/bots.getBotCommands for reference.
type BotsGetBotCommandsRequest struct {
	// Scope field of BotsGetBotCommandsRequest.
	Scope BotCommandScopeClass
	// LangCode field of BotsGetBotCommandsRequest.
	LangCode string
}

// BotsGetBotCommandsRequestTypeID is TL type id of BotsGetBotCommandsRequest.
const BotsGetBotCommandsRequestTypeID = 0xe34c0dd6

func (g *BotsGetBotCommandsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Scope == nil) {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *BotsGetBotCommandsRequest) String() string {
	if g == nil {
		return "BotsGetBotCommandsRequest(nil)"
	}
	type Alias BotsGetBotCommandsRequest
	return fmt.Sprintf("BotsGetBotCommandsRequest%+v", Alias(*g))
}

// FillFrom fills BotsGetBotCommandsRequest from given interface.
func (g *BotsGetBotCommandsRequest) FillFrom(from interface {
	GetScope() (value BotCommandScopeClass)
	GetLangCode() (value string)
}) {
	g.Scope = from.GetScope()
	g.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsGetBotCommandsRequest) TypeID() uint32 {
	return BotsGetBotCommandsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsGetBotCommandsRequest) TypeName() string {
	return "bots.getBotCommands"
}

// TypeInfo returns info about TL type.
func (g *BotsGetBotCommandsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.getBotCommands",
		ID:   BotsGetBotCommandsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Scope",
			SchemaName: "scope",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *BotsGetBotCommandsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotCommands#e34c0dd6 as nil")
	}
	b.PutID(BotsGetBotCommandsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *BotsGetBotCommandsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode bots.getBotCommands#e34c0dd6 as nil")
	}
	if g.Scope == nil {
		return fmt.Errorf("unable to encode bots.getBotCommands#e34c0dd6: field scope is nil")
	}
	if err := g.Scope.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.getBotCommands#e34c0dd6: field scope: %w", err)
	}
	b.PutString(g.LangCode)
	return nil
}

// GetScope returns value of Scope field.
func (g *BotsGetBotCommandsRequest) GetScope() (value BotCommandScopeClass) {
	return g.Scope
}

// GetLangCode returns value of LangCode field.
func (g *BotsGetBotCommandsRequest) GetLangCode() (value string) {
	return g.LangCode
}

// Decode implements bin.Decoder.
func (g *BotsGetBotCommandsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotCommands#e34c0dd6 to nil")
	}
	if err := b.ConsumeID(BotsGetBotCommandsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.getBotCommands#e34c0dd6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *BotsGetBotCommandsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode bots.getBotCommands#e34c0dd6 to nil")
	}
	{
		value, err := DecodeBotCommandScope(b)
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotCommands#e34c0dd6: field scope: %w", err)
		}
		g.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.getBotCommands#e34c0dd6: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for BotsGetBotCommandsRequest.
var (
	_ bin.Encoder     = &BotsGetBotCommandsRequest{}
	_ bin.Decoder     = &BotsGetBotCommandsRequest{}
	_ bin.BareEncoder = &BotsGetBotCommandsRequest{}
	_ bin.BareDecoder = &BotsGetBotCommandsRequest{}
)

// BotsGetBotCommands invokes method bots.getBotCommands#e34c0dd6 returning error if any.
//
// See https://core.telegram.org/method/bots.getBotCommands for reference.
func (c *Client) BotsGetBotCommands(ctx context.Context, request *BotsGetBotCommandsRequest) ([]BotCommand, error) {
	var result BotCommandVector

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []BotCommand(result.Elems), nil
}
