package util

import "time"

const (

	//AppointLockDelayTime 预约锁过期延迟时间
	AppointLockDelayTime = 20 * time.Second
	//AppointLockUpdateDelayTime 预约更新延迟时间，小于预约锁过期延迟时间
	AppointLockUpdateDelayTime = 10 * time.Second
	//AppointLockExpireTime 预约锁过期时间
	AppointLockExpireTime = time.Minute

	//AppointmentExpirationRemind 预约到期短信messageType
	AppointmentExpirationRemind = "appointment.expiration_remind"
)

// AppointLockExpired 判断预约锁是否过期
func AppointLockExpired(score int64) bool {
	return time.Now().Add(AppointLockDelayTime).After(time.Unix(score, 0))
}

// AppointLockUpdateExpired 判断预约锁更新是否过期
func AppointLockUpdateExpired(score int64) bool {
	return time.Now().Add(AppointLockUpdateDelayTime).After(time.Unix(score, 0))
}

// BeyondAppointCancelTime 超过可提前取消时间
func BeyondAppointCancelTime(h float64, t time.Time) bool {
	return time.Now().After(t.Add(-time.Duration(h*60) * time.Hour))
}
