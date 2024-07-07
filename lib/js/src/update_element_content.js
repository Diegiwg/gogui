function updateElementContent(targetId, html) {
    const target = document.getElementById(targetId);
    if (!target) {
        console.log("target with id "+ targetId + " not found");
        return;
    }

    const frag = document.createRange().createContextualFragment(html);
    if (!frag) {
        console.log("new contextual fragment could not be created");
        return;
    }

    // target.replaceChildren(frag);

    target.parentElement.replaceChild(frag, target);
}
window.updateElementContent = updateElementContent;
