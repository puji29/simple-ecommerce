package config

const (
	//query user
	InsertUser  = `insert into users (username,userrole,address, email, passwordhash) values ($1,$2,$3,$4,$5) RETURNING id, created_at;`
	SelectUsers = `select id, username, userrole,address,email, passwordhash, created_at, updated_at FROM users ORDER BY created_at`
	UpdateUser  = `UPDATE users SET username = $2, address=$3, email=$4, passwordhash= $5 WHERE id=$1 RETURNING updated_at, created_at`
	DeleteUser  = `DELETE FROM users WHERE id=$1`

	//query product
	InsertProduct = `INSERT INTO product (product_name, description, price,stock_quantity, category_id,image_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at`
	SelectProduct = `SELECT id,product_name, description, price,stock_quantity,created_at, updated_at, category_id,image_id FROM product ORDER BY created_at`
	UpdateProduct = `UPDATE product SET product_name= $2, description= $3, price= $4,stock_quantity=$5, category_id=$6,image_id=$7 WHERE id=$1 RETURNING updated_at `
	DeleteProduct = `DELETE FROM product WHERE id=$1`
)
