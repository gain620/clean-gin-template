package jwt

type JwtPayload struct {
	Exp           int      `json:"exp"`
	Iat           int      `json:"iat"`
	AuthTime      int      `json:"auth_time"`
	Jti           string   `json:"jti"`
	Iss           string   `json:"iss"`
	Aud           string   `json:"aud"`
	Sub           string   `json:"sub"`
	Typ           string   `json:"typ"`
	Azp           string   `json:"azp"`
	SessionState  string   `json:"session_state"`
	AtHash        string   `json:"at_hash"`
	Acr           string   `json:"acr"`
	EmailVerified bool     `json:"email_verified"`
	Roles         []string `json:"roles"`
	//	"roles": [
	//	"project1_DL",
	//	"project4_DE",
	//	"project2_DE",
	//	"project4_RO",
	//	"project3_SR"
	//]
}
