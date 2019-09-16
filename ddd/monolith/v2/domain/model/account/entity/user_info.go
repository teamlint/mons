package entity

// UserInfo 用户聚合
type UserInfo struct {
	User
	Roles []*Role `json:"roles"`
}
