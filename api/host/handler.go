package host

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"jintmp/app"
	"jintmp/api"
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