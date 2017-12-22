/*
 * Believe it or not, you can declare and use functions in EJS templates too.
 */

var ejs = require("./js/ejs/ejs.js");
var file = require("file");
var utils = require("utils");

function fileLoader(filePath) {
  return utils.toString(file.read(filePath));
}

ejs.fileLoader = fileLoader; 

var data = {
  users: [
    { name: 'Tobi', age: 2, species: 'zengming' },
    { name: 'Loki', age: 2, species: 'ferret' },
    { name: 'Jane', age: 6, species: 'ferret' }
  ]
};

var path = './js/ejs/functions.ejs';
var content = fileLoader(path);
var ret = ejs.compile(content, {filename: path})(data);
console.log(ret);
console.log('ejs version:', ejs.VERSION);
ret;

