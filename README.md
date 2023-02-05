# 接口文档

### before

1.所有必选项如未填写，统一返回（之后省略）：

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

**请求路径**

```http
POST /user/register
```

**请求头**

无

**请求参数**

| 名称             | 位置 | 类型   | 必选 | 说明     |
| ---------------- | ---- | ------ | ---- | -------- |
| username         | body | string | 是   | 用户名   |
| password         | body | string | 是   | 密码     |
| confirm_password | body | string | 是   | 确认密码 |

**返回参数**

无

**返回示例**

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
        "user_id": 20,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjAiLCJyb2xlIjoidXNlciIsImV4cCI6MTY3NTU5MzI3NywiaXNzIjoiWUpYIn0.c7BTz66NqGsfB4xr9EcQamUlrMgGAET3wHrbM_LxZUk",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU1OTMyNzcsImlzcyI6IllKWCJ9.1VbdA9knu256dhcADNZ15HPqu-uxJjuyiQ7owgsf4og"
    }
}
```

### 刷新token

**请求路径**

```http
POST /user/auth/refresh
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称          | 位置 | 类型   | 必选 | 说明      |
| ------------- | ---- | ------ | ---- | --------- |
| refresh_token | body | string | 是   | 刷新token |

**返回参数**

| 字段名        | 类型          | 说明          |
| ------------- | ------------- | ------------- |
| token         | Bearer $token | 新的验证token |
| refresh_token | Bearer $token | 新的刷新token |

**返回示例**

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

**请求路径**

```http
GET /home/search
```

**请求头**

无

**请求参数**

| 名称    | 位置  | 类型   | 必选 | 说明                                                   |
| ------- | ----- | ------ | ---- | ------------------------------------------------------ |
| keyword | query | string | 是   | 搜索关键词                                             |
| sort_by | query | string | 否   | (`price`/`sales`/`rating`)(价格/销量/评分) 默认`sales` |
| order   | query | string | 否   | `desc`/`asc`(降序/升序) 默认`desc`                     |

**返回参数**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 搜索商品信息的集合 |

**返回示例**

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

**请求路径**

```http
GET /home/category
```

**请求头**

无

**请求参数**

| 名称     | 位置  | 类型   | 必选 | 说明       |
| -------- | ----- | ------ | ---- | ---------- |
| category | query | string | 是   | 分类关键词 |

**返回参数**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 分类商品信息的集合 |

**返回示例**

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
            "price": "4000.00",
            "sales": "5",
            "rating": "5.0",
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

## 账户相关

### 充值余额

**请求路径**

```http
POST /auth/individual/recharge/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置  | 类型      | 必选 | 说明     |
| ------- | ----- | --------- | ---- | -------- |
| user_id | path  | string    | 是   | 用户id   |
| account | query | float/int | 是   | 充值金额 |

**返回参数**

无

**返回示例**

| status | info               | 说明                   |
| ------ | ------------------ | ---------------------- |
| 200    | “recharge success” | 充值余额成功           |
| 400    | "invalid charge"   | 充值金额错误（不合法） |
| 500    | "internal error"   | 数据库增删查改错误     |

```json
{
    "status": 200,
    "info": "recharge success"
}
```

### 查看余额

**请求路径**

```http
GET /auth/individual/balance/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名   | 类型   | 说明   |
| -------- | ------ | ------ |
| user_id  | int    | 用户id |
| username | string | 用户名 |
| balance  | float  | 余额   |

**返回示例**

| status | info                   | 说明               |
| ------ | ---------------------- | ------------------ |
| 200    | “view balance success” | 查看余额成功       |
| 500    | "internal error"       | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view balance success",
    "data": {
        "user_id": 20,
        "username": "小3",
        "balance": 40
    }
}
```

## 个人相关

### 修改个人信息

**请求路径**

```http
PUT /auth/individual/modify/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 字段名    | 位置 | 类型   | 必选 | 说明                |
| --------- | ---- | ------ | ---- | ------------------- |
| user_id   | path | string | 是   | 用户id              |
| nickname  | body | string | 否   | 昵称                |
| gender    | body | string | 否   | 性别（0保密1男2女） |
| phone_num | body | string | 否   | 电话号码            |
| year      | body | string | 否   | 生日年份            |
| month     | body | string | 否   | 生日月份            |
| day       | body | string | 否   | 生日哪一天          |
| avatar    | body | string | 否   | 头像`url`           |

**返回参数**

无

**返回示例**

| status | info                         | 说明                 |
| ------ | ---------------------------- | -------------------- |
| 200    | “change information success” | 修改个人信息成功     |
| 400    | "fail to update"             | 所有非必选项均为填写 |
| 500    | "internal error"             | 数据库增删查改错误   |

