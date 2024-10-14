package reqvalidator

import (
	"auth-service-test/pkg/consts"
	"regexp"
	"unicode"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2/log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("phone", validatePhoneNumber); err != nil {
		log.Fatalf("Couldn't register phone number validator, err=%v", err)
	}
	if err := validate.RegisterValidation("email", validateEmail); err != nil {
		log.Fatalf("Couldn't register email validator, err=%v", err)
	}
	if err := validate.RegisterValidation("nonZalgoText", validateText); err != nil {
		log.Fatalf("Couldn't register text validator, err=%v", err)
	}
	if err := validate.RegisterValidation("role", validateRole); err != nil {
		log.Fatalf("Couldn't register role validator, err=%v", err)
	}
}

func ReadRequest(c *fiber.Ctx, request interface{}) error {
	if err := c.BodyParser(request); err != nil {
		return err
	}

	return validate.StructCtx(c.Context(), request)
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	phoneRegex := `^(8|\+7|7)(\d{10})$` // Allows numbers starting with +7 or 8 or 7 and 10 digits

	return regexp.MustCompile(phoneRegex).MatchString(phoneNumber)
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	return regexp.MustCompile(emailRegex).MatchString(email)
}

func validateText(fl validator.FieldLevel) bool {
	text := fl.Field().String()

	for len(text) > 0 {
		runeValue, size := utf8.DecodeRuneInString(text)

		if unicode.Is(unicode.Mn, runeValue) {
			return false
		}

		text = text[size:]
	}

	return true
}

func validateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()

	return role == consts.UserRole || role == consts.AdminRole
}
