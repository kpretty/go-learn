package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type tbl struct {
	TblId            sql.NullInt32
	CreateTime       sql.NullInt32
	DbId             sql.NullInt32
	LastAccessTime   sql.NullInt32
	OWNER            sql.NullString
	OwnerType        sql.NullString
	RETENTION        sql.NullInt32
	SdId             sql.NullInt32
	TblName          sql.NullString
	TblType          sql.NullString
	ViewExpandedText sql.NullString
	ViewOriginalText sql.NullString
	IsRewriteEnabled sql.NullString
}

// go 有类似 java 的 JDBC 设计，使用 database/sql 包

func MySqlCRUD(sqlStmt string, args ...any) {
	db, err := sql.Open("mysql", "root:980729@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	// 关闭数据库连接
	defer func() { _ = db.Close() }()
	// 用同一个 string 接收数据
	var row tbl
	// 查询数据
	query, err := db.Query(sqlStmt, args...)
	if err != nil {
		panic(err)
	}
	defer func() { _ = query.Close() }()
	for query.Next() {
		err := query.Scan(
			&row.TblId,
			&row.CreateTime,
			&row.DbId,
			&row.LastAccessTime,
			&row.OWNER,
			&row.OwnerType,
			&row.RETENTION,
			&row.SdId,
			&row.TblName,
			&row.TblType,
			&row.ViewExpandedText,
			&row.ViewOriginalText,
			&row.IsRewriteEnabled,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(row)
	}

	// 预处理
	prepare, err := db.Prepare(sqlStmt)
	if err != nil {
		panic(err)
	}
	rows, err := prepare.Query(args...)
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		_ = rows.Scan(
			&row.TblId,
			&row.CreateTime,
			&row.DbId,
			&row.LastAccessTime,
			&row.OWNER,
			&row.OwnerType,
			&row.RETENTION,
			&row.SdId,
			&row.TblName,
			&row.TblType,
			&row.ViewExpandedText,
			&row.ViewOriginalText,
			&row.IsRewriteEnabled,
		)
		fmt.Println(row)
	}

}
