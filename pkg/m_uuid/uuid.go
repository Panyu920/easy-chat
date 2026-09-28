package muuid

import (
	"github.com/google/uuid"
	"strings"
)

// GenUUIDv7 生成UUIDv7 无横杠字符串
func GenUUIDv7() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	s := id.String()
	return strings.ReplaceAll(s, "-", ""), nil
}
