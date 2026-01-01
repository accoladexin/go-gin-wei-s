



# 0.总览





# 1.目录

go-shop/
└── services/
    └── user-service/
        ├── cmd/
        │   └── main.go
        ├── internal/
        │   ├── handler/
        │   ├── service/
        │   ├── repository/
        │   └── config/
        ├── proto/
        │   └── user.proto
        ├── go.mod

```sh
PS E:\codeO\01-go-gin-wei\go-shop> cd .\services\
PS E:\codeO\01-go-gin-wei\go-shop\services> cd .\user-service\
PS E:\codeO\01-go-gin-wei\go-shop\services\user-service> go mod init go-shop/user-service
go: creating new go.mod: module go-shop/user-service
go: to add module requirements and sums:
        go mod tidy
PS E:\codeO\01-go-gin-wei\go-shop\services\user-service>

```

```
protoc --go_out=. --go-grpc_out=. proto/user.proto



protoc    --go_out=.   --go-grpc_out=.  proto/product.proto

protoc  --go_out=.   --go-grpc_out=.   proto/*.proto
```



海外服务器

```
version: '3.8'

services:
  kong:
    image: kong:3.6
    container_name: kong-gateway
    user: root  # 避免 Rocky Linux 下的权限挂载问题
    environment:
      KONG_DATABASE: "off"
      KONG_DECLARATIVE_CONFIG: /kong/kong.yml
      KONG_PROXY_ACCESS_LOG: /dev/stdout
      KONG_PROXY_ERROR_LOG: /dev/stderr
      KONG_ADMIN_ACCESS_LOG: /dev/stdout
      KONG_ADMIN_ERROR_LOG: /dev/stderr
      KONG_ADMIN_LISTEN: 0.0.0.0:8001
    ports:
      - "8000:8000"     # 外部访问接口 (HTTP)
      - "8001:8001"   # 管理接口 (Admin API)
      - "8443:8443"    # 外部访问接口 (HTTPS)
    volumes:
      - ./kong.yml:/kong/kong.yml
    restart: always

networks:
  default:
    name: kong-net

```

```
sudo firewall-cmd --permanent --add-port=8001/tcp
sudo firewall-cmd --permanent --add-port=8000/tcp
sudo firewall-cmd --permanent --add-port=8001/udp
sudo firewall-cmd --permanent --add-port=8000/udp
sudo firewall-cmd --permanent --add-port=8443/udp
sudo firewall-cmd --permanent --add-port=8443/udp
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https

sudo firewall-cmd --reload
```





















Go 语言中的 `&` 和 `*` 是初学者最容易混淆的地方。要系统掌握它们，只需要记住一句话：**`&` 是把变量变地址，`*` 是把地址变变量（或者声明指针类型）。**

我们可以从三个维度来彻底拆解：

---

### 一、 核心定义：`&` (取地址) 与 `*` (取值)

#### 1. `&` (取地址符)
它作用于一个**变量**，返回这个变量在内存中的**地址**。
*   想象变量是一个房间，`&` 就是拿到这个房间的“门牌号”。

#### 2. `*` (取值符/解引用符)
它作用于一个**地址（指针）**，返回这个地址对应的**具体数值**。
*   想象你手里有一个“门牌号”，`*` 就是拿着钥匙开门进去，拿到房间里的东西。

**代码示例：**
```go
func main() {
    age := 18
    
    // & 取地址
    ptr := &age      // ptr 现在存储的是 age 的内存地址（比如 0xc0000140a8）
    
    fmt.Println(ptr)  // 输出：0xc0000140a8
    
    // * 取值
    value := *ptr    // 拿着地址去找值
    fmt.Println(value) // 输出：18
}
```

---

### 二、 `*` 的双重身份（最容易混淆点）

在 Go 中，`*` 出现在不同的地方，含义完全不同：

#### 1. 出现在【类型】前面：表示“这是一个指针类型”
用于声明变量、函数参数或返回值。
```go
var p *int    // p 是一个变量，它只能用来存 int 类型的“地址”
func test(p *Product) {} // 参数必须传地址
```

#### 2. 出现在【变量】前面：表示“我要取这个地址里的值”
这是一个**动作**，前提是这个变量必须已经是一个指针。
```go
fmt.Println(*p) // 报错！如果 p 还没赋值地址（是 nil），这叫“空指针解引用”
```

---

### 三、 为什么 Go 随处可见 `&` 和 `*`？（实战场景）

在写业务（比如你的 `ProductService`）时，为什么要频繁用它们？

#### 1. 修改外部变量（传引用）
Go 函数的参数传递默认是**值拷贝**。
```go
func changeName(n string) {
    n = "New Name" // 只改了副本，原变量没变
}

func changeNameReal(n *string) {
    *n = "New Name" // 顺着地址改了原变量的值
}
```

#### 2. 结构体初始化与单例
当你看到 `return &ProductService{...}` 时：
*   **不带 `&`**：`return ProductService{}` 返回的是整个结构体的所有数据。如果结构体有 100 个字段，就会在内存里复制 100 个字段。
*   **带 `&`**：`return &ProductService{}` 返回的是一个 8 字节的内存地址。不管结构体多大，传输效率都极高。

#### 3. 表示“空”的概念 (nil)
*   基本类型 `int`, `string` 永远有默认值（0, ""），它们不能为 `nil`。
*   但**指针类型**可以为 `nil`。
*   在微服务里，如果我们查不到某个产品，通常返回 `(*Product, error)`。如果查不到，第一个参数就返回 `nil`，这非常符合逻辑。

---

### 四、 快速口诀

如果你在写代码时纠结了，背这个口诀：

1.  **想拿地址？用 `&`**
    *   `ptr := &user`
2.  **想定义一个存地址的类型？用 `*`**
    *   `var ptr *User`
3.  **手里有地址，想改里面的值？用 `*`**
    *   `*ptr = newUser`
4.  **函数签名里带 `*`？**
    *   返回时记得加 `&`：`return &Product{}`
    *   调用时记得加 `&`：`service.Update(&user)`

---

### 五、 回到你之前的疑问
```go
func NewProductService(userClient *UserClient) *ProductService {
    return &ProductService{
        userClient: userClient,
    }
}
```
*   **参数 `userClient *UserClient`**：我不想要你的 Client 副本，请把你的 Client 地址告诉我，我直接引用它。
*   **返回值类型 `*ProductService`**：这个函数结束后，我会把创建好的 Service 的地址给你。
*   **实现 `&ProductService{...}`**：我现在内存里造一个 Service，并把它的地址（`&`）吐出去。

**这样，整个系统中只有这一份 `ProductService` 数据，所有人都通过地址访问它，这就是高效的单例模式。**