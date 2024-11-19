// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func LoginSignup() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"wrapper\"><div class=\"card-switch\"><label class=\"switch\"><input type=\"checkbox\" class=\"toggle\"> <span class=\"slider\"></span> <span class=\"card-side\"></span><div class=\"flip-card__inner\"><div class=\"flip-card__front\"><div id=\"login-error\"></div><form class=\"flip-card__form\" hx-post=\"/login\" hx-target=\"#main\"><input class=\"flip-card__input\" name=\"email\" placeholder=\"Email\" type=\"email\" required> <input class=\"flip-card__input\" name=\"password\" placeholder=\"Contraseña\" type=\"password\" required><div class=\"radio-inputs\"><label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"1\" checked=\"\"> <span class=\"name\">Cliente</span></label> <label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"2\"> <span class=\"name\">Propietario</span></label> <label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"3\"> <span class=\"name\">Administrador</span></label></div><button type=\"submit\" class=\"flip-card__btn\">LogIn</button></form></div><div class=\"flip-card__back\"><div class=\"title-err\" id=\"signupErr\"></div><form class=\"flip-card__form\" hx-post=\"/signup\" hx-target=\"#main\"><input class=\"flip-card__input\" placeholder=\"Nombre\" type=\"name\" required> <input class=\"flip-card__input\" name=\"email\" placeholder=\"Email\" type=\"email\" required> <input class=\"flip-card__input\" name=\"password\" placeholder=\"Contraseña\" type=\"password\" required><div class=\"radio-inputs\"><label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"1\" checked=\"\"> <span class=\"name\">Cliente</span></label> <label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"2\"> <span class=\"name\">Propietario</span></label> <label class=\"radio\"><input type=\"radio\" name=\"userType\" value=\"3\"> <span class=\"name\">Administrador</span></label></div><button class=\"flip-card__btn\">SignUp</button></form></div></div></label></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate