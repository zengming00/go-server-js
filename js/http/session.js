// session.start() // 可选操作，如果没有调用，在set和get时会自动调用
// session.end()
session.set('aa', 'helloworld')
console.log(session.get('aa'))