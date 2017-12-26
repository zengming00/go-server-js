/*
 * This example demonstrates how to use Array.prototype.forEach() in an EJS
 * template.
 */

// var ejs = require("./js/ejs/ejs.js");
var ejs = require("./js/ejs/ejs.min.js");
var file = require("file");
var utils = require("utils");
var myUtils = require('./js/utils.js');


ejs.fileLoader = myUtils.fileLoader;

var data = {
  names: ['foo', 'bar', '"baz']
};


ejs.renderFile('./js/ejs/list.ejs', data, function (err, html) {
  if (err) {
    console.log(err);
    return
  }
  response.header().set('Content-Type', 'text/html; charset=utf-8');
  response.write(html);
});
