/*
 * Believe it or not, you can declare and use functions in EJS templates too.
 */

var ejs = require("./js/ejs/ejs.js");
var file = require("file");
var utils = require("utils");

function fileLoader(filePath) {
  var r = file.read(filePath);
  if(r.err){
    throw r.err
  }
  return utils.toString(r.data);
}

ejs.fileLoader = fileLoader;

var data = {
  users: [
    { name: 'Tobi', age: 2, species: 'zengming' },
    { name: 'Loki', age: 2, species: 'ferret' },
    { name: 'Jane', age: 6, species: new Date().toString() },
  ],
  version: ejs.VERSION,
};

var path = './js/ejs/functions.ejs';
var content = fileLoader(path);
var func = ejs.compile(content, { filename: path });
var ret = func(data);
response.write(ret)

