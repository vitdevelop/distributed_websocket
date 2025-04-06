document.getElementById("send_message").addEventListener("click", sendMessage, false);
document.getElementById("message").addEventListener("keyup", event => {
    if (event.key !== "Enter") return;
    document.getElementById("send_message").click();
    event.preventDefault();
});

function appendMessage(userMessage, outside) {
    const message = userMessage.message
    const user = userMessage.user;
    const message_template = `<div class="chat chat-${outside ? 'start' : 'end'} m-2">
                        <div class="chat-header">
                            <div>${user.name}</div>
                            <div class="text-xs uppercase font-semibold opacity-60">${user.instance}</div>
                        </div>
                        <div class="chat-image avatar">
                            <div class="w-10 rounded-full">
                                <img
                                        src="${user.image}"
                                        alt=""/>
                            </div>
                        </div>
                        <div class="chat-bubble bg-${outside ? 'gray' : 'green'}-200">${message}</div>
                    </div>`;
    let temp = document.createElement('template');
    temp.innerHTML = message_template;

    const messages = document.getElementById("messages");
    messages.appendChild(temp.content.firstChild);
    messages.scrollTo(0, messages.scrollHeight);
}

function appendConnectedUser(user) {
    const message_template = `<li id="user_${user.id}" class="list-row">
                        <div><img class="size-10 rounded-box object-cover"
                                  src="${user.image}" alt=""/></div>
                        <div>
                            <div>${user.name}</div>
                            <div class="text-xs uppercase font-semibold opacity-60">${user.instance}</div>
                        </div>
                    </li>`;
    let temp = document.createElement('template');
    temp.innerHTML = message_template;

    const messages = document.getElementById("connectedUsers");
    messages.appendChild(temp.content.firstChild);
    messages.scrollTo(0, messages.scrollHeight);
}

function removeConnectedUser(user) {
    document.getElementById("user_" + user.id).remove();
}

// -----------------------------------------------------------------------------------------

let socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = function (event) {
    console.log('Connection established!');
};

socket.onmessage = function (event) {
    handleWsMessage(event.data);
    // var messages = document.getElementById('messages');
    // messages.innerHTML += '<p>' + event.data + '</p>';
};

let currentUser = null;

function handleWsMessage(data) {
    let json_data = JSON.parse(data);
    switch (json_data.command) {
        case 1:
            //connected users
            for (let user of json_data.data) {
                appendConnectedUser(user);
            }
            break;
        case 2:
            //message
            appendMessage(json_data.data, true);
            break;
        case 3:
            //current user
            currentUser = json_data.data;
            break;
        case 4:
            //disconnect user
            let disconnectedUser = json_data.data;
            removeConnectedUser(disconnectedUser);
            break;
    }
}

function sendMessage() {
    let input = document.getElementById("message");
    const text = input.value;
    appendMessage({user: currentUser, message: text}, false);

    socket.send(JSON.stringify({message: text}));
    input.value = '';
}