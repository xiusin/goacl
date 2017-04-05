### 根据Phalcon的Acl组件编写的GoAcl,实现了基础的ACL验证功能

#### 使用方式 
```
go get -u github.com/igoper/goacl
```

```
package main
import "github.com/igoper/goacl"
func main() {
	acl := goacl.NewAcl()
	acl.AddRole(goacl.NewRole("ADMIN", "超级管理员"))
	acl.SetDefaultAccess(goacl.DENY)
	resource := goacl.NewResource("admin", "管理员权限")
	accessList := goacl.AccessList{"index", "edit", "update","delete"}
	acl.AddResource(resource, accessList)
	acl.Allow("ADMIN", "admin", "index", "edit")
	acl.Deny("ADMIN", "admin", "delete")
	allow1 := acl.IsAllowed("ADMIN", "admin", "index") //设置了允许权限
	allow2 := acl.IsAllowed("ADMIN", "admin", "delete")//无法得到验证权限 走默认权限
	allow3 := acl.IsAllowed("ADMIN", "user", "delete")
	allow4 := acl.IsAllowed("MEMBER", "user", "delete")
	fmt.Println(allow1, allow2, allow3, allow4)
}
```
#### 输出 : 
```
//设置默认权限为DENY时
true false false false 
//设置默认权限为ALLow时
true true true true
```
