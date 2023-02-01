# 接口文档

### before

1.所有必选项如未填写，统一返回（之后省略s）：

```json
{
    "status": 400,
    "info": "params error"
}
```

2.所有带请求头的错误返回示例集合（之后省略）：

| status | info                   | 说明                 |
| ------ | ---------------------- | -------------------- |
| 2003   | “请求头中auth为空”     | 请求头中auth为空     |
| 2004   | "请求头中auth格式有误" | 请求头中auth格式有误 |
| 2005   | "无效的Token"          | token已过期          |

## 用户相关

### 注册

**请求路径**：

```http
POST user/register
```

**请求头：**

无

**请求参数：**

| 名称             | 位置 | 类型   | 必选 | 说明     |
| ---------------- | ---- | ------ | ---- | -------- |
| username         | body | string | 是   | 用户名   |
| password         | body | string | 是   | 密码     |
| confirm_password | body | string | 是   | 确认密码 |

**返回参数：**

无

**返回示例：**

| status | info                 | 说明                         |
| ------ | -------------------- | ---------------------------- |
| 200    | “register success”   | 登录成功                     |
| 400    | "different password" | 密码不一致                   |
| 400    | "user has existed"   | 用户名已存在                 |
| 500    | "internal error"     | 数据库增删查改或生成盐值错误 |

```json
{
    "status": 200,
    "info": "register success"
}
```

### 登录

**请求路径**：

```http
POST user/login
```

**请求头：**

无

**请求参数：**

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| username | body | string | 是   | 用户名 |
| password | body | string | 是   | 密码   |

**返回参数：**

| 字段名        | 类型          | 说明      |
| ------------- | ------------- | --------- |
| uid           | string        | 用户id    |
| token         | Bearer $token | 验证token |
| refresh_token | Bearer $token | 刷新token |

**返回示例：**

| status | info               | 说明               |
| ------ | ------------------ | ------------------ |
| 200    | “login success”    | 登录成功           |
| 400    | "user don't exist" | 用户不存在         |
| 400    | “wrong password"   | 密码错误           |
| 500    | "internal error"   | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "login success",
    "data": {
        "uid": 9,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI5IiwiZXhwIjoxNjczODc4OTczLCJpc3MiOiJZSlgifQ.TVAW3nNgqc8ROnBQACOkLxnu1qqJ9DcUGX684-5pHlY",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU0MTEzNzMsImlzcyI6IllKWCJ9.4GhEDBx6agW2wK_o1-gFuCIifRyWqMwLSKR2VXWueJ8"
    }
}
```