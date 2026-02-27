# Go User Center 后端说明

这是一个用于学习 Go Web 开发的后端项目，提供用户中心基础能力：注册、登录、查询个人信息、更新资料、修改密码。

## 一、技术栈

- **Go 1.22**
- **Gin**：HTTP Web 框架
- **GORM**：ORM
- **MySQL 8**：数据存储（当前配置复用 `mysql-production`）
- **JWT**：鉴权与登录态（Bearer Token）
- **Viper**：配置加载（`config.yaml` + 环境变量）
- **bcrypt**：密码哈希加密（`golang.org/x/crypto/bcrypt`）
- **Docker**：容器化部署

## 二、项目结构

```text
backend/
├── cmd/
│   └── server/
│       └── main.go                  # 程序入口、路由注册、中间件、启动服务
├── internal/
│   ├── config/
│   │   └── config.go                # 配置读取（Viper + Env）
│   ├── db/
│   │   └── db.go                    # MySQL 连接与连接池
│   ├── middleware/
│   │   └── auth.go                  # JWT 鉴权中间件
│   ├── modules/
│   │   ├── auth/
│   │   │   └── handler.go           # 注册/登录接口
│   │   └── user/
│   │       ├── handler.go           # 用户资料、改密接口
│   │       └── model.go             # User 模型
│   └── utils/
│       ├── jwt.go                   # JWT 签发与解析
│       └── password.go              # bcrypt 密码工具
├── .env.example                     # 环境变量示例
├── config.yaml                      # 默认配置（可被环境变量覆盖）
├── Dockerfile                       # 后端镜像构建
├── go.mod
└── go.sum
```

## 三、接口概览

基础前缀：`/api`

### 1) 注册
- `POST /api/auth/register`
- Body:
```json
{
  "email": "user@example.com",
  "password": "abc12345",
  "nickname": "tester"
}
```

### 2) 登录
- `POST /api/auth/login`
- Body:
```json
{
  "email": "user@example.com",
  "password": "abc12345"
}
```
- 返回 `accessToken`

### 3) 获取当前用户信息
- `GET /api/me`
- Header: `Authorization: Bearer <accessToken>`

### 4) 更新个人资料
- `PUT /api/me`
- Header: `Authorization: Bearer <accessToken>`
- Body（可选字段）：
```json
{
  "nickname": "新昵称",
  "avatarUrl": "https://example.com/avatar.png",
  "bio": "个人简介"
}
```

### 5) 修改密码
- `PUT /api/me/password`
- Header: `Authorization: Bearer <accessToken>`
- Body:
```json
{
  "oldPassword": "old-pass",
  "newPassword": "new-pass-123"
}
```

## 四、配置说明

项目支持 `config.yaml` + 环境变量覆盖。

关键环境变量：

- `DB_DSN`：MySQL 连接串（必填）
- `SERVER_ADDR`：监听地址（默认 `:18601`）
- `JWT_SECRET`：JWT 密钥（生产环境请使用高强度随机值）
- `JWT_ISSUER`：JWT 签发者
- `JWT_ACCESSTTL`：Access Token 有效期（如 `30m`）

示例：

```env
DB_DSN=go_uc:change_me@tcp(mysql-production:3306)/go_user_center?charset=utf8mb4&parseTime=True&loc=Local
SERVER_ADDR=:18601
JWT_SECRET=replace_with_strong_secret
JWT_ISSUER=go-user-center
JWT_ACCESSTTL=30m
```

## 五、本地运行

```bash
cd backend
go mod tidy
DB_DSN="go_uc:change_me@tcp(mysql-production:3306)/go_user_center?charset=utf8mb4&parseTime=True&loc=Local" \
JWT_SECRET="replace_with_strong_secret" \
SERVER_ADDR=":18601" \
go run ./cmd/server
```

## 六、Docker 运行

建议在项目根目录通过 `docker-compose` 启动（已与前端联动）。

```bash
cd /root/go-user-center
docker-compose up -d --build
```

后端端口：`18601`

---

如果你后续想继续学习，我可以再给这套后端加：
- 统一错误码体系
- 参数校验错误国际化
- 日志链路追踪（request id）
- 刷新 token / 退出登录黑名单
- 单元测试与集成测试模板
