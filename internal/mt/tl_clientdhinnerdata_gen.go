// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// ClientDHInnerData represents TL type client_DH_inner_data#6643b654.
type ClientDHInnerData struct {
	// Nonce field of ClientDHInnerData.
	Nonce bin.Int128
	// ServerNonce field of ClientDHInnerData.
	ServerNonce bin.Int128
	// RetryID field of ClientDHInnerData.
	RetryID int64
	// GB field of ClientDHInnerData.
	GB []byte
}

// ClientDHInnerDataTypeID is TL type id of ClientDHInnerData.
const ClientDHInnerDataTypeID = 0x6643b654

// Encode implements bin.Encoder.
func (c *ClientDHInnerData) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode client_DH_inner_data#6643b654 as nil")
	}
	b.PutID(ClientDHInnerDataTypeID)
	b.PutInt128(c.Nonce)
	b.PutInt128(c.ServerNonce)
	b.PutLong(c.RetryID)
	b.PutBytes(c.GB)
	return nil
}

// Decode implements bin.Decoder.
func (c *ClientDHInnerData) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode client_DH_inner_data#6643b654 to nil")
	}
	if err := b.ConsumeID(ClientDHInnerDataTypeID); err != nil {
		return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field nonce: %w", err)
		}
		c.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field server_nonce: %w", err)
		}
		c.ServerNonce = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field retry_id: %w", err)
		}
		c.RetryID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode client_DH_inner_data#6643b654: field g_b: %w", err)
		}
		c.GB = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ClientDHInnerData.
var (
	_ bin.Encoder = &ClientDHInnerData{}
	_ bin.Decoder = &ClientDHInnerData{}
)