```json
{
    "status": 200,
    "info": "change information success"
}
```

### 查看个人信息

**请求路径**

```http
GET /auth/individual/information/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名    | 类型   | 说明                |
| --------- | ------ | ------------------- |
| user_id   | int    | 用户id              |
| username  | string | 用户名              |
| nickname  | string | 昵称                |
| gender    | string | 性别（0保密1男2女） |
| phone_num | string | 电话号码            |
| email     | string | 邮箱                |
| year      | string | 生日年份            |
| month     | string | 生日月份            |
| day       | string | 生日哪一天          |
| avatar    | string | 头像`url`           |

**返回示例**

| status | info                       | 说明               |
| ------ | -------------------------- | ------------------ |
| 200    | “view information success” | 查看个人信息成功   |
| 500    | "internal error"           | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view information success",
    "data": {
        "user_id": 18,
        "username": "小1",
        "nickname": "YIYI",
        "gender": "0",
        "phone_num": "15688888888",
        "email": "@qq.com",
        "year": "2023",
        "month": "02",
        "day": "04",
        "avatar": ""
    }
}
```

## 地址相关

### 添加地址

**请求路径**

```http
POST /auth/individual/address/add/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称                | 位置 | 类型   | 必选 | 说明       |
| ------------------- | ---- | ------ | ---- | ---------- |
| user_id             | path | string | 是   | 用户id     |
| name                | body | string | 是   | 收货人姓名 |
| phone               | body | string | 是   | 收货人电话 |
| province            | body | string | 是   | 省份       |
| city                | body | string | 是   | 城市       |
| street_or_community | body | string | 是   | 街道或小区 |

**返回参数**

无

**返回示例**

| status | info                  | 说明               |
| ------ | --------------------- | ------------------ |
| 200    | “add address success” | 添加地址成功       |
| 500    | "internal error"      | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "add address success"
}
```

### 查看地址

**请求路径**

```http
GET /auth/individual/address/view/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名  | 类型         | 说明       |
| ------- | ------------ | ---------- |
| address | 复杂数据类型 | 地址的集合 |

**返回示例**

| status | info                   | 说明               |
| ------ | ---------------------- | ------------------ |
| 200    | “view balance success” | 查看余额成功       |
| 500    | "internal error"       | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view address success",
    "address": [
        {
            "address_id": "3",
            "user_id": "20",
            "recipient_name": "小3",
            "recipient_phone": "17723888888",
            "province": "江苏",
            "city": "盐城",
            "state_or_community": "**街道"
        },
        {
            "address_id": "4",
            "user_id": "20",
            "recipient_name": "小3",
            "recipient_phone": "17723888888",
            "province": "江苏",
            "city": "南京",
            "state_or_community": "**街道"
        }
    ]
}
```

### 修改地址

**请求路径**

```http
PUT /auth/individual/address/update/:address_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称                | 位置 | 类型   | 必选 | 说明       |
| ------------------- | ---- | ------ | ---- | ---------- |
| address_id          | path | string | 是   | 地址id     |
| name                | body | string | 是   | 收货人     |
| phone               | body | string | 是   | 收货电话   |
| province            | body | string | 是   | 省份       |
| city                | body | string | 是   | 城市       |
| street_or_community | body | string | 是   | 街道或小区 |

**返回参数**

无

**返回示例**

| status | info                     | 说明               |
| ------ | ------------------------ | ------------------ |
| 200    | “update address success” | 修改地址成功       |
| 500    | "internal error"         | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "update address success"
}
```

### 删除地址

**请求路径**

```http
DELETE /auth/individual/address/delete/:address_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| address_id | path | string | 是   | 地址id |

**返回参数**

无

**返回示例**

| status | info                     | 说明               |
| ------ | ------------------------ | ------------------ |
| 200    | “delete address success” | 删除地址成功       |
| 500    | "internal error"         | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "delete address success"
}
```

## 商品相关

### 查看商品款式

**请求路径**

```http
GET /product/style/:product_id
```

**请求头**

无

**请求参数**

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| product_id | path | string | 是   | 商品id |

**返回参数**

| 字段名 | 类型         | 说明               |
| ------ | ------------ | ------------------ |
| style  | 复杂数据类型 | 商品款式信息的集合 |

**返回示例**

