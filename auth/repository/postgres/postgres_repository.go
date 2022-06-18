package postgres

import (
	"context"
	"database/sql"

	"github.com/questizen/core-system/domain"
	"github.com/questizen/core-system/helpers"
)

type postgresRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewPostgreAuthRepository(Conn *sql.DB) domain.AuthRepository {
	return &postgresRepository{Conn}

}

func (m *postgresRepository) CreateUser(ctx context.Context, user *domain.AuthUser) (domain.User, error) {
	sql := `INSERT INTO users (email, password, phone_number, name)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, name, email, phone_number, created_date`

	var resultUser domain.User
	tx, err := m.Conn.BeginTx(ctx, nil)
	helpers.PanicIfErr(err)
	defer helpers.CommitOrRollback(tx)

	err = tx.QueryRowContext(ctx, sql, user.Email, user.Password, user.PhoneNumber, user.Name).Scan(&resultUser.UserID,
		&resultUser.Name, &resultUser.Email, &resultUser.PhoneNumber, &resultUser.CreatedDate)
	helpers.PanicIfErr(err)

	return resultUser, nil
}

func (m *postgresRepository) GetUser(ctx context.Context) ([]domain.User, error) {
	rows, err := m.Conn.QueryContext(ctx, "SELECT user_id, name, email, phone_number, created_date FROM users")
	helpers.PanicIfErr(err)

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.PhoneNumber, &user.CreatedDate)
		helpers.PanicIfErr(err)
		users = append(users, user)
	}

	return users, nil
}
