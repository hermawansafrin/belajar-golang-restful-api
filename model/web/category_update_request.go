package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`                 // request dari client
	Name string `validate:"required,max=200,min=1" json:"name"` // request dari client
}
