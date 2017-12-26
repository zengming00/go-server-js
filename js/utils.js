var file = require('file');

exports.log = function (data) {
  console.log("log:", data);
  console.log("log json:", JSON.stringify(data, null, 2));
}

exports.fileLoader = function (filePath) {
  var r = file.read(filePath);
  if (r.err) {
    throw err;
  }
  return utils.toString(r.data);
}

exports.toInt = function (v, def) {
  def = def || 0;
  v = parseInt(v);
  return isNaN(v) ? def : v;
}