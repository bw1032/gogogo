module gogo

go 1.12

//download	download modules to local cache(下载依赖包)
//edit	    edit go.mod from tools or scripts（编辑go.mod)
//graph	    print module requirement graph (打印模块依赖图)
//verify	initialize new module in current directory（在当前目录初始化mod）
//tidy	    add missing and remove unused modules(拉取缺少的模块，移除不用的模块)
//vendor	make vendored copy of dependencies(将依赖复制到vendor下)
//verify	verify dependencies have expected content (验证依赖是否正确）
//why	    explain why packages or modules are needed(解释为什么需要依赖)

require (
	golang.org/x/exp v0.0.0-20211214223157-bafe2e20209a
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)
