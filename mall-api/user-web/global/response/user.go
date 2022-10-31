package response

type UserResponse struct {
	Id       int32  `json:"id"`
	NickName string `json:"name"`
	Birthday string `json:"birthday"`
	Mobile   string `json:"mobile"`
	Gender   string `json:"gender"`
}