| status | info                   | 说明               |
| ------ | ---------------------- | ------------------ |
| 200    | “search style success” | 查看款式成功       |
| 500    | "internal error"       | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "search style success",
    "style": [
        {
            "style_id": "1",
            "product_id": "2",
            "product": "手表",
            "style": "红色"
        },
        {
            "style_id": "2",
            "product_id": "2",
            "product": "手表",
            "style": "蓝色"
        }
    ]
}
```

### 查看商品详情

**请求路径**

```http
GET /product/detail/:product_id
```

**请求头**

无

**请求参数**

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| product_id | path | string | 是   | 商品id |

**返回参数**

| 字段名 | 类型         | 说明               |
| ------ | ------------ | ------------------ |
| detail | 复杂数据类型 | 商品详情信息的集合 |

**返回示例**

| status | info                          | 说明               |
| ------ | ----------------------------- | ------------------ |
| 200    | “view product detail success” | 查看商品详情成功   |
| 500    | "internal error"              | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view product detail success",
    "detail": {
        "detail_id": "1",
        "name": "手机",
        "seller": "皓皓",
        "category": "电子类",
        "price": "4000.00",
        "stock": "500",
        "description": "速度快，特别好用",
        "image": " ",
        "product_id": "10"
    }
}
```

## 购物车相关

### 加入购物车

**请求路径**

```http
POST /auth/cart/add/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称     | 位置  | 类型   | 必选 | 说明     |
| -------- | ----- | ------ | ---- | -------- |
| user_id  | path  | string | 是   | 用户id   |
| name     | query | string | 是   | 商品名   |
| quantity | query | string | 是   | 商品数量 |

**返回参数**

无

**返回示例**

| status | info                  | 说明               |
| ------ | --------------------- | ------------------ |
| 200    | “add to cart success” | 加入购物车成功     |
| 400    | "repeated name"       | 商品名重复         |
| 400    | "convert err"         | 数字格式错误       |
| 500    | "internal error"      | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "add to cart success"
}
```

### 查看购物车

**请求路径**

```http
GET /auth/cart/view/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名 | 类型         | 说明         |
| ------ | ------------ | ------------ |
| cart   | 复杂数据类型 | 购物车的集合 |

**返回示例**

| status | info                | 说明               |
| ------ | ------------------- | ------------------ |
| 200    | “view cart success” | 查看购物车成功     |
| 500    | "internal error"    | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view cart success",
    "cart": [
        {
            "cart_id": "2",
            "user_id": "20",
            "name": "手机",
            "unit_price": "4000.00",
            "quantity": "2",
            "price": "8000.00",
            "image": ""
        },
        {
            "cart_id": "3",
            "user_id": "20",
            "name": "洗面奶",
            "unit_price": "70.00",
            "quantity": "5",
            "price": "350.00",
            "image": ""
        }
    ]
}
```

### 删除商品

**请求路径**

```http
DELETE /auth/cart/delete/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置  | 类型   | 必选 | 说明   |
| ------- | ----- | ------ | ---- | ------ |
| user_id | path  | string | 是   | 用户id |
| name    | query | string | 是   | 商品名 |

**返回参数**

无

**返回示例**

| status | info                  | 说明                |
| ------ | --------------------- | ------------------- |
| 200    | “delete cart success” | 删除购物车成功      |
| 400    | "not exist name"      | 商品名不在g购物车中 |
| 500    | "internal error"      | 数据库增删查改错误  |

```json
{
    "status": 200,
    "info": "delete cart success"
}
```

## 评价相关

### 发布评价

**请求路径**

```http
POST /auth/review/add/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称       | 位置 | 类型   | 必选 | 说明     |
| ---------- | ---- | ------ | ---- | -------- |
| user_id    | path | string | 是   | 用户id   |
| product_id | body | string | 是   | 商品id   |
| content    | body | string | 是   | 评价内容 |
| rating     | body | string | 是   | 评分     |

**返回参数**

无

**返回示例**

| status | info                  | 说明               |
| ------ | --------------------- | ------------------ |
| 200    | “give review success” | 发布评价成功       |
| 500    | "internal error"      | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "give review success"
}
```

### 查看评价

**请求路径**

```http
GET /review/view/:product_id
```

**请求头**

无

**请求参数**

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| product_id | path | string | 是   | 商品id |

**返回参数**

| 字段名     | 类型         | 说明               |
| ---------- | ------------ | ------------------ |
| collection | 复杂数据类型 | 商品收藏信息的集合 |

**返回示例**

| status | info                  | 说明               |
| ------ | --------------------- | ------------------ |
| 200    | “view review success” | 查看评价成功       |
| 500    | "internal error"      | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view review success",
    "review": [
        {
            "review_id": "1",
            "user_id": "20",
            "name": "手机",
            "content": "不喜欢",
            "create_time": "2023-02-02T20:57:22+08:00",
            "rating": "3.0",
            "product_id": "10"
        },
        {
            "review_id": "2",
            "user_id": "20",
            "name": "手机",
            "content": "我特别喜欢",
            "create_time": "2023-02-02T21:32:53+08:00",
            "rating": "5.0",
            "product_id": "10"
        }
    ]
}
```

