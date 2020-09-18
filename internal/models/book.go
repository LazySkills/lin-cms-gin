/** Created By wene<354007048@qq.com> . Date at 2020/6/9 */
package models

import jtime "lin-cms-gin/internal/pkg/time"

// Book [...]
type Book struct {
	Model
	Title      string    `gorm:"column:title;type:varchar(50);not null;comment:'书名'"`
	Author     string    `gorm:"column:author;type:varchar(30);not null;comment:'作者'"`
	Summary    string    `gorm:"column:summary;type:varchar(1000);not null;comment:'简介'"`
	Image      string    `gorm:"column:image;type:varchar(50);not null;comment:'封面'"`
	DeleteTime jtime.JSONTime  `gorm:"column:delete_time;type:datetime(3);default:NULL;comment:'删除时间'"`
}

func GetBook(pageNum int, pageSize int, maps interface {}) (books []Book) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&books)

	return
}

func GetBookTotal(maps interface {}) (count int64){
	db.Model(&Book{}).Where(maps).Count(&count)

	return
}

func ExistBookByName(name string) bool {
	var book Book
	db.Select("id").Where("title = ?", name).First(&book)
	if book.ID > 0 {
		return true
	}

	return false
}

func AddBook(title string, author string, summary string, image string) bool{
	db.Create(&Book{
		Title : title,
		Author : author,
		Summary : summary,
		Image : image,
	})

	return true
}

func ExistBookByID(id int) (book Book) {
	db.Where("id = ?", id).First(&book)

	return
}

func DeleteBook(id int) bool {
	db.Where("id = ?", id).Delete(&Book{})

	return true
}

func EditBook(id int, data interface {}) bool {
	db.Model(&Book{}).Where("id = ?", id).Updates(data)

	return true
}

//func (book *Book) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (book *Book) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}

