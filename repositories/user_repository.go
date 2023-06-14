package repositories

import (
	"database/sql"
	"time"

	"github.com/buniekbua/gousers/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	query := `insert into users
	(first_name, last_name, email, password, created_at, modified_at)
	values ($1, $2, $3, $4, $5, $6)`

	_, err := ur.db.Exec(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		time.Now(),
		time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `SELECT id, first_name, last_name, email, created_at, modified_at 
	FROM users WHERE id = $1;`

	row := ur.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.CreatedAt, &user.ModifiedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) UpdateUser(id int, user *models.User) error {
	query := `UPDATE users
	SET first_name = $1,
		last_name =  $2,
		email = $3,
		password = $4,
		modified_at = $5
	WHERE id = $6;`

	_, err := ur.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, time.Now(), id)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	rows, err := ur.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.ModifiedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	_, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
//query := `SELECT * FROM users WHERE email = $1`

//}
