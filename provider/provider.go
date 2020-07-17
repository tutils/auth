package provider

type Provider interface {
	String() string
	Register()
	Login()
	Logout()
}
