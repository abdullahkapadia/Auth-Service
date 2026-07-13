package repository

import (
	"context"

	"github.com/adullahkapadia/auth-service/internal/database"
	"github.com/adullahkapadia/auth-service/internal/model"
)

type UserRepository struct{}

func (r * UserRepository) Create(user *model.User) error{

	query := ` INSERT INTO users
	(id,name,email,password)
	VALUES($1,$2,$3,$4) `

	_, err := database.DB.Exec(
		context.Background(),
		query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)

	return err
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {

	query := `
	SELECT
	id,
	name,
	email,
	password,
	created_at,
	updated_at
	FROM users
	WHERE email=$1
	`

	user := &model.User{}

	err := database.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(id string) (*model.User, error) {

	query := `
	SELECT
	id,
	name,
	email,
	password,
	created_at,
	updated_at
	FROM users
	WHERE id=$1
	`

	user := &model.User{}

	err := database.DB.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(user *model.User) error {

	query := `
	UPDATE users
	SET
	name=$1,
	email=$2,
	updated_at=NOW()
	WHERE id=$3
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		user.Name,
		user.Email,
		user.ID,

		
	)

	return err
}


func (r *UserRepository) Delete(id string) error {

	query := `
	DELETE FROM users
	WHERE id=$1
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		id,
	)

	return err
}
func (r *UserRepository) Exists(email string) (bool, error) {

    query := `
        SELECT EXISTS(
            SELECT 1
            FROM users
            WHERE email=$1
        )
    `

    var exists bool

    err := database.DB.QueryRow(
        context.Background(),
        query,
        email,
    ).Scan(&exists)

    return exists, err
}