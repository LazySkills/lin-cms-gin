package v1


type AddBookValidator struct {
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Author          string `form:"author" valid:"Required;MaxSize(255)"`
	Summary       string `form:"summary" valid:"Required;MaxSize(65535)"`
	Image string `form:"image" valid:"Required;MaxSize(255)"`
}


type UpdateBookValidator struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	Title         string `form:"title" valid:"Required;MaxSize(100)"`
	Author        string `form:"author" valid:"Required;MaxSize(255)"`
	Summary       string `form:"summary" valid:"Required;MaxSize(65535)"`
	Image 		  string `form:"image" valid:"Required;MaxSize(255)"`
}