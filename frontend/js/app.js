'use strict'
var util = require("util")

var app = (function () {

    var getWSAddress = function () {
        var loc = window.location, wsurl
        if (loc.protocol === "https:") {
            wsurl = "wss:"
        } else {
            wsurl = "ws:"
        }
        wsurl += "//" + loc.host + "/ws"

        return wsurl
    }



    var hideToggle = function (element) {
        if (element.style.display === "none") {
            element.style.display = "block"
        } else {
            element.style.display = "none"
        }
    }

    return {
        init: function () {
            var socket = new WebSocket(getWSAddress())

            socket.onmessage = function (event) {
                var msg = JSON.parse(event.data)

                console.log(msg.dataType)
                switch(msg.dataType) {
                    case "uplink":
                        console.log(msg.data)
                        break
                }

            }
        },
    }

})()

app.init()
