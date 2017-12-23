var sql = require('sql')

var ddl = 'CREATE TABLE IF NOT EXISTS [users] ([id] INTEGER PRIMARY KEY AUTOINCREMENT, [username] TEXT, [realname] TEXT, [password] TEXT, [created_at] DATETIME);'

var db = sql.new("sqlite3", "./test.db")
var ret = db.query(ddl)
console.log(ret)
db.close()

/*
//插入数据
stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
checkErr(err)

res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
checkErr(err)

id, err := res.LastInsertId()
checkErr(err)

fmt.Println(id)
//更新数据
stmt, err = db.Prepare("update userinfo set username=? where uid=?")
checkErr(err)

res, err = stmt.Exec("astaxieupdate", id)
checkErr(err)

affect, err := res.RowsAffected()
checkErr(err)

fmt.Println(affect)

//查询数据
rows, err := db.Query("SELECT * FROM userinfo")
checkErr(err)

for rows.Next() {
    var uid int
    var username string
    var department string
    var created string
    err = rows.Scan(&uid, &username, &department, &created)
    checkErr(err)
    fmt.Println(uid)
    fmt.Println(username)
    fmt.Println(department)
    fmt.Println(created)
}

//删除数据
stmt, err = db.Prepare("delete from userinfo where uid=?")
checkErr(err)

res, err = stmt.Exec(id)
checkErr(err)

affect, err = res.RowsAffected()
checkErr(err)

fmt.Println(affect)

db.Close()
*/