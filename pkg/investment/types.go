package investment

import (
	"errors"
)

// CreditAssigner is implemented by any value that has a Assign method.
// The Assign method is used to calculate the valid credits with the parameter
// passed to the method.
// Returns 0 in the int32 values if the assignation can not be done with the
// error errRemainderExists because the investment can not be assigned to get a
// remainder of 0 between all the assignations and valid credits
// Returns int32 > 0 values and error equals to nil when is possible to assign
// an investment between all the possible credits with a remainder of 0
type CreditAssigner interface {
	Assign(int32) (int32, int32, int32, error)
}

var (
	// ErrRemainderExists informs about the possible investment into all the
	// credits
	ErrRemainderExists = errors.New("remainder is not 0")
)

// Request is the struct to handle the request of the endpoint
type Request struct {
	Investment int32 `json:"investment"`
}

// Response is the struct to handle the response of the endpoint
type Response struct {
	CreditType300 int32 `json:"credit_type_300"`
	CreditType500 int32 `json:"credit_type_500"`
	CreditType700 int32 `json:"credit_type_700"`
}
