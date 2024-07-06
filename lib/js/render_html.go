package lib_js

const RenderHtml = `
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
function renderHtml(widget) {
    const app = document.querySelector("#app");
    if (is_null(app)) {
        return Logger.log("app not found");
    }

    const root = document.createElement(widget.tag);
    root.id = widget.id;
    root.innerHTML = widget.content;

    for (const child of widget.children) {
        root.appendChild(renderHtml(child));
    }

    app.replaceChildren(root);
}
window.renderHtml = renderHtml;
`
