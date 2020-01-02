package main

import "github.com/labstack/echo"
import "fmt"
import "html/template"
import "io"

type ND struct {
	Room string
	Msg string
	Sender string
}

var DATA = map[string][]ND{}
var Rooms = []string{}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render (w io.Writer,name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w,name,data)
}

func main(){
	DATA["TestRoom"]=append(DATA["TestRoom"],ND{
		Room:"TestRoom",
		Msg:"TestMsg",
		Sender:"TestSender",
	})

	Rooms=[]string{"TestRoom"}
	e:=echo.New()

	renderer:=&TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	
	e.Renderer=renderer
	
	e.Debug=true

	e.POST("/add",func(c echo.Context) error {
		room:=c.FormValue("room")
		msg:=c.FormValue("msg")
		sender:=c.FormValue("sender")
		
		if(room==""||msg==""||sender==""){
			return c.String(200,"")
		}

		fmt.Println(room+" "+msg+" "+sender);
	
		DATA[room]=append(DATA[room],ND{
			Room:room,
			Msg:msg,
			Sender:sender,
		})

		if(!Contains(Rooms,room)){
			Rooms=append(Rooms,room)
		}
		return c.String(200,"");
	})

	e.GET("/",func(c echo.Context) error {
		room:=c.QueryParam("room")
	
		return c.Render(200,"view.html",map[string]interface{}{
			"Rooms":Rooms,
			"Data":DATA[room],
		})
	})
	e.Start(":8989")
}

func Contains(a []string,b string) bool {
	for _,v := range a{
		if(v==b){
			return true
		}
	}
	return false
}
