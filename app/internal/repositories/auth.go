package repositories

import (
	"database/sql"
)

type auth struct {
	db *sql.DB
}

func newAuth(db *sql.DB) *auth {
	return &auth{db: db}
}

func (a *auth) GetUserById(id int) (UserEntity, error) {
	row := a.db.QueryRow("SELECT * FROM users WHERE id = $1",
		id,
	)

	var user UserEntity

	if err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Salt); err != nil {
		return user, err
	}

	return user, nil
}

func (a *auth) GetUserByLogin(login string) (UserEntity, error) {
	row := a.db.QueryRow("SELECT * FROM users WHERE login = $1",
		login,
	)

	var user UserEntity

	if err := row.Scan(&user.Id, &user.Login, &user.Password, &user.Salt); err != nil {
		return user, err
	}

	return user, nil
}

func (a *auth) AddUser(u UserEntity) (UserEntity, error) {
	var id int
	row := a.db.QueryRow("INSERT INTO users (login, password, salt) VALUES ($1, $2, $3) RETURNING id",
		u.Login,
		u.Password,
		u.Salt,
	)

	if err := row.Scan(&id); err != nil {
		return UserEntity{}, err
	}

	return a.GetUserById(id)
}

func (a *auth) GetRefreshTokenById(userId int) (TokenEntity, error) {
	row := a.db.QueryRow("SELECT * FROM tokens WHERE user_id = $1",
		userId,
	)

	var token TokenEntity

	if err := row.Scan(&token.UserId, &token.RefreshToken, &token.RefreshTokenExpiresAt); err != nil {
		return token, err
	}

	return token, nil
}

func (a *auth) UpdateRefreshTokenById(t TokenEntity) error {
	_, err := a.db.Exec("INSERT INTO tokens (refresh_token, refresh_token_expires_at, user_id) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET refresh_token=$1, refresh_token_expires_at=$2",
		t.RefreshToken,
		t.RefreshTokenExpiresAt,
		t.UserId,
	)

	return err
}
