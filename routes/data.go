package routes

import "coronalive/controller"

func init() {
	Router.HandleFunc("/getAll", controller.GetAll)
}
