package user

import "shopping/utils/hash"

type Service struct {
	repo Repository
}

// 实例化service
func NewUserService(repo Repository) *Service {
	repo.Migration()
	return &Service{
		repo: repo,
	}
}

// 创建用户
func (this *Service) Create(user *User) error {

	// 无效用户名
	if ValidateUserName(user.Username) == false {
		return ErrorUsernameInvalid
	}

	// 无效密码
	if ValidatePassword(user.Password) == false {
		return ErrorPasswordError
	}

	// 用户名已经存在
	_, err := this.repo.GetByName(user.Username)
	if err == nil {
		return ErrorUserExistWithName
	}

	// 创建用户
	err = this.repo.Create(user)
	return err
}

// 查询用户
func (this *Service) GetUser(username string, password string) (User, error) {
	user, err := this.repo.GetByName(username)
	if err != nil {
		return User{}, ErrorUserNotFound
	}

	match := hash.CheckPasswordHash(password+user.Salt, user.Password)
	if !match {
		return User{}, ErrorPasswordError
	}
	return user, nil
}

// 更新用户
func (this *Service) UpdateUser(user *User) error {
	return this.repo.Update(user)
}
