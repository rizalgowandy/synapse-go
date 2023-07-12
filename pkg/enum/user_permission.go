package enum

type UserPermission string

var (
	UserPermissionUnverified     UserPermission = "UNVERIFIED"
	UserPermissionReceive        UserPermission = "RECEIVE"
	UserPermissionSendAndReceive UserPermission = "SEND-AND-RECEIVE"
	UserPermissionLocked         UserPermission = "LOCKED"
	UserPermissionClosed         UserPermission = "CLOSED"
	UserPermissionMakeItGoAway   UserPermission = "MAKE-IT-GO-AWAY"
)
