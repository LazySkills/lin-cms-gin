package v1


type AddBookValidator struct {
	Title         string `form:"title" binding:"required,max=100"`
	Author          string `form:"author" binding:"required,max=255"`
	Summary       string `form:"summary" binding:"required,max=65535"`
	Image string `form:"image" binding:"required;max=255"`
}


type UpdateBookValidator struct {
	ID            int    `form:"id" binding:"required,min=1"`
	Title         string `form:"title" binding:"required,max=100"`
	Author        string `form:"author" binding:"required,max=255"`
	Summary       string `form:"summary" binding:"required,max=65535"`
	Image 		  string `form:"image" binding:"required,max=255"`
}
