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
        let targetId = app;

        /** @type {EventRequest} */
        const event = JSON.parse(msg.data);
        switch (event.action) {
            case "render-html":
                if (event.data.targetId !== "app") {
                    targetId = document.getElementById(event.data.targetId);
                    if (!targetId) {
                        console.log(
                            "target with id " +
                                event.data.targetId +
                                " not found"
                        );
                        return;
                    }
                }

                window.renderHtml(targetId, event.data.widget, true);
                break;

            case "delete-widget":
                window.deleteWidget(event.data);
                break;

            case "update":
                window.updateElementContent(event.data);
                break;
        }
    };
