package gogui_js

const ButtonHandler = `
window.buttonActionTrigger = (btn) => {
  const data = {
    id: btn.id
  }

  socket.send(JSON.stringify({
    action: "button-click",
    data: data,
  }));
};
`
