package helpers

import "github.com/go-playground/validator/v10"

func FormValidationError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " wajib diisi!"
	case "email":
		return fe.Field() + " harus diisi dengan format email yang valid!"
	case "min":
		return fe.Field() + " minimal " + fe.Param() + " karakter!"
	case "max":
		return fe.Field() + " maksimal " + fe.Param() + " karakter!"
	case "alphanum":
		return fe.Field() + " hanya boleh berisi huruf dan angka!"
	case "numeric":
		return fe.Field() + " hanya boleh berisi angka!"
	case "eqfield":
		return fe.Field() + " harus sama dengan " + fe.Param() + "!"
	case "alphanumunicode":
		return fe.Field() + " harus berisi karakter, huruf dan angka!"
	case "gt":
		return fe.Field() + " harus diatas 8 tahun!"
	default:
		return fe.Field() + " tidak valid!"
	}
}