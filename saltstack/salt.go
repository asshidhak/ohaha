package saltstack

import "net/http"

var saltBaseUrl = "http://192.168.3.121:8000"

type SaltStack struct {
	Eath string
	User string
	Passwd string
	Token string
}

func (ss *SaltStack) Connect() {
	response,err := http.Get(saltBaseUrl+"/login")
	if err != nil {
		return
	}

}