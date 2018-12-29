// js

var notifier = document.getElementById("notifier");
var comments = document.getElementById("comments");
var submit = document.getElementById("submit");

var box = document.getElementById("hidden_box");
var link = document.getElementById("link");
var paste = document.getElementById("paste");

submit.onclick = function () {
    var xhr = new XMLHttpRequest();
    var url = "/api/task/submit";
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState != 4) {
            return;
        }
        var json = JSON.parse(xhr.responseText);
        if (xhr.status == 200) {
            console.debug("id:", json.id + ", link:" + json.link);
            link.textContent = json.link;
            var pngLink = '<img src="' + json.link + '">';
            paste.textContent = pngLink;
            box.hidden = false;
        } else {
            console.debug("err:", json);
        }
    }
    var d = new Date();
    var data = JSON.stringify({
        "notifier": notifier.value,
        "comments": comments.value,
        "time": d.toJSON()
        });
    console.debug(data)
    xhr.send(data);
};