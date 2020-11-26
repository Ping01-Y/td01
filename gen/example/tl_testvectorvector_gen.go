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

// TestVectorVector represents TL type testVectorVector#69e8846c.
type TestVectorVector struct {
	// Value field of TestVectorVector.
	Value [][]string
}

// TestVectorVectorTypeID is TL type id of TestVectorVector.
const TestVectorVectorTypeID = 0x69e8846c

// Encode implements bin.Encoder.
func (t *TestVectorVector) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorVector#69e8846c as nil")
	}
	b.PutID(TestVectorVectorTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, row := range t.Value {
		b.PutVectorHeader(len(row))
		for _, v := range row {
			b.PutString(v)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorVector) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorVector#69e8846c to nil")
	}
	if err := b.ConsumeID(TestVectorVectorTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorVector#69e8846c: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorVector#69e8846c: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			innerLen, err := b.VectorHeader()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorVector#69e8846c: field value: %w", err)
			}
			var row []string
			for innerIndex := 0; innerIndex < innerLen; innerLen++ {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode testVectorVector#69e8846c: field value: %w", err)
				}
				row = append(row, value)
			}
			t.Value = append(t.Value, row)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorVector.
var (
	_ bin.Encoder = &TestVectorVector{}
	_ bin.Decoder = &TestVectorVector{}
)
