module myPackage

go 1.14

require (
	github.com/uesrname/project/utils v0.0.0 // 在github上没有此项目
	github.com/uesrname/project/plugins v0.0.0 // 在github上没有此项目
	github.com/uesrname/project/common v0.0.0 // 在github上没有此项目
)

replace github.com/uesrname/project/utils => ./utils // 使用替换的方法，这里指向本地目录
replace github.com/uesrname/project/common => ./common // 使用替换的方法，这里指向本地目录
replace github.com/uesrname/project/plugins => ./plugins // 使用替换的方法，这里指向本地目录