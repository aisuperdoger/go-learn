package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// SimpleORM 简易ORM结构体
type SimpleORM struct {
    db *sql.DB
}

// 假设的User结构体
type User struct {
    ID   int    `orm:"id"`
    Name string `orm:"name"`
}

// Insert 模拟插入操作
func (o *SimpleORM) Insert(entity interface{}) error {
    // 1. 获取结构体类型和值的反射对象
    t := reflect.TypeOf(entity).Elem()  // 获取指针指向的类型
    v := reflect.ValueOf(entity).Elem() // 获取指针指向的值

    tableName := t.Name() // 简化处理：直接用结构体名作表名
    var fields []string
    var placeholders []string
    var args []interface{}

    // 2. 遍历结构体字段，解析标签，收集字段名和值
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        ormTag := field.Tag.Get("orm") // 获取orm标签值
        if ormTag != "" {
            fields = append(fields, ormTag)
            placeholders = append(placeholders, "?")
            fieldValue := v.Field(i).Interface() // 获取字段值
            args = append(args, fieldValue)
        }
    }

    // 3. 动态生成SQL
    sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
        tableName,
        strings.Join(fields, ", "),
        strings.Join(placeholders, ", "))

    // 4. 执行SQL
    _, err := o.db.Exec(sql, args...)
    return err
}

// First 模拟查询单条记录
func (o *SimpleORM) First(dest interface{}, where string, args ...interface{}) error {
    // 1. 获取目标结构体信息
    destValue := reflect.ValueOf(dest).Elem()
    destType := destValue.Type()

    tableName := destType.Name() // 简化处理

    // 2. 生成查询SQL
    sql := fmt.Sprintf("SELECT * FROM %s WHERE %s LIMIT 1", tableName, where)

    // 3. 执行查询
    row := o.db.QueryRow(sql, args...)

    // 4. 准备扫描结果的容器
    // 获取结构体所有字段的指针，用于rows.Scan
    var scanArgs []interface{}
    for i := 0; i < destType.NumField(); i++ {
        field := destType.Field(i)
        if ormTag := field.Tag.Get("orm"); ormTag != "" {
            // 获取字段的指针
            fieldAddr := destValue.Field(i).Addr().Interface()
            scanArgs = append(scanArgs, fieldAddr)
        }
    }

    // 5. 将数据库结果扫描到结构体中
    return row.Scan(scanArgs...)
}

// 使用示例
func main3() {
    db, _ := sql.Open("mysql", "user:password@/dbname") // 需替换为真实DSN
    orm := &SimpleORM{db: db}

    user := &User{Name: "Alice"}
    // 插入
    err := orm.Insert(user)
    if err != nil {
        log.Fatal(err)
    }

    // 查询
    var fetchedUser User
    err = orm.First(&fetchedUser, "id = ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Fetched User: %+v\n", fetchedUser)
}