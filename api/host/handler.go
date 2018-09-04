package host

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/asshidhak/ohaha/app"
	"github.com/asshidhak/ohaha/api"
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
	cmd := exec.Command("git --version")

	outbyte, err := cmd.Output()
	if err != nil {
		log.Println(err)
		c.JSON(400,err)
		return
	}
	log.Println(string(outbyte))
	c.JSON(200, string(outbyte))
	return

}