## 收藏相关

### 添加收藏

**请求路径**

```http
POST /auth/collection/add/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置  | 类型   | 必选 | 说明   |
| ------- | ----- | ------ | ---- | ------ |
| user_id | path  | string | 是   | 用户id |
| name    | query | string | 是   | 商品名 |

**返回参数**

无

**返回示例**

| status | info                     | 说明               |
| ------ | ------------------------ | ------------------ |
| 200    | “add collection success” | 添加收藏成功       |
| 400    | "repeated name"          | 收藏的商品名重复   |
| 500    | "internal error"         | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "add collection success"
}
```

### 查看收藏

**请求路径**

```http
GET /auth/collection/view/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名     | 类型         | 说明               |
| ---------- | ------------ | ------------------ |
| collection | 复杂数据类型 | 商品收藏信息的集合 |

**返回示例**

| status | info                      | 说明               |
| ------ | ------------------------- | ------------------ |
| 200    | “view collection success” | 查看收藏成功       |
| 500    | "internal error"          | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view collection success",
    "data": [
        {
            "collection_id": "1",
            "user_id": "10",
            "name": "手机"
        },
        {
            "collection_id": "3",
            "user_id": "10",
            "name": "手表"
        }
    ]
}
```

### 删除收藏

**请求路径**

```http
DELETE /auth/collection/delete/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置  | 类型   | 必选 | 说明   |
| ------- | ----- | ------ | ---- | ------ |
| user_id | path  | string | 是   | 用户id |
| name    | query | string | 是   | 商品名 |

**返回参数**

无

**返回示例**

| status | info                        | 说明               |
| ------ | --------------------------- | ------------------ |
| 200    | “delete collection success” | 删除收藏成功       |
| 400    | "not exist name"            | 商品名不在收藏中   |
| 500    | "internal error"            | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "delete collection success"
}
```

## 订单相关

### 生成订单

**请求路径**

```http
POST /order/auth/add/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称         | 位置  | 类型         | 必选 | 说明               |
| ------------ | ----- | ------------ | ---- | ------------------ |
| user_id      | path  | string       | 是   | 用户id             |
| address_id   | query | string       | 是   | 地址id             |
| pay_products | query | []string数组 | 是   | 订单内商品名的集合 |

**返回参数**

无

**返回示例**

| status | info                 | 说明                         |
| ------ | -------------------- | ---------------------------- |
| 200    | “gen order success”  | 生成订单成功                 |
| 400    | "no product in cart" | 购物车中没有添加要结算的商品 |
| 500    | "internal error"     | 数据库增删查改错误           |

```json
{
    "status": 200,
    "info": "gen order success"
}
```

### 结算订单

**请求路径**

```http
POST /order/auth/settle/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称           | 位置  | 类型   | 必选 | 说明              |
| -------------- | ----- | ------ | ---- | ----------------- |
| user_id        | path  | string | 是   | 用户id            |
| order_id       | query | string | 是   | 订单id            |
| payment        | query | float  | 是   | 支付金额(2位小数) |
| payment_method | query | string | 是   | 支付方式          |

**返回参数**

无

**返回示例**

| status | info                  | 说明                   |
| ------ | --------------------- | ---------------------- |
| 200    | “settle bill success” | 结算订单成功           |
| 400    | "balance not enough"  | 余额不足               |
| 400    | "wrong payment"       | 支付金额不合法         |
| 400    | "wrong total price"   | 支付金额与应付金额不等 |
| 500    | "internal error"      | 数据库增删查改错误     |

```json
{
    "status": 200,
    "info": "settle bill success"
}
```

### 查看所有订单

**请求路径**

