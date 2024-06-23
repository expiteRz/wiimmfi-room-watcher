const ws = new WebSocket("ws://localhost:24051/ws");

ws.onopen = (event) => {
    console.log("connection established");
};

ws.onerror = (error) => {
    console.log("connection failed: ", error.data);
};

ws.onmessage = (event) => {
    console.log(event.data);
};

ws.onclose = () => {
    console.log("connection closed");
};