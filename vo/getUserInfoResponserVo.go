package vo

type GetUserInfoResponsetVo struct {
	Name string `json:"name" form:"name" binding:"required"`
}
