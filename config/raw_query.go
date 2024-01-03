package config

const (
	InsertUser  = `insert into users (username,userrole,address, email, passwordhash) values ($1,$2,$3,$4,$5) RETURNING id, created_at;`
	SelectUsers = `select id, username, userrole,address,email, passwordhash, created_at, updated_at FROM users ORDER BY created_at`
	UpdateUser  = `UPDATE users SET username = $2, address=$3, email=$4, passwordhash= $5 WHERE id=$1 RETURNING updated_at, created_at`
	DeleteUser  = `DELETE FROM users WHERE id=$1`
)
