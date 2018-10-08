var express = require("express");
var https = require("https");
var fs = require("fs");
var bodyParser = require("body-parser");
var request = require("request");

const config = require("./config");

//ProvincialCourt HaveCapital 
const rootURL = "https://localhost:8080/"

//var endpoint = "https://localhost:8080";
var endpoint = "https://localhost:8080";
var app = express();
app.set("views", "./views");
app.set("view engine", "ejs");
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.use(express.static(__dirname + "/views"));
app.use(function(req, res, next) {
  res.header("Access-Control-Allow-Origin", "*");
  res.header(
    "Access-Control-Allow-Headers",
    "Origin, X-Requested-With, Content-Type, Accept"
  );
  next();
});

var options = {
  pfx: fs.readFileSync("./cert/Cert.pfx"),
  passphrase: config.certPassword()
};

//Main : index
app.get("/", (req, res) => {
  res.render("index", { active: "index" });
});

//Main : courtfees
app.get("/courtfees", (req, res) => {
  res.render("courtfees", { active: "courtfees" });
});


//Main : arbitration
app.get("/arbitration", (req, res) => {
  res.render("arbitration", { active: "arbitration" });
});

app.get("/arbitration/:feeCapital/amountPerson/:amountPerson", (req, res) => {
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: rootURL+"arbitration?feeCapital=" + req.params.feeCapital+"&amountPerson="+ req.params.amountPerson,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

app.get("/courtfees/provincialcourt/havecapital/:feeCapital", (req, res) => {
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: rootURL+"provincialcourt/havecapital?feeCapital=" + req.params.feeCapital,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//ProvincialCourt NonCapital 
app.get("/courtfees/provincialcourt/noncapital/:feeCapital", (req, res) => {

  var URI = rootURL+"provincialcourt/noncapital?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//ProvincialCourt mortgage 
app.get("/courtfees/provincialcourt/mortgage/:feeCapital", (req, res) => {

  var URI = rootURL+"provincialcourt/mortgage?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//ProvincialCourt appeal or supreme 
app.get("/courtfees/provincialcourt/appeal_or_supreme/:feeCapital", (req, res) => {

  var URI = rootURL+"provincialcourt/appealorsupreme?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});


//ProvincialCourt non_and_havecapital 
///courtfees/provincialcourt/non_and_havecapital/0
app.get("/courtfees/provincialcourt/non_and_havecapital/:feeCapital", (req, res) => {

  var URI = rootURL+"provincialcourt/nonandhavecapital?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});





//ProvincialCourt compensation 
app.get("/courtfees/provincialcourt/compensation/:feeCapital", (req, res) => {

  var URI = rootURL+"provincialcourt/compensation?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

app.get("/courtfees/kwaengcourt/havecapital/:feeCapital", (req, res) => {
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: rootURL+"kwaengcourt/havecapital?feeCapital=" + req.params.feeCapital,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//kwaengcourt NonCapital 
app.get("/courtfees/kwaengcourt/noncapital/:feeCapital", (req, res) => {

  var URI = rootURL+"kwaengcourt/noncapital?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//kwaengcourt mortgage 
app.get("/courtfees/kwaengcourt/mortgage/:feeCapital", (req, res) => {

  var URI = rootURL+"kwaengcourt/mortgage?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});

//kwaengcourt appeal or supreme 
app.get("/courtfees/kwaengcourt/appeal_or_supreme/:feeCapital", (req, res) => {

  var URI = rootURL+"kwaengcourt/appealorsupreme?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});


//kwaengcourt non_and_havecapital 
///courtfees/kwaengcourt/non_and_havecapital/0
app.get("/courtfees/kwaengcourt/non_and_havecapital/:feeCapital", (req, res) => {

  var URI = rootURL+"kwaengcourt/nonandhavecapital?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});





//kwaengcourt compensation 
app.get("/courtfees/kwaengcourt/compensation/:feeCapital", (req, res) => {

  var URI = rootURL+"kwaengcourt/compensation?feeCapital=" + req.params.feeCapital
  console.log(URI)
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: URI,
      body: JSON.stringify({
        feeCapital: req.params.feeCapital
      }),
      rejectUnauthorized: false
    },
    (error, response, body) => {
      if (error) {
        res.json(JSON.parse(error));
        return console.dir(error);
      }
      //console.dir(JSON.parse(body));
      res.json(JSON.parse(body));
    }
  );
});


https.createServer(options, app).listen(443);
console.log("Start server on port : 443");
