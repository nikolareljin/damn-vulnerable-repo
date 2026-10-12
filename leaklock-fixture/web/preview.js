// Fixture for Leak Lock: intentionally insecure code. Never deploy or copy.
const { exec } = require("child_process");

function showPreview(element, markup) {
    element.innerHTML = markup;
}

function runFormula(text) {
    return eval(text);
}

function thumbnail(file) {
    exec("convert " + file + " -resize 64x64 thumb.png");
}

module.exports = { showPreview, runFormula, thumbnail };
