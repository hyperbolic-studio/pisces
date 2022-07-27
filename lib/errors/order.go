package errors

import "fmt"

//ErrNoOrderInquiryProvided is returned when there wasn't an inquiry for an
//order even though it is required. Inquiry is required because it gives the
//order context.
type ErrNoOrderInquiryProvided struct {
	OrderID int64
}

func (err *ErrNoOrderInquiryProvided) Error() string {
	return fmt.Sprintf("inquiry is required on an order for order %d and there wasn't one provided", err.OrderID)
}
