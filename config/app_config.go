package config

const (
	ApiGroup = "/api/v1"

	//Users
	UserPost   = "/users/"
	UserList   = "/users/"
	UserUpdate = "/users/:id"
	UserDelete = "/users/:id"
	//product
	ProductPost             = "/products/"
	ProductGet              = "/products/"
	ProductPut              = "/products/:id"
	ProductDelete           = "/products/:id"
	ProductGetById          = "/products/:id"
	ProductGetByProductName = "/product/:productName"
	//category
	//orderItem
	//orderTable
	OrderTablePost   = "/orderTable/"
	OrderTableList   = "/orderTable/"
	OrderTablePut    = "/orderTable/:id"
	OrderTableDelete = "/orderTable/:id"
	OrderTableByID   = "/orderTable/:id"

	//image
	ImageInsert = "/image"
	ImageGet    = "/image"
	ImagePut    = "/image/:id"
	ImageById   = "/image/:id"
	ImageDelete = "/image/:id"
)
