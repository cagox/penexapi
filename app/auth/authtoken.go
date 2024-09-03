package auth

type AuthToken struct {
	Token_ID int    `json:"auth_token"`
	Token    string `json:"token"`
	User_ID  int    `json:"user_id"`
}
