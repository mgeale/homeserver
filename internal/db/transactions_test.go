package db

import (
	"log"
	"os"
	"testing"
)

func TestTransactionModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name            string
		wantTransaction *Transaction
		wantError       error
	}{
		{
			name:      "Valid ID",
			wantError: nil,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			query := &Query{
				Filters: &Filter{
					Field: Field("Amount"),
					Kind:  Equal,
					Value: 100.00,
				},
				Sort: Sort{
					Field:     Field("created"),
					Direction: Ascending,
				},
				Limit: 1,
			}

			transactions, err := m.Get(query)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if len(transactions) != 1 {
				t.Errorf("want at only 1 balance")
			}
		})
	}
}

func TestTransactionModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		transaction *Transaction
		wantError   error
	}{
		{
			name: "Valid Transaction",
			transaction: &Transaction{
				Name:   "TNS-00999",
				Amount: 111.00,
				Date:   "2018-12-23 17:25:22",
				Type:   "Repayment",
			},
			wantError: nil,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			_, err := m.Insert(tt.name, tt.transaction.Amount, tt.transaction.Date, tt.transaction.Type)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestTransactionModelUpdate(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name        string
		transaction *Transaction
		wantError   error
	}{
		{
			name: "Valid ID",
			transaction: &Transaction{
				ID:     "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
				Name:   "TNS-00999",
				Amount: 111.00,
				Date:   "2018-12-23 17:25:22",
				Type:   "Repayment",
			},
			wantError: nil,
		},
		{
			name: "Non-existent ID",
			transaction: &Transaction{
				ID:     "99999999",
				Name:   "TNS-00999",
				Amount: 111.00,
				Date:   "2018-12-23 17:25:22",
				Type:   "Repayment",
			},
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			err := m.Update(tt.transaction.ID, tt.transaction.Name, tt.transaction.Amount, tt.transaction.Date, tt.transaction.Type)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}

func TestTransactionModelDelete(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name      string
		transID   string
		wantError error
	}{
		{
			name:      "Valid ID",
			transID:   "1c0d2b44-b0ce-11ed-b95f-dca632bb7cae",
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			transID:   "22222222",
			wantError: ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			err := m.Delete(tt.transID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}
		})
	}
}
