package course

type Course struct {
	ID              int     `json:"id"`
	CID             string  `json:"cid"`
	Name            string  `json:"name"`
	Locale          *int    `json:"locale"`
	Validity        *int    `json:"validity"`
	Price           int     `json:"price"`
	DiscountPercent *int    `json:"discount_percent"`
	IsPublic        bool    `json:"is_public"`
	IsOpen          bool    `json:"is_open"`
	Description     *int    `json:"description"`
	Thumbnail       *string `json:"thumbnail"`
	StartsAt        *int    `json:"starts_at"`
	EndsAt          *int    `json:"ends_at"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}
