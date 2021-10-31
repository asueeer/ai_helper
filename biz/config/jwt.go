package config

var (
	JwtDefaultConfig = JwtConfig{
		JwtKey:  "nearby123",
		Issuer:  "be_nearby",
		Subject: "be_nearby_subject",
	}
)

type JwtConfig struct {
	JwtKey  string
	Issuer  string
	Subject string
}
