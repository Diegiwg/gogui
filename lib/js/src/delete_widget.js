/**
 * @param {string} widgetId
 */
function deleteWidget(widgetId) {
    const widget = document.getElementById(widgetId);
    if (widget) {
        widget.remove();
    }
}
window.deleteWidget = deleteWidget;
