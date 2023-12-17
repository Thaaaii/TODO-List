package backend

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Tasks    []Task `json:"tasks"`
}
