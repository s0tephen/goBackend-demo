package text

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID() string {
	uuids := uuid.New()
	return strings.ReplaceAll(uuids.String(), "-", "")
}
