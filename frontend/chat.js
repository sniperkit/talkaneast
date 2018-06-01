var socket = new WebSocket('ws://80.211.54.149:2148/ws');



function init() {
    console.log("Talka Neast");
    socket.onmessage = logMessage
    document.getElementById("msg")
        .addEventListener("keyup", function(event) {
            event.preventDefault();
            if (event.keyCode === 13) {
                document.getElementById("sendmsg").click();
            }
        }
    );

    document.getElementById("nickname")
    .addEventListener("keyup", function(event) {
        event.preventDefault();
        if (event.keyCode === 13) {
            document.getElementById("setnick").click();
        }
    }
);
}

function sendFunc() {
    let message = document.getElementById("msg").value;
    socket.send(JSON.stringify({"event": "Message", "data": { "message": message }}));
    document.getElementById("msg").value = "";
}

function newNick() {
    const nick = document.getElementById("nickname").value
    socket.send(JSON.stringify({"event": "SetNick", "data": { "nickname": nick }}))
    document.getElementById("modalnick").classList.remove("is-active");
}

function openNickModel() {
    document.getElementById("modalnick").classList.add("is-active");
}

function logMessage(event) {
    console.log(event.data)
    let data = JSON.parse(event.data)
    document.getElementById("log").innerHTML += getMessageHTML(data["data"]);
    var elem = document.getElementById('log');
    elem.scrollTop = elem.scrollHeight;
}

function getMessageHTML(message) {
    return "<div class=\"card\"><strong>"+ message["username"] +":</strong><p>"+ message["message"] + "</p></div>"
}