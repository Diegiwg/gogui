/**
 * @typedef WidgetAttribute
 * @property {string} name
 * @property {string} value
 */

/**
 * @typedef WidgetEvent
 * @property {string} name
 */

/**
 * @typedef Widget
 * @property {string} id
 * @property {string} tag
 * @property {WidgetAttribute[]} attributes
 * @property {WidgetEvent[]} events
 * @property {string} content
 * @property {string} style
 * @property {Widget[]} children
 * @property {Widget} parent
 */


// /**
//  * @type {Widget[]}
//  */
// var widgetPool = [];
// window.widgetPool = widgetPool;

/**
 * @param {HTMLElement} root
 * @param {Widget} widget
 */
function renderHtml(root, widget, replace = false, parent = null) {
    // TODO: move this for proper check function

    // Check if widget already exists
    const existing = document.getElementById(widget.id);
    if (existing) {
        // Check if widget is the same
        if (existing.innerHTML === widget.content) {
            return;
        }

        existing.innerHTML = widget.content;
        return;
    }

    const el = document.createElement(widget.tag);
    el.id = widget.id;
    el.style = widget.style;
    el.innerHTML = widget.content;

    if (parent) {
        el.setAttribute("data-parent", parent.id);
    }

    for (const attr of widget.attributes) {
        el.setAttribute(attr.name, attr.value);
    }

    for (const event of widget.events) {
        el.addEventListener(event.name, () => {
            console.log({
                event: event.name,
                target: widget.tag,
                id: widget.id,
                action: widget.tag + "-" + event.name + "-" + widget.id,
            });

            socket.send(
                JSON.stringify({
                    id: widget.id,
                    action: widget.tag + "-" + event.name + "-" + widget.id,
                    data: {},
                })
            );
        });
    }

    for (const child of widget.children) {
        console.log(child);
        renderHtml(el, child);
    }

    if (replace) {
        root.replaceChildren(el);
    } else {
        root.appendChild(el);
    }
}
window.renderHtml = renderHtml;
