package user

type User struct {
	Name             string   `json:"name"`
	EstablishmentIds []string `json:"establishmentIds"`
}