```http
GET /order/auth/view/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

**返回参数**

| 字段名       | 类型         | 说明           |
| ------------ | ------------ | -------------- |
| order        | 复杂数据类型 | 订单的集合     |
| order_detail | 复杂数据类型 | 订单明细的集合 |

**返回示例**

| status | info                 | 说明               |
| ------ | -------------------- | ------------------ |
| 200    | “view order success” | 查看收藏成功       |
| 500    | "internal error"     | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view order success",
    "order": [
        {
            "order_id": "2",
            "order_number": "1675523974c1d12fbb8d79c0cd4a70",
            "order_time": "2023-02-04T23:19:35+08:00",
            "status": "待支付",
            "payment_method": "",
            "payment_amount": 20350,
            "payment_time": "0001-01-01T00:00:00Z",
            "recipient_name": "小3",
            "recipient_address": "江苏南京**街道",
            "recipient_phone": "17723888866",
            "user_id": "20"
        },
        {
            "order_id": "3",
            "order_number": "167557834539164d579eebb18a36e0",
            "order_time": "2023-02-05T14:25:45+08:00",
            "status": "待支付",
            "payment_method": "",
            "payment_amount": 20350,
            "payment_time": "0001-01-01T00:00:00Z",
            "recipient_name": "小3",
            "recipient_address": "江苏南京**街道",
            "recipient_phone": "17723888866",
            "user_id": "20"
        }
    ],
    "order_detail": [
        [
            {
                "order_detail_id": "2",
                "order_id": "2",
                "name": "手机",
                "quantity": "5",
                "price": "20000.00"
            },
            {
                "order_detail_id": "3",
                "order_id": "2",
                "name": "洗面奶",
                "quantity": "5",
                "price": "350.00"
            }
        ],
        [
            {
                "order_detail_id": "4",
                "order_id": "3",
                "name": "手机",
                "quantity": "5",
                "price": "20000.00"
            },
            {
                "order_detail_id": "5",
                "order_id": "3",
                "name": "洗面奶",
                "quantity": "5",
                "price": "350.00"
            }
        ]
    ]
}
```

### 更新订单状态

**请求路径**

```http
PUT /order/auth/update/:order_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称     | 位置  | 类型   | 必选 | 说明             |
| -------- | ----- | ------ | ---- | ---------------- |
| order_id | path  | string | 是   | 订单id           |
| status   | query | string | 是   | 改变后的订单状态 |

**返回参数**

无

**返回示例**

| status | info                          | 说明               |
| ------ | ----------------------------- | ------------------ |
| 200    | “update order status success” | 修改订单状态成功   |
| 500    | "internal error"              | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "update order status success"
}
```

### 按订单状态查看

**请求路径**

```http
GET /order/auth/search/:user_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称    | 位置  | 类型   | 必选 | 说明     |
| ------- | ----- | ------ | ---- | -------- |
| user_id | path  | string | 是   | 用户id   |
| status  | query | string | 是   | 订单状态 |

**返回参数**

| 字段名       | 类型         | 说明           |
| ------------ | ------------ | -------------- |
| order        | 复杂数据类型 | 订单的集合     |
| order_detail | 复杂数据类型 | 订单明细的集合 |

**返回示例**

| status | info                   | 说明               |
| ------ | ---------------------- | ------------------ |
| 200    | “search order success” | 商品分类成功       |
| 500    | "internal error"       | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "search order success",
    "order": [
        {
            "order_id": "2",
            "order_number": "1675523974c1d12fbb8d79c0cd4a70",
            "order_time": "2023-02-04T23:19:35+08:00",
            "status": "待支付",
            "payment_method": "",
            "payment_amount": 20350,
            "payment_time": "0001-01-01T00:00:00Z",
            "recipient_name": "小3",
            "recipient_address": "江苏南京**街道",
            "recipient_phone": "17723888866",
            "user_id": "20"
        },
        {
            "order_id": "3",
            "order_number": "167557834539164d579eebb18a36e0",
            "order_time": "2023-02-05T14:25:45+08:00",
            "status": "待支付",
            "payment_method": "",
            "payment_amount": 20350,
            "payment_time": "0001-01-01T00:00:00Z",
            "recipient_name": "小3",
            "recipient_address": "江苏南京**街道",
            "recipient_phone": "17723888866",
            "user_id": "20"
        }
    ],
    "order_detail": [
        [
            {
                "order_detail_id": "2",
                "order_id": "2",
                "name": "手机",
                "quantity": "5",
                "price": "20000.00"
            },
            {
                "order_detail_id": "3",
                "order_id": "2",
                "name": "洗面奶",
                "quantity": "5",
                "price": "350.00"
            }
        ],
        [
            {
                "order_detail_id": "4",
                "order_id": "3",
                "name": "手机",
                "quantity": "5",
                "price": "20000.00"
            },
            {
                "order_detail_id": "5",
                "order_id": "3",
                "name": "洗面奶",
                "quantity": "5",
                "price": "350.00"
            }
        ]
    ]
}
```

### 删除订单

**请求路径**

```http
DELETE /order/auth/delete/:order_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| order_id | path | string | 是   | 订单id |

**返回参数**

无

**返回示例**

| status | info                   | 说明               |
| ------ | ---------------------- | ------------------ |
| 200    | “delete order success” | 删除订单成功       |
| 500    | "internal error"       | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "delete order success"
}
```

## 店铺相关

### 注册

**请求路径**

```http
POST /seller/register
```

**请求头**

无

**请求参数**

| 名称             | 位置 | 类型   | 必选 | 说明     |
| ---------------- | ---- | ------ | ---- | -------- |
| seller           | body | string | 是   | 用户名   |
| password         | body | string | 是   | 密码     |
| confirm_password | body | string | 是   | 确认密码 |

