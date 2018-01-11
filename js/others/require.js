var fmt = require('fmt')
var os = require('os')
var utils = require('utils')
var file = require('file')

var r = cache.get("a")

if (!r.ok) {
  console.log('set a')
  cache.set("a", 11.9, 10)
}

fmt.printf("%s %d %T \n", 'hello', 123, { a: true })

var list = require_list()
var str = JSON.stringify(list, null, 2)

console.log('require_list:', str)



var fs = require_get('fs')
if (!fs) {
  var myExports = {
    foo: function () {
      console.log('foooooo');
    },
    readFileSync: function (name) {
      var r = file.read(name);
      if (r.err) {
        throw r.err
      }
      return utils.toString(r.value);
    },
    existsSync: function (name) {
      var r = os.stat(name)
      return !r.err;
    }
  }
  var myModule = {
    exports: myExports,
    a: true,
    b: 1234,
    c: 'asdf',
    d: [1, 4, 6]
  };
  require_set('fs', myModule);
}
var fs = require('fs');
fs.foo();
var name = './test/a.js';
var r = fs.existsSync('asdf')
console.log('existsSync:', r)
var r = fs.readFileSync(name);
console.log(r.toString())

