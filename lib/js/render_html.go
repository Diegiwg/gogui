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
 */

/**
 * @param {Widget} widget
 */
function renderHtml(root, widget, replace = false) {
    const el = document.createElement(widget.tag);
    el.id = widget.id;
    el.style = widget.style
    el.innerHTML = widget.content;

    for (const attr of widget.attributes) {
        el.setAttribute(attr.name, attr.value);
    }

    for (const event of widget.events) {
        el.addEventListener(event.name, () => {
            console.log({
                event: event.name,
                target: widget.tag,
                id: widget.id
            });
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
`
