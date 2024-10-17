package main
import ("Log")
func main(){
	handlers:=new(handler.Handler)
	srv:= new(todo.Server)
	if err:= srv.Run(port: "8000",handlers.InitRoutes()); err !=nil {
		log.Fatalf(format:"error occured while running http server:%s",err.Error())
	}
	
}