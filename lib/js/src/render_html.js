/**
 * @typedef Widget
 * @property {string} id
 * @property {string} tag
 * @property {string} content
 * @property {Widget[]} children
 */

/**
 * @param {Widget} widget
 */
function renderHtml(root, widget, replace = false) {
    const el = document.createElement(widget.tag);
    el.id = widget.id;
    el.innerHTML = widget.content;

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
