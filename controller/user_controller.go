package controller

import (
	"github.com/kataras/iris/v12"
	"main/model"
	"main/service"
	"main/utils"
)

type UserController struct {
	Service *service.UserService
}

func (userController *UserController) GetBy(id uint) *model.User {
	return userController.Service.GetById(id)
}

func (userController *UserController) GetAll() *[]model.User {
	return userController.Service.GetAll()
}

func (userController *UserController) GetListBySex(sex string, ctx iris.Context) iris.Map {
	page, _ := ctx.URLParamInt("page")
	pageSize, _ := ctx.URLParamInt("pageSize")
	return userController.Service.GetListBySex(page, pageSize, sex)
}

func (userController *UserController) GetList(ctx iris.Context) iris.Map {
	page, _ := ctx.URLParamInt("page")
	pageSize, _ := ctx.URLParamInt("pageSize")
	var user model.User
	ctx.ReadForm(&user)
	return userController.Service.GetList(utils.NewPagination(page, pageSize), &user)
}

func (userController *UserController) GetListPage(pagination *utils.Pagination, user *model.User, ctx iris.Context) iris.Map {

	return userController.Service.GetList(pagination, user)
}

func (userController *UserController) Post(user *model.User, ctx iris.Context) *model.User {
	/*var user model.User
	ctx.ReadForm(&user)*/
	return userController.Service.Add(user)
}

func (userController *UserController) Put(user *model.User) *model.User {
	return userController.Service.Update(user)
}

func (userController *UserController) DeleteBy(id uint) *model.User {
	return userController.Service.Delete(id)
}

func (userController *UserController) GetJoinList(pagination *utils.Pagination, user *model.User) iris.Map {

	return userController.Service.JoinSelect(pagination, user)
}
