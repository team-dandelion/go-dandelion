# MongoDB 使用示例

## 概述

参照 MySQL 的实现方式，完成了 MongoDB 的初始化封装。现在你可以通过 application 包来使用 MongoDB。

**重要说明**: MongoDB 使用**副本集（Replica Set）**架构，而不是传统的主从复制。

## 配置示例

### 1. 单节点配置

```yaml
mongo:
  max_pool_size: 100              # 最大连接池大小
  min_pool_size: 10               # 最小连接池大小  
  max_conn_idle_time: 1800        # 最大连接空闲时间（秒）
  connect_timeout: "10s"          # 连接超时时间
  socket_timeout: "30s"           # Socket 超时时间
  server_selection_timeout: "30s" # 服务器选择超时时间
  log_level: 1                    # 日志级别
  slow_threshold: "100ms"         # 慢查询阈值
  hosts:                          # MongoDB主机列表
    - "localhost:27017"
  database: "myapp_db"            # 数据库名
  username: "mongo_user"          # 用户名
  password: "mongo_password"      # 密码
  auth_db: "admin"                # 认证数据库
  read_preference: "primary"      # 读偏好
```

### 2. 副本集配置

```yaml
mongo:
  max_pool_size: 100
  min_pool_size: 10
  max_conn_idle_time: 1800
  connect_timeout: "10s"
  socket_timeout: "30s"
  server_selection_timeout: "30s"
  log_level: 1
  slow_threshold: "100ms"
  hosts:                          # 副本集所有节点
    - "mongo1:27017"
    - "mongo2:27017"
    - "mongo3:27017"
  database: "myapp_db"
  username: "mongo_user"
  password: "mongo_password"
  auth_db: "admin"
  replica_set: "rs0"              # 副本集名称
  read_preference: "secondaryPreferred" # 读偏好：优先从Secondary读取
```

## 读偏好说明

MongoDB支持以下读偏好设置：

- **primary**: 只从 Primary 读取（默认）
- **primaryPreferred**: 优先从 Primary 读取，Primary 不可用时从 Secondary 读取
- **secondary**: 只从 Secondary 读取
- **secondaryPreferred**: 优先从 Secondary 读取，Secondary 不可用时从 Primary 读取
- **nearest**: 从延迟最低的节点读取

## 基础使用

### 1. 使用基础 MongoDB 连接

```go
package main

import (
    "context"
    "fmt"
    "github.com/team-dandelion/go-dandelion/application"
    "go.mongodb.org/mongo-driver/bson"
)

func main() {
    // 初始化应用（包含MongoDB初始化）
    application.Init()
    
    // 创建 MongoDB 实例
    mongoDB := &application.MongoDB{}
    
    // 获取客户端
    client := mongoDB.GetClient()
    if client == nil {
        fmt.Println("无法连接到MongoDB")
        return
    }
    
    // 获取数据库
    db := mongoDB.GetDatabase("myapp_db")
    if db == nil {
        fmt.Println("无法连接到MongoDB数据库")
        return
    }
    
    // 获取集合
    collection := mongoDB.GetCollection("myapp_db", "users")
    if collection == nil {
        fmt.Println("无法获取集合")
        return
    }
    
    // 插入文档
    user := bson.M{
        "name":  "张三",
        "email": "zhangsan@example.com",
        "age":   25,
    }
    
    result, err := collection.InsertOne(context.TODO(), user)
    if err != nil {
        fmt.Printf("插入失败: %v\n", err)
        return
    }
    
    fmt.Printf("插入成功，ID: %v\n", result.InsertedID)
    
    // 查询文档
    var foundUser bson.M
    err = collection.FindOne(context.TODO(), bson.M{"name": "张三"}).Decode(&foundUser)
    if err != nil {
        fmt.Printf("查询失败: %v\n", err)
        return
    }
    
    fmt.Printf("查询结果: %+v\n", foundUser)
    
    // 关闭连接
    mongoDB.Close()
}
```

### 2. 使用应用级 MongoDB 管理

```go
package main

import (
    "fmt"
    "github.com/team-dandelion/go-dandelion/application"
    "github.com/team-dandelion/go-dandelion/database/mongo"
)

func main() {
    // 初始化应用
    application.Init()
    
    // 定义应用配置获取函数
    configFunc := func(appKey ...string) ([]mongo.AppConfig, error) {
        return []mongo.AppConfig{
            {
                AppKey:     "shop",
                MongoHosts: []string{"localhost:27017"},
                Database:   "shop_db",
                Username:   "shop_user",
                Password:   "shop_pass",
                AuthDB:     "admin",
            },
            {
                AppKey:     "user",
                MongoHosts: []string{"mongo1:27017", "mongo2:27017", "mongo3:27017"},
                Database:   "user_db",
                Username:   "user_user",
                Password:   "user_pass",
                AuthDB:     "admin",
                ReplicaSet: "rs0",
            },
        }, nil
    }
    
    // 定义应用变更处理函数
    changeFunc := func(appKey string, changeType mongo.ChangeType) error {
        fmt.Printf("应用 %s 配置发生变更，类型: %d\n", appKey, changeType)
        // 这里可以处理配置变更后的业务逻辑
        return nil
    }
    
    // 初始化应用级MongoDB管理
    application.InitAppMongoDB(configFunc, changeFunc)
    
    // 创建 MongoDB 实例
    mongoDB := &application.MongoDB{}
    
    // 使用不同应用的MongoDB连接
    shopClient := mongoDB.GetClient("shop")
    if shopClient != nil {
        fmt.Println("成功获取shop应用的MongoDB客户端")
    }
    
    userDB := mongoDB.GetDatabase("user_db", "user")
    if userDB != nil {
        fmt.Println("成功获取user应用的数据库")
    }
    
    shopCollection := mongoDB.GetCollection("shop_db", "products", "shop")
    if shopCollection != nil {
        fmt.Println("成功获取shop应用的商品集合")
    }
}
```

