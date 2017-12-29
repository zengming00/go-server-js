
var header = response.header()
header.set('location', '/public')

response.writeHeader(302)

