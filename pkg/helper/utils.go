package helper

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func ConvertSliceToPostgresArray(slice []string) string {
	arrayString := "{"
	for i, value := range slice {
		// Escape any double quotes in the string value
		value = strings.Replace(value, `"`, `\"`, -1)
		// Append each string value to the arrayString
		arrayString += `"` + value + `"`
		// Add a comma separator except for the last element
		if i < len(slice)-1 {
			arrayString += ","
		}
	}
	arrayString += "}"
	return arrayString
}

func ParsePostgresArray(src string) []string {
	// Remove curly braces from the string
	src = strings.Trim(src, "{}")
	// Split the string by comma to get individual values
	values := strings.Split(src, ",")
	// Trim whitespace from each value
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() != "" {
				return false
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 {
				return false
			}
		case reflect.Bool:
			if field.Bool() != false {
				return false
			}
		// Add cases for other types as needed
		default:
			// Handle other types if necessary
		}
	}
	return true
}

func IntToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func CombineAndUniqueWithExclusion(a, b []int64, exclude ...int64) []int64 {
	// Create a map to store unique elements
	unique := make(map[int64]struct{})

	// Add elements of slice a to the map
	for _, val := range a {
		if !contains(exclude, val) {
			unique[val] = struct{}{}
		}
	}

	// Add elements of slice b to the map
	for _, val := range b {
		if !contains(exclude, val) {
			unique[val] = struct{}{}
		}
	}

	// Extract keys from the map to form the result slice
	result := make([]int64, 0, len(unique))
	for key := range unique {
		result = append(result, key)
	}

	return result
}

// Helper function to check if a slice contains a specific value
func contains(slice []int64, val int64) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func IsOrderValueValid(val string) bool {
	orders := []string{
		"asc",
		"desc",
	}

	for _, c := range orders {
		if c == val {
			return true
		}
	}
	return false
}

func PlaceholdersString(n int) string {
	placeholders := make([]string, n)
	for i := range placeholders {
		placeholders[i] = "$" + strconv.Itoa(i+1) + ""
	}
	return strings.Join(placeholders, ", ")
}

func ValidateStruct(req interface{}) error {
	validate := validator.New()
	return validate.Struct(req)
}

func IdIsInteger(id string) bool {
	var idPattern = regexp.MustCompile(`^[0-9]+$`)
	return idPattern.MatchString(id)
}

func ValidateURL(fl validator.FieldLevel) bool {
	url, ok := fl.Field().Interface().(string)
	if !ok {
		// Field is not a string
		return false
	}
	// Define the regex pattern
	pattern := `^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$`
	// Match the regex pattern
	matched, _ := regexp.MatchString(pattern, url)
	return matched
}

func ValidateNIPForIT(fl validator.FieldLevel) bool {
	nipInt := fl.Field().Int()
	nip := strconv.Itoa(int(nipInt))

	return ValidatePrefixIT(nip)
}

func ValidatePrefixIT(nip string) bool {
	// Check the length
	if len(nip) != 13 {
		return false
	}

	// Check the prefix "615"
	if nip[:3] != "615" {
		return false
	}

	return true
}

func ValidatePrefixNurse(nip string) bool {
	// Check the length
	if len(nip) != 13 {
		return false
	}

	// Check the prefix "303"
	if nip[:3] != "303" {
		return false
	}

	return true
}

func ValidateNIPForNurse(fl validator.FieldLevel) bool {
	nipInt := fl.Field().Int()
	nip := strconv.Itoa(int(nipInt))

	return ValidatePrefixNurse(nip)
}

// ValidateNIP checks if the NIP meets the specific criteria
func ValidateNIP(fl validator.FieldLevel) bool {
	nipInt := fl.Field().Int()
	nip := strconv.Itoa(int(nipInt))

	// Check the length
	if len(nip) != 13 {
		return false
	}

	// Check the fourth digit for gender
	genderDigit := nip[3]
	if genderDigit != '1' && genderDigit != '2' {
		return false
	}

	// Check the year part (fifth and sixth digits)
	yearStr := nip[4:6]
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 0 || year > (time.Now().Year()%100) {
		return false
	}

	// Check the month part (seventh and eighth digits)
	monthStr := nip[6:8]
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		return false
	}

	// Check the random digits part (ninth to thirteenth digits)
	randomDigits := nip[8:]
	if len(randomDigits) != 5 {
		return false
	}
	if _, err := strconv.Atoi(randomDigits); err != nil {
		return false
	}

	return true
}

func GenerateULID() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	return ulid.MustNew(ms, entropy).String()

}
