package user

type User struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	EstablishmentIds []string `json:"establishmentIds"`
}
