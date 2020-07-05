package input

import "param/validation"

type Config struct {
	Money float64 `valid:"Required;" json:"money"`
}
func (u *Config) Valid(v *validation.Validation) {
	if u.Money > 10000 {
		v.SetError("money", "请输入正确的金额!")
	}
}