**返回参数**

无

**返回示例**

| status | info                 | 说明           |
| ------ | -------------------- | -------------- |
| 200    | “register success”   | 注册成功       |
| 400    | "different password" | 密码不一致     |
| 400    | "seller has existed" | 卖家名已存在   |
| 500    | "internal error"     | 数据库增删查改 |

```json
{
    "status": 200,
    "info": "register success"
}
```

### 登录

**请求路径**

```http
POST /seller/login
```

**请求头**

无

**请求参数**

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| username | body | string | 是   | 用户名 |
| password | body | string | 是   | 密码   |

**返回参数**

| 字段名        | 类型          | 说明      |
| ------------- | ------------- | --------- |
| uid           | string        | 用户id    |
| token         | Bearer $token | 验证token |
| refresh_token | Bearer $token | 刷新token |

**返回示例**

| status | info               | 说明               |
| ------ | ------------------ | ------------------ |
| 200    | “login success”    | 登录成功           |
| 400    | "seller not exist" | 用户不存在         |
| 400    | “wrong password"   | 密码错误           |
| 500    | "internal error"   | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "login success",
    "data": {
        "seller_id": 4,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiNCIsInJvbGUiOiJzZWxsZXIiLCJleHAiOjE2NzU1OTMxMTQsImlzcyI6IllKWCJ9.r9Ll_xqZIM5oX3zgcj48dgy-XEIMZGIm1H0345W9_5o",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU1OTMxMTQsImlzcyI6IllKWCJ9.7wRUQ5IwheEKp_qrmUdWdQBarZ7nXUgl2wkptGfINHs"
    }
}
```

### 刷新token

**请求路径**

```http
POST /seller/auth/refresh
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称          | 位置 | 类型   | 必选 | 说明      |
| ------------- | ---- | ------ | ---- | --------- |
| refresh_token | body | string | 是   | 刷新token |

**返回参数**

| 字段名        | 类型          | 说明          |
| ------------- | ------------- | ------------- |
| token         | Bearer $token | 新的验证token |
| refresh_token | Bearer $token | 新的刷新token |

**返回示例**

| status | info                              | 说明                |
| ------ | --------------------------------- | ------------------- |
| 200    | “refresh token success”           | 刷新令牌成功        |
| 2005   | "无效的Token"                     | refresh_token已过期 |
| 400    | "invalid refresh token signature" | 签名认证错误        |

```json
{
    "status": 200,
    "info": "refresh token success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiIiwicm9sZSI6IiIsImV4cCI6MTY3NTU5MDY3OCwiaXNzIjoiWUpYIn0.rgizSWRrU83W0I-AgOS-AEL7cgcRCoLMNDe72aW05xc",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU1OTA2NzgsImlzcyI6IllKWCJ9.F-rsyfFQsJJhf3EneKzWgU1sWZ2fJid6E-VDEUdx4Dw"
    }
}
```

### 上架商品

**请求路径**

```http
POST /seller/auth/add/:seller_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称      | 位置 | 类型   | 必选 | 说明          |
| --------- | ---- | ------ | ---- | ------------- |
| seller_id | path | string | 是   | 卖家id        |
| name      | body | string | 是   | 商品名        |
| price     | body | string | 是   | 商品价格      |
| sales     | body | string | 是   | 商品销量      |
| rating    | body | string | 是   | 商品评分      |
| category  | body | string | 是   | 商品种类      |
| image     | body | string | 否   | 商品图片`url` |

**返回参数**

无

**返回示例**

| status | info                  | 说明               |
| ------ | --------------------- | ------------------ |
| 200    | “add product success” | 上架商品成功       |
| 400    | "invalid seller id"   | 卖家id不合法       |
| 400    | "product has existed" | 商品重复添加       |
| 500    | "internal error"      | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "add product success"
}
```

### 查看上架商品

**请求路径**

```http
GET /seller/auth/view/:seller_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称      | 位置 | 类型   | 必选 | 说明   |
| --------- | ---- | ------ | ---- | ------ |
| seller_id | path | string | 是   | 卖家id |

**返回参数**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 搜索商品信息的集合 |

**返回示例**

| status | info                    | 说明               |
| ------ | ----------------------- | ------------------ |
| 200    | “view products success” | 查看上架商品成功   |
| 500    | "internal error"        | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view products success",
    "product": [
        {
            "product_id": 14,
            "seller_id": 3,
            "seller": "小3",
            "name": "羽毛球",
            "price": "50.00",
            "sales": "688",
            "rating": "4.1",
            "category": "体育类",
            "image": " "
        }
    ]
}
```

### 修改商品信息

