package uhybrid
// 定义包名和导入所需的包和库，其中testing包是用于写测试代码的。
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)
// 定义一个测试函数，函数名以Test开头，后面跟着需要测试的函数名或功能名称。
// 函数名为TestAPI。这个函数接受一个*testing.T类型的参数，这个参数用于输出测试结果和错误信息。

func TestAPI(t *testing.T) {

// 创建一个测试服务器，并定义了一个匿名函数作为服务器的处理逻辑。
// 这个匿名函数接受两个参数：一个http.ResponseWriter类型的参数w，用于写入响应头和响应体，以及一个*http.Request类型的参数r，表示HTTP请求。
// 使用http.HandlerFunc函数将匿名函数转换成http.Handler类型的对象，然后将这个对象传递给httptest.NewServer函数，创建一个测试服务器。

server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 根据请求的路径分发请求到不同的处理函数中。在这个例子中，我们有两个处理函数：loginHandler用于处理登录请求，userHandler用于处理获取用户信息请求。
// 如果请求的路径不是/login或/user，则返回404响应。

switch r.URL.Path {
case "/login":
    loginHandler(w, r)
case "/user":
    userHandler(w, r)
default:
    http.NotFound(w, r)
}

}))
// 创建完测试服务器后，使用defer关键字延迟关闭服务器。这样，在测试函数结束时，服务器就会自动关闭，释放资源。
defer server.Close()

// 这部分代码创建了一个登录请求，并使用http.Post函数将请求发送到测试服务器的/login路径。
// http.Post函数接受三个参数：请求的URL、请求的MIME类型和请求的body。
// 将请求体序列化成JSON格式，并使用bytes.NewBuffer函数将其转换为io.Reader类型，作为请求的body。
// http.Post函数返回一个*http.Response类型的对象，表示服务器的响应。我们使用defer关键字延迟关闭响应体的Body，以便及时释放资源。

loginReq := map[string]string{
    "username": "testuser",
    "password": "password",
}
loginReqJSON, _ := json.Marshal(loginReq)
loginResp, err := http.Post(server.URL+"/login", "application/json", bytes.NewBuffer(loginReqJSON))
if err != nil {
    t.Fatalf("Failed to send login request: %v", err)
}
defer loginResp.Body.Close()
// 检查服务器的响应状态码是否为200。如果状态码不是200，就使用t.Errorf函数输出测试错误信息。
if loginResp.StatusCode != http.StatusOK {
    t.Errorf("Expected status 200, got %d", loginResp.StatusCode)
}
// 解析登录响应体中的JSON数据，并将其存储到一个结构体中。在这个例子中，我们将JSON数据解码为一个带有Token字段的结构体。
// 我们使用json.NewDecoder函数创建一个JSON解码器，并使用Decode方法将响应体中的JSON数据解码为结构体。
// 如果解码失败，就使用t.Fatalf函数输出测试错误信息，并终止测试函数的运行。

var loginRespData struct {
    Token string `json:"token"`
}

if err := json.NewDecoder(loginResp.Body).Decode(&loginRespData); err != nil {
    t.Fatalf("Failed to decode login response: %v", err)
}
// 检查服务器的响应状态码是否为200。如果状态码不是200，就使用t.Errorf函数输出测试错误信息。
userReq := map[string]string{
    "token": loginRespData.Token,
}
userReqJSON, _ := json.Marshal(user这部分代码创建了一个获取用户信息的请求，并使用`loginRespData.Token`中存储的token作为请求参数。与之前的登录请求类似，我们将请求体序列化为JSON格式，并使用`http.Post`函数将请求发送到测试服务器的`/user`路径。我们使用`defer`关键字延迟关闭响应体的`Body`，以便及时释放资源。

// ```go
// if userResp.StatusCode != http.StatusOK {
//     t.Errorf("Expected status 200, got %d", userResp.StatusCode)
// }

// 解析获取用户信息响应体中的JSON数据，并将其存储到一个map中。
// 我们使用json.NewDecoder函数创建一个JSON解码器，并使用Decode方法将响应体中的JSON数据解码为map。
// 如果解码失败，就使用t.Fatalf函数输出测试错误信息，并终止测试函数的运行。
var userData map[string]interface{}

if err := json.NewDecoder(userResp.Body).Decode(&userData); err != nil {
    t.Fatalf("Failed to decode user response: %v", err)
}
// 检查获取用户信息响应体中的用户名字段是否为testuser。如果用户名字段不是testuser，就使用t.Errorf函数输出测试错误信息。
if userData["username"] != "testuser" {
    t.Errorf("Expected username to be testuser, got %v", userData["username"])
}
// 检查获取用户信息响应体中的邮箱字段是否为testuser@example.com。如果邮箱字段不是testuser@example.com，就使用t.Errorf函数输出测试错误信息。
if userData["email"] != "testuser@example.com" {
    t.Errorf("Expected email to be testuser@example.com, got %v", userData["email"])
}}