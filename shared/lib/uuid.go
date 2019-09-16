package lib

import "github.com/hashicorp/go-uuid"

// UUID uuid字符串
func UUID() string {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return id
}
