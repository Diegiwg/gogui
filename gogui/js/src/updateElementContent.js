function updateElementContent(targetId, html) {
    const target = document.getElementById(targetId);
    if (!target) {
        console.log(`target with id ${targetId} not found`);
        return;
    }

    const frag = document.createRange().createContextualFragment(html);
    if (!frag) {
        console.log("new contextual fragment could not be created");
        return;
    }

    console.info("function::updateElementContent -> content in target updated");
    target.replaceChildren(frag);
}
window.updateElementContent = updateElementContent;
