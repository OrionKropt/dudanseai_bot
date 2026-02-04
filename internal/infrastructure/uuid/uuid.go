package uuid

import (
	"github.com/google/uuid"
)

type UUID uuid.UUID

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

func Generate(base []byte) UUID {
	id := uuid.NewSHA1(uuid.NameSpaceDNS, base)
	return UUID(id)
}
