package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"` //格式做validate
	PassWord string `form:"password" json:"password" binding:"required,min=8,max=16"`
}
