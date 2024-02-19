package utils

const (
	ResponseStatusSuccess = "SUCCESSFUL"
	ResponseStatusFail    = "FAILED"
	ResponseStatusProcess = "PROCESSING"
	ResponseStatusPending = "PENDING"
)

func ResponseWithData(status string, data map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status": status,
		"data":   data,
	}
}

func ResponseWithMessage(status string, message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}
