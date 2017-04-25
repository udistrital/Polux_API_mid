package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:CuposController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:CuposController"],
		beego.ControllerComments{
			Method: "Obtener",
			Router: `/Obtener`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:FechasController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:FechasController"],
		beego.ControllerComments{
			Method: "ObtenerFechas",
			Router: `/ObtenerFechas`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:SeleccionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:SeleccionController"],
		beego.ControllerComments{
			Method: "Seleccionar",
			Router: `/Seleccionar`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:VerificarRequisitosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/Polux_API_mid/controllers:VerificarRequisitosController"],
		beego.ControllerComments{
			Method: "Registrar",
			Router: `/Registrar`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
