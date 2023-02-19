package request

type Todo struct {
	Title string `json:"title" bindings:"required,max=255"`
}
