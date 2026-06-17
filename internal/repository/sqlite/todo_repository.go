package sqlite

import (
	"database/sql"
	"todo_cli/internal/domain"

	_ "modernc.org/sqlite"
)

type Repository struct {
	db *sql.DB
}

func New(path string) (*Repository, error) {

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	r := &Repository{db: db}
	if err := r.migrate(); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Repository) migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER NOT NULL PRIMARY KEY,

		title TEXT NOT NULL,

		done INTEGER NOT NULL,

		created_at DATETIME NOT NULL,

		updated_at DATETIME NOT NULL
	);
	`
	_, err := r.db.Exec(query)
	return err
}

func (r *Repository) GetAll() ([]domain.Todo, error) {

	rows, err := r.db.Query(`
		SELECT
			id,
			title,
			done,
			created_at,
			updated_at
		FROM todos
		ORDER BY created_at DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []domain.Todo

	for rows.Next() {

		var todo domain.Todo

		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Done,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *Repository) Create(todo domain.Todo) error {

	_, err := r.db.Exec(
		`
		INSERT INTO todos(
			id,
			title,
			done,
			created_at,
			updated_at
		)
		VALUES(?,?,?,?,?)
		`,
		todo.ID,
		todo.Title,
		todo.Done,
		todo.CreatedAt,
		todo.UpdatedAt,
	)

	return err
}

func (r *Repository) Update(todo domain.Todo) error {

	_, err := r.db.Exec(
		`
		UPDATE todos
		SET
			title = ?,
			done = ?,
			updated_at = ?
		WHERE id = ?
		`,
		todo.Title,
		todo.Done,
		todo.UpdatedAt,
		todo.ID,
	)

	return err
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(
		`
		DELETE FROM todos
		WHERE id = ?
		`,
		id,
	)

	return err
}