**请求路径**

```http
PUT /seller/auth/update/:seller_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称      | 位置 | 类型   | 必选 | 说明          |
| --------- | ---- | ------ | ---- | ------------- |
| seller_id | path | string | 是   | 卖家id        |
| name      | body | string | 是   | 商品名        |
| price     | body | string | 否   | 商品价格      |
| sales     | body | string | 否   | 商品销量      |
| rating    | body | string | 否   | 商品评分      |
| category  | body | string | 否   | 商品种类      |
| image     | body | string | 否   | 商品图片`url` |

**返回参数**

无

**返回示例**

| status | info                     | 说明                 |
| ------ | ------------------------ | -------------------- |
| 200    | “update product success” | 修改商品信息成功     |
| 400    | "unknown name"           | 商品名未输入         |
| 400    | "invalid seller id"      | 卖家id不合法         |
| 400    | "product not exist"      | 商品名不存在         |
| 400    | "fail to update"         | 所有非必选项均未输入 |
| 500    | "internal error"         | 数据库增删查改错误   |

```json
{
    "status": 200,
    "info": "update product success"
}
```

### 下架商品

**请求路径**

```http
DELETE /seller/auth/delete/:seller_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称      | 位置  | 类型   | 必选 | 说明   |
| --------- | ----- | ------ | ---- | ------ |
| seller_id | path  | string | 是   | 卖家id |
| name      | query | string | 是   | 商品名 |

**返回参数**

无

**返回示例**

| status | info                     | 说明               |
| ------ | ------------------------ | ------------------ |
| 200    | “delete product success” | 下架商品成功       |
| 400    | "product not exist"      | 商品名不存在       |
| 500    | "internal error"         | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "delete product success"
}
```

## 店铺详情相关

### 商品排序

**请求路径**

```http
GET /store/sort/:seller_id
```

**请求头**

无

**请求参数**

| 名称      | 位置  | 类型   | 必选 | 说明                                                   |
| --------- | ----- | ------ | ---- | ------------------------------------------------------ |
| seller_id | path  | string | 是   | 卖家id                                                 |
| sort_by   | query | string | 否   | (`price`/`sales`/`rating`)(价格/销量/评分) 默认`sales` |
| order     | query | string | 否   | `desc`/`asc`(降序/升序) 默认`desc`                     |

**返回参数**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 搜索商品信息的集合 |

**返回示例**

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
            "product_id": 10,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "手机",
            "price": "4000.00",
            "sales": "5",
            "rating": "5.0",
            "category": "电子类",
            "image": " "
        },
        {
            "product_id": 11,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "洗面奶",
            "price": "70.00",
            "sales": "66",
            "rating": "4.0",
            "category": "日用品",
            "image": ""
        },
        {
            "product_id": 12,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "ipad",
            "price": "1999.99",
            "sales": "280",
            "rating": "4.7",
            "category": "电子类",
            "image": " "
        },
        {
            "product_id": 13,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "牙刷",
            "price": "9.99",
            "sales": "404",
            "rating": "4.6",
            "category": "日用品",
            "image": " "
        }
    ]
}
```

### 商品分类

**请求路径**

```http
GET /store/category/:seller_id
```

**请求头**

无

**请求参数**

| 名称      | 位置  | 类型   | 必选 | 说明       |
| --------- | ----- | ------ | ---- | ---------- |
| seller_id | path  | string | 是   | 卖家id     |
| category  | query | string | 是   | 分类关键词 |

**返回参数**

| 字段名  | 类型         | 说明               |
| ------- | ------------ | ------------------ |
| product | 复杂数据类型 | 分类商品信息的集合 |

**返回示例**

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
            "product_id": 11,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "洗面奶",
            "price": "70.00",
            "sales": "66",
            "rating": "4.0",
            "category": "日用品",
            "image": ""
        },
        {
            "product_id": 13,
            "seller_id": 2,
            "seller": "皓皓",
            "name": "牙刷",
            "price": "9.99",
            "sales": "404",
            "rating": "4.6",
            "category": "日用品",
            "image": " "
        }
    ]
}
```

## 店铺公告相关

### 更新公告

**请求路径**

```http
PUT /auth/announcement/update/:seller_id
```

**请求头**

| 字段名        | 必选 | 类型          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

**请求参数**

| 名称      | 位置 | 类型   | 必选 | 说明     |
| --------- | ---- | ------ | ---- | -------- |
| seller_id | path | string | 是   | 卖家id   |
| title     | body | string | 是   | 公告标题 |
| content   | body | string | 是   | 公告内容 |

**返回参数**

无

**返回示例**

| status | info                          | 说明               |
| ------ | ----------------------------- | ------------------ |
| 200    | “update announcement success” | 更新公告成功       |
| 500    | "internal error"              | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "update announcement success"
}
```

