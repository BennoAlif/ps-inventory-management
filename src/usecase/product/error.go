package productusecase

import "errors"

var ErrProductNotFound = errors.New("product not found")
var ErrTotalPriceNotMatch = errors.New("total price not match")
var ErrChangeNotMatch = errors.New("change not match")
var ErrStockNotAvailable = errors.New("stock not available")
var ErrCustomerNotFound = errors.New("customer not found")
