/*
Create time at 2023年3月18日0018下午 20:32:11
Create User at Administrator
*/

package main

import (
	"JWT-Test-Golang/initvariable"
	"JWT-Test-Golang/routes"
)

func main() {
	initvariable.InitEnv()
	initvariable.InitVariable()
	routes.Initrouter()

}
