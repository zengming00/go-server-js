var sql = require('sql')
var url = require('url')
var file = require("file");
var utils = require('utils')
var ejs = require("./js/ejs/ejs.min.js");

function fileLoader(filePath) {
  return utils.toString(file.read(filePath));
}

var html;

try {
  var db = sql.new("mysql", "root:root@/test2?parseTime=true&loc=" + url.queryEscape("Asia/Shanghai"))
  // var db = sql.new("mysql", "root:root@/test2")
  var rows = db.query("select * from users")

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


  var path = './js/ejs/sql.ejs';
  var content = fileLoader(path);
  var func = ejs.compile(content, { filename: path });
  html = func({ datas: rowDatas });
} catch (e) {
  log(e)
  throw e;
}

function log(data) {
  console.log("log:", data)
  console.log("log:", JSON.stringify(data, null, 2))
}

html;