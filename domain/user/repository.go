package user

import (
	"gorm.io/gorm"
	"log"
)

type Repository struct {
	db *gorm.DB
}

// 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (this *Repository) Migration() {
	err := this.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户
func (this *Repository) Create(user *User) error {
	result := this.db.Create(user)
	return result.Error
}

// 根据用户名查询用户
func (this *Repository) GetByName(name string) (User, error) {
	var user User
	err := this.db.Where("UserName = ?", name).Where("IsDeleted = ?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 更新用户
func (this *Repository) Update(user *User) error {
	return this.db.Save(&user).Error
}

// 添加测试数据
func (this *Repository) InsertSampleData() {
	user := NewUser("admin", "admin")
	user.IsAdmin = true
	this.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password,
		}).FirstOrCreate(&user)

	user = NewUser("aiewing", "aiewing")
	user.IsAdmin = false
	this.db.Where(User{Username: user.Username}).Attrs(
		User{
			Username: user.Username, Password: user.Password,
		}).FirstOrCreate(&user)
}
