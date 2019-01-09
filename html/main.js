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
    $(".tooltiptext").each(function(){
        this.innerHTML = 'Click to copy!';
    })
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
    var taskID = $('.taskID')[0].value;
    $.ajax({
        url: '/api/task/resume?id=' + taskID,
        type: 'post',
        success: function (data) {
            console.info(data.responseText);
            r.innerHTML = data.responseText;
        },
        error: function (data){
            console.error(data.responseText);
            r.innerHTML = data.responseText;
        }
    });
    r.hidden = false;
};

$('#status')[0].onclick = function () {
    var r = $("#status_box")[0];
    var children = r.children;
    var taskID = $('.taskID')[1].value;
    var d = new Date();
    $.ajax({
        url: '/api/task/get?id=' + taskID,
        type: 'get',
        success: function (data) {
            children["_state"].innerText = data.state;
            children["_comments"].innerText = data.comments;
            d.setTime(Date.parse(data.submitAt));
            children["_submit"].innerText = d.toLocaleString();
            var events = children["_events"];
            data.events.forEach(e => {
                d.setTime(Date.parse(e.time));
                var row = events.insertRow(-1);
                row.insertCell(0).innerHTML = d.toLocaleString();
                row.insertCell(1).innerHTML = e.ip;
                row.insertCell(2).innerHTML = e.ua;
            });
            $("#error_box")[0].hidden = false;
            r.hidden = false;
        },
        error: function (data){
            console.error(data);
            r.hidden = true;
            $("#error").empty();
            $("#error").append(data.responseText);
            $("#error_box")[0].hidden = false;
        }
    });
}

// select all
$("input").each(function(){
    this.onclick = function(){ this.select();}
})