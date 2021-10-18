package model

type ServiceTypeValue = string

var ServiceType = struct {
	DynamicLobby ServiceTypeValue
	StaticLobby  ServiceTypeValue

	Dynamic ServiceTypeValue
	Static  ServiceTypeValue
}{
	DynamicLobby: "DYNAMIC_LOBBY",
	StaticLobby:  "STATIC_LOBBY",
	Dynamic:      "DYNAMIC",
	Static:       "STATIC",
}

func IsDynamic(s string) bool {
	return s == ServiceType.DynamicLobby || s == ServiceType.Dynamic
}

func IsStatic(s string) bool {
	return s == ServiceType.StaticLobby || s == ServiceType.Static
}

func IsLobby(s string) bool {
	return s == ServiceType.DynamicLobby || s == ServiceType.StaticLobby
}

type Service struct {
	Name      string           `json:"name"`
	Type      ServiceTypeValue `json:"type"`
	MaxMemory int              `json:"maxMemory"`
	Versions
}
