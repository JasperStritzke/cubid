package model

type Template struct {
	Name    string
	Proxy   bool
	Version VersionValue
	Group   string `json:"-"`
}

func (t Template) ProxyAsString() string {
	var proxyString = "Server"
	if t.Proxy {
		proxyString = "Proxy"
	}

	return proxyString
}
