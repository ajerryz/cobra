# cobra
Cobra 是 Go 语言中最流行的命令行工具开发框架之一，被广泛用于构建复杂的命令行应用（如 Kubernetes 的 kubectl、Hugo、Docker CLI 等）。它提供了命令（Command）、标志（Flag）、参数（Argument）的完整支持，以及自动生成帮助文档、bash/zsh 补全脚本等功能。

# 核心概念
Cobra 的核心模型由三个部分组成：

- Command: 命令行的基本操作单元（如 git clone 中的 clone），可包含子命令。
- Flag: 命令的可选参数（如 ls -l 中的 -l），分为`持久标志`（对当前命令及所有子命令生效）和`本地标志`（仅对当前命令生效）。
- Argument: 命令的必选参数（如 git clone <repo-url> 中的 <repo-url>）。

# 安装与初始化
```shell
# 1.安装 Cobra 库
go get -u github.com/spf13/cobra@latest


# 2.安装 Cobra 代码生成工具（推荐）
go install github.com/spf13/cobra-cli@latest
# 2.1
mkdir cobra-demo && cd cobra-demo
go mod init github.com/yourname/cobra-demo
cobra-cli init  # 生成根命令结构
```

# 与viper集成
Cobra 常与配置管理库 Viper（同一作者）配合使用，实现 “标志> 环境变量 > 配置文件” 的优先级读取。
