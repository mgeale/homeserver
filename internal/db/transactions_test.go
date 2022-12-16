package db

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestTransactionModelGet(t *testing.T) {
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	tests := []struct {
		name            string
		transactionID   int
		wantTransaction *Transaction
		wantError       error
	}{
		{
			name:          "Valid ID",
			transactionID: 1,
			wantTransaction: &Transaction{
				ID:      1,
				Name:    "name",
				Amount:  100.00,
				Date:    "2018-12-23 17:25:22",
				Type:    "Repayment",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
			},
			wantError: nil,
		},
		{
			name:            "Zero ID",
			transactionID:   0,
			wantTransaction: nil,
			wantError:       ErrRecordNotFound,
		},
		{
			name:            "Non-existent ID",
			transactionID:   22222222,
			wantTransaction: nil,
			wantError:       ErrRecordNotFound,
		},
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := TransactionModel{db, infoLog, errorLog}

			transaction, err := m.Get(tt.transactionID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if !reflect.DeepEqual(transaction, tt.wantTransaction) {
				t.Errorf("want %v; got %v", tt.wantTransaction, transaction)
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

			id, err := m.Insert(tt.name, tt.transaction.Amount, tt.transaction.Date, tt.transaction.Type)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			if id == 0 {
				t.Error("want valid id; got 0")
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
				ID:     1,
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
				ID:     99999999,
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
		transID   int
		wantError error
	}{
		{
			name:      "Valid ID",
			transID:   1,
			wantError: nil,
		},
		{
			name:      "Non-existent ID",
			transID:   22222222,
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