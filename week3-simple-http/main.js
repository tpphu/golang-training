// Khong phai define package

// import express
var express = require('express')


function main() {
    // Step 1: new app/engine
    var app = express()
    // 1.1 port
    var port = 3000

    // Step 2: route & hannder
    app.get('/ping', function (req, res) {
        res.status(200).json({
            "message": "pong"
        })
    })
    // Step 3: run
    app.listen(port, () => console.log(
        `Example app listening at http://localhost:${port}`
    ))
}

main();
