package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:""`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
// ListTasks muestra todas las tareas en la lista.
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas!")
		return
	}

	for _, task := range tasks {

		status:=" "
		if task.Complete{
			status = "✓"
		}

	fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}

}
// AddTask agrega una nueva tarea a la lista.
func AddTask(tasks []Task, name string)[]Task{
	newTask := Task{
		ID: GetNextID(tasks),
		Name: name,
		Complete: false,
	}
	return append(tasks,newTask)
}

// DeleteTask elimina una tarea de la lista por su ID.
func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks{
		if task.ID == id {
			return append(tasks[:i], tasks[i + 1:]...)
		}
	}
	return tasks
}

// CompleteTask marca una tarea como completada por su ID.
func CompleteTask(tasks []Task, id int)[]Task{
	for i, task := range tasks{
		if task.ID == id {
			tasks[i].Complete = true
			break
		}
	}
	return tasks
}

// SaveTasks guarda las tareas en el archivo proporcionado.
func SaveTasks(file *os.File, tasks []Task){
	bytes,err := json.Marshal(tasks)
	if err != nil{
		panic(err)
	}

	_, err = file.Seek(0, 0)
	if err != nil{
		panic(err)
	}

	err = file.Truncate(0)
	if err != nil{
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil{
		panic(err)
	}

	err = writer.Flush()
	if err != nil{
		panic(err)
	}
}

// GetNextID obtiene el siguiente ID para una nueva tarea en la lista.
func GetNextID(tasks []Task) int{
	if len(tasks) == 0{
		return 1
	}
	return tasks[len(tasks) - 1].ID + 1
}
