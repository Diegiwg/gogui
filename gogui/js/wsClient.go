package gogui_js

const WsClient = `
  const socket = new WebSocket("ws://%s/ws");
  window.socket = socket;

  socket.onmessage = (msg) => {
    window.location.reload()
};
`
