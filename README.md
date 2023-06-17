# Tank War

坦克大战小游戏

- w、a、s、d 控制移动
- 空格键发射炮弹 装填3s

### 第三方库

- kitex  https://github.com/cloudwego/kitex
- hertz  https://github.com/cloudwego/hertz
- quic https://github.com/quic-go/quic-go
- gui tcell https://github.com/gdamore/tcell
- 监听键盘 https://github.com/eiannone/keyboard
- snowflake https://github.com/bwmarrin/snowflake
- consul https://github.com/hashicorp/consul/api
- viper https://github.com/spf13/viper
- gorm https://gorm.io

### 技术栈

| 功能               | 实现          |
| ------------------ | ------------- |
| HTTP框架           | hertz         |
| RPC 框架           | kitex         |
| 数据库             | redis、msyql  |
| 消息队列           | RabbitMQ      |
| 服务治理           | OpenTelemetry |
| 服务发现与配置中心 | Paseto        |
| 加密               | Snowflake     |

### 服务架构

![服务器架构](https://github.com/MulL1on/tank_war/blob/master/img/%E6%9C%8D%E5%8A%A1%E5%99%A8%E6%9E%B6%E6%9E%84.png)

- 客户端-大厅服务器(agent server)通过http通信
- 客户端-游戏服务器(game server)通过quic通信
- 大厅服务器-用户服务(user server)通过rpc通信

#### 大厅服务器

大厅服务器使用heartz框架开发，通过操作redis维护以个房间列表提供创建房间、浏览房间列表、加入房间的功能，通过调用kitex rpc提供用户注册、用户登录的功能。传输数据通过protobuf序列化。

#### 用户服务器

用户服务器使用kitex开发，通过操作mysql管理用户信息。传输数据通过thrift序列化。

#### 游戏服务器

客户端可以通过大厅服务器返回的游戏服务器地址、端口等信息，通过quic连接到游戏服务器，进行游戏。

游戏结束后，游戏服务器会将对局结果发布到rabbitmq的队列中，以便数据服务器（data server）进行处理，持久化。

### 更快的通讯方式

QUIC在UDP的基础上进行了扩展和改进，旨在提供更快的连接建立和数据传输速度。

与传统的TCP协议相比，QUIC 具有以下几个特点，这些特点有助于提高其传输速度：

1. 连接建立：QUIC使用一个称为"握手"的过程来建立连接。相比TCP的三次握手，QUIC采用了更高效的方式来减少往返时间。这意味着QUIC可以更快地建立连接，减少了初始延迟。
2. 多路复用：QUIC支持在单个连接上同时传输多个数据流。这意味着可以同时发送和接收多个请求和响应，而无需等待前一个请求完成。多路复用减少了连接建立的开销和延迟，提高了数据传输的效率。
3. 错误恢复：QUIC具有内置的错误恢复机制。当发生数据包丢失或网络故障时，QUIC可以更快地恢复连接，并重新传输丢失的数据，而无需等待重新建立连接。
4. 拥塞控制：QUIC使用自己的拥塞控制算法，可以更快地适应网络的变化。QUIC可以更精确地检测网络拥塞并调整发送数据的速率，以避免网络拥塞导致的性能下降。
5. 加密：QUIC内置了传输层加密，这有助于保护数据的安全性和隐私。

### 异步IO

通过消息队列的发布订阅模式实现异步的数据库io，对局结果异步持久化到数据库，提高游戏服务器的性能。

### 高可用

处理游戏逻辑的游戏服务器可以独立出来，一个游戏服务器可以负责多个房间的处理，当房间数超过一定限度时，游戏服务器将更新服务meta的status为busy，用户将不能建立在这个游戏服务器上创建房间，直到游戏服务器进行的房间数少于一定值，将status恢复为free。可以提高游戏服务器的性能和可用性。

- **游戏服务器**

```go
func updateServiceMeta(status string) {
	cfg := api.DefaultConfig()
	cfg.Address = net.JoinHostPort(
		config.GlobalConsulConfig.Host,
		strconv.Itoa(config.GlobalConsulConfig.Port))
	client, err := api.NewClient(cfg)
	if err != nil {
		klog.Infof("updatetServiceMeta new consul client error :%v", err)
		return
	}

	check := &api.AgentServiceCheck{
		TTL:                            consts.ConsulCheckTTL,
		DeregisterCriticalServiceAfter: consts.ConsulCheckDeregisterCriticalServiceAfter,
		Timeout:                        consts.ConsulCheckTimeout,
	}

	r := &api.AgentServiceRegistration{
		ID:      "game_srv:" + net.JoinHostPort(config.GlobalServerConfig.Host, strconv.Itoa(config.GlobalServerConfig.Port)),
		Address: config.GlobalServerConfig.Host,
		Name:    "game_srv",
		Port:    config.GlobalServerConfig.Port,
		Check:   check,
		Meta:    map[string]string{"status": status},
	}
	err = client.Agent().ServiceRegister(r)

	if err != nil {
		klog.Infof("updateServiceMeta error :%v", err)
		return
	}

	klog.Infof("updateServiceMeta success")

}
```

- **大厅服务器**

```go
for _, service := range services {
		if service.Service.Meta["status"] == "free" {
			r.Host = service.Service.Address
			r.Port = int32(service.Service.Port)
			r.RoomID = sf.Generate().Int64()

			//保存房间信息到redis
			jsonData, err := json.Marshal(r)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}

			err = config.Rdb.HSet(ctx, consts.Room, req.Name, jsonData).Err()
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
```

### 高并发

游戏服务器可以无脑扩容，大厅服务器只会在健康且status为free的游戏服务器上创建房间。理论上可以同时运行成百上千个房间。

### Consul

服务中心和配置中心

kv对详细：

- agent_server

  ```yaml
  {
    "name": "api",
    "port": 8888,
    "paseto": {
      "pub_key": "YOUR_KEY",
      "implicit": "YOUR-IMPLICIT"
    },
    "otel": {
      "endpoint": ":4317"
    },
    "user_srv":{
      "name": "user_srv"
    },
    "redis":{
      "host": "127.0.0.1",
      "port": 6379,
      "username":"",
      "password":"123456",
      "db":0
  }
  }
  ```

- game_server

  ```yaml
  {
    "name": "game_srv",
    "host":"localhost",
    "otel": {
      "endpoint": ":4317"
    },
    "user_srv":{
      "name": "user_srv"
    },
      "redis":{
      "host": "127.0.0.1",
      "port": 6379,
      "username":"",
      "password":"123456",
      "db":0
  	},
      "rabbitmq": {
      "host": "127.0.0.1",
      "port": 5672,
      "username": "guest",
      "password": "guest"
    }
  }
  ```

- user_srv

  ```yaml
  {
    "name": "user_srv",
    "paseto": {
      "pub_key": "YOUR_KEY",
      "implicit": "YOUR-IMPLICIT"
    },
    "mysql": {
      "host": "127.0.0.1",
      "port": 3306,
      "user": "root",
      "password": "123456",
      "db": "tank_war",
      "salt": "liangweijian"
    },
    "otel": {
      "endpoint": "localhost:4317"
    }
  }
  ```

### 最终结果

1. 启动环境

   ```
   docker compose up -d
   ```

2. 启动用户服务

   ```
   make user
   ```

3. 启动大厅服务器

   ```
   make api
   ```

4. 启动游戏服务器

   ```
   make game
   ```

5. 启动数据服务器

   ```
   make data
   ```

6. 启动游戏客户端

   ```
   //client 文件夹下
   go run main.go
   ```

![image-20230617195409959](https://github.com/MulL1on/tank_war/blob/master/img/image-20230617195409959.png)

按下对应数据使用功能

### 游戏界面

![image-20230617195831705](https://github.com/MulL1on/tank_war/blob/master/img/image-20230617195831705.png)

- @为岩石
- 箭头为坦克
- *为炮弹





