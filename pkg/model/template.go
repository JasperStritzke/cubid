package model

type Template struct {
	Name    string
	Proxy   bool
	Version VersionValue
	Group   string `json:"_"`
}
