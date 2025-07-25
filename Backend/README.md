# BBS

## api文档

### 登录

#### URL: api/auth/login  POST

请求体---json

```json
{
  "username": "admin",
  "password": "123456"
}
```

响应示例

```json
{
  "code": 200,
  "data": {
    "Username": "admin",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTM1MTY4NTMsInVzZXJuYW1lIjoiYWRtaW4ifQ.7V0IYWjZ4kRkPf3267ozNmA6qyVFrxkFo7YU9GakjEI"
  },
  "msg": "登录成功！"
}
```

###注册

#### URL: api/auth/register POST

请求体---json

```json
{
  "username": "admin",
  "password": "123456"
}

```

响应示例

```json
{
  "code": 200,
  "data": {
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "username": "admin",
    "password": "admin"
  },
  "msg": "注册成功！"
}
```

###获取分类列表

#### URL: /api/cate/getAllCate GET


响应示例

```json
{
  "code": 200,
  "data": [
    {
      "ID": 1,
      "Name": "科技",
      "Icon": "🚗"
    },
    {
      "ID": 2,
      "Name": "游戏",
      "Icon": "🎮"
    }
  ],
  "msg": "获取成功"
}
```