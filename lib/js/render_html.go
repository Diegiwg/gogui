package lib_js

const RenderHtml = `
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

/**
 * @param {HTMLElement} root
 * @param {Widget} widget
 */
function renderHtml(root, widget, replace = false, parent = null) {
    if (!isNewRender(root, widget)) return;

    const el = document.createElement(widget.tag);
    el.id = widget.id;
    el.style = widget.style;
    el.innerHTML = widget.content;

    el.setAttribute("data-widget", JSON.stringify(widget));

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

    renderChildren(el, widget);

    if (replace) {
        root.replaceChildren(el);
    } else {
        root.appendChild(el);
    }
}
window.renderHtml = renderHtml;

function isNewRender(root, widget) {
    const existing = document.getElementById(widget.id);
    if (!existing) return true;

    const existingWidget = JSON.parse(existing.getAttribute("data-widget"));
    if (existingWidget !== widget) return true;

    return false;
}
window.isNewRender = isNewRender;

function renderChildren(root, widget) {
    for (const child of widget.children) {
        renderHtml(root, child);
    }
}
window.renderChildren = renderChildren;
`
