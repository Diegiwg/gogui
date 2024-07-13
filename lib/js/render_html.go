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
 * @property {string} data
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

        try {
            document.getElementById(widget.id)[attr.name] = attr.value;
        } catch (error) {
            // ignore
        }
    }

    for (const event of widget.events) {
        el.addEventListener(event.name, (e) => {
            const data = collectEventData(e, event.data);
            console.log(data);

            window.socket.send(
                JSON.stringify({
                    id: widget.id,
                    action: widget.tag + "-" + event.name + "-" + widget.id,
                    data,
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

/**
 * @param {CustomEvent} event
 * @param {Array<string>} targetData
 * @returns
 */
function collectEventData(event, targetData) {
    const collectedData = {};
    if (!Array.isArray(targetData)) return collectedData;

    for (const target of targetData) {
        let parts = [];
        if (target.includes(".")) {
            parts = target.split(".");
        }

        if (parts.length === 0) {
            collectedData[target] = event[target];
            continue;
        }

        let inner = event;
        for (const part of parts) {
            let has = inner[part];
            if (!has) continue;

            inner = inner[part];
        }

        collectedData[target] = inner;
    }

    return collectedData;
}
`
