package utils

import (
	"log"
	"time"
)

func GetStringPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func GetPointerFromString(value string) *string {
	return &value
}

func ChooseString(newVal string, oldVal string) string {
	if newVal == "" {
		return oldVal
	}
	return newVal
}

func ChoosePointerString(newVal string, oldVal *string) *string {
	if newVal == "" {
		return oldVal
	}
	return &newVal
}

func ChooseTimeReq(req *time.Time, oldValue time.Time) time.Time {
	if req != nil {
		return *req
	}
	return oldValue
}

func ChooseTimeReqFromString(req *string, oldValue time.Time) time.Time {
	if req != nil && *req != "" {
		parsedTime, err := time.Parse("2006-01-02", *req)
		if err != nil {
			log.Printf("Error parsing POExpiryDate: %v", err)
			return oldValue
		}
		return parsedTime
	}
	return oldValue
}

func ChooseStringReqPointer(req *string, oldValue string) string {
	if req != nil {
		return *req
	}
	return oldValue
}

func ChoosePointerUint(req *uint, oldValue uint) uint {
	if req != nil {
		return *req
	}
	return oldValue
}

func ChoosePointerFloat64(req *float64, oldValue float64) float64 {
	if req != nil {
		return *req
	}
	return oldValue
}

func ChoosePointerFloat64ReturnPointer(req *float64, oldValue float64) *float64 {
	if req != nil {
		return req
	}
	return &oldValue
}

func ChoosePointerFloat32(req *float32, oldValue float32) float32 {
	if req != nil {
		return *req
	}
	return oldValue
}

func ChooseBoolToInt(req *bool, oldValue int) int {
	if req != nil {
		if *req {
			return 1
		}
		return 0
	}
	return oldValue
}
