var express = require("express");
var https = require("https");
var http = require("http");
var fs = require("fs");
var path = require("path");
var request = require("request");

var endpoint = "https://localhost:8080";
var app = express();
app.set("views", "./views");
app.set("view engine", "ejs");

// This line is from the Node.js HTTPS documentation.
var options = {
  key: fs.readFileSync("./key.pem"),
  cert: fs.readFileSync("./cert.pem"),
  passphrase: "xibPPk27"
};

// Create a service (the app object is just a callback).
app.get("/", (req, res) => {
  res.render("index");
});

app.get("/havecapital", (req, res) => {
  var req = https.request(
    {
      host: "localhost",
      port: 8080,
      path: '/courtfees',
      method: 'GET',
      rejectUnauthorized: false,
      requestCert: true,
      agent: false
    },
    function(res) {
      var body = [];
      res.on("data", function(data) {
        body.push(data);
      });

      res.on("end", function() {
        console.log(body.join(""));
      });

      console.log(body)
    }
  );
  req.end();

  req.on("error", function(err) {
    console.log(err);
  });


});

// async function fetchHaveCapital() {
//   let response = await fetch(endpoint + "/courtfees");
//   let responseJson = await response.json();
//   let fromServer = responseJson.myString;
//   console.log(fromServer);
// }

https.createServer(options, app).listen(443);
console.log("Start server on port : 443");
