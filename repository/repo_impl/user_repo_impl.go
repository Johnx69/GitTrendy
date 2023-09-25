package repo_impl

import (
	custom_error "GitTrendy/custome_error"
	"GitTrendy/db"
	"GitTrendy/model"
	"GitTrendy/repository"
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)

	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, custom_error.UserConflict
			}
		}
		return user, custom_error.SignUpFail
	}

	return user, nil

}