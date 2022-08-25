package helpers

type Response struct {
	StatusCode   int         `json:"statusCode,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Token        interface{} `json:"token_sekarang,omitempty"`
	RefreshToken interface{} `json:"token_refresh,omitempty"`
}
