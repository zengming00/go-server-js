function test() {
  try {
    var a = { haha: "hehe", b: 1234, c: [1, true] }
    console.log("test(): ", a, Error);
    for (var i=0; i<100000; i++){

    }
    throw new Error("hehe")
  } catch (e) {
    console.log(e)
    console.log(JSON.stringify(e))
  }
}
// 中文
exports.test = test;