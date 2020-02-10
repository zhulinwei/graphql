# graphql

## 什么是GraphQL
GraphQL是一种用于API的查询语言，它提供了完整且易于理解的数据描述，并使客户可以准确地询问他们需要什么，而不会得到任何多余的东西。它之所以强大，是因为它为我们提供了客户端和服务器都可以理解的共享类型系统，同时还为我们提供了惊人的可重用性。

一个GraphQL服务是通过定义类型和类型上的字段来创建的，如类型示例如下
```graphql

type Query {
    User(id: Int!): User
}

type User {
    id: Int
    name: String
    email: String
    phone: String
    status: UserStatusEnum
}

enum UserStatusEnum {
    EnableUser
    DisableUser
}

```
然后给每个类型上的每个字段提供解析函数，如：
```
function QueyrUser (userId) {
    retuen service.userDao.find(userId)
}
```
最后再通过查询拿到精确的用户数据，如：

```
{
    user(id: 10) {
        name
    }
}
```

在搭建GraphQL服务时，我们需要考虑使用哪个模块，这里以graphql-go和gengql两个Go中的Graphql的实现展开对比。

## graphql-go
我们以上述的类型为例，一步步演示graphql-go中的代码实现：

### 定义类型

```golang
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "user info",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "user id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "username",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "user email",
		},
		"phone": &graphql.Field{
			Type:        graphql.String,
			Description: "user phone",
		},
		"status": &graphql.Field{
			Type:        UserStatusEnumType,
			Description: "user status",
		},
	},
})

var UserStatusEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserStatusEnum",
	Description: "user status enum",
	Values: graphql.EnumValueConfigMap{
		"EnableUser": &graphql.EnumValueConfig{
			Value:       model.EnableStatus,
			Description: "user enable",
		},
		"DisableUser": &graphql.EnumValueConfig{
			Value:       model.DisableStatus,
			Description: "user disable",
		},
	},
})
```
### 执行查询

```golang
var UserField = &graphql.Field{
	Type: UserType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (result interface{}, err error) {
        if id, ok := params.Args["id"].(string); ok {
            retuen service.userDao.find(userId), nil
		}
	},
}
```

## gengql
gengql可以自动生成代码，快速构建应用程序，我们先创建项目框架：
``` sh
# 创建项目文件夹
mkdir gqlgen-demo
cd gqlgen-demo
# 初始化为Go模块
go mod init gqlgen-demo
# 初始化项目框架
go run github.com/99designs/gqlgen init
```

这样就创建了一个空框架，其中包含文件如下：
- server/server.go: GraphQL服务器的入口点
- gqlgen.yml: gqlgen配置文件，用于控制所生成代码的旋钮
- generated.go: GraphQL执行运行时，生成的大部分代码
- models_gen.go: 生成GraphQL所需的模型，通常我们会使用自己的模型来覆盖它们
- resolver.go: 执行查询文件，应用程序代码所在的位置，generated.go将调用此方法以获取用户请求的数据
- schema.graphql: 存放GraphQL类型的文件

### 配置文件
在这里我们依然复用第一节中的Graphql类型示例，将其放入schema.graphql中，同时我们还可以自定义一下生成的文件其存放的路径，如：
```
schema:
- "pkg/schema/**/*.graphql"
exec:
  filename: pkg/util/generated.go
model:
  filename: pkg/model/models_gen.go
resolver:
  filename: pkg/controller/resolver.go
  type: Resolver
autobind: []
```

随后项目更新框架：
```sh
go run github.com/99designs/gqlgen
```

### 添加解析
gengql已经帮我们生成GraphQL服务器所需的基本代码，这个时候需要关注resolver.go文件，其内容如下：


```golang
package controller
 
import (
    "context"
    "graphql/pkg/model"
    "graphql/pkg/util"
)
 
// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.
 
type Resolver struct{}
 
func (r *Resolver) Query() util.QueryResolver {
    return &queryResolver{r}
}
 
type queryResolver struct{ *Resolver }
 
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
    panic("not implemented")
}
```
我们现在只需要添加查询时具体的实现即可，在这里是修改上述文件第19行中方法的内容，如：
```golang
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
    return service.findUserById(id), nil
}
```

## 总结
在使用graphql-go时发现一些比较突出的问题，比如在在定义类型时：
- 代码复杂详细程度高，繁琐且容易出错
- 大量使用接口和反射，使得编译时丧失类型安全性
在执行查询时：
- 我们需要处理自己从map[string]interface{}拆包解析参数
- 依赖关系处理麻烦（无法较好的实现依赖注入）
相比之下gengql在可用性和安全性上更佳：
- 自动生成代码，避免繁杂的类型定义工作
- 使用静态生成所有字段绑定和json序列化，替代反射
- 不需要手动拆包解析参数
- 查询执行文件resolver.go可以方便实现依赖注入
