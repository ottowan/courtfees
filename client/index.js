var express = require("express");
var https = require("https");
var http = require("http");
var fs = require("fs");
var path = require("path");
var request = require("request");

const fetch = require("node-fetch");


const config = require("./config");

var endpoint = "https://localhost:8080";
var app = express();
app.set("views", "./views");
app.set("view engine", "ejs");

app.use(express.static(__dirname + "/views"));

// This line is from the Node.js HTTPS documentation.
// var options = {
//   key: fs.readFileSync("./cert/key.pem"),
//   cert: fs.readFileSync("./cert/cert.pem"),
//   passphrase: "xibPPk27"
// };

var options = {
  pfx: fs.readFileSync('./cert/CertCOJ.pfx'),
  passphrase: config.certPassword()
};


// Create a service (the app object is just a callback).
app.get("/", (req, res) => {
  res.render("index", {active : "index"});
});

// Create a service (the app object is just a callback).
app.get("/courtfees", (req, res) => {
  res.render("courtfees", {active : "courtfees"});
});
// Create a service (the app object is just a callback).
app.get("/arbitration", (req, res) => {
  res.render("arbitration", {active : "arbitration"});
});


https.createServer(options, app).listen(443);
console.log("Start server on port : 443");
