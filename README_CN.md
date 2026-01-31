# cliz

[English](README.md) | 简体中文

一个现代化、轻量级且直观的 Go CLI 框架。cliz 提供简单的 API，用于创建带有子命令、标志、验证以及高级功能如继承、位置参数和自定义错误处理的命令行界面。

## 功能特性

### 核心功能
- **子命令支持**: 轻松创建层次化的子命令系统
- **灵活的标志系统**: 支持各种类型的标志和验证器
- **结构标签集成**: 通过结构体标签快速定义标志
- **位置参数**: 支持按位置传递参数
- **标志验证**: 内置丰富的验证器，支持自定义验证规则
- **继承机制**: 子命令可以继承父命令的标志

### 类型支持
- 基本类型: `string`, `int`, `uint`, `bool`, `float32`, `float64`
- 切片类型: `[]string`, `[]int`, `[]uint`, `[]bool`, `[]float32`, `[]float64`
- 整数类型: `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32`, `uint64`及其切片形式

### 验证器
- `Required`: 必选标志
- `Range`: 数值范围验证
- `Length`: 字符串长度验证
- `Pattern`: 正则表达式验证
- `In`: 枚举值验证
- `Contains`: 子字符串验证
- `Email`: 邮箱格式验证
- `URL`: URL格式验证
- `Alpha`: 只包含字母
- `AlphaNum`: 只包含字母和数字
- 自定义验证器

## 安装

```bash
go get github.com/zkep/cliz
```

## 快速开始

### 基本用法

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("myapp", "示例 CLI 应用", "1.0.0")

	var name string
	var age int
	var verbose bool

	app.String("name", "您的姓名", &name, cliz.Required())
	app.Int("age", "您的年龄", &age, cliz.Range(18, 99))
	app.Bool("verbose", "启用详细输出", &verbose)

	app.Action(func() error {
		fmt.Printf("您好，%s！您今年 %d 岁。\n", name, age)
		if verbose {
			fmt.Println("已启用详细模式")
		}
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### 子命令

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("myapp", "带子命令的 CLI 应用", "1.0.0")

	var name string
	app.String("name", "您的姓名", &name)

	greet := app.NewSubCommand("greet", "问候某人")
	greet.Action(func() error {
		fmt.Printf("您好，%s！\n", name)
		return nil
	})

	farewell := app.NewSubCommand("farewell", "说再见")
	farewell.Action(func() error {
		fmt.Printf("再见，%s！\n", name)
		return nil
	})

	app.AddCommand(greet)
	app.AddCommand(farewell)

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### 结构标签

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("struct_example", "使用结构体标志的 CLI", "1.0.0")

	type Config struct {
		Name     string `name:"name" description:"您的姓名" validate:"required"`
		Age      int    `name:"age" description:"您的年龄" validate:"range=18-99"`
		Email    string `name:"email" description:"您的邮箱" validate:"email"`
		Verbose  bool   `name:"verbose" description:"启用详细输出"`
	}

	var config Config
	app.AddFlags(&config)

	app.Action(func() error {
		fmt.Printf("姓名: %s\n", Config.Name)
		fmt.Printf("年龄: %d\n", Config.Age)
		fmt.Printf("邮箱: %s\n", Config.Email)
		fmt.Printf("详细模式: %t\n", Config.Verbose)
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

### 位置参数

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zkep/cliz"
)

func main() {
	app := cliz.NewCli("positional_example", "带位置参数的 CLI", "1.0.0")

	type Args struct {
		Source string `position:"0"`
		Dest   string `position:"1"`
	}

	var args Args
	app.AddPositionalArgs(&args)

	app.Action(func() error {
		fmt.Printf("从 %s 复制到 %s\n", Args.Source, Args.Dest)
		return nil
	})

	if err := app.Run(os.Args...); err != nil {
		log.Fatal(err)
	}
}
```

## 高级特性

### 继承标志

子命令可以继承父命令的标志：

```go
parent := cliz.NewCommand("parent", "父命令")
parent.String("name", "您的姓名", &name)

child := parent.NewSubCommand("child", "子命令")
child.InheritFlags(parent)
```

### 预运行回调

在命令执行前执行的回调函数：

```go
app.PreRun(func(c *cliz.Cli) error {
	fmt.Println("在执行命令之前")
	return nil
})
```

### 自定义 Banner

```go
app.Banner(func(c *cliz.Cli) string {
	return fmt.Sprintf("自定义 Banner - %s v%s", c.Name(), c.Version())
})
```

### 默认命令

当没有指定命令时执行的默认命令：

```go
defaultCmd := cliz.NewCommand("default", "默认命令")
defaultCmd.Action(func() error {
	fmt.Println("这是默认命令")
	return nil
})

app.DefaultCommand(defaultCmd)
```

## API 文档

### 主要类型

#### `Cli`
- `NewCli(name, description, version string) *Cli`: 创建新的 CLI 应用
- `Run(args ...string) error`: 执行 CLI 应用
- `PrintHelp()`: 打印帮助信息
- `NewSubCommand(name, description string) *Command`: 创建子命令
- `AddFlags(flags any) *Cli`: 从结构添加标志
- `Action(callback Action) *Cli`: 设置命令执行回调
- `PreRun(callback func(*Cli) error)`: 设置预运行回调
- `DefaultCommand(defaultCommand *Command) *Cli`: 设置默认命令

#### `Command`
- `NewCommand(name, description string) *Command`: 创建新命令
- `AddSubCommand(cmd *Command) *Command`: 添加子命令
- `String(name, description string, variable *string, validators ...Validator) *Command`: 添加字符串标志
- `Int(name, description string, variable *int, validators ...Validator) *Command`: 添加整数标志
- `Bool(name, description string, variable *bool, validators ...Validator) *Command`: 添加布尔标志
- `Float64(name, description string, variable *float64, validators ...Validator) *Command`: 添加浮点数标志
- `StringSlice(name, description string, variable *[]string, validators ...Validator) *Command`: 添加字符串切片标志
- `IntSlice(name, description string, variable *[]int, validators ...Validator) *Command`: 添加整数切片标志
- `BoolSlice(name, description string, variable *[]bool, validators ...Validator) *Command`: 添加布尔切片标志
- `AddPositionalArgs(args any) *Command`: 添加位置参数
- `InheritFlags(parent *Command) *Command`: 继承父命令的标志

### 验证器

- `Required() Validator`: 必选验证
- `Range(min, max any) Validator`: 范围验证
- `Length(min, max int) Validator`: 长度验证
- `Pattern(regex string) Validator`: 正则表达式验证
- `In(values ...string) Validator`: 枚举验证
- `Contains(substr string) Validator`: 子字符串验证
- `Email() Validator`: 邮箱验证
- `URL() Validator`: URL 验证
- `Alpha() Validator`: 字母验证
- `AlphaNum() Validator`: 字母数字验证

每个验证器都支持 `WithMessage(msg string)` 方法来自定义错误信息。

## 示例

更多示例请查看 [_examples](_examples) 目录：

- [basic](_examples/basic/main.go): 基本用法
- [subcommands](_examples/subcommands/main.go): 子命令示例
- [struct](_examples/struct/main.go): 结构标志示例
- [positional](_examples/positional/main.go): 位置参数示例
- [advanced](_examples/advanced/main.go): 高级验证器示例
- [inherit](_examples/inherit/main.go): 继承示例
- [banner](_examples/banner/main.go): Banner 示例
- [prerun](_examples/prerun/main.go): 预运行回调示例

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！

