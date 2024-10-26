# GIT WEBHOOK


## 项目介绍

`git-webhook` 是一个使用 Go 语言编写的服务置自动拉取最新代码。该服务支持基于 `push` 或 `tag` 的更新操作，可以使用基本认证或 SSH 密钥进行 Git 操作，适用于自动化部署和持续集成（CI）场景。

### 主要功能

- 接收 GitHub Webhook 请求
- 自动拉取代码或克隆仓库
- 支持使用 Git 用户名/密码或 SSH 密钥进行身份验证
- 使用日志记录操作和错误信息
- 配置灵活，支持 YAML 格式

## 项目结构
```
git-webhook/
├── config/               # 配置加载模块
│   └── config.go         # 配置加载代码
├── git/                  # Git 操作模块
│   ├── git.go            # Git 操作接口
│   ├── gitfactory.go     # Git 工厂模块
│   ├── clone.go          # Clone 操作
│   ├── pull.go           # Pull 操作
│   └── auth.go           # 认证策略模块
├── logger/               # 日志管理模块
│   └── logger.go         # 日志管理代码
├── webhook/              # Webhook 处理模块
│   └── webhook.go        # Webhook 处理代码
├── config.yml            # 配置文件
└── main.go               # 主程序入口
```

## 配置说明

在项目根目录下，您需要创建一个 `config.yml` 配置文件，示例配置如下：

```yaml
secret: "your_secret_here"      # GitHub Webhook 秘钥
project_path: "/path/to/project" # 本地项目存储路径
log_file_path: "/path/to/logfile.log" # 日志文件路径
port: 5000                       # 服务监听端口
repo_url: "https://github.com/yourusername/yourrepository.git" # Git 仓库 URL
use_tag: false                   # 是否基于标签拉取
branch: "main"                   # 拉取的分支
git_username: "your_username"    # Git 用户名（可选）
git_password: "your_password"    # Git 密码（可选）
ssh_key_path: "/path/to/key"     # SSH 密钥路径（可选）
```
## 启动说明

1. **安装 Go 环境**
   请确保您已安装 Go 语言环境。可以通过以下命令检查 Go 是否已安装：

   ```bash
   go version
   ```
2. **获取项目代码**
   使用 `git clone` 将项目克隆到本地：

   ```bash
   git clone https://github.com/yourusername/git-webhook-auto-pull.git
   cd git-webhook-auto-pull
   ```
3. **构建项目**
   在项目根目录下运行以下命令以构建可执行文件：

   ```bash
   
   # Windows
    set GOOS=windows 
    set GOARCH=amd64 
    go build -o webhook_server.exe
   
    # macOS
    export GOOS=darwin
    export GOARCH=amd64
    go build -o webhook_server

    # Linux
    export GOOS=linux
    export GOARCH=amd64 
    go build -o webhook_server
   ```
4. **启动服务**
   使用以下命令启动服务，并指定配置文件路径：

   ```bash
   ./webhook_server -config=config.yml
   ```
   该服务将监听您在配置文件中指定的端口。

## 调试

- **查看日志**
  日志文件路径在配置文件中定义。您可以在日志文件中查看服务的运行状态和可能的错误信息。
- **使用 curl 测试 Webhook**
  您可以使用以下 `curl` 命令来模拟 GitHub Webhook 的请求：

    ```bash  
    payload='{"ref": "refs/heads/main", "repository": {"full_name": "example/repo"}}'
    secret="9W1tB7hedsMtL"
    signature="sha256=$(echo -n $payload | openssl dgst -sha256 -hmac "$secret" | sed 's/^.* //')"
    curl -X POST http://localhost:5000/webhook -d "$payload" -H "Content-Type: application/json" -H "X-Hub-Signature-256: $signature"
    ```
  请根据实际情况替换 `your_signature` 和请求数据。

## 贡献

欢迎对本项目进行贡献！如果您有任何建议或改进，请提交 Issue 或 Pull Request。

## 许可证

本项目使用 MIT 许可证，详情请参见 LICENSE 文件。

### 使用说明
- 将上述内容保存为 `README.md` 文件，放在项目根目录下。
- 您可以根据需要进一步自定义或扩展文档内容，确保信息的准确性和完整性。

