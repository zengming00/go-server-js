var m = require("./b.js");
m.test();
var a = { a: 123 };
console.log("aaaa", 1, true, false, [11, 'asdf', true], a)
console.log(JSON.stringify(a))

function foo(locals, escapeFn, include, rethrow
  /**/) {
  var __line = 1
    , __lines = "<h1>Users</h1>\r\n\r\n<% function user(user) { %>\r\n  <li><strong><%= user.name %></strong> is a <%= user.age %> year old <%= user.species %>.</li>\r\n<% } %>\r\n\r\n<ul>\r\n  <% users.map(user) %>\r\n</ul>\r\n"
    , __filename = "d:\\workspace\\ejs\\examples\\functions.ejs";
  try {
    var __output = [], __append = __output.push.bind(__output);
    with (locals || {}) {
      ; __append("<h1>Users</h1>\r\n\r\n")
        ; __line = 3
        ; function user(user) {
          ; __append("\r\n  <li><strong>")
            ; __line = 4
            ; __append(escapeFn(user.name))
            ; __append("</strong> is a ")
            ; __append(escapeFn(user.age))
            ; __append(" year old ")
            ; __append(escapeFn(user.species))
            ; __append(".</li>\r\n")
            ; __line = 5
            ;
        }
      ; __append("\r\n\r\n<ul>\r\n  ")
        ; __line = 8
        ; users.map(user)
        ; __append("\r\n</ul>\r\n")
        ; __line = 10
    }
    return __output.join("");
  } catch (e) {
    rethrow(e, __lines, __filename, __line, escapeFn);
  }
}

function escapeFn(str) {
  return str;
}

function rethrow(e, __lines, __filename, __line, escapeFn) {
  console.log(e);
}

var datas = [
  { name: 'zengming', age: 20, species: '111' },
  { name: 'kathy', age: 10, species: '222' },
];
var str = foo({ users: datas }, escapeFn, null, rethrow);
str;
// console.log(str)