### 3. 配置变更通知

```go
package main

import (
    "github.com/team-dandelion/go-dandelion/application"
    "github.com/team-dandelion/go-dandelion/database/mongo"
)

func main() {
    // 初始化应用
    application.Init()
    
    // 上报应用MongoDB配置变更（通常在管理后台调用）
    err := application.AppMongoDBChange("shop", mongo.Updated)
    if err != nil {
        panic(err)
    }
    
    // 其他服务订阅到变更消息后，会自动重新连接
}
```

## 配置说明

### 必需配置

- `hosts`: MongoDB主机列表，格式为 ["host:port"]
- `database`: 数据库名
- `username`: MongoDB用户名（如果启用认证）
- `password`: MongoDB密码（如果启用认证）

### 可选配置

- `auth_db`: 认证数据库，默认为 "admin"
- `replica_set`: 副本集名称，使用副本集时必须配置
- `read_preference`: 读偏好设置
- `max_pool_size`: 最大连接池大小，默认100
- `min_pool_size`: 最小连接池大小，默认10
- `max_conn_idle_time`: 最大连接空闲时间（秒），默认1800
- `connect_timeout`: 连接超时时间，默认"10s"
- `socket_timeout`: Socket超时时间，默认"30s"
- `server_selection_timeout`: 服务器选择超时时间，默认"30s"
- `log_level`: 日志级别，1=Silent, 2=Error, 3=Warn, 4=Info
- `slow_threshold`: 慢查询阈值，默认"100ms"

## MongoDB 架构说明

### 副本集 vs 主从复制

| 特性 | MySQL 主从 | MongoDB 副本集 |
|------|-----------|---------------|
| 架构模式 | 主从复制 | 副本集 |
| 故障转移 | 手动 | 自动 |
| 读写分离 | 手动配置 | 读偏好设置 |
| 数据一致性 | 最终一致 | 强一致性（可配置） |
| 配置复杂度 | 中等 | 简单 |

### 连接字符串示例

```
# 单节点
mongodb://username:password@localhost:27017/database?authSource=admin

# 副本集
mongodb://username:password@mongo1:27017,mongo2:27017,mongo3:27017/database?authSource=admin&replicaSet=rs0&readPreference=secondaryPreferred
```

## API 说明

### MongoDB 结构体方法

- `GetClient(appKeys ...string) *mongo.Client`: 获取MongoDB客户端
- `GetDatabase(dbName string, appKeys ...string) *mongo.Database`: 获取数据库实例
- `GetCollection(dbName, collectionName string, appKeys ...string) *mongo.Collection`: 获取集合实例
- `Close()`: 关闭连接

### 全局函数

- `InitAppMongoDB(configFunc, changeFunc)`: 初始化应用级MongoDB管理
- `AppMongoDBChange(appKey, changeType)`: 上报MongoDB配置变更

## 注意事项

1. **副本集概念**: MongoDB使用副本集而非传统主从复制
2. **必需配置**: 使用前必须配置 `hosts` 和 `database`
3. **副本集配置**: 使用副本集时必须指定 `replica_set` 名称
4. **读偏好**: 根据业务需求选择合适的读偏好策略
5. **认证配置**: 如果MongoDB启用了认证，需配置用户名、密码和认证数据库
6. **错误处理**: 务必检查返回的客户端、数据库、集合是否为nil
7. **连接管理**: 适时调用Close()方法释放资源
8. **应用隔离**: 不同应用使用不同的appKey来隔离MongoDB连接
9. **配置变更**: 配置变更需要依赖Redis来进行消息通知

## 对比 MySQL 使用

| 功能 | MySQL | MongoDB |
|------|-------|---------|
| 基础连接 | `DB{}.GetDB()` | `MongoDB{}.GetClient()` |
| 获取数据库 | 包含在连接中 | `MongoDB{}.GetDatabase()` |
| 获取集合/表 | 通过ORM操作 | `MongoDB{}.GetCollection()` |
| 应用级管理 | `InitAppDB()` | `InitAppMongoDB()` |
| 配置变更通知 | `AppDBChange()` | `AppMongoDBChange()` |
| 连接关闭 | 自动管理 | `MongoDB{}.Close()` |

## 注意事项

1. **必需配置**: 使用前必须配置 `primary` 连接信息
2. **配置完整性**: 确保MongoDB服务器地址、端口、认证信息正确
3. **错误处理**: 务必检查返回的客户端、数据库、集合是否为nil
4. **连接管理**: 适时调用Close()方法释放资源
5. **应用隔离**: 不同应用使用不同的appKey来隔离MongoDB连接
6. **配置变更**: 配置变更需要依赖Redis来进行消息通知
7. **认证数据库**: `auth_db` 通常设置为 "admin"，确保用户有正确的认证权限 