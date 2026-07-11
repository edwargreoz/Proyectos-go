package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//definir rutas manuales
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/Tareas", Tareas)
	http.HandleFunc("/Saludo", Saludar)

	fmt.Println("Servidor inciado en el puerto 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Scanln()

}

//HANDLER para la ruta principal

func Inicio(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "¡Bienvenido a la API de Gestion de Tareas!")
}

func Tareas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "listando Todas las Tareas")
	case "POST":
		fmt.Fprintf(w, "Crear una nueva Tarea")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo No permitido")
	}
}

// Handler para verificar salud del servidor
func Saludar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "ok", "message": "Servidor funcionando
	correctamente"}`)
}
