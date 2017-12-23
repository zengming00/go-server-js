var sql = require('sql')
var url = require('url')
var utils = require('utils')

try {
  // var db = sql.new("mysql", "root:root@/test2?parseTime=true&loc=" + url.queryEscape("Asia/Shanghai"))
  var db = sql.new("mysql", "root:root@/test2")
  var rows = db.query("select * from users")

  while (rows.next()) {
    var data = rows.scan();
    utils.print(JSON.stringify(data, null, 2))
    utils.print("\r\n")
  }

  var err = rows.err()
  if (err) {
    log(err)
  }

} catch (e) {
  log(e)
}

function log(data) {
  console.log("log:", data)
  console.log("log:", JSON.stringify(data, null, 2))
}

