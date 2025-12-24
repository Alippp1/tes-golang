package dto

type PurchasingDetailResponse struct {
	ItemID   uint    `json:"item_id"`
	ItemName string  `json:"item_name"`
	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
	SubTotal float64 `json:"sub_total"`
}

type PurchasingResponse struct {
	ID         uint                       `json:"id"`
	Date       string                     `json:"date"`
	SupplierID uint                       `json:"supplier_id"`
	Supplier   string                     `json:"supplier"`
	GrandTotal float64                    `json:"grand_total"`
	UserID     uint                       `json:"user_id"`
	User       string                     `json:"user"`
	Details    []PurchasingDetailResponse `json:"details"`
}

type CreatePurchasingItemRequest struct {
	ItemID uint `json:"item_id" validate:"required"`
	Qty    int  `json:"qty" validate:"required,min=1"`
}

type CreatePurchasingRequest struct {
	Date       string                        `json:"date" validate:"required"`
	SupplierID uint                          `json:"supplier_id" validate:"required"`
	Items      []CreatePurchasingItemRequest `json:"items" validate:"required,min=1"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}