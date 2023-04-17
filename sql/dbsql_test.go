package sql

import "testing"

func TestMySqlCRUD(t *testing.T) {
	MySqlCRUD("select * from hive.TBLS where DB_ID = ?", 1)
}
