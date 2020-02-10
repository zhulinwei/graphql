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

```json
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
缺点：
- 代码复杂详细程度高，繁琐且容易出错
- 大量使用接口和反射，使得编译时丧失类型安全性

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

缺点：
- 我们需要处理自己从map[string]interface{}拆包解析args
- id可能不是字符串
- 依赖关系处理麻烦


## gengql
在定义完GraphQL类型后，gengql可以自动生成代码，快速构建应用程序，在这里我们依然复用第一节中的Graphql类型示例。

### 配置文件
我们首先需要定义gengql自动生成代码的配置文件，示例如下：
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

定义完成后执行命令：
```sh
go run github.com/99designs/gqlgen
```

经过上述步骤gqlgen已经自动帮我们生成好Graphql类型的定义

### 添加解析

我们只需要在resolver文件中添加user的具体实现

```golang
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
    // do something
}

```