### 查看公告

**请求路径**

```http
GET /announcement/view/:seller_id
```

**请求头**

无

**请求参数**

| 名称      | 位置 | 类型   | 必选 | 说明   |
| --------- | ---- | ------ | ---- | ------ |
| seller_id | path | string | 是   | 卖家id |

**返回参数**

| 名称            | 类型   | 必选 | 说明     |
| --------------- | ------ | ---- | -------- |
| announcement_id | string | 是   | 公告id   |
| seller_id       | string | 是   | 卖家id   |
| title           | string | 是   | 公告标题 |
| content         | string | 是   | 公告内容 |

**返回示例**

| status | info                        | 说明               |
| ------ | --------------------------- | ------------------ |
| 200    | “view announcement success” | 查看公告成功       |
| 500    | "internal error"            | 数据库增删查改错误 |

```json
{
    "status": 200,
    "info": "view announcement success",
    "data": {
        "announcement_id": "1",
        "seller_id": "3",
        "title": "2.3最新公告",
        "content": "我特别喜欢"
    }
}
```



## 复杂数据类型

##### 搜索/分类/查看商品信息的集合

| 名称       | 类型   | 说明          |
| ---------- | ------ | ------------- |
| product_id | int    | 商品id        |
| seller_id  | int    | 卖家id        |
| seller     | string | 卖家          |
| name       | string | 商品名称      |
| price      | string | 商品价格      |
| sales      | string | 商品销量      |
| rating     | string | 商品评分      |
| category   | string | 商品种类      |
| image      | string | 商品图片`url` |

##### 商品款式信息的集合

| 名称       | 类型   | 说明   |
| ---------- | ------ | ------ |
| style_id   | string | 款式id |
| product_id | string | 商品id |
| product    | string | 商品名 |
| style      | string | 款式   |

##### 商品详情信息的集合

| 名称        | 类型   | 说明          |
| ----------- | ------ | ------------- |
| detail_id   | string | 商品详情id    |
| name        | string | 商品名        |
| seller      | string | 卖家          |
| category    | string | 商品类别      |
| price       | string | 商品价格      |
| stock       | string | 商品库存      |
| description | string | 商品描述      |
| image       | string | 商品图片`url` |
| product_id  | string | 商品id        |

##### 商品评价信息的集合

| 名称        | 类型   | 说明     |
| ----------- | ------ | -------- |
| review_id   | string | 评价id   |
| user_id     | string | 用户id   |
| name        | string | 商品名   |
| content     | string | 评价内容 |
| create_time | string | 评价时间 |
| rating      | string | 评分     |
| product_id  | string | 商品id   |

##### 商品收藏信息的集合

| 名称          | 类型   | 说明   |
| ------------- | ------ | ------ |
| collection_id | string | 收藏id |
| user_id       | string | 用户id |
| name          | string | 商品名 |

##### 地址的集合

| 名称                | 类型   | 说明       |
| ------------------- | ------ | ---------- |
| address_id          | string | 评价id     |
| user_id             | string | 用户id     |
| recipient_name      | string | 收货人     |
| recipient_phone     | string | 收获电话   |
| province            | string | 省份       |
| city                | string | 城市       |
| street_or_community | string | 街道或小区 |

##### 购物车的集合

| 名称       | 类型   | 说明          |
| ---------- | ------ | ------------- |
| cart_id    | string | 购物车id      |
| user_id    | string | 用户id        |
| name       | string | 商品名        |
| unit_price | string | 商品单价      |
| quantity   | string | 数量          |
| price      | string | 商品总价      |
| image      | string | 商品图片`url` |

##### 订单的集合

| 名称              | 类型   | 说明                                           |
| ----------------- | ------ | ---------------------------------------------- |
| order_id          | string | 订单id                                         |
| order_number      | string | 订单号                                         |
| order_time        | string | 下单时间                                       |
| status            | string | 订单状态                                       |
| payment_method    | string | 支付方式                                       |
| payment_amount    | float  | 应付金额                                       |
| payment_time      | string | 支付时间（ "0001-01-01T00:00:00Z" 表示未支付） |
| recipient_name    | string | 收货人                                         |
| recipient_address | string | 收货地址                                       |
| recipient_phone   | string | 收货电话                                       |
| user_id           | string | 用户id                                         |

##### 订单明细的集合

| 名称            | 类型   | 说明       |
| --------------- | ------ | ---------- |
| order_detail_id | string | 订单明细id |
| order_id        | string | 订单id     |
| name            | string | 商品名     |
| quantity        | string | 购买数量   |
| price           | string | 该商品总价 |

