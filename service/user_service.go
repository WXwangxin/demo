package service

import (
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"main/model"
	"main/utils"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (service *UserService) GetById(id uint) *model.User {
	var user model.User
	service.db.Model(&model.User{}).Where("id=?", id).First(&user)
	return &user
}

func (service *UserService) GetAll() *[]model.User {
	var users []model.User
	service.db.Model(&model.User{}).Find(&users)
	return &users
}

func (service *UserService) GetListBySex(page, pageSize int, sex string) iris.Map {
	query := service.db.Model(&model.User{}).Where(model.User{Sex: sex})
	var users []model.User
	data := utils.WrapPage(page, pageSize, &users, query)
	return data
}

func (service *UserService) GetList(pagination *utils.Pagination, user *model.User) iris.Map {
	query := service.db.Model(&model.User{}).Where(user)
	var users []model.User
	data := utils.WrapPage(pagination.Page, pagination.PageSize, &users, query)
	return data
}

func (service *UserService) Add(user *model.User) *model.User {
	service.db.Create(&user)
	return user
}

func (service *UserService) Update(user *model.User) *model.User {
	service.db.Updates(&user)
	return user
}

func (service *UserService) Delete(id uint) *model.User {
	var user model.User
	service.db.Where("id=?", id).First(&user)
	service.db.Delete(&user)
	return &user
}

func (service *UserService) JoinSelect(pagination *utils.Pagination, user *model.User) iris.Map {
	//service.db.Row("SELECT * from t_users u join dogs d on u.age = d.age") 使用count后，不能再次执行，且不支持offset和limit
	//service.db.Model(model.User{})  会导致其它表的非整型字段，转换为uint数组，导致乱码
	query := service.db.Table("t_users usr").
		Joins("join dogs dog on usr.age=dog.age").
		Where(&user).
		Select("usr.id,usr.name,dog.id as dog_id,dog.name as dog_name,dog.age as dog_age,dog.breed")
	var users []iris.Map
	data := utils.WrapPage(pagination.Page, pagination.PageSize, &users, query)
	return data
}
