package osssrv

import (
	"fmt"
	"flag"
	"github.com/gorilla/mux"
)


func NewCmd() *cmd{
	c := new(cmd)
	c.actions = make(map[string]cmdAction)
	return c
}

type cmd struct {
	Router string
	actions map[string]cmdAction
	Flags struct{
		Addr string
		Debug bool
	}
}

type cmdAction interface{
	run()
}

func (c *cmd) Run()  {
	c.runAction()
}

func (c *cmd) Install(){
	c.Register("run-server", Server{
		Addr : c.Flags.Addr,
		Router: mux.NewRouter(),
		Debug: bool(c.Flags.Debug),
	})
}


func (c *cmd) Register(name string, action cmdAction)  {
	if _,ok := c.actions[name]; ok {
		panic(fmt.Sprintf("%s has already exist\n", name))
	}
	c.actions[name] = action
}

func (c *cmd) Parse()  {
	flag.StringVar(&c.Flags.Addr, "addr", "localhost:8080", "服务器监听地址")
	flag.BoolVar(&c.Flags.Debug, "debug", false, "是否开启调试")
	flag.BoolVar(&c.Flags.Debug, "d", false, "是否开启调试")

	flag.Parse()
	c.Router = flag.Arg(0)
}

func (c *cmd) runAction()  {
	if action, ok := c.actions[c.Router];ok{
		action.run()
	}else{
		fmt.Printf("ossfile [flag] {action}\n")
	}
}



