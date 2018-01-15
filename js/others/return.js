var types = require('types');

var a = types.retNil();
// console.log('a:', a === null) // false
console.log('a:', a === undefined) // ture

var b = types.retNull();
// console.log(b === null) // true
console.log(b === undefined) // false

var c = types.retUndefined()
// console.log(c === null) // false
console.log(c === undefined)  // true
console.log('=a:', c === a)  // true
console.log('=b:', c === b)  // true