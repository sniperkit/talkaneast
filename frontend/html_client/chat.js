const socket = new WebSocket('ws://80.211.54.149:2148/ws');



function init() {
    console.log("Talka Neast");
    socket.onmessage = handleEvent;
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

function scrollToBottom(){
    const div = document.getElementById('log');
    div.scrollTop = div.scrollHeight - div.clientHeight;
 }

function sendFunc() {
    const message = document.getElementById("msg").value;
    socket.send(JSON.stringify({"event": "Message", "data": { "message": message }}));
    document.getElementById("msg").value = "";
}

function newNick() {
    const nick = document.getElementById("nickname").value
    socket.send(JSON.stringify({"event": "SetNick", "data": { "nickname": nick }}))
    document.getElementById("modalnick").classList.remove("is-active");
    socket.send(JSON.stringify({"event": "ListChannels", "data": {}}));
}

function openNickModel() {
    document.getElementById("modalnick").classList.add("is-active");
}

function handleEvent(event) {
    console.log(event.data)
    const data = JSON.parse(event.data)
    switch(data["event"]) {
        case "Message":
            handleMessage(data);
            break;
        case "ChannelsList":
            addChannels(data);
            break;
        case "Notification":
            handleMessage(data);
            break;
    }

}

function addChannels(data) {
    document.getElementById("channels").innerHTML = "";
    data["data"].forEach((element, index) => {
        document.getElementById("channels").innerHTML += "<button id=\"channel"+element["name"] + "\" class=\"button\">#" + element["name"] + "</button><br/>"
        document.getElementById("channel"+element["name"]).addEventListener("click", function(){
            joinChannel(element["name"])
        });
    });
}

function joinChannel(channel) {
    console.log(channel);
    document.getElementById("log").innerHTML = "";
    socket.send(JSON.stringify({"event": "SetChannel", "data": { "channel": channel }}))

}

function handleMessage(data) {
    document.getElementById("log").innerHTML += getMessageHTML(data["data"]);
    const div = document.getElementById('log');
    div.scrollTop = div.scrollHeight
}

function getMessageHTML(message) {
    return "<div class=\"card\"><strong>"+ message["username"] +":</strong><p>"+ message["message"] + "</p></div>"
}