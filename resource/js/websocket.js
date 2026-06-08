export default class SolidWebSocket {
    constructor(url) {
        this.url = url;
        this.socket = new WebSocket(window.location.origin + this.url);
    }

    send(data) {
        this.socket.send(data);
    }

    on(event, callback) {
        switch (event) {
            case 'open':
                this.socket.onopen = callback;
                break;

            case 'message':
                this.socket.onmessage = callback;
                break;

            case 'close':
                this.socket.onclose = callback;
                break;

            case 'error':
                this.socket.onerror = callback;
                break;

            default:
                console.warn(`Unsupported event type: ${event}`);
        }
    }
}