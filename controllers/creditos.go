package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/Polux_API_mid/golog"
)

type CreditosController struct {
	beego.Controller
}

func (c *CreditosController) URLMapping() {
	c.Mapping("ObtenerMinimo", c.ObtenerMinimo)
}

// Get ...
// @Title ObtenerMinimo
// @Description get creditos
// @Success 200 {object} make(map[string]string)
// @Failure 403
// @router /ObtenerMinimo [get]
func (c *CreditosController) ObtenerMinimo() {
	var comprobacion string = ""
	//consultar las reglas
	reglasBase := CargarReglasBase("RequisitosModalidades")
	fmt.Println(reglasBase)

	comprobacion = "minimo_numero_creditos_posgrado(Y)."
	r := golog.Obtener(reglasBase, comprobacion)
	var m = make(map[string]string)
	m["minimo_creditos_posgrado"] = r

	comprobacion = "minimo_numero_creditos_profundizacion(Y)."
	r = golog.Obtener(reglasBase, comprobacion)
	m["minimo_creditos_profundizacion"] = r

	fmt.Println(m)

	c.Data["json"] = m
	c.ServeJSON()

}