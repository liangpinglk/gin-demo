package openapi

type CommonInfoRes struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}
type UserInfo struct {
	name     string `json:"name"`
	password string `json:"password"`
}
type CreateUserRes struct {
	CommonInfoRes
	data UserInfo
}
