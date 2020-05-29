package common

var Permissionissions = []map[string]interface{}{
	{
		"name":           "异常页面权限",
		"roleId":         "admin",
		"permissionId":   "exception",
		"permissionName": "异常页面权限",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "仪表盘",
		"roleId":         "admin",
		"permissionId":   "dashboard",
		"permissionName": "仪表盘",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "分析页",
		"roleId":         "test",
		"permissionId":   "dashboard",
		"permissionName": "仪表盘",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "权限管理",
		"roleId":         "admin",
		"permissionId":   "permission",
		"permissionName": "权限管理",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "角色管理",
		"roleId":         "admin",
		"permissionId":   "role",
		"permissionName": "角色管理",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "用户管理",
		"roleId":         "admin",
		"permissionId":   "users",
		"permissionName": "用户管理",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "超级模块",
		"roleId":         "admin",
		"permissionId":   "support",
		"permissionName": "超级模块",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "帐户设置",
		"roleId":         "admin",
		"permissionId":   "user",
		"permissionName": "帐户设置",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
	{
		"name":           "系统管理",
		"roleId":         "admin",
		"permissionId":   "system",
		"permissionName": "系统管理",
		"actionData":     `[{"action":"add","defaultCheck":false,"describe":"新增"},{"action":"query","defaultCheck":false,"describe":"查询"},{"action":"get","defaultCheck":false,"describe":"详情"},{"action":"update","defaultCheck":false,"describe":"修改"},{"action":"delete","defaultCheck":false,"describe":"删除"}]`,
		"actionEntitySet": []map[string]interface{}{{
			"action":       "add",
			"describe":     "新增",
			"defaultCheck": false,
		}, {
			"action":       "query",
			"describe":     "查询",
			"defaultCheck": false,
		}, {
			"action":       "get",
			"describe":     "详情",
			"defaultCheck": false,
		}, {
			"action":       "update",
			"describe":     "修改",
			"defaultCheck": false,
		}, {
			"action":       "delete",
			"describe":     "删除",
			"defaultCheck": false,
		}},
		"actionList": nil,
		"dataAccess": nil,
	},
}
