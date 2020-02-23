var socket = new WebSocket("ws://localhost:8080/ws")

let connect = cb => {
    console.log("Attempting to connect...");

    socket.onopen = () => {
        console.log("Connection successful!");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("Closed Socket Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMessage = msg => {
    console.log("Sending Message: ", msg);
};

export { connect, sendMessage };