package user

import (
	"net/http"
	//"strconv"
	"errors"
	"strings"
	//"time"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/crypto"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
	//"github.com/20326/vega/app/handler/common"
	uuid "github.com/satori/go.uuid"
)

type LoginUser struct {
	ID       uint64 `json:"id"`
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func TestAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	// clear session

	data := map[string]interface{}{}
	data["stepCode"] = 0
	result.Result = data
	print(result)

}

func Step2CodeAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	user, err := srv.Users.Find(c, 1)
	if nil == user {
		log.Warn().Err(err).Msg("can not get user by id")
		result.Error(err)

		return
	}
	// clear session

	data := map[string]interface{}{}
	data["stepCode"] = 0
	result.Result = data
	print(result)
}

func RegisterAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	log.Info().Str("atcion", "register").Msg("User SignUp [" + c.Request.URL.String() + "]")

	arg := &LoginUser{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}
	if "" == arg.Name {
		arg.Name = arg.Phone
	}

	userName := arg.Name
	password := arg.Password

	data := map[string]interface{}{}

	srv := service.FromContext(c)
	user, _ := srv.Users.FindName(c, userName)
	if nil != user {

		// hash the password
		user = &model.User{
			Username: userName,
			Avatar:   "/favicon.ico",
			Password: crypto.HashAndSalt([]byte(password)),
		}

		if err := srv.Users.Create(c, user); nil != err {
			result.Error(err)
		}
	}

	// init session and save
	data["next"] = "/"
	result.Result = data
}

func LoginAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	log.Info().Msg("User Login [" + c.Request.URL.String() + "]")

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}
	log.Info().Msgf("User Login %v", arg)

	userName := arg["username"].(string)
	password := arg["password"].(string)

	srv := service.FromContext(c)
	user, err := srv.Users.FindName(c, userName)
	if nil != err {
		log.Warn().Msg("can not get user by name [" + userName + "]")
		result.Error(err)
		return
	}

	// check password
	success := crypto.ComparePasswords(user.Password, []byte(password))
	if !success {
		log.Warn().Msg("Invalid Password [" + userName + "]")
		result.Error(err)
		return
	}

	log.Info().Msgf("User: %v", user)

	user.Token = uuid.NewV4().String()
	user.LoginIP = c.Request.RemoteAddr
	// user.LoginAt = time.Now()

	// update user action
	if err := srv.Users.Update(c, user); nil != err {
		log.Error().Err(err).Str("token", user.Token).Msg("update user action")
	}

	// new session data
	//session := &pkg.SessionData{
	//	UID:       user.ID,
	//	UName:     user.Nickname,
	//	UNickname: user.Nickname,
	//	UAvatar:   user.AvatarURL,
	//	UPhone:    user.Phone,
	//	TOKEN:     user.Token,
	//	// URoles:    user.Role,
	//
	//}

	// save session
	//log.Error().Msgf("saves session: %+v", session)
	//if err := session.Save(c); nil != err {
	//	log.Error().Err(err).Msg("saves session failed")
	//	c.Status(http.StatusInternalServerError)
	//}

	referer := c.Request.URL.Query().Get("referer")
	if "" == referer || !strings.Contains(referer, "://") {
		// add server for referer
	}
	if strings.HasSuffix(referer, "/") {
		referer = referer[:len(referer)-1]
	}
	log.Info().Msg("user sign referer [" + referer + " username:" + userName + " password: " + password + "]")

	data := map[string]interface{}{}
	data["next"] = referer
	// data["token"] = user.Token

	result.Msg = "success"
	result.Result = data
}

func LogoutAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	//session := sessions.Default(c)
	//// get user
	//srv := service.FromContext(c)
	//user := srv.Users.Get(session.UID)
	//if nil == user {
	//	log.Error().Msg("session illegal")
	//	result.Code = errors.CodeErr
	//	result.Msg = "session illegal"
	//	return
	//}
	//
	//user.Token = ""
	//if err := s.Users.Update(user); nil != err {
	//	log.Error().Err(err).Str("token", user.Token).Msg("update user action")
	//}

	// clear session
	defaultSession := sessions.Default(c)
	defaultSession.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	defaultSession.Clear()
	if err := defaultSession.Save(); nil != err {
		log.Error().Err(err).Msg("saves session failed")
	}

	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	data["next"] = "/login"
	result.Result = data
}

func ChangePasswordAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	arg := map[string]interface{}{}
	if err := c.BindJSON(&arg); nil != err {
		result.Error(err)

		return
	}

	userName := arg["username"].(string)
	oldPassword := arg["old_password"].(string)
	password := arg["password"].(string)

	srv := service.FromContext(c)
	user, err := srv.Users.FindName(c, userName)
	if nil != err {
		log.Warn().Msg("can not get user by name [" + userName + "]")
		result.Error(err)
		result.Result = map[string]interface{}{}
		return
	}

	// check password
	success := crypto.ComparePasswords(user.Password, []byte(oldPassword))
	if !success {
		log.Warn().Msg("Invalid Old Password [" + userName + "]")
		result.Error(errors.New("invalid password"))
		return
	} else {
		user.Password = crypto.HashAndSalt([]byte(password))
		_ = srv.Users.Update(c, user)
	}

	own, err := srv.Users.Find(c, user.ID)
	if nil == own {
		log.Warn().Msg("can not get user by name [" + userName + "]")

		data := map[string]interface{}{}
		data["msg"] = "error"
		data["errorMsg"] = "can not get user Blog by name [" + userName + "]"
		result.Result = data

		return
	}

	// session save

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
	user, err := srv.Users.FindName(c, userName)
	if nil == user {
		log.Warn().Msg("can not get user by name [" + userName + "]")
		data := map[string]interface{}{}
		result.Error(err)
		result.Result = data
		return
	}
	// send email or sms

	// session save

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

	userName := arg["username"].(string)

	srv := service.FromContext(c)
	user, err := srv.Users.FindName(c, userName)
	if nil == user {
		log.Warn().Msg("can not get user by name [" + userName + "]")
		data := map[string]interface{}{}
		result.Error(err)
		result.Result = data
		return
	}
	// send email or sms

	// session save

	data := map[string]interface{}{}
	data["msg"] = "ok"
	data["errorMsg"] = ""
	result.Result = data
}

func UserInfoAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	log.Info().Msg("User Profile [" + c.Request.URL.String() + "]")

	// {
	//   "id': '4291d7da9005377ec9aec4a71ea837f',
	//   'name': '天野远子',
	//   'username': 'admin',
	//   'password': '',
	//   'avatar': '/avatar2.jpg',
	//   'status': 1,
	//   'telephone': '',
	//   'lastLoginIp': '27.154.74.117',
	//   'lastLoginTime': 1534837621348,
	//   'creatorId': 'admin',
	//   'createTime': 1497160610259,
	//   'merchantCode': 'TLif2btpzg079h15bk',
	//   'deleted': 0,
	//   'roleId': 'admin',
	//   'role': {}
	//  }
	// get name from session
	// TODO
	// session := pkg.GetSession(c)

	//user := srv.User.GetUserWithRole(session.UID)
	//if nil == user {
	//	log.Warn().Msg("can not get user by name [" + session.UName + "]")
	//	result.Code = pkg.CodeErr
	//	result.Msg = "can not get user by name [" + session.UName + "]"
	//	return
	//}
	//
	//// var permissions map[uint64]*model.ConsolePermissionission
	//roleModels := srv.Roles.ConsoleGetRoles()
	//log.Warn().Msgf("roleModels %+v", roleModels)
	//
	//
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
	//
	//data := map[string]interface{}{
	//	"id":            strconv.Itoa(int(user.ID)),
	//	"name":          user.Name,
	//	"username":      user.Name,
	//	"password":      "",
	//	"avatar":        user.AvatarURL,
	//	"status":        user.Status,
	//	"lastLoginIp":   user.LoginIP,
	//	"lastLoginTime": user.LoginAt,
	//	"roleId":        "admin",
	//	"role": map[string]interface{}{
	//		"permissions": common.Permissionissions,
	//	},
	//}
	data := map[string]interface{}{}
	log.Error().Msgf("UserInfoAction Data:", data)
	result.Msg = "Success"
	result.Result = data
}
