// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// A simple object containing a number; for testing only
type TestInt struct {
	// Number
	Value int32
}

// TestIntTypeID is TL type id of TestInt.
const TestIntTypeID = 0xddbd2c09

// Encode implements bin.Encoder.
func (t *TestInt) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testInt#ddbd2c09 as nil")
	}
	b.PutID(TestIntTypeID)
	b.PutInt32(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestInt) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testInt#ddbd2c09 to nil")
	}
	if err := b.ConsumeID(TestIntTypeID); err != nil {
		return fmt.Errorf("unable to decode testInt#ddbd2c09: %w", err)
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode testInt#ddbd2c09: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TestInt.
var (
	_ bin.Encoder = &TestInt{}
	_ bin.Decoder = &TestInt{}
)
