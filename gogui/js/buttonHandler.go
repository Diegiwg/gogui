package gogui_js

const ButtonHandler = `
window.buttonActionTrigger = (btn) => {
  fetch("button?actionId=" + btn.id).finally(() => window.location.reload());
};
`
