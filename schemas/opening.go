package schemas

import(
	"gorm.io/gorm"
)

// Database structure 
type Opening struct{
	gorm.Model 
	Role string 
	Comapany string 
	Location string 
	Remote bool 
	Link string 
	Salary int64
}