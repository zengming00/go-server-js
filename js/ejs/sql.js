var sql = require('sql')
var url = require('url')
var file = require('file');
var utils = require('utils')
var ejs = require('./js/ejs/ejs.min.js');
var myUtils = require('./js/utils.js');



try {
  // var db = sql.new("mysql", "root:root@/test2?parseTime=true&loc=" + url.queryEscape("Asia/Shanghai"))
  // var db = sql.new("mysql", "root:root@/test2")
  var r = sql.open("sqlite3", "./js/db/test.db")
  if (r.err) {
    throw r.err;
  }

  var db = r.db;
  r = db.query("select * from users")
  if (r.err) {
    throw r.err;
  }

  var rows = r.rows;
  var rowDatas = [];
  while (rows.next()) {
    var data = rows.scan();
    rowDatas.push(data);
  }

  var err = rows.err()
  if (err) {
    log(err)
    throw err;
  }
  db.close()
  log(rowDatas)
  var path = './js/ejs/sql.ejs';
  var content = myUtils.fileLoader(path);
  var func = ejs.compile(content, { filename: path });
  var html = func({ datas: rowDatas });
  response.write(html)
} catch (e) {
  log(e)
}

function log(data) {
  console.log("log:", data)
  console.log("log:", JSON.stringify(data, null, 2))
}
