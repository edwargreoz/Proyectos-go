package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Estructuramos los datos para una Tarea

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}

// Saludar al servidor
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Applicaction/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "ok",
	"message": "Servidro funcionando"}`)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bienvenido al servicio de Peticiones de Gestiond de Tareas")
}

// Almacen temporal de memoria
var tasks []Task
var nextID = 1

// Handler mejorado para tareas con manejo completo.
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		handleGetTasks(w, r)

	case "POST":
		handleCreateTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Metodo no permitido",
		})
	}
}

// GET /tasks - Listar tareas con filtros opcionales

func handleGetTasks(w http.ResponseWriter, r *http.Request) {
	//Lectura del parametro de consulta

	queryParams := r.URL.Query()
	Completed := queryParams.Get("completed")
	limit := queryParams.Get("limit")

	//Aplicacion de filtros

	filteredTasks := tasks

	//filtrar por estado completo

	if Completed != "" {
		isCompleted, err := strconv.ParseBool(Completed)
		if err == nil {
			var filtered []Task
			for _, task := range tasks {
				if task.Completed == isCompleted {
					filtered = append(filtered, task)
				}
			}
			filteredTasks = filtered
		}
	}

	//Aplicar limite

	if limit != "" {
		if limitNum, err := strconv.Atoi(limit); err == nil && limitNum > 0 {
			if limitNum < len(filteredTasks) {
				filteredTasks = filteredTasks[:limitNum]
			}
		}

	}

	//leer headers importantes
	userAgent := r.Header.Get("User-Agent")
	log.Printf("Cliente conectado: %s", userAgent)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tasks": filteredTasks,
		"total": len(filteredTasks),
	})
}

// POST /tasks - Crear nueva tarea
func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	//Verificar el tipo de contenido
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Content-type debe ser application/json",
		})
		return
	}

	// leer y parsear body
	var newTask Task

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields() // Rechazar campos desconocidos

	if err := decoder.Decode(&newTask); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"error": "Json Invalido: " + err.Error(),
		})
		return
	}

	//Validar datos

	if strings.TrimSpace(newTask.Title) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "El titulo es obligatorio",
		})
		return
	}

	//crear tarea
	newTask.ID = nextID
	newTask.CreatedAt = time.Now()
	newTask.Completed = false //Iniciar siempre como no completada
	nextID++

	tasks = append(tasks, newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Tarea creada exitosamente",
		"task":    newTask,
	})

}

// Handler para manejar una tarea en especifico por ID
func taskByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Extraer id de la url (routing manual basico)

	path := strings.TrimPrefix(r.URL.Path, "/tasks/")

	if path == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID de tarea requerido",
		})
		return
	}
	id, err := strconv.Atoi(path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "ID debe ser un número valido",
		})
		return
	}

	var foundTask *Task

	for i := range tasks {
		if tasks[i].ID == id {
			foundTask = &tasks[i]
			break
		}
	}

	for foundTask == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Tarea no encontrada",
		})
		return
	}

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(foundTask)
	case "PUT":
		handleUpdateTask(w, r, foundTask)
	case "DELETE":
		handleDeleteTask(w, r, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"erro": "Metodo no permitido",
		})

	}

}

// PUT /tasks/{id} - Actualizar tarea
func handleUpdateTask(w http.ResponseWriter, r *http.Request, task *Task) {

	var updates Task

	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Json invalido",
		})
		return
	}

	if updates.Description != "" {
		task.Description = updates.Description
	}
	task.Completed = updates.Completed

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Tarea Actualizada Exitosamente",
		"task":    task,
	})
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request, id int) {
	for i, task := range tasks {
		if task.ID == id {
			//eliminar tarea del slice

			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Tarea eliminada exitosamente",
			})
			return
		}
	}
}

func main() {
	//Inicializar con algunas tareas de ejemplo

	tasks = []Task{
		{
			ID:          1,
			Title:       "Estudiar Go",
			Description: "Completar el tutorial de APIS",
			Completed:   false,
			CreatedAt:   time.Now().Add(-24 * time.Hour),
		},
		{
			ID:          2,
			Title:       "Hacer ejercicio",
			Description: "30 miinutitos de ful cardio",
			Completed:   true,
			CreatedAt:   time.Now().Add(-12 * time.Hour),
		},
	}
	nextID = 3

	//Registrar Rutas

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskByIDHandler)
	http.HandleFunc("/headlth", healthHandler)

	fmt.Println("Servidor iniciado en puerto 8080")
	fmt.Println("Rutas disponibles")

	fmt.Println(" GET /tasks - Listar tareas")
	fmt.Println(" POST /tasks - Crear tarea")
	fmt.Println(" GET /tasks/{id} - Obtener tarea por ID")
	fmt.Println(" PUT /tasks/{id} - Actualizar tarea")
	fmt.Println(" DELETE /tasks/{id} - Eliminar tarea")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
