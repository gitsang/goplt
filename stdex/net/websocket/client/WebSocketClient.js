
let WebSocketClient = function () {
    let WebSocketClient = {}

    function DefaultWebSocket(host, handler) {
        let _host = host;
        let _handler = handler;
        let _isOpen = false;
        let _socket = new WebSocket(_host);
        _socket.binaryType = "arraybuffer";

        _socket.onopen = function(even) {
            _isOpen = true;
            _handler.onConnect(even);
        };

        _socket.onclose = function(even) {
            _isOpen = false;
            _handler.onDisconnect({host:_host, event:even});
        };

        _socket.onerror = function(err) {
            _isOpen = false;
            _handler.onDisconnect({host:_host, event:err});
        };

        _socket.onmessage = function(e) {
            let data = e.data;
            _handler.onMsg(data);
        };

        this.send = function(data) {
            _socket.send(data);
        };

        this.close = function() {
            _socket.close(1000, "normal");
        };

        this.isOpen = function() {
            return _isOpen
        }
    }

    try {
        WebSocketClient.EngineSocket = DefaultWebSocket;
    } catch (e) {
        console.error("WebSocketClient error ", e)
    }

    return WebSocketClient;
}();