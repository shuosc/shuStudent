# API Reference

## 模型

### 学生基本信息

```json
{
  "id":   学生证号,
  "name": 姓名,
  "mail": 邮箱,
  "phoneNumber": 手机号
}
```

## web api

- `GET /ping`

  检查服务是否可用，应该直接返回`pong`。

- `GET /student`

  获得jwt token指示的学生的信息
  
- `POST /student`、`PUT /student`
 
  设置学生信息