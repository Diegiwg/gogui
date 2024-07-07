const socket = new WebSocket("ws://%s/ws");
window.socket = socket;

/**
 * @typedef EventRequest
 * @property {string} action
 * @property {string} data
 */

socket.onopen = () => {
    socket.send(JSON.stringify({ action: "html-content" }));
};

socket.onmessage =
    /** @param {{data: string}} msg */
    (msg) => {
        console.log("received: " + msg.data);

        /** @type {EventRequest} */
        const event = JSON.parse(msg.data);
        switch (event.action) {
            case "render-html":
                window.renderHtml(app, event.data, true);
                break;
            case "update":
                window.updateElementContent(event.data);
                break;
        }
    };
