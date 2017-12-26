
var sql = require('sql');
var time = require('time');
var utils = require('./js/utils.js');
var types = require('types');

var action = request.formValue('action');


var r = sql.open("sqlite3", "./js/db/test.db");
if (r.err) {
    throw r.err;
}
var db = r.db;

if (action === 'init') {
    var ddl = 'CREATE TABLE IF NOT EXISTS [users] ('
        + '[id] INTEGER PRIMARY KEY AUTOINCREMENT,'
        + '[firstname] TEXT,'
        + '[lastname] TEXT,'
        + '[phone] TEXT,'
        + '[email] TEXT,'
        + '[created_at] DATETIME'
        + ');';

    r = db.exec(ddl);
    utils.log(r)
    r = db.exec("INSERT INTO `users` VALUES ('3', 'fname1', 'lname1', '(000)000-0000', 'name1@gmail.com', CURRENT_TIMESTAMP);");
    utils.log(r);
    r = db.exec("INSERT INTO `users` VALUES ('4', 'zengming', 'lname1', '(000)000-0000', 'zengming00@qq.com', CURRENT_TIMESTAMP);");
    utils.log(r);

} else if (action === 'get_users') {
    var page = utils.toInt(request.formValue('page'), 1);
    var page_size = utils.toInt(request.formValue('page_size'), 10);
    var offset = (page - 1) * page_size;

    var result = {
        total: 0,
        rows: [],
    }

    var data = count(db, "select count(*) from users");
    result.total = data[0];
    result.rows = query(db, "select * from users limit " + offset + "," + page_size + ";");
    output(result);

} else if (action === 'remove_user') {
    var id = utils.toInt(request.formValue('id'), 0);
    var s = "delete from users where id=$id";
    /*
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

} else if (action === 'update_user') {
    var id = utils.toInt(request.formValue('id'), 0);
    var firstname = request.formValue('firstname');
    var lastname = request.formValue('lastname');
    var phone = request.formValue('phone');
    var email = request.formValue('email');

    var s = "update users set firstname='$firstname',lastname='$lastname',phone='$phone',email='$email' where id=$id";
    // //更新数据
    // stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    // checkErr(err)

    // res, err = stmt.Exec("astaxieupdate", id)
    // checkErr(err)

    // affect, err := res.RowsAffected()
    // checkErr(err)

    // fmt.Println(affect)

} else if (action === 'save_user') {
    var firstname = request.formValue('firstname');
    var lastname = request.formValue('lastname');
    var phone = request.formValue('phone');
    var email = request.formValue('email');

    var s = "insert into users(firstname,lastname,phone,email) values('$firstname','$lastname','$phone','$email')";

    // //插入数据
    // stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
    // checkErr(err)

    // res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
    // checkErr(err)

    // id, err := res.LastInsertId()
    // checkErr(err)

    // fmt.Println(id)
}

db.close();

function output(data) {
    response.write(JSON.stringify(data, null, 2));
}

function count(db, msql) {
    var ret = [];
    var r = db.query(msql);
    if (r.err) {
        throw r.err;
    }
    while (r.rows.next()) {
        var n = types.newInt();
        var err = r.rows.scan(n)
        if (err) {
            throw err;
        }
        n = types.intValue(n);
        ret.push(n);
    }
    var err = r.rows.err();
    if (err) {
        throw err;
    }
    err = r.rows.close();
    if (err) {
        throw err;
    }
    return ret;
}

function query(db, msql) {
    var ret = [];
    var r = db.query(msql);
    if (r.err) {
        throw r.err;
    }
    while (r.rows.next()) {
        var id = types.newInt();
        var firstname = types.newString();
        var lastname = types.newString();
        var phone = types.newString();
        var email = types.newString();
        var created_at = types.newString();

        var err = r.rows.scan(id, firstname, lastname, phone, email, created_at);
        if (err) {
            throw err;
        }

        ret.push({
            id: types.intValue(id),
            firstname: types.stringValue(firstname),
            lastname: types.stringValue(lastname),
            phone: types.stringValue(phone),
            email: types.stringValue(email),
            created_at: types.stringValue(created_at),
        });
    }
    var err = r.rows.err();
    if (err) {
        throw err;
    }
    err = r.rows.close();
    if (err) {
        throw err;
    }
    return ret;
}