package api

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	golog "github.com/MindorksOpenSource/Go-Log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime/trace"
)

//开启Trace
// @Summary start trace
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /start_trace [get]
func StartTrace(c *gin.Context) {
	golog.ConfigureTimer()
	golog.ConfigureCallingFunction()
	appG := app.Gin{C: c}
	f, err := os.Create("./runtime/trace.out")
	if err != nil {
		panic(err)
	}

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	golog.D("Start trace")
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//停止Trace
// @Summary stop trace
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /stop_trace [get]
func StopTrace(c *gin.Context) {
	golog.ConfigureTimer()
	golog.ConfigureCallingFunction()
	appG := app.Gin{C: c}

	trace.Stop()
	golog.D("Stop trace")
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
