package validator

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Funksiya uuid ni valid yoki yo'qligini tekshiradi
func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

// Funksiyaga yuborilgan url ni tekshiradi link bulgan holatda true aks holatda false qaytaradi
func IsUrl(url string) bool {
	urlRegex := regexp.MustCompile(`^(https?:\/\/)?([\da-z.-]+)\.([a-z.]{2,6})(\/[\w.-]*)*\/?$`)
	result := urlRegex.MatchString(url)
	return result
}

// Funksiyaga yuborilgan emailni haqiqatdan email yoki yo'qligini tekshiradi
func IsEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return false
	} else {
		return true
	}
}

// Agar time o'zgaruvchi null bo'lsa funksiya true qaytaradi aks holda false
func IsTimeNull(date1 time.Time) bool {
	return date1.IsZero()
}

// Funksiyaga kelgan string farmatidagi o'zgaruvchilarni to'g'riligini tekshiradi
func IsEmpty(value string) bool {
	return value == ""
}
