/*
Пакет server содержит хендлеры для обработки запросов к базе данных и пути API
/tasks – создание задачи;
/tasks – получение списка всех задач;
/tasks/:id – обновление задачи;
/tasks/:id – удаление задачи.
*/
package server

import (
	"time"

	"github.com/fevse/srtodo/internal/storage"
	"github.com/gofiber/fiber/v2"
)

// CreateTask предназначена для внесения новой задачи в базу данных
func CreateTask(c *fiber.Ctx) error {
	task := &storage.Task{}
	err := c.BodyParser(task)
	if err != nil {
		return c.Status(400).SendString("Wrong request format")
	}

	query := "INSERT INTO tasks (id, title, description, status) VALUES($1, $2, $3, $4)"

	_, err = storage.Storage.Exec(query, task.ID, task.Title, task.Description, task.Status)

	if err != nil {
		return c.Status(500).SendString("Create task error")
	}

	return c.Status(201).SendString("Task has created successfully")
}

// ShowTasks предназначена для получения списка задач
func ShowTasks(c *fiber.Ctx) error {
	query := "SELECT * FROM tasks"
	rows, err := storage.Storage.Query(query)
	if err != nil {
		return c.Status(500).SendString("DB executing query error")
	}
	defer rows.Close()

	var tasks []storage.Task

	for rows.Next() {
		var task storage.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			return c.Status(500).SendString("Data scan error")
		}
		tasks = append(tasks, task)
	}
	return c.JSON(tasks)
}

// UpdateTask предназначена для обновления выбранной задачи
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := &storage.Task{}

	err := c.BodyParser(task)
	if err != nil {
		return c.Status(400).SendString("Wrong request format")
	}

	query := "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5"

	_, err = storage.Storage.Exec(query, task.Title, task.Description, task.Status, time.Now(), id)
	if err != nil {
		return c.Status(500).SendString("Update task error")
	}

	return c.SendString("Task has updated successfully")
}

// DeleteTask предназначена для удаления выбранной задачи
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")
	query := "DELETE FROM tasks WHERE id = $1"

	_, err := storage.Storage.Exec(query, id)
	if err != nil {
		return c.Status(500).SendString("Delete task error")
	}

	return c.SendString("Task has deleted successfully")
}
