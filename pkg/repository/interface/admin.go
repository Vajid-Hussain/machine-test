package interfaceRepository

type IAdminRepository interface {
	FetchAdminPassword(string) (string, error)
	FetchAdminID(string) (string, error)
}
