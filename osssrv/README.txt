package osssrv

import (
	"fmt"
	"flag"
	"crypto/x509/pkix"
)

func NewCmd() *cmd{
	c := new(cmd)
	c.actions = make(map[string]cmdAction)
	return c
}

type cmd struct {
	router string
	actions map[string]cmdAction
}

type cmdAction interface{
	run()
}

func (c cmd) Install(){
	c.Register("run-server", Server{
		Addr : "localhost:8080",
	})
}

func (c cmd) Parse()  {
	flag.Parse()
	c.router = flag.Arg(0)
}

func (c cmd) RunAction(name pkix.Name)  {
	if action, ok := c.actions[c.router];ok{
		action.run()
	}else{
		fmt.Printf("ossfile {action} [flag]\n")
	}
}


func (c cmd) Register(router string, action cmdAction)  {
	if _,ok := c.actions[router]; ok {
		panic(fmt.Sprintf("%s has already exist\n", router))
	}
	c.actions[router] = action
}
