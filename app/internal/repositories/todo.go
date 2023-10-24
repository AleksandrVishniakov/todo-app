package repositories

import (
	"database/sql"
	"time"
)

type todo struct {
	db *sql.DB
}

func newToDo(db *sql.DB) *todo {
	return &todo{db: db}
}

func (t *todo) GetAllMessagesById(userId int) ([]ToDoEntity, error) {
	rows, err := t.db.Query("SELECT * FROM todos WHERE user_id = $1",
		userId,
	)

	var todos []ToDoEntity

	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo ToDoEntity
		err := rows.Scan(&todo.UserId, &todo.Id, &todo.ToDoText, &todo.ToDoColor, &todo.ToDoDate)
		if err != nil {
			return todos, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *todo) UpdateMessageById(id int, messageText, messageColor string, messageDate time.Time) error {
	_, err := t.db.Exec("UPDATE todos SET message_text=$1, message_color=$2, message_date=$3 WHERE id=$4",
		messageText,
		messageColor,
		messageDate,
		id,
	)

	return err
}

func (t *todo) DeleteMessageById(id int) error {
	_, err := t.db.Exec("DELETE FROM todos WHERE id=$1",
		id,
	)

	return err
}

func (t *todo) AddMessage(userId int, todo ToDoEntity) (ToDoEntity, error) {
	var id int
	row := t.db.QueryRow("INSERT INTO todos (user_id, message_text, message_date, message_color) VALUES ($1, $2, $3, $4) RETURNING id",
		userId,
		todo.ToDoText,
		todo.ToDoDate,
		todo.ToDoColor,
	)

	if err := row.Scan(&id); err != nil {
		return ToDoEntity{}, err
	}

	todo.Id = id
	todo.UserId = userId

	return todo, nil
}
