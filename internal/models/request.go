package models

// ErrorResponse - структура для ошибок
type ErrorResponse struct {
	Errors string `json:"errors"`
}

// AuthRequest - структура запроса для аутентификации
type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse - структура ответа для аутентификации
type AuthResponse struct {
	Token string `json:"token"`
}

// SendCoinRequest - структура запроса для отправки монет
type SendCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
}

// InfoResponse - структура ответа с информацией о пользователе
type InfoResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

// InventoryItem - структура для элемента инвентаря
type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

// CoinHistory - структура для истории транзакций монет
type CoinHistory struct {
	Received []Transaction `json:"received"`
	Sent     []Transaction `json:"sent"`
}

// Transaction - структура для транзакции монет
type Transaction struct {
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
	Amount   int    `json:"amount"`
}
