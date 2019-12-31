# graphql

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