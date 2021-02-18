window.onload = function() {
	var socket = new WebSocket("ws://localhost:3000/hello");
	socket.onmessage = m => {
		div = document.createElement("div");
		div.innerText = JSON.parse(m.data);
		document.body.append(div);
	};
}