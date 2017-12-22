/*
 * This example demonstrates how to use Array.prototype.forEach() in an EJS
 * template.
 */

// var ejs = require("./js/ejs/ejs.js");
var ejs = require("./js/ejs/ejs.min.js");
var file = require("file");
var utils = require("utils");
var redis = require('redis')

function fileLoader(filePath) {
  return utils.toString(file.read(filePath));
}


var data = {
  names: [ejs.VERSION],
};

var conn = redis.dial('tcp', ':6379')
var v = conn.do('keys', '*')
if (Array.isArray(v)) {
  v.forEach(function (item) {
    var str = utils.toString(item);
    data.names.push(str);
  });
}
conn.close()

var path = './js/ejs/list.ejs';
var content = fileLoader(path);
ejs.compile(content, { filename: path })(data);
