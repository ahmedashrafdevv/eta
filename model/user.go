package model

type UserLoginRequest struct {
	Username uint   `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

type User struct {
	EmpName     string `json:"emp_name"`
	EmpCode     uint   `json:"emp_code"`
	EmpPassword string `json:"emp_password"`
	SecLevel    int    `json:"sec_level"`
	FixEmpStore int    `json:"store_code"`
}
