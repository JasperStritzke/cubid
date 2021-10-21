package model

type Template struct {
	Name     string
	Version  VersionValue
	Group    string `json:"-"`
	Enabled  bool   `json:"enabled"`
	DataFile string `json:"-"`
}
