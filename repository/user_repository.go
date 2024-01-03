package repository

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(payload entity.User) (entity.User, error)
	List() ([]entity.User, error)
	UpdateUser(data entity.User) (entity.User, error)
	Delete(id string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

// Delete implements UserRepository.
func (u *userRepository) Delete(id string) (entity.User, error) {
	_, err := u.db.Exec(config.DeleteUser, id)
	if err != nil {
		log.Println("error deleteting users: ", err)
		return entity.User{}, nil
	}
	return entity.User{}, nil
}

// UpdateUser implements UserRepository.
func (u *userRepository) UpdateUser(data entity.User) (entity.User, error) {
	//tidak ada pembaruan jika semua field kosong
	if data.Username == "" && data.Addres == "" && data.Email == "" && data.PasswordHash == "" {
		return data, nil
	}
	var hashedPassword string

	if data.PasswordHash != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(data.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			log.Println("bcrypt.GeneratedPassword:", err)
			return entity.User{}, err
		}
		hashedPassword = string(hashed)
	}

	if err := u.db.QueryRow(config.UpdateUser, data.ID, data.Username, data.Addres, data.Email, hashedPassword).Scan(&data.Updated_at, &data.Created_at); err != nil {
		log.Println("QueryRow.UPdate :", err)
		return entity.User{}, err
	}
	return data, nil

}

// list implements UserRepository.
func (u *userRepository) List() ([]entity.User, error) {
	var users []entity.User

	rows, err := u.db.Query(config.SelectUsers)
	if err != nil {
		log.Println("userQuery:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.UserRole,
			&user.Addres,
			&user.Email,
			&user.PasswordHash,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			log.Println("userRepository.Next:", err.Error())
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil

}

// Create implements UserRepository.
func (u *userRepository) Create(payload entity.User) (entity.User, error) {
	if err := u.db.QueryRow(config.InsertUser, payload.Username, payload.UserRole, payload.Addres, payload.Email, payload.PasswordHash).Scan(&payload.ID, &payload.Created_at); err != nil {
		return entity.User{}, err
	}
	return payload, nil
}

func NewUserREpository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
