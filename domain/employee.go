package domain

type Product struct {
	productCode        string `gorm:"primaryKey;type:varchar(15)"`
	productName        string `gorm:"type:varchar(70);"`
	productScale       string `gorm:"type:varchar(10);"`
	productDescription string `gorm:"type:text"`
	productVendor      string `gorm:"type:varchar(50);"`
	quantityInStock    string `gorm:"type:smallint(6);"`
	buyPrice           string `gorm:"type:decimal(10,2);"`
	MSRP               string `gorm:"type:decimal(10,2);"`
}
