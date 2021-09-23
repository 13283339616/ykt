package vo

type GetUserInfoRequestVo struct {
	Name string `json:"name" form:"name" binding:"required"`
}
