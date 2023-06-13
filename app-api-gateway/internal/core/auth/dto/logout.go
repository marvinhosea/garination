package dto

const urlRegex = `^(http|https):\/\/[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

// AuthLogoutRequest ...
type AuthLogoutRequest struct {
	Name string
}

func (r AuthLogoutRequest) Validate() error {
	return nil
}

// AuthLogoutResponse ...
type AuthLogoutResponse struct {
	State string `json:"state"`
}
