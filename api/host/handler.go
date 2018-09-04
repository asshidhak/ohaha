package host

import (
	"github.com/asshidhak/ohaha/api"
	"github.com/asshidhak/ohaha/app"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
)

type HostReques struct {
	Name string `json:"name"`
}

func HandleHostList(c *gin.Context) {

	rows, err := app.Db.Query("select `env`, `region` from host limit 100")
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorRes(err))
		return
	}
	hosts := []Host{}
	for rows.Next() {
		host := Host{}
		rows.Scan(&host.Env,&host.Region)
		hosts = append(hosts, host)
	}
	c.JSON(http.StatusOK, api.Response{hosts, ""})
}


func HandleCmd(c *gin.Context) {
	//cmd := exec.Command("sh","-c", "git clone https://gitlab.xgimi.com/devops/tagTest.git")
	cmd := exec.Command("sh","-c", "sh ")
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		c.String(400,"%v",err)
		return
	}

	cmd.Wait()
}