package validator

import (
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"time"
)

func New(options govalidator.Options) *govalidator.Validator {
	return govalidator.New(options)
}

func init() {
	govalidator.AddCustomRule("timestamp", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "" {
			_, err := time.Parse(time.RFC3339, val)
			if err != nil {
				return fmt.Errorf("invalid value for from_created_at: %v", field)
			}

			return nil
		}
		return nil
	})

	govalidator.AddCustomRule("order", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "" {
			if val != "asc" && val != "desc" {
				return fmt.Errorf("invalid value for sort (use desc or asc): %v", field)
			}
			return nil
		}
		return nil
	})
}
