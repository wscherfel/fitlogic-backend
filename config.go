package fitlogic

import "time"

const (
	Secret = "FitLogic random secret"

	// JWTExpiration is default expiration time of a socket
	// currently its set to 5 days
	JWTExpiration = time.Hour * 24 * 5
)
