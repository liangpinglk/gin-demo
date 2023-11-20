package openapi

type CommonInfoRes struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}
type CreateUserInfo struct {
	name     string `json:"name"`
	password string `json:"password"`
}
type ListUserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type CreateUserRes struct {
	CommonInfoRes
	data CreateUserInfo
}

type ListUserRes struct {
	CommonInfoRes
	data []ListUserInfo
}

type UpdateInfo struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UpdateUserRes struct {
	CommonInfoRes
	data UpdateInfo
}
