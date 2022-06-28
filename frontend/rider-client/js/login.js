$(document).ready(function() {
    $("#login-button").click(login)
    $("#register-button").click(register)
})

function login() {
    $.ajax({
        url: "http://localhost:8080/api/user/logout",
        xhrFields: {
            withCredentials: true
        }
    })
    $("#longitude").val(randomNum(119.3660541, 119.4153504))
    $("#latitude").val(randomNum(32.3322146, 32.3987607))
    // $("#longitude").val(119.39537596307396)
    // $("#latitude").val(32.3361488487744)
    $.ajax({
        url: "http://localhost:8080/api/user/login",
        type: "get",
        xhrFields: {
            withCredentials: true
        },
        data: $("#login-form").serialize(),
        success: function(res) {
            if(res.code === '0') {
                window.location.href = "../rider-client/index.html"
            }
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function register() {
    $.ajax({
        url: "http://localhost:8080/api/user/register",
        type: "get",
        xhrFields: {
            withCredentials: true
        },
        data: $("#login-form").serialize(),
        success: function(res) {
            if(res.code === '0') {
                login()
            }
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function randomNum(minNum, maxNum){
    return Math.random()*(maxNum-minNum)+minNum
}