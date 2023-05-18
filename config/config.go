package config

const (
	DeviceId    = "device_id"
	AccessToken = "access_token"
	PrivateKey  = "private_key"
	LicenseKey  = "license_key"
	ClientId    = "ClientId"
)

type Context struct {
	DeviceId    string
	AccessToken string
	PrivateKey  string
	LicenseKey  string
	ClientId    string
}
