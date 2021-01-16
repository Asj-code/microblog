package main
 
import (
    "log"
    "github.com/asj-code/microblog/handlers"
    "github.com/asj-code/microblog/bd"
)
func main(){
    if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
    handlers.Manejadores()
 
}
