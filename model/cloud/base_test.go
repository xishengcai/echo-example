package cloud

import (
	"testing"
)

func TestLoadMySql(t *testing.T) {
	LoadCloudSQL()
	var tables = make([]string, 10)
	if err := db.Raw("show tables").Scan(&tables).Error; err != nil {
		t.Fatal(err)
	}
	t.Log("tables: ")
}
