package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID				primitive.ObjectID		`json:"_id" bson:"_id"`
	FirstName		*string					`json:"first_name"`
	SecondName		*string					`json:"second_name"`
	Email			*string					`json:"email"`
	Created_at		time.Time				`json:"created_at"`
	UserCart		[]ProductUser			`json:"usercart" bson:"usercart"`
	Order_status	[]Order					`json:"orders" bson:"orders"`
	phone			*string					`json:"phone"`
	password		*string					`json:"password"`
	Token			*string					`json:"token"`
	Refresh_Token	*string					`json:"refresh_token"`
	Updated_at		time.Time				`json:"updated_at"`
	address_Details	[]Address				`json:"address" bson:"address"`
}
type Product struct{
	Product_id				primitive.ObjectID	`bson:"_id"`
	Product_Name			*string				`json:"product_name" bson:"product_name"`
	Image					*string				`json:"image" bson:"image"`
	price					*uint64				`json:"price" bson:"price"`
	Rating					*uint8				`json:"rating" bson:"rating"`
}

type Order struct{
	Order_ID			primitive.ObjectID		`bson:"_id"`
	Order_cart			[]ProductUser			`json:"order_cart" bson:"order_cart"`
	Ordered_At			time.Time				`json:"order_at" bson:"order_at"`
	Payment_Method      Payment					`json:"payment_method" bson:"payment_method"`
	price				*int					`json:"price"bson:"price"`
	Discount			int						`json:"discount"bson:"discount"`
}

type ProductUser struct{
	
}
type Address struct{
	Address_id			primitive.ObjectID		`bson:"_id"`
	City				*string					`json:"city_name" bson:"city_name"`
	Town				*string					`json:"town_name" bson:"town_name"`			 
	PostalCode			*string					`json:"postal_code"bson:"postal_code"`
	BoxNumber			*string					`json:"box_number"bson:"box_number"`
}
type Payment struct{
	Digital bool
	COD bool
}