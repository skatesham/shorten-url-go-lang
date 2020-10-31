package config

import "os"

// UserKey key for user in envirolment variable
const UserKey = "USER_DB"
// DbNameKey key for user in envirolment variable
const DbNameKey = "DB_NAME"
// PasswdKey key for user in envirolment variable
const PasswdKey = "PASSWD"
// HostKey key for user in envirolment variable
const HostKey = "HOST"
// SslModeKey key for user in envirolment variable
const SslModeKey = "SSL_MODE"

// SetUpProfileConfiguration ...
func SetUpProfileConfiguration() {

	profile := os.Getenv("profile")

	if profile == "" {
		os.Setenv(UserKey, "postgres")
		os.Setenv(DbNameKey, "shorten")
		os.Setenv(PasswdKey, "shan")
		os.Setenv(HostKey, "172.21.0.2")
		os.Setenv(SslModeKey, "disable")
	}
}