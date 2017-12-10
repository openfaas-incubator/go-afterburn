package function

import (
	"fmt"
)

func Handle(req []byte) []byte {
	return []byte(fmt.Sprintf("Hello there %s\n", string(req)))
}
