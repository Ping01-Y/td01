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

// Bytes represents TL type bytes#e937bb82.
type Bytes struct {
}

// BytesTypeID is TL type id of Bytes.
const BytesTypeID = 0xe937bb82

// Encode implements bin.Encoder.
func (b *Bytes) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bytes#e937bb82 as nil")
	}
	buf.PutID(BytesTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (b *Bytes) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bytes#e937bb82 to nil")
	}
	if err := buf.ConsumeID(BytesTypeID); err != nil {
		return fmt.Errorf("unable to decode bytes#e937bb82: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Bytes.
var (
	_ bin.Encoder = &Bytes{}
	_ bin.Decoder = &Bytes{}
)
