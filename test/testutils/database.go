package testutils

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"testing"

// 	_ "github.com/lib/pq"
// )

// var TestDB *sql.DB

// func SetupTestDB(t *testing.T) {
// 	connStr := os.Getenv("TEST_DATABASE_URL")
// 	if connStr == "" {
// 		t.Fatal("TEST_DATABASE_URL environment variable not set")
// 	}
// 	var err error
// 	TestDB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		t.Fatalf("failed to open test database: %v", err)
// 	}
// 	if err = TestDB.Ping(); err != nil {
// 		t.Fatalf("failed to connect to test database: %v", err)
// 	}
// 	fmt.Println("Connected to test database successfully")
// }

// func TeardownTestDB() {
// 	if TestDB != nil {
// 		TestDB.Close()
// 	}
// }

