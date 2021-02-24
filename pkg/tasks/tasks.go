package tasks

import "DailyDo-backend/pkg/database"

// Tasks struct
type Tasks struct {
	Store database.Store
}

// GetTask method
func (t *Tasks) GetTask(ID int) (string, error) {
	result, err := t.Store.Get(ID)
	if err != nil {
		return "", err
	}

	return result, nil
}
