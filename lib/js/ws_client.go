package lib_js

const WsClient = `
const socket = new WebSocket("ws://%s/ws");
window.socket = socket;

socket.onmessage =
    /** @param {{data: string}} msg */
    (msg) => {
        /** @type {string} */
        let action;
        /** @type {string} */
        let id;
        /** @type {string} */
        let data;

        [action, data] = msg.data.split("|");
        [action, id] = action.split(":");

        switch (action) {
            case "update-element-content":
                updateElementContent(id, data);
                break;
        }
    };
`
