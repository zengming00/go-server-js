exports.log = function (data) {
  console.log("log:", data)
  console.log("log json:", JSON.stringify(data, null, 2))
}

