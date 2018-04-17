package main

import "ossfile/osssrv"

func main()  {
	cmd := osssrv.NewCmd()
	cmd.Parse()
	cmd.Install()
	cmd.Run()
}