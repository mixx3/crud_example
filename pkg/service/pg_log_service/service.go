package pg_log_service

import "gorm.io/gorm"

type pgService struct {
	session *gorm.Session
}
