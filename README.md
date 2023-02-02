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
POST /user/register
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
POST /user/login
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
### 刷新token

##### 请求路径：

```http
POST /user/auth/refresh
```

**请求头：**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数：**

| 名称          | 位置 | 类型   | 必选 | 说明      |
| ------------- | ---- | ------ | ---- | --------- |
| refresh_token | body | string | 是   | 刷新token |

**返回参数：**

| 字段名        | 类型          | 说明          |
| ------------- | ------------- | ------------- |
| token         | Bearer $token | 新的验证token |
| refresh_token | Bearer $token | 新的刷新token |

##### 返回示例：

| status | info                              | 说明                |
| ------ | --------------------------------- | ------------------- |
| 200    | “refresh token success”           | 刷新令牌成功        |
| 2005   | "无效的Token"                     | refresh_token已过期 |
| 400    | "invalid refresh token signature" | 签名认证错误        |

```json
{
  "status": 200,
  "info": "login success",
  "data": {
    "uid": 20,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjAiLCJyb2xlIjoidXNlciIsImV4cCI6MTY3NTMyNzM5OCwiaXNzIjoiWUpYIn0.Oqrbuk4KCDmL1oSleNIZ-vNkAlqEISyhVz6jTEGwNZ8",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU0MTAxOTgsImlzcyI6IllKWCJ9.NnTH31uDdoyhV3vDWNKAaWZDHkAK7CClipKzl5Hp2yg"
  }
}
```

## 主页相关

### 搜索商品（规则排序）

##### 请求路径：

```http
GET /home/search
```

**请求头：**

无

**请求参数：**

| 名称    | 位置  | 类型   | 必选 | 说明                                           |
| ------- | ----- | ------ | ---- | ---------------------------------------------- |
| keyword | query | string | 是   | 搜索关键词                                     |
| sort_by | query | string | 否   | (price/sales/rating)(价格/销量/评分) 默认sales |
| order   | query | string | 否   | `desc`/`asc`(降序/升序) 默认desc               |

**返回参数：**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 搜索商品信息的集合 |

##### 返回示例：

| status | info                      | 说明               |
| ------ | ------------------------- | ------------------ |
| 200    | “search products success” | 搜索商品成功       |
| 500    | "internal error"          | 数据库增删查改错误 |

```json
{
  "status": 200,
  "info": "search products success",
  "product": [
    {
      "product_id": 2,
      "seller_id": 1,
      "seller": " ",
      "name": "手表",
      "price": "95",
      "sales": "20",
      "rating": "4.6",
      "category": "奢侈品",
      "image": ""
    },
    {
      "product_id": 10,
      "seller_id": 2,
      "seller": "皓皓",
      "name": "手机",
      "price": "4000",
      "sales": "5",
      "rating": "5",
      "category": "电子类",
      "image": " "
    }
  ]
}
```

### 商品分类

##### 请求路径：

```http
GET /home/category
```

**请求头：**

无

**请求参数：**

| 名称     | 位置  | 类型   | 必选 | 说明       |
| -------- | ----- | ------ | ---- | ---------- |
| category | query | string | 是   | 分类关键词 |

**返回参数：**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 分类商品信息的集合 |

##### 返回示例：

| status | info                        | 说明               |
| ------ | --------------------------- | ------------------ |
| 200    | “category products success” | 商品分类成功       |
| 500    | "internal error"            | 数据库增删查改错误 |

```json
{
  "status": 200,
  "info": "category products success",
  "product": [
    {
      "product_id": 10,
      "seller_id": 2,
      "seller": "皓皓",
      "name": "手机",
      "price": "4000",
      "sales": "5",
      "rating": "5",
      "category": "电子类",
      "image": " "
    },
    {
      "product_id": 12,
      "seller_id": 2,
      "seller": " ",
      "name": "ipad",
      "price": "1999.99",
      "sales": "280",
      "rating": "4.7",
      "category": "电子类",
      "image": " "
    }
  ]
}
```



## 复杂数据类型

##### 搜索/分类商品信息的集合

| 名称       | 类型   | 说明          |
| ---------- | ------ | ------------- |
| product_id | int    | 商品id        |
| seller_id  | int    | 卖家id        |
| seller     | string | 卖家          |
| name       | string | 商品名称      |
| price      | string | 商品价格      |
| sales      | string | 商品销量      |
| rating     | string | 商品评分      |
| category   | string | 商品图片`url` |

