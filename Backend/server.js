var Express = require("express");
var parser = require("body-parser");
var app = Express();
var mongo = require("mongodb");
app.use(parser.urlencoded());
app.use(parser.json());

app.post('/login', function(req, res){

})

app.listen(3000, function() {
	console.log("listening on port 3000");
})