
## 简书

安装Go环境，本文适合类unix系统安装 
Go的开发环境安装配置，还是比较简洁的。我们使用的是源码安装，会生成相应版本的Go源码，方便读源码学习 。



winX系统应该与安装个软件无差别，最多配置下环境变量。笔者没有试过。

## 1 官方安装地址

地址1: `https://golang.org/dl/`

地址2: `https://golang.google.cn/dl/`

地址1 经常需要高科技上网，地址2正常能上网即可打开

还有很多其他的平台的以及CDN的下载地址，我们就不一一列举了。 找到一个即可。

---
## 2 安装步骤
### 2.1.	下载
用wget命令下载，可以从官网地址上选择一个对应系统的版本，右键复制链接，跟在wget后面即可。下面示例的命令可直接用于运行。

例：$ `wget https://golang.google.cn/dl/go1.15.5.linux-amd64.tar.gz`
 

### 2.2.	安装到指定目录
源码安装到指定目录。下面例子安装目录为 `/usr/local/go`

例：$ `tar -C /usr/local  -xzf  go1.15.5.linux-amd64.tar.gz`
 


### 2.3.	添加环境变量
添加环境变量的意思就是，让Go编译器可执行文件，全局可用。
#### •	GOROOT添加
添加GO全局环境变量，全局导出Go。
##### 配置文件修改

$ `vi  /etc/profile`

此命令打开的是全局配置文件。将光标移动到文件最末尾（快捷键：shift+g），添加下面的变量（换行插入的命令是o）。
```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

保存并退出：

1> 按：	ESC键；

2> 输入：	:wq。

##### 导入， 全局可使用go命令

$ `source /etc/profile`

##### 判断是否安装成功，输入命令，则安装完成。

$ `go version`

输出：

$ go version go1.15.5 linux/amd64

#### •	GOPATH添加
GOPATH是Go工作环境变量，一般添加在开发者用户下，比如多人在同一台服务器上作业。又有各自的工作账号以及工作目录。
添加开发者用户 ，如果有用户里此步骤可以省略。用开发者账户的好处就是，多人协同在同一台服务器上作战。如果机器是自己的直接在root账户编程也是没问题的。
我们为系统添加一个boy用户（自定义开发用户，即使）。

##### 添加用户的
添加用户命令 $ `useradd boy`；其中boy可以改成自己的，若是用root往后的步骤可以sheng

设置用户密码 $ `passwd 123456`

切换到boy用户下 $ `su boy` 

若是使用root账号开发，可省略 添加用户这一部分。

##### 配置GOPATH。

$  `vi  ~/.bash_profile`

在export PATH 前添加如下内容。
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export GOPROXY=https://goproxy.cn #代理
export GO111MODULE=on
```

保存并退出：

1> 导入： $ `source  ~/.bash_profile`

2> 测试： $ `echo $GOPATH`

通过这三个步骤Go的环境我们就安装好了。

---
## 3 运行hello, world!
在boy用户下创建一个Go项目文件夹hello，在这个工程下添加一个hello.go 的文件。将第一章的“hello，world”代码粘贴到这个文件里。用户boy可以替换成你自己的用户。
执行前先确认是否在项目文件下。

CODE hello.go
```go
package main
import "fmt"
func main() {
	fmt.Println("Hello,world!")
}
```

$ `pwd`
输出
/home/boy/hello

### 3.1. 直接运行使用go run命令
$ `go run hello.go`

输出:

Hello，world!

### 3.2.	编译+运行 使用go build命令
若是第一次使用它，先用命令go mod 初始化我们的项目。下面的命令常用格式，并不是严格要求格式必须这样。命令的斜线部分是可以自定义的。

#### • 初始化项目命令 
$ `go mod init github.com/xxx/hello` 

命令go mod 解析如下：

1> github.com:	远程仓库地址；

2> xxx:			你在github仓库的用户名；

3> hello:		项目名。

#### • 编译+执行

$ `go build && ./hello`

输出:
Hello，world!

还会生成一个 hello 的可执行文件下次再运行 可以直接是用。

$ `./hello`

#### • 补充 初始化项目命令，以下命令都是有效的。

$ `go mod init github.com/xxx/path_a/path_b/hello`

$ `go mod init hello`

---
### 完成



