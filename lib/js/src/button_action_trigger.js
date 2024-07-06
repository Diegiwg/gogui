function buttonActionTrigger(btn) {
    socket.send(
        JSON.stringify({
            id: btn.id,
            action: "click",
            data: {},
        })
    );
}
window.buttonActionTrigger = buttonActionTrigger;
