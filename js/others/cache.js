cache.set("a", 1234)
cache.set("b", "bbbb")

get("a")
get("b")
cache.del("b")
get("b")

var r = cache.get("c")
if (!r.ok) {
  console.log("set c")
  cache.set("c", Date.now(), 10)
} else {
  console.log("get c:", r.value)
}

function get(key) {
  var r = cache.get(key)
  if (r.ok) {
    console.log(key + ":" + r.value)
  } else {
    console.log(key + " not exists")
  }
}

get('i')
cache.add("i", 3, 10)
get('i')
cache.sub("i", 2, 10)
get('i')

console.log("--------------------------------------")