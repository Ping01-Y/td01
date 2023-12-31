package crypto

import (
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRSAPublicDecrypt(t *testing.T) {
	testKey := []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAyr+18Rex2ohtVy8sroGPBwXD3DOoKCSpjDqYoXgCqB7ioln4eDCF
fOBUlfXUEvM/fnKCpF46VkAftlb4VuPDeQSS/ZxZYEGqHaywlroVnXHIjgqoxiAd
192xRGreuXIaUKmkwlM9JID9WS2jUsTpzQ91L8MEPLJ/4zrBwZua8W5fECwCCh2c
9G5IzzBm+otMS/YKwmR1olzRCyEkyAEjXWqBI9Ftv5eG8m0VkBzOG655WIYdyV0H
fDK/NWcvGqa0w/nriMD6mDjKOryamw0OP9QuYgMN0C9xMW9y8SmP4h92OAWodTYg
Y1hZCxdv6cs5UnW9+PWvS+WIbkh+GaWYxwIDAQAB
-----END RSA PUBLIC KEY-----`)
	data := []uint8{
		0x02, 0x05, 0x1a, 0x4a, 0x58, 0x12, 0xc3, 0xa6,
		0xb1, 0xa4, 0xb5, 0x61, 0x2e, 0x45, 0x29, 0x12,
		0xaa, 0xeb, 0x77, 0xc2, 0x12, 0x19, 0x37, 0x64,
		0x3d, 0xb7, 0xfa, 0x6a, 0x7c, 0xe0, 0xaa, 0xc3,
		0xc4, 0xcc, 0xe6, 0x6c, 0xdb, 0x55, 0xa6, 0x33,
		0x8d, 0x38, 0x45, 0xb6, 0x35, 0xb0, 0x61, 0x10,
		0x87, 0x28, 0x33, 0x40, 0x2a, 0x24, 0x22, 0x47,
		0xda, 0x84, 0x5a, 0x9f, 0x16, 0x8e, 0xd2, 0x07,
		0x08, 0x25, 0x32, 0xc8, 0x55, 0x42, 0x20, 0x7f,
		0x58, 0x52, 0x2f, 0x33, 0x1e, 0x46, 0x4f, 0x86,
		0x76, 0x00, 0xa3, 0x49, 0x15, 0xe9, 0x6e, 0xff,
		0x41, 0x5f, 0xc9, 0xbd, 0xee, 0x1e, 0x6c, 0x12,
		0x49, 0x82, 0xe0, 0x0c, 0xe6, 0x96, 0x3b, 0x06,
		0xa3, 0x59, 0xc4, 0x66, 0x12, 0xfd, 0xd6, 0x36,
		0xdd, 0x88, 0xe0, 0x90, 0x2e, 0x45, 0x26, 0x3e,
		0x83, 0xbf, 0xec, 0xe0, 0x65, 0xbf, 0x1c, 0xef,
		0xf0, 0x67, 0x27, 0x5b, 0x9b, 0xe0, 0x6c, 0x15,
		0xd1, 0xa6, 0x46, 0x5b, 0x36, 0x70, 0x12, 0x85,
		0x44, 0x19, 0x9a, 0x28, 0xd4, 0x49, 0xd3, 0x19,
		0xc6, 0x04, 0xfc, 0xd8, 0xca, 0x1a, 0xc3, 0x6a,
		0xcd, 0x55, 0x84, 0x19, 0xaa, 0x0d, 0x6f, 0x3d,
		0x74, 0x91, 0xaf, 0xa3, 0x3a, 0x80, 0x1a, 0x23,
		0x46, 0x0c, 0xfe, 0x3d, 0xda, 0x6a, 0xa9, 0x00,
		0x3c, 0xb8, 0x1b, 0x3a, 0x2a, 0xcd, 0xb1, 0x6c,
		0x2b, 0xef, 0xa6, 0x83, 0x1f, 0x94, 0x04, 0x2c,
		0x35, 0x65, 0x31, 0xd0, 0x5e, 0x72, 0x84, 0xdd,
		0xda, 0x70, 0x64, 0x76, 0x08, 0xa7, 0x85, 0x5b,
		0x0b, 0xd9, 0xc4, 0xe8, 0x1a, 0x7b, 0x2d, 0xea,
		0xd4, 0x40, 0x39, 0xed, 0xd6, 0x9a, 0xd3, 0xa6,
		0x0e, 0x9d, 0x40, 0x3c, 0xb1, 0x26, 0x63, 0x6a,
		0xed, 0xfc, 0xfb, 0xe5, 0x97, 0x5c, 0x6b, 0x25,
		0x9c, 0x2e, 0xc3, 0xd3, 0x49, 0xd0, 0x15, 0x80,
	}
	expected := []byte{
		0x2f, 0x1a, 0x77, 0xcc, 0x57, 0xa4, 0xad, 0xb2,
		0x8f, 0xad, 0x2e, 0x49, 0x05, 0xaf, 0x88, 0x4a,
		0x3e, 0x15, 0x0a, 0xc7, 0xf4, 0x96, 0xdc, 0xe5,
		0x05, 0x7e, 0xc9, 0xd3, 0x5e, 0x3c, 0x86, 0x62,
		0x81, 0xb5, 0xd4, 0x11, 0x45, 0x3d, 0xf2, 0xa7,
		0x14, 0xe9, 0x15, 0xa4, 0x56, 0x60, 0x85, 0xd8,
		0xb2, 0xd2, 0x0f, 0x5c, 0xd6, 0x12, 0xfb, 0x85,
		0x98, 0xf2, 0x04, 0x5c, 0x75, 0xc3, 0x36, 0x69,
		0x42, 0x63, 0x72, 0xbc, 0xe6, 0xe0, 0x54, 0x86,
		0x6f, 0x42, 0xce, 0x36, 0xeb, 0xbc, 0x51, 0xa2,
		0x81, 0xeb, 0xcf, 0xd9, 0x74, 0x69, 0x05, 0x03,
		0x2c, 0xbe, 0xad, 0xc2, 0xd0, 0x08, 0x49, 0xd5,
		0x19, 0x5b, 0x8d, 0x35, 0x83, 0x45, 0x84, 0x3e,
		0xa5, 0x9b, 0x2b, 0x79, 0x76, 0x59, 0x7b, 0x3e,
		0x2a, 0x8f, 0xc1, 0xe0, 0x3e, 0x70, 0x68, 0x1f,
		0x7f, 0x0d, 0x92, 0x6f, 0x10, 0x85, 0x0c, 0x5c,
		0xa2, 0xdf, 0x74, 0x8a, 0x7d, 0x9c, 0x04, 0xa8,
		0xd0, 0x96, 0x39, 0xf1, 0xf4, 0xdb, 0x15, 0x22,
		0x38, 0xa0, 0x67, 0x6e, 0x87, 0x36, 0x2e, 0xce,
		0xef, 0x63, 0x15, 0xd9, 0x37, 0x7c, 0x59, 0xb1,
		0x62, 0x06, 0x83, 0x7c, 0x2e, 0xe2, 0x93, 0x65,
		0xad, 0xdb, 0x67, 0x1b, 0x0f, 0xe9, 0x81, 0x61,
		0xbb, 0x0d, 0x8c, 0x2b, 0x6f, 0x25, 0x58, 0x0d,
		0x24, 0x59, 0xff, 0xb2, 0x56, 0x62, 0x84, 0x0b,
		0xa6, 0x20, 0xee, 0xe3, 0x03, 0xb1, 0x7d, 0x39,
		0x91, 0x65, 0x6b, 0x5d, 0x3a, 0x94, 0xc8, 0xd1,
		0x01, 0x73, 0xa2, 0xd7, 0x38, 0x28, 0x9d, 0x0f,
		0x75, 0x0a, 0xad, 0xf4, 0x2e, 0x6f, 0x65, 0x2f,
		0x5f, 0xaa, 0x4c, 0x6e, 0x8b, 0x61, 0x60, 0xfb,
		0xab, 0xb1, 0x48, 0xa1, 0x32, 0x76, 0xa0, 0x88,
		0x42, 0x6c, 0xf5, 0x1e, 0x23, 0x3d, 0x03, 0x47,
		0x48, 0x09, 0x3d, 0xf9, 0x92, 0x1d, 0xcd, 0x66,
	}

	keys, err := ParseRSAPublicKeys(testKey)
	require.NoError(t, err)
	require.Len(t, keys, 1)
	key := keys[0]

	actual, err := RSAPublicDecrypt(key, data)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

	_, err = RSAPublicDecrypt(key, nil)
	assert.ErrorIs(t, err, rsa.ErrVerification)
}
