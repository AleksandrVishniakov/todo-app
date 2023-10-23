package repositories

import (
	"database/sql"
	"github.com/spf13/viper"
	"time"
)

var todosTable = viper.GetString("db.tables.todos")

type todo struct {
	db *sql.DB
}

func newToDo(db *sql.DB) *todo {
	return &todo{db: db}
}

func (t *todo) GetAllMessagesById(userId int) ([]ToDoEntity, error) {
	rows, err := t.db.Query("SELECT * FROM $1 WHERE user_id = $2",
		todosTable,
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
	_, err := t.db.Exec("UPDATE $1 SET message_text=$2, message_color=$3, message_date=$4 WHERE id=$5",
		todosTable,
		messageText,
		messageColor,
		messageDate,
		id,
	)

	return err
}

func (t *todo) DeleteMessageById(id int) error {
	_, err := t.db.Exec("DELETE FROM $1 WHERE id=$2",
		todosTable,
		id,
	)

	return err
}
