package pay


// Env is the environment type
type Env string

const (
	SANDBOX = iota
	PRODUCTION
)

// Service is an Mpesa Service
type Service struct {
	AppKey    string
	AppSecret string
	Env       int
}

// New return a new Mpesa Service
func New(appKey, appSecret string, env int) (Service, error) {
	return Service{appKey, appSecret, env}, nil
}


func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://api.safaricom.co.ke/"
	}
	return "https://sandbox.safaricom.co.ke/"
}
