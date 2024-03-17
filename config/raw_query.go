package config

const (
	//query user
	InsertUser  = `insert into users (username,userrole,address, email, passwordhash) values ($1,$2,$3,$4,$5) RETURNING id, created_at;`
	SelectUsers = `select id, username, userrole,address,email, passwordhash, created_at, updated_at FROM users ORDER BY created_at`
	UpdateUser  = `UPDATE users SET username = $2, address=$3, email=$4, passwordhash= $5 WHERE id=$1 RETURNING updated_at, created_at`
	DeleteUser  = `DELETE FROM users WHERE id=$1`

	//query product
	InsertProduct              = `INSERT INTO product (product_name, description, price,stock_quantity, category_id,image_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id,created_at`
	SelectProduct              = `SELECT id,product_name, description, price,stock_quantity,created_at, updated_at, category_id,image_id FROM product ORDER BY created_at`
	SelectProductById          = `SELECT id,product_name, description, price,stock_quantity,created_at, updated_at, category_id,image_id FROM product WHERE id=$1`
	SelectProductByProductName = `SELECT id,product_name, description, price,stock_quantity,created_at, updated_at, category_id,image_id FROM product WHERE product_name=$1`
	UpdateProduct              = `UPDATE product SET product_name= $2, description= $3, price= $4,stock_quantity=$5, category_id=$6,image_id=$7 WHERE id=$1 RETURNING updated_at `
	DeleteProduct              = `DELETE FROM product WHERE id=$1`

	//query orderTable
	InsertOrderTable     = `INSERT INTO order_table (user_id, order_date, total_amount) VALUES ($1,$2,$3) RETURNING id,created_at`
	SelectOrderTable     = `SELECT id, user_id,order_date,total_amount,created_at,updated_at FROM order_table ORDER BY created_at`
	UpdateOrderItem      = `UPDATE order_Table SET order_date=$2, total_amount=$3 WHERE id=$1 RETURNING updated_at`
	DeleteOrderTable     = `DELETE FROM order_table WHERE id=$1`
	SelectOrderTableById = `SELECT id,user_id, order_date, total_amount, created_at, updated_at FROM order_table WHERE id=$1`

	//query image
	InsertImage     = `INSERT INTO images (image) Values ($1) RETURNING id,created_at`
	SelectAllImage  = `SELECT id, image, created_at,updated_at FROM images ORDER BY created_at`
	ImageUpdate     = `UPDATE images SET image=$2 WHERE id=$1 RETURNING updated_at`
	SelectImageById = `SELECT id, image, created_at, updated_at FROM images WHERE id=$1`
	DeleteImage     = `DELETE FROM images WHERE id=$1`

	//category
	InsertCategory = `INSERT INTO category (category_name) VALUES ($1) RETURNING id,created_at`
)
