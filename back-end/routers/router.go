package routers

import (
	"github.com/astaxie/beego"
	"github.com/tongyuehong1/design-back-end/back-end/controllers"
)

func init() {
	beego.Router("/user/register", &controllers.UserController{}, "post:Register")
	beego.Router("/user/login", &controllers.UserController{}, "post:Login")
	beego.Router("/user/class", &controllers.UserController{}, "post:Classes")

	beego.Router("/student/add", &controllers.StudentController{}, "post:Add")
	beego.Router("/student/insert", &controllers.StudentController{}, "post:Insert")
	beego.Router("/student/change", &controllers.StudentController{}, "post:Modify")
	beego.Router("/student/getall", &controllers.StudentController{}, "post:GetAll")
	beego.Router("/student/getleader", &controllers.StudentController{}, "post:GetLeaders")
	beego.Router("/student/getone", &controllers.StudentController{}, "post:GetOne")
	beego.Router("/student/upavatar", &controllers.StudentController{}, "post:UpAvatar")
	beego.Router("/student/delete", &controllers.StudentController{},"post:Delete")

	beego.Router("/teacher/change", &controllers.TeacherController{}, "Post:ChangeTech")
	beego.Router("/teacher/get", &controllers.TeacherController{}, "Post:GetTeacher")
	beego.Router("/teacher/upavatar", &controllers.StudentController{}, "post:TechAvatar")


	beego.Router("/message/publish", &controllers.MessageController{}, "Post:Publish")
	beego.Router("/message/show", &controllers.MessageController{}, "Post:Show")

	beego.Router("/grade/all", &controllers.GradeController{}, "Post:GetAll")
	beego.Router("/grade/one", &controllers.GradeController{}, "Post:GetOne")
	beego.Router("/grade/add", &controllers.GradeController{}, "Post:AddGrade")



}
