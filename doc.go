/*

Kwiscale is a middleware/framework that is designed to work "as" WSGI implementation from Python.

The main goal is to let you develop "handlers" that have HTTP verbs as Go method.

Example:

    // __/handlers/myhandler.go

    package handlers

    import "github.com/metal3d/kwiscale"

    type MyHandler struct {
        kwiscale.RequestHandler `route:"/foo"`
    }

    // You MUST give a constructor that returns IRequestHandler
    func (t* MyHandler) New() kwiscale.IRequestHandler {
        return new(MyHandler)
    }

    // reply to GET on /foo route
    func (h *MyHandler) Get () {
        h.Write("Hello !")
    }


Note that New() method is required, we use it as factory to get always a new instance of handler to not interfer with other connections.
This is a fix from previous version of kwiscale

Now, in your main file:

    // __/main.go

    package main
    import "github.com/metal3d/kwiscale"
    import "./handlers"

    func main() {
        kwiscale.AddHandler(&handlers.MyHandler{})
        kwiscale.Server(":8080")
    }


Launching "go run main.go" starts server, you can visit http://127.0.0.1:8080/foo 

Kwiscale appends some features that are needed to easilly develop web application or REST API. 

Template:

You can override a base template to get "main" frame, then interact on blocks. This functionnality
that is a simili from Jinja (Yython) or Twig (PHP) or Swig (NodeJS).

Example:

    <!-- /templates/main.html -->
    {{define "CONTENT"}}{{end}}
    <html>
        <head></head>
        <body>
            {{ template "CONTENT" . }}
        </body>
    </html>

    <!-- /template/home/default.html -->
    {{ override "main.html" }}
    {{ define "CONTENT" }}
        This is the content from my home handler: {{ .msg }}
    {{ end }}

The handler can be:

    func (h *MyHomeHandler) Get() {
        h.Render("home/default.html", map[string]string{
            "msg" : "Hello"
        })
    }

As you can see (at this time) we don't use "template/" prefix while calling "Render" function. And Inside "templates/home/default.html"
we override "main.html". The path is always use from "templates" directory.

There is a very beta system that creates forms from a basic struct.

Example:
    
    type AuthForm struct {
        Name        string `form:"text,Give a name,required"`
        Password    string `form:"password,Type your password,required"`
    }

    //...

    auth := AuthForm{}
    form := kwiscale.CreateForm(auth)

"form" variable contains the full form. If any of field as a value, it will be inserted in th "value" attribute.

Select tag can be use, but it is in early stage.



*/
package kwiscale