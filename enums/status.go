package enums

const (
	StatusPublish = "publish"
	StatusDraft   = "draft"
	StatusThrash  = "thrash"
)

var validStatuses = map[string]bool{
	StatusPublish: true,
	StatusDraft:   true,
	StatusThrash:  true,
}

func IsValidStatus(status string) bool {
	return validStatuses[status]
}
