// js

// input
var notifier = document.getElementById("notifier");
var comments = document.getElementById("comments");
var taskID = document.getElementById("taskID");

// buttons
var submit = document.getElementById("submit");
var resume = document.getElementById("resume");
var status = document.getElementById("status");

// boxes
var submit_task_box = document.getElementById("submit_task_box");
var link = document.getElementById("link");
var clipboardLink = new ClipboardJS('#link');
var paste = document.getElementById("paste");
var clipboardImg = new ClipboardJS('#paste');
var err_box = document.getElementById("error_box");
var error = document.getElementById("error");
var result_box = document.getElementById("result_box");

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
            err_box.hidden = true;
            console.debug("id:", json.id + ", link:" + json.link);
            link.value = json.link;
            var pngLink = '<img src="' + json.link + '">';
            paste.value = pngLink;
            submit_task_box.hidden = false;
        } else {
            submit_task_box.hidden = true;
            error.textContent = json.err;
            err_box.hidden = false;
            console.debug("err:", json);
        }
    }
    var d = new Date();
    var data = JSON.stringify({
        "notifier": notifier.value,
        "comments": comments.value,
        "offset": d.getTimezoneOffset()
        });
    xhr.send(data);

    // after submit new task, reset the tooltip
    var tooltip = document.getElementsByClassName("tooltip");
    for (i = 0; i < tooltip.length; i++) {
        tooltip[i].children[1].textContent = "Click to copy!";
    };
};

link.onclick = function (){
    this.nextElementSibling.textContent = "Copied!";
};

paste.onclick = function (){
    this.nextElementSibling.textContent = "Copied!";
};

function toggleID(id, hidden) {
    var x = document.getElementById(id);
    if (hidden) {
        x.hidden = true;
    }else{
        if (x.hidden) {
            x.hidden = false;
        } else {
            x.hidden = true;
        }
    }
}

function toggle(id, others=[]) {
    toggleID(id);
    others.forEach(element => {
        toggleID(element, true)
    });
    reset();
}

function reset() {
    err_box.hidden = true;
    submit_task_box.hidden = true;
}

resume.onclick = function () {
    var xhr = new XMLHttpRequest();
    var url = "/api/task/resume?id=" + taskID.value;
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    xhr.onreadystatechange = function () {
        if (xhr.readyState != 4) {
            return;
        }
        if (xhr.status == 200) {
            result_box.textContent = xhr.responseText;
        } else {
            result_box.textContent = xhr.responseText;
        }
        result_box.hidden = false;
        if (xhr.responseText){
            console.debug("result:", xhr.responseText);
        }
    }
    xhr.send();
};