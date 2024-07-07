package lib_js

const Base = `
const app = document.querySelector("#app");
if (is_null(app)) throw new Error("app not found");
window.app = app;
`
