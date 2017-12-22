/*
 * This example demonstrates how to use Array.prototype.forEach() in an EJS
 * template.
 */

// var ejs = require("./js/ejs/ejs.js");
var ejs = require("./js/ejs/ejs.min.js");
var file = require("file");
var utils = require("utils");

function fileLoader(filePath) {
  return utils.toString(file.read(filePath));
}

ejs.fileLoader = fileLoader;

var data = {
  names: ['foo', 'bar', '"baz']
};

var result;

ejs.renderFile('./js/ejs/list.ejs', data, function (err, html) {
  if (err) {
    console.log(err);
    return
  }
  console.log(html);
  result = html;
});

console.log('ejs version:', ejs.VERSION)

result;