const MESSAGE_BUFFER_LENGTH = 3;

function update_message_buffer(e, c, m) {
	div = document.createElement("div");
	div.innerHTML = `<hr/><h3>From: ${m.Author}</h3><div>${m.Content}</div>`;

	let el = document.getElementById(e);
	el.prepend(div);

	let n = el.childNodes;
	n.forEach((x, i) => {
		if (i >= MESSAGE_BUFFER_LENGTH) el.removeChild(x);
	});

	document.getElementById(`${e}_count`).innerHTML = c;
}

function post_comment() {
	var f = document.forms["addMessage"];
	var m = {
		recipient: f.recipient.value,
		content: f.message.value,
	}
	server.send(JSON.stringify(m));
	f.recipient.value = "";
	f.message.value = "";
}

function server_socket(url, onMessage) {
	var socket = new WebSocket(url);
	socket.onerror = function(error) {
		console.log(`error for ${url}: ${error.message}`);
	};
	socket.onmessage = onMessage;
	return socket;
}

var server = null;
window.onload = function() {
	server = server_socket("ws://localhost:3000/register", m => {
		var client_id = JSON.parse(m.data);
		var public_total = 0;
		var private_total = 0;

		document.getElementById("id_banner").innerText = client_id;
		document.getElementById("public_list_count").innerText = public_total;
		document.getElementById("private_list_count").innerText = private_total;

		server.onmessage = function(m) {
			var d = JSON.parse(m.data);

			switch (d.shift().toLowerCase()) {
			case "broadcast":
				public_total++;
				console.log(`m.data = ${m.data}`);
				update_message_buffer("public_list", public_total, d[0]);
				break;

			case "private":
				private_total++;
				console.log(`m.data = ${m.data}`);
				update_message_buffer("private_list", private_total, d[0]);
				break;
			}
		}
	})
}