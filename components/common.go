package components

import (
	"cin/base"
	"cin/controllers"
	"reflect"
	"strings"
)

// 组件内通用的方法封装
type commonFunc struct {
}

// 组件内通用的方法
var common = new(commonFunc)

// 获取控制器记录map
func (cf *commonFunc) getControllerDict(controllerList []base.ControllerInterface) (map[string]reflect.Type, map[string]map[string]bool) {
	controllerDict := map[string]reflect.Type{}
	controllerActionDict := map[string]map[string]bool{}

	excludeMethodNameDict := cf.getControllerExcludeMethodNameDict()

	for _, controllerInterface := range controllerList {
		t := reflect.TypeOf(controllerInterface)
		controllerType := t.Elem() // 获取实体
		controllerName := controllerType.Name()
		controllerName = strings.TrimSuffix(controllerName, "Controller")
		controllerName = strings.ToLower(controllerName)
		controllerDict[controllerName] = t

		// 保存控制器啊所有方法名字
		controllerActionDict[controllerName] = map[string]bool{}
		methodLen := t.NumMethod()
		for i := 0; i < methodLen; i++ {
			method := t.Method(i)
			methodName := method.Name

			// 判断是否是排除的方法名字
			if _, exclude := excludeMethodNameDict[methodName]; exclude {
				continue
			}

			controllerActionDict[controllerName][methodName] = true
		}
	}

	return controllerDict, controllerActionDict
}

// 获取控制器排除的方法名字
func (cf *commonFunc) getControllerExcludeMethodNameDict() map[string]bool {
	excludeDict := map[string]bool{}

	t := reflect.TypeOf(new(controllers.BaseController))
	methodLen := t.NumMethod()
	for i := 0; i < methodLen; i++ {
		method := t.Method(i)
		methodName := method.Name
		excludeDict[methodName] = true
	}

	return excludeDict
}
