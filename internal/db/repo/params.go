package repo

type AddAccountParams struct {
	UserID  int     `json:"user_id"`
	Deposit float64 `json:"deposit"`
}

type AddAccountNewUserParams struct {
	Name    string  `json:"name"`
	Deposit float64 `json:"deposit"`
}

type GetAccountParams struct {
	AccountNumber string
}

type ExecTransferParams struct {
	FromAccountNumber string
	ToAccountNumber   string
	Amount            float64
}

type GetTransfersParams struct {
	AccountNumber string
}

type AddUserParams struct {
	Name string `json:"name"`
}

type GetUserParams struct {
	ID int `json:"id"`
}
