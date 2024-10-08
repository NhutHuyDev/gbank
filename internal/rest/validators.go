package rest

import (
	"github.com/NhutHuyDev/sgbank/pkg/utils"
	"github.com/go-playground/validator/v10"
)

var currencyValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}

	return false
}
