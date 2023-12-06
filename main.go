package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// dummy tugas
var tasks = []Task{
	{ID: 1, Title: "Tugas Golang", Done: false},
	{ID: 2, Title: "Tugas PHP", Done: false},
	{ID: 3, Title: "Tugas ScriptJawa", Done: false},
}

func main() {
	e := echo.New()

	e.Use(loggingMiddleware)

	e.GET("/tasks", getTasks)
	e.POST("/tasks", addTask)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	e.Start(":8888")
}

// read
func getTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, tasks)
}

// create
func addTask(c echo.Context) error {
	var newTask Task
	if err := c.Bind(&newTask); err != nil {
		return err
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	return c.JSON(http.StatusCreated, newTask)
}

// put
func updateTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}
	var input struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = input.Done
			tasks[i].Title = input.Title
			return c.JSON(http.StatusOK, tasks[i])
		}
	}
	return c.String(http.StatusNotFound, "Task not found")
}

// delete
func deleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid task ID")
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.String(http.StatusNotFound, "Task not found")
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Printf("Request: %s %s\n", c.Request().Method, c.Request().URL.Path)
		return next(c)
	}
}
