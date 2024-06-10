document.addEventListener('DOMContentLoaded', function() {
    // Replace 'ws://localhost:8080' with the appropriate URL if your server is running on a different host or port
    var ws = new WebSocket('ws://localhost:8080');

    ws.onopen = function() {
        console.log('Connected to the WebSocket server');
        // Example: Send a message to the server once the connection is open
        ws.send('Hello, server!');
    };

    ws.onmessage = function(event) {
        // Log messages received from the server
        console.log('Message from server:', event.data);
    };

    ws.onerror = function(error) {
        // Handle any errors that occur
        console.log('WebSocket Error:', error);
    };

    ws.onclose = function(event) {
        // Handle the connection closing
        console.log('WebSocket connection closed:', event);
    };
});
