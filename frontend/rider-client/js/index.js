let MAP = null
let orders = new Map()
let points = new Map()
let windows = new Map()

function translate(lon, lat) {
    let truePoint = GPS.gcj_encrypt(lat,lon)
    truePoint = GPS.bd_encrypt(truePoint.lat, truePoint.lon)
    return truePoint
}

function start() {
    $.ajax({
        url: "http://localhost:8080/api/rider/grab",
        type: "get",
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            console.log(res)
            $("#start").attr("disabled", true)
            $("#stop").attr("disabled", false)
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function stop() {
    $.ajax({
        url: "http://localhost:8080/api/rider/unGrab",
        type: "get",
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            console.log(res)
            $("#start").attr("disabled", false)
            $("#stop").attr("disabled", true)

        },
        error: function(err) {
            console.log(err)
        }
    })
}

function logout() {
    $.ajax({
        url: "http://localhost:8080/api/user/logout",
        type: "get",
        xhrFields: {
            withCredentials: true,
        },
        success: function(res) {
            if(res.code === '0') {
                console.log(res)
                WS.close()
                window.location.href = "../rider-client/login.html"
            }
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function accept(uid) {
    $.ajax({
        url: "http://localhost:8080/api/rider/accept",
        type: "get",
        data: {
            uuid: uid
        },
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            console.log(res)
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function cancel(uid) {
    $.ajax({
        url: "http://localhost:8080/api/rider/cancel",
        type: "get",
        data: {
            uuid: uid,
        },
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            console.log(res)
        },
        error: function(err) {
            console.log(err)
        }
    })
}

function complete(uid) {
    $.ajax({
        url: "http://localhost:8080/api/rider/complete",
        type: "get",
        data: {
            uuid: uid,
        },
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            console.log(res)
        },
        error: function(err) {
            console.log(err)
        }
    })
}


$(document).ready(function() {

    $("#logout-button").click(logout)
    $("#start").click(start)
    $("#stop").click(stop)

    WS.connect()


    const map = new BMap.Map("map")
    const point = translate(119.3968391418457, 32.35741571571183)
    map.centerAndZoom(new BMap.Point(point.lon, point.lat), 14)
    //添加地图类型控件
    map.addControl(new BMap.MapTypeControl({
        mapTypes:[BMAP_NORMAL_MAP, BMAP_HYBRID_MAP]
    }))
    // 设置地图显示的城市 此项是必须设置的
    //开启鼠标滚轮缩放
    map.enableScrollWheelZoom(true)
    MAP = map

    $.ajax({
        url: "http://localhost:8080/api/rider/getOrders",
        xhrFields: {
            withCredentials: true
        },
        success: function(res) {
            let temp = JSON.parse(res)
            for(let i = 0;i < temp.length;i++) {
                console.log(temp[i])
                    orders.set(temp[i].Name, temp[i])
                    const src = translate(temp[i].Longitude, temp[i].Latitude)
                    const point = new BMap.Point(src.lon, src.lat)
                    const marker = new BMap.Marker(point)
                    points.set(temp[i].Name, marker)
                    MAP.addOverlay(marker)
                    const opts = {
                        width: 200,
                        height: 100,
                        title: temp[i].Name
                    }
                    const context = "<p>" + "未被接取" + "</p><button type='button' onclick='accept(\"" + temp[i].Name + "\")'>接取订单</button>"
                    const infoWindow = new BMap.InfoWindow(context, opts)
                    windows.set(temp[i].Name, infoWindow)
                    marker.addEventListener("click", function() {
                        MAP.openInfoWindow(infoWindow, point)
                    })
            }
        },
        error: function(err) {
            console.log(err)
        }
    })

    WS.client.onmessage = function(event) {
        const data = event.data
        // console.log(datas)
        const json = JSON.parse(data)
        switch(json.flag) {
            case "delete":
                console.log(json)
                console.log(points.get(json.order_id))
                MAP.removeOverlay(points.get(json.order_id))
                orders.delete(json.order_id)
                break
            case "append": {
                const order = JSON.parse(json.order_id)
                orders.set(order.uuid, order)
                const src = translate(order.src.longitude, order.src.latitude)
                const point = new BMap.Point(src.lon, src.lat)
                const marker = new BMap.Marker(point)
                points.set(order.uuid, marker)
                map.addOverlay(marker)
                const opts = {
                    width: 200,
                    height: 100,
                    title: order.uuid
                }
                const context = "<p>" + "未被接取" + "</p><button type='button' onclick='accept(\"" + order.uuid +  "\")'>接取订单</button>"
                const infoWindow = new BMap.InfoWindow(context, opts)
                windows.set(order.uuid, infoWindow)
                marker.addEventListener("click", function() {
                    MAP.openInfoWindow(infoWindow, point)
                })
                break
            }
            case "accept": {
                $("#stop").attr("disabled", true)
                $("#start").attr("disabled", false)
                const window = windows.get(json.order_id)
                window.setContent("<p>" + "已被接取" + "</p><button type='button' onclick='complete(\"" + json.order_id +  "\")'>完成订单</button><button type='button' onclick='cancel(\"" + json.order_id +  "\")'>取消订单</button>")
                break
            }
            case "cancel": {
                const window = windows.get(json.order_id)
                window.setContent("<p>" + "未被接取" + "</p><button type='button' onclick='accept(\"" + json.order_id +  "\")'>接取订单</button>")
                break
            }
        }
    }
})
