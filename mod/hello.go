package mod

type SayRequest struct {
	Name string `json:"name"`
}

type SayResponse struct {
	Reply string `json:"reply"`
}

type RunRequest struct {
}

type RunResponse struct {
	Something string
}
