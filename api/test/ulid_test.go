package test

import (
	"encoding/base64"
	"fmt"
	"github.com/codelite7/momentum/api/ent/schema/nulid"
	"github.com/gofrs/uuid/v5"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNulid(t *testing.T) {
	const prefix = "user"
	id := nulid.MustNew("user")
	fmt.Println(id.Uid.String())
	fmt.Println(id.Prefix)
	encoded := id.String()
	fmt.Println(encoded)
	decoded, err := nulid.FromString(encoded)
	require.NoError(t, err)
	require.Equal(t, id.Uid.String(), decoded.Uid.String())
	require.Equal(t, id.Prefix, decoded.Prefix)
	fmt.Println(decoded)
}

func TestUlidEncode(t *testing.T) {
	id, err := uuid.NewV7()
	fmt.Println(id.String())
	require.NoError(t, err)
	base64 := encodeBase64(id)
	fmt.Println(base64)
	decoded, err := decodeBase64(base64)
	fmt.Println(decoded.String())
	require.NoError(t, err)
	require.EqualValues(t, id.String(), decoded.String())
}

func encodeBase64(id uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(id.Bytes())
}

func decodeBase64(id string) (uuid.UUID, error) {
	bytes, err := base64.RawURLEncoding.DecodeString(id)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.FromBytes(bytes)
}
