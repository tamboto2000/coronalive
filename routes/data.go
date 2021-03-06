package routes

import "github.com/tamboto2000/coronalive/controller"

func init() {
	Router.HandleFunc("/getAll", controller.GetAll)
	Router.HandleFunc("/getAllByProvince", controller.GetByProvince)
	Router.HandleFunc("/getNationalSummary", controller.GetNationalSummary)
}
