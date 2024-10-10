package helpers

type Message struct {
	Status string
}

type ResponseParams struct {
	Message string
	Data    any
}

type UriIDParam struct {
	ID int `uri:"id" binding:"required,number"`
}

type ListQueryParams struct {
	OrderBy  string `form:"order-by"`
	OrderDir string `form:"order-dir"`
	Search   string `form:"search"`
	Limit    int    `form:"limit" validate:"number"`
	Skip     int    `form:"skip" validate:"number"`
}
