package core

func (u *User) IsRetryAuth() bool {
	return u.e == ERROR_NEED_RELOGIN && u.retryConfig.User.IsAutoReAuth
}

func (u *User) ForceRetryAuth() (*User, error) {
	u.e = nil // force nil
	return LoadConfig(*u.retryConfig).LoginPW(u.retryConfig.User.StaffId, u.retryConfig.User.Password)
}
