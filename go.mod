module github.com/EDDYCJY/go-gin-example

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/MindorksOpenSource/Go-Log v0.0.0-20181118110108-6bbeddb1fc76
	github.com/alauda/go-redis-client v0.0.0-20190714061509-7d5ca20d0820 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/apenella/go-ansible v1.1.0
	github.com/astaxie/beego v1.12.3
	github.com/chasex/redis-go-cluster v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/gitstliu/go-redis-cluster v0.0.0-20190226073442-d274d87c0bfa // indirect
	github.com/go-ini/ini v1.62.0
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/spf13/viper v1.8.0 // indirect
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/tealeg/xlsx v1.0.5
	github.com/ugorji/go v1.2.6 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210616094352-59db8d763f22 // indirect
	golang.org/x/tools v0.1.3 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => ./go-blog/pkg/conf
	github.com/EDDYCJY/go-gin-example/docs => ./go-blog/service/docs
	github.com/EDDYCJY/go-gin-example/middleware => ./go-blog/middleware
	github.com/EDDYCJY/go-gin-example/models => ./go-blog/models
	github.com/EDDYCJY/go-gin-example/pkg/app => ./go-blog/pkg/app
	github.com/EDDYCJY/go-gin-example/pkg/e => ./go-blog/pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => ./go-blog/pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => ./go-blog/pkg/util
	github.com/EDDYCJY/go-gin-example/routers => ./go-blog/routers
	github.com/EDDYCJY/go-gin-example/routers/api => ./go-blog/routers/api
	github.com/EDDYCJY/go-gin-example/service/tag_service => ./go-blog/service/tag_service
)
