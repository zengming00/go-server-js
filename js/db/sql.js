var sql = require('sql')
var url = require('url')
var utils = require('utils')
var types = require('types')

try {
  // var db = sql.new("mysql", "root:root@/test2?parseTime=true&loc=" + url.queryEscape("Asia/Shanghai"))
  var reUse = true;
  var r = sql.open("mysql", "root:root@/test2", reUse);
  if (r.err) {
    throw r.err;
  }
  var db = r.value;
  if (!r.isReUse) {
    db.setMaxOpenConns(100);
  }
  var stat = db.stats();
  response.write(JSON.stringify(stat, null, 2))
  r = db.query("select * from users")
  if (r.err) {
    throw r.err;
  }
  var rows = r.value;
  while (rows.next()) {
    var data = getData(rows)
    utils.print(JSON.stringify(data, null, 2))
    utils.print("\r\n")
  }
  var err = rows.err()
  if (err) {
    throw err;
  }
  err = rows.close();
  if (err) {
    throw err;
  }
  // err = db.close();
  // if (err) {
  //   throw err;
  // }
} catch (e) {
  console.log("ee: %j", e)
  throw e;
}

function getData(rows) {
  var arr = [
    types.newString(),
    types.newString(),
    types.newString(),
    types.newString(),
    types.newString(),
    types.newString(),
    types.newString(),
    types.newString(),
  ];
  var err = rows.scan.apply(rows, arr)
  if (err) {
    throw err;
  }
  response.write(JSON.stringify(arr, null, 2))
  return {
    // id: types.intValue(arr[0]),
    a: types.stringValue(arr[1]),
  }
}


