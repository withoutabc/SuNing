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
    "uid": 9,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI5IiwiZXhwIjoxNjczODc4OTczLCJpc3MiOiJZSlgifQ.TVAW3nNgqc8ROnBQACOkLxnu1qqJ9DcUGX684-5pHlY",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDU0MTEzNzMsImlzcyI6IllKWCJ9.4GhEDBx6agW2wK_o1-gFuCIifRyWqMwLSKR2VXWueJ8"
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

### 查看余额

## 个人相关

### 修改个人信息

### 查看个人信息

## 地址相关

### 添加地址

### 查看地址

### 修改地址

### 删除地址

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

### 查看购物车

### 删除商品

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

## 店铺相关

### 注册

### 登录

### 刷新token

### 上架商品

### 查看上架商品

### 修改商品信息

### 删除商品

## 店铺详情相关

### 商品分类

### 商品排序

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
| title     | body | string | 否   | 公告标题 |
| content   | body | string | 否   | 公告内容 |

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