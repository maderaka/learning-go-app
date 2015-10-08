package api

const (
	ApiIdHeaderName     = "X-RT-ID"
	ApiUserHeaderName   = "X-RT-USER"
	ApiSecretHeaderName = "X-RT-SECRET"
)

var clients = map[string]map[string]string{
	"website": map[string]string{
		ApiUserHeaderName:   "1234",
		ApiSecretHeaderName: "rakatejaa@gmail.com",
	},
	"mobile": map[string]string{
		ApiUserHeaderName:   "12345",
		ApiSecretHeaderName: "rakatejaa@gmail.com",
	},
}
