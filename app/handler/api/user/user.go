package user

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	//"strconv"
	//"time"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/crypto"
	"github.com/20326/vega/pkg/render"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func Step2CodeAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	// TODO
	result.Result = map[string]interface{}{
		"stepCode": 0,
	}
	print(result)
}

func RegisterAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	var err error
	srv := service.FromContext(c)
	log := srv.GetLogger()

	// user validate
	user := &model.User{}
	if err = c.BindJSON(&user); nil != err {
		result.Error(err)
		return
	}

	if err = user.Validate(); nil != err {
		//password and username
		result.Error(err)
		return
	}

	// find user
	out, err := srv.Users.FindName(c, user.Username)
	if nil != err && out != nil {
		result.Error(errors.New("not available username "))
		return
	}

	// create user
	password := crypto.HashAndSalt([]byte(user.Password))
	if err := srv.Users.Create(c, &model.User{
		Username:     user.Username,
		PasswordHash: password,
		Phone:        user.Phone,
	}); nil != err {
		result.Error(err)
	}

	log.WithFields(logrus.Fields{
		"action":   "UserRegister",
		"username": user.Username,
	}).Info("register success")

	// init session and save
	result.Result = map[string]interface{}{
		"next": "/",
	}
}

func LoginAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	log := srv.GetLogger()

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}

	userName := arg["username"].(string)
	password := arg["password"].(string)
	// find user by name
	user, err := srv.Users.FindName(c, userName)
	if nil != err {
		result.Error(err)
		return
	}
	// check password
	if !crypto.ComparePasswords(user.PasswordHash, []byte(password)) {
		result.Error(errors.New("username and password do not match"))
		return
	}

	// update user action
	t := time.Now()
	if err := srv.Users.Updates(c, user, map[string]interface{}{
		"token":    uuid.NewV4().String(),
		"login_ip": c.Request.RemoteAddr,
		"login_at": &t,
	}); nil != err {
		log.WithError(err).Info("update user action")
	}

	// update session data
	session := &model.SessionData{
		UID:       user.ID,
		Username:  user.Username,
		Roles:     "admin", // TODO interface
		Token:     user.Token,
		Anonymous: false,
	}

	if err := session.Save(c); nil != err {
		result.Error(errors.New("internal error of login"))
		render.JSON(c.Writer, result, http.StatusInternalServerError)
	}

	// referer
	referer := c.Request.URL.Query().Get("referer")
	if "" == referer || !strings.Contains(referer, "://") {
		// add server for referer
	}
	if strings.HasSuffix(referer, "/") {
		referer = referer[:len(referer)-1]
	}

	log.WithFields(logrus.Fields{
		"action":  "UserLogin",
		"referer": referer,
	}).Info("login success")

	data := map[string]interface{}{}
	data["next"] = referer
	data["token"] = user.Token

	result.Msg = "success"
	result.Result = data
}

func LogoutAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	log := srv.GetLogger()

	session := &model.SessionData{}
	user, err := session.Get(c, srv.Users)
	if nil != err {
		result.Error(errors.New("session illegal"))
		return
	}

	// delete session
	if err := session.Delete(c, user, srv.Users); nil != err {
		result.Error(errors.New("internal error of login"))
	}

	log.WithFields(logrus.Fields{
		"action":   "UserLogout",
		"username": user.Username,
	}).Info("logout success")

	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	data["next"] = "/login"
	result.Result = data
}

func ChangePasswordAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	log := srv.GetLogger()

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}

	oldPassword := arg["old_password"].(string)
	password := arg["password"].(string)

	session := &model.SessionData{}
	user, err := session.Get(c, srv.Users)
	if nil != err {
		result.Error(errors.New("session illegal"))
		return
	}

	// check password
	session.Token = uuid.NewV4().String()
	if !crypto.ComparePasswords(user.PasswordHash, []byte(oldPassword)) {
		result.Error(errors.New("invalid old password"))
		return
	} else {
		if err = user.Validate(); nil != err {
			//password and username
			result.Error(err)
			return
		}
		err = srv.Users.Updates(c, user, map[string]interface{}{
			"password_hash": crypto.HashAndSalt([]byte(password)),
			"token":    session.Token,
		})
	}

	// session update save
	if err := session.Save(c); nil != err {
		result.Error(errors.New("internal error"))
	}

	log.WithFields(logrus.Fields{
		"action":   "UserChangePassword",
		"username": user.Username,
	}).Info("change password success")

	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	result.Result = data
}

func ForgetPasswordAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}

	userName := arg["username"].(string)

	srv := service.FromContext(c)
	log := srv.GetLogger()

	user, err := srv.Users.FindName(c, userName)
	if nil == user {
		result.Error(err)
		return
	}

	// send email or sms

	// session delete ?

	log.WithFields(logrus.Fields{
		"action":   "UserForgetPassword",
		"username": user.Username,
	}).Info("forget password")

	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	result.Result = data
}

func ResetPasswordAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	log := srv.GetLogger()

	userName := arg["username"].(string)
	user, err := srv.Users.FindName(c, userName)
	if nil == user {
		result.Error(err)
		return
	}

	log.WithFields(logrus.Fields{
		"action":   "UserResetPassword",
		"username": user.Username,
	}).Info("forget password")

	// session delete ?
	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	result.Result = data
}

func UserInfoAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	log := srv.GetLogger()

	session := &model.SessionData{}
	user, err := session.Get(c, srv.Users)
	if nil != err {
		result.Error(errors.New("session illegal"))
		return
	}

	// var permissions map[uint64]*model.ConsolePermissionission
	//roleModels := srv.Roles.ConsoleGetRoles()
	//log.Warn().Msgf("roleModels %+v", roleModels)

	//for _, roleModel := range roleModels {
	//
	//	log.Warn().Msgf("roleModel Permissionissions:  %s  ", len(roleModel.Permissionissions))
	//	//for _, permission := range roleModel.Permissionissions {
	//	//}
	//
	//	//permissionModels := service.Permissionission.GetPermissionissions()
	//	//
	//	//for _, permissionModel := range permissionModels {
	//	//	permission := &model.ConsolePermissionission{
	//	//		ID:             permissionModel.ID,
	//	//		Name:           permissionModel.Name,
	//	//		//PermissionissionId:   permissionModel.Object,
	//	//		PermissionissionName: permissionModel.Name,
	//	//		//ActionEntitySet:     resourceModel.Actions,
	//	//		RoleID:         roleModel.Name,
	//	//		Status:         1,
	//	//	}
	//	//
	//	//	actionData, err := json.Marshal(resourceModel.Actions)
	//	//	permission.ActionData = string(actionData)
	//	//
	//	//	log.Warn().Msgf("roleModels err:  %s  ", err, permission.Actions)
	//	//	if nil != err {
	//	//		continue
	//	//	}
	//	//
	//	//	permissions = append(permissions, permission)
	//	//}
	//}

	log.WithFields(logrus.Fields{
		"action":   "UserInfo",
		"username": user.Username,
	}).Info("user get info success")

	data := map[string]interface{}{
		"id":            strconv.Itoa(int(user.ID)),
		"name":          user.Username,
		"username":      user.Username,
		"password":      "",
		"avatar":        user.Avatar,
		"status":        user.Status,
		"lastLoginIp":   user.LoginIP,
		"lastLoginTime": user.LoginAt,
		"roleId":        "admin",
		"role": map[string]interface{}{
			"permissions": Permissionissions,
		},
	}
	result.Msg = "success"
	result.Result = data
}
