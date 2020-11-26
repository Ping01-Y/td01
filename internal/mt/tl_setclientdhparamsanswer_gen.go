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

// DhGenOk represents TL type dh_gen_ok#3bcbf734.
type DhGenOk struct {
	// Nonce field of DhGenOk.
	Nonce bin.Int128
	// ServerNonce field of DhGenOk.
	ServerNonce bin.Int128
	// NewNonceHash1 field of DhGenOk.
	NewNonceHash1 bin.Int128
}

// DhGenOkTypeID is TL type id of DhGenOk.
const DhGenOkTypeID = 0x3bcbf734

// Encode implements bin.Encoder.
func (d *DhGenOk) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_ok#3bcbf734 as nil")
	}
	b.PutID(DhGenOkTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash1)
	return nil
}

// Decode implements bin.Decoder.
func (d *DhGenOk) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_ok#3bcbf734 to nil")
	}
	if err := b.ConsumeID(DhGenOkTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_ok#3bcbf734: field new_nonce_hash1: %w", err)
		}
		d.NewNonceHash1 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswer.
func (d DhGenOk) construct() SetClientDHParamsAnswer { return &d }

// Ensuring interfaces in compile-time for DhGenOk.
var (
	_ bin.Encoder = &DhGenOk{}
	_ bin.Decoder = &DhGenOk{}

	_ SetClientDHParamsAnswer = &DhGenOk{}
)

// DhGenRetry represents TL type dh_gen_retry#46dc1fb9.
type DhGenRetry struct {
	// Nonce field of DhGenRetry.
	Nonce bin.Int128
	// ServerNonce field of DhGenRetry.
	ServerNonce bin.Int128
	// NewNonceHash2 field of DhGenRetry.
	NewNonceHash2 bin.Int128
}

// DhGenRetryTypeID is TL type id of DhGenRetry.
const DhGenRetryTypeID = 0x46dc1fb9

// Encode implements bin.Encoder.
func (d *DhGenRetry) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_retry#46dc1fb9 as nil")
	}
	b.PutID(DhGenRetryTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash2)
	return nil
}

// Decode implements bin.Decoder.
func (d *DhGenRetry) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_retry#46dc1fb9 to nil")
	}
	if err := b.ConsumeID(DhGenRetryTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_retry#46dc1fb9: field new_nonce_hash2: %w", err)
		}
		d.NewNonceHash2 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswer.
func (d DhGenRetry) construct() SetClientDHParamsAnswer { return &d }

// Ensuring interfaces in compile-time for DhGenRetry.
var (
	_ bin.Encoder = &DhGenRetry{}
	_ bin.Decoder = &DhGenRetry{}

	_ SetClientDHParamsAnswer = &DhGenRetry{}
)

// DhGenFail represents TL type dh_gen_fail#a69dae02.
type DhGenFail struct {
	// Nonce field of DhGenFail.
	Nonce bin.Int128
	// ServerNonce field of DhGenFail.
	ServerNonce bin.Int128
	// NewNonceHash3 field of DhGenFail.
	NewNonceHash3 bin.Int128
}

// DhGenFailTypeID is TL type id of DhGenFail.
const DhGenFailTypeID = 0xa69dae02

// Encode implements bin.Encoder.
func (d *DhGenFail) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dh_gen_fail#a69dae02 as nil")
	}
	b.PutID(DhGenFailTypeID)
	b.PutInt128(d.Nonce)
	b.PutInt128(d.ServerNonce)
	b.PutInt128(d.NewNonceHash3)
	return nil
}

// Decode implements bin.Decoder.
func (d *DhGenFail) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dh_gen_fail#a69dae02 to nil")
	}
	if err := b.ConsumeID(DhGenFailTypeID); err != nil {
		return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field nonce: %w", err)
		}
		d.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field server_nonce: %w", err)
		}
		d.ServerNonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode dh_gen_fail#a69dae02: field new_nonce_hash3: %w", err)
		}
		d.NewNonceHash3 = value
	}
	return nil
}

// construct implements constructor of SetClientDHParamsAnswer.
func (d DhGenFail) construct() SetClientDHParamsAnswer { return &d }

// Ensuring interfaces in compile-time for DhGenFail.
var (
	_ bin.Encoder = &DhGenFail{}
	_ bin.Decoder = &DhGenFail{}

	_ SetClientDHParamsAnswer = &DhGenFail{}
)

// SetClientDHParamsAnswer represents SetClientDHParamsAnswer generic type.
//
// Example:
//  g, err := DecodeSetClientDHParamsAnswer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *DhGenOk: // dh_gen_ok#3bcbf734
//  case *DhGenRetry: // dh_gen_retry#46dc1fb9
//  case *DhGenFail: // dh_gen_fail#a69dae02
//  default: panic(v)
//  }
type SetClientDHParamsAnswer interface {
	bin.Encoder
	bin.Decoder
	construct() SetClientDHParamsAnswer
}

// DecodeSetClientDHParamsAnswer implements binary de-serialization for SetClientDHParamsAnswer.
func DecodeSetClientDHParamsAnswer(buf *bin.Buffer) (SetClientDHParamsAnswer, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DhGenOkTypeID:
		// Decoding dh_gen_ok#3bcbf734.
		v := DhGenOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswer: %w", err)
		}
		return &v, nil
	case DhGenRetryTypeID:
		// Decoding dh_gen_retry#46dc1fb9.
		v := DhGenRetry{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswer: %w", err)
		}
		return &v, nil
	case DhGenFailTypeID:
		// Decoding dh_gen_fail#a69dae02.
		v := DhGenFail{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswer: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode SetClientDHParamsAnswer: %w", bin.NewUnexpectedID(id))
	}
}
