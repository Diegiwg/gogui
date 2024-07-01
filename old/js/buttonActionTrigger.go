package gogui_js

const ButtonActionTrigger = `
function buttonActionTrigger(btn) {
  const data = {
    id: btn.id
  }

  socket.send(JSON.stringify({
    action: "button-click",
    data: data,
  }));
};
window.buttonActionTrigger = buttonActionTrigger;
`
