package goacl

type AclLevel bool

type AccessList []string

const DENY = false

const ALLOW = true

type Acl struct {
	//记录默认访问权限
	_roles          []Role
	_defaultAccess  AclLevel            //记录添加的角色信息
	_rolesName      map[string]bool     //同步roleName到map内, 以便查询
	_resources      []Resource          //同步添加的资源信息
	_resourcesNames map[string]bool     //同步添加的资源信息到map内,以便查询
	_accessList     map[string]bool     //所有的权限列表 map的key为resource!access
	_access         map[string]AclLevel //设定的权限访问类型
}

var empty map[string]bool = map[string]bool{}

func NewAcl() *Acl {
	return &Acl{
		_defaultAccess:  DENY,
		_roles:          []Role{},
		_rolesName:      empty,
		_resourcesNames: empty,
		_resources:      []Resource{},
		_accessList:     empty,
		_access:         map[string]AclLevel{},
	}
}

//设置默认权限
func (a *Acl) SetDefaultAccess(defaultAccess AclLevel) *Acl {
	a._defaultAccess = defaultAccess
	return a
}

//取得默认权限
func (a *Acl) GetDefaultAccess() AclLevel {
	return a._defaultAccess
}

//添加角色是否存在
func (a *Acl) AddRole(role Role) *Acl {
	if !a.IsRole(role) {
		a._roles = append(a._roles, role)
		a._rolesName[role.GetName()] = true
	}
	return a
}

//检查Role是否存在
func (a *Acl) IsRole(role Role) bool {
	for _, v := range a._roles {
		if v.GetName() == role.GetName() {
			return true
		}
	}
	return false
}

//添加资源到ACL
func (a *Acl) AddResource(resource Resource, accessList AccessList) *Acl {
	if !a.IsResource(resource) {
		a._resources = append(a._resources, resource)
		a._resourcesNames[resource.GetName()] = true
	}

	for _, v := range accessList {
		accessKey := resource.GetName() + "!" + v
		a._accessList[accessKey] = true
	}
	return a
}

//从资源中移除动作
func (a *Acl) RemoveResourceAccess(resourceName string, list AccessList) *Acl {
	for _, v := range list {
		accessKey := resourceName + "!" + v
		delete(a._resourcesNames, accessKey)
	}
	return a
}

//为角色设置某个资源的访问权限
func (a *Acl) Allow(roleName, resourceName string, access ... string) {
	a.setAccessLevel(ALLOW, roleName, resourceName, access...)
}

//限制角色访问某个资源的动作权限
func (a *Acl) Deny(roleName, resourceName string, access ... string) {
	a.setAccessLevel(DENY, roleName, resourceName, access...)
}

//设置权限是否可访问
func (a *Acl) setAccessLevel(level AclLevel, roleName, resourceName string, access ... string) {
	a.checkRoleAndResource(roleName, resourceName)
	for _, v := range access {
		resourceListKey := resourceName + "!" + v
		if _, ok := a._accessList[resourceListKey]; !ok {
			panic("the access " + v + " not in resource " + resourceName)
		}
		a._access[roleName+"!"+resourceListKey] = level
	}
}

//检查角色和资源是否存在
func (a *Acl) checkRoleAndResource(roleName, resourceName string) {
	if !a.IsRole(NewRole(roleName, "")) {
		panic("the role :" + roleName + " is not exists in ACL")
	}
	if !a.IsResource(NewResource(resourceName, "")) {
		panic(" the resource :" + resourceName + " is not exists in ACLa")
	}
}

//检查是否存在资源
func (a *Acl) IsResource(resource Resource) bool {
	_, ok := a._resourcesNames[resource.GetName()]
	if ok {
		return true
	}
	return false
}

//检查是否允许访问
func (a *Acl) IsAllowed(roleName, resourceName, access string) bool {
	//如果系统内没有设置角色信息,该角色访问所有的未知操作均为默认权限
	if !a.IsRole(NewRole(roleName, "")) {
		return a.GetDefaultAccess() == ALLOW
	}
	//根据设置的权限进行判定
	accessKey := roleName + "!" + resourceName + "!" + access
	if result, ok := a._access[accessKey]; ok {
		return result == ALLOW
	}
	//最终判定
	return a.GetDefaultAccess() == ALLOW
}
