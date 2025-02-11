// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Layout() templ.Component {
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
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>Chattr</title><script src=\"https://unpkg.com/htmx.org@1.9.10\"></script><script src=\"https://unpkg.com/htmx.org/dist/ext/ws.js\"></script><script async crossorigin=\"anonymous\" data-clerk-publishable-key=\"pk_test_dW5pcXVlLWJpc29uLTk4LmNsZXJrLmFjY291bnRzLmRldiQ\" src=\"https://unique-bison-98.clerk.accounts.dev/npm/@clerk/clerk-js@5/dist/clerk.browser.js\" type=\"text/javascript\"></script><script>\n            window.addEventListener('load', async function () {\n                await Clerk.load()\n\n                if (Clerk.user) {\n                document.getElementById('auth').innerHTML = `\n                    <div id=\"user-button\"></div>\n                `\n\n                const userButtonDiv = document.getElementById('user-button')\n\n                Clerk.mountUserButton(userButtonDiv)\n                } else {\n                document.getElementById('auth').innerHTML = `\n                    <div id=\"sign-in\"></div>\n                `\n\n                const signInDiv = document.getElementById('sign-in')\n\n                Clerk.mountSignIn(signInDiv)\n                }\n            })\n        </script><style>\n            #auth {\n                padding: 1rem;\n            }\n            #user-button {\n                min-width: 200px;\n            }\n            #sign-in {\n                min-height: 600px;\n            }\n        </style><link rel=\"stylesheet\" href=\"/static/styles.css\"></head><body><div id=\"auth\"></div><div id=\"container\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
