package nulid

import (
	"database/sql/driver"
	"encoding/base64"
	"fmt"
	"github.com/gofrs/uuid"
	"strings"
)

type Nulid struct {
	Uid    uuid.UUID
	Prefix string
}

var Nil = Nulid{}

const separator = ":"

func New(prefix string) (Nulid, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return Nil, err
	}

	return Nulid{Uid: id, Prefix: prefix}, nil
}

func MustNew(prefix string) Nulid {
	id, err := New(prefix)
	if err != nil {
		panic(err)
	}
	return id
}

func (n Nulid) String() string {
	builder := strings.Builder{}
	builder.Write([]byte(n.Prefix))
	builder.Write([]byte(separator))
	builder.Write(n.Uid.Bytes())
	return base64.RawURLEncoding.EncodeToString([]byte(builder.String()))
}

func FromString(value string) (Nulid, error) {
	bytes, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return Nil, err
	}
	byteString := string(bytes)
	parts := strings.SplitN(byteString, separator, 2)
	if len(parts) != 2 {
		return Nil, fmt.Errorf("invalid nulid format: %s", value)
	}
	uid, err := uuid.FromBytes([]byte(parts[1]))
	if err != nil {
		return Nil, err
	}
	return Nulid{Uid: uid, Prefix: string([]byte(parts[0]))}, nil
}

func (n Nulid) Value() (driver.Value, error) {
	return n.Uid.Value()
}

func (n *Nulid) Scan(src interface{}) error {
	var uid uuid.UUID
	err := uid.Scan(src)
	if err != nil {
		return err
	}
	n.Uid = uid
	return nil
}
