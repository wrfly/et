// js

// boxes
var link = document.getElementById("link");
var clipboardLink = new ClipboardJS('#link');
var paste = document.getElementById("paste");
var clipboardImg = new ClipboardJS('#paste');
var error = document.getElementById("error");

$("#submit")[0].onclick = function () {
    var r = $("#result_box")[0];
    var d = new Date();
    $.ajax({
        url: '/api/task/submit',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify({
            "notifier": $("#notifier").val(),
            "comments": $("#comments").val(),
            "offset": d.getTimezoneOffset()
            }),
        success: function (data) {
            $("#error_box")[0].hidden = true;
            console.debug("id:", data.id + ", link:" + data.link);
            $("#link")[0].value = data.link;
            var pngLink = '<img src="' + data.link + '">';
            $("#paste")[0].value = pngLink;
            $('#submit_task_box')[0].hidden = false;
        },
        error: function (data) {
            $('#submit_task_box')[0].hidden = true;
            $("#error").empty();
            $("#error").append(data.responseJSON.err);
            $("#error_box")[0].hidden = false;
        }
    });

    // after submit new task, reset the tooltip
    var tooltip = $("#tooltip");
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
    $("#error_box")[0].hidden = true;
    $('#submit_task_box')[0].hidden = true;
}

$('#resume')[0].onclick = function () {
    var r = $("#result_box")[0];
    $.ajax({
        url: '/api/task/resume?id=' + $('#taskID').val(),
        type: 'post',
        success: function (data) {
            console.info(data);
            r.hidden = false;
            r.html = data.responseJSON;
        },
        error: function (data){
            console.info(data);
            r.hidden = false;
            r.html = data.responseJSON;
        }
    });
};