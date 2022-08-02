package users

import "fiber-simple-api/database"

type UsersRepository struct {
	database *database.Database
}

func NewUsersRepository(database *database.Database) *UsersRepository {
	return &UsersRepository{
		database: database,
	}
}

func (repository *UsersRepository) GetUsers() ([]*User, error) {
	var users []*User

	rows, err := repository.database.Query(`select id, name, email from users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository *UsersRepository) GetUser(id int) (*User, error) {
	user := new(User)
	row := repository.database.QueryRow(`select id, name, email from users where id = ? limit 1`, id)
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return user, err
	}
	return user, err
}

func (repository *UsersRepository) GetUserByEmail(email string) (*User, error) {
	user := new(User)
	row := repository.database.QueryRow(`select id, name, email, passwordfrom users where email = ? limit 1`, email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (repository *UsersRepository) InsertUser(user *User) error {
	tx, err := repository.database.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"insert into users (name, email, password) values (?, ?, ?)",
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repository *UsersRepository) DeleteUser(Id int) error {
	tx, err := repository.database.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		"delete from users where id = ?",
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
