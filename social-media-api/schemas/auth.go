package schemas

// type Signup struct {
// 	name     string `valid:"required,alphanum,stringlength(2|50)" json:"name"`
// 	email    string `valid:"required,email" json:"email"`
// 	password string `valid:"required,stringlength(8|),pattern(^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@#$!%^&*()_+\-=])[A-Za-z\d@#$!%^&*()_+\-=]{8,}$) json:"password"`
// }

type SignupJsonSchema struct {
	Name     string `json:"name"`
	Email    string ` json:"email"`
	Password string `json:"password"`
}

type LoginSchema struct {
	Email    string `json:"name"`
	Password string `json:"password"`
}
