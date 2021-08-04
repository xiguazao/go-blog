package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/ansible_service"
	golog "github.com/MindorksOpenSource/Go-Log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnsiblePlayExecutorForm struct {
	// smart
	Connection string `form:"connection" valid:"Required;`
	// root
	User string `form:"user" valid:"Required;"`
	//"10.100.2.118,10.100.2.164"
	Inventory string `form:"inventory" valid:"Required;"`
	// {"ansible-runner/site.yml", "ansible-runner/site2.yml"}
	Playbooks string `form:"playbooks"`
}

// 执行 ansible
// @Summary Ansible 执行
// @Produce  json
// @Param connection query string true "Connection"
// @Param user query string true "User"
// @Param inventory query string true "Inventory"
// @Param playbooks query string false "Playbooks"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/ansible [post]
func AnsiblePlayExecutor(c *gin.Context) {
	golog.ConfigureTimer()
	golog.ConfigureCallingFunction()

	var (
		appG = app.Gin{C: c}
		form AnsiblePlayExecutorForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	tagService := ansible_service.ConnectOptions{
		Connection: form.Connection,
		User:       form.User,
		Inventory:  form.Inventory,
		Playbooks:  []string{form.Playbooks},
	}
	_, err := tagService.PlayBookExecutor()
	if err != nil {

		golog.E(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR_ANSIBLE_EXECUTE, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
