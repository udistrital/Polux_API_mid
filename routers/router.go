// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/Polux_API_mid/controllers"
)

func init() {

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/verificarRequisitos",
			beego.NSInclude(
				&controllers.VerificarRequisitosController{},
			),
		),
		beego.NSNamespace("/seleccion",
			beego.NSInclude(
				&controllers.SeleccionController{},
			),
		),
		beego.NSNamespace("/cupos",
			beego.NSInclude(
				&controllers.CuposController{},
			),
		),
		beego.NSNamespace("/fechas",
			beego.NSInclude(
				&controllers.FechasController{},
			),
		),
		beego.NSNamespace("/creditos",
			beego.NSInclude(
				&controllers.CreditosController{},
			),
		),
		beego.NSNamespace("/evaluadores",
			beego.NSInclude(
				&controllers.EvaluadoresController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
