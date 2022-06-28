<template>
  <div>
    <div class="flex xs9">
      <div id="map" />
    </div>

    <div class="flex">
      <div class="manage-list">
        <va-card color="divider">
          <va-card-title>
            系统自动化
          </va-card-title>
          <va-card-content>
            <va-chip class="mb-2 mr-2 add-button"
              small="true"
              color="#87CEEB"
              @click="generate()"
            >
              <span class="white-color">
                {{ generateStr }}
              </span>
            </va-chip>
            <va-chip class="mb-2 mr-2 add-button"
              small="true"
              color="#87CEEB"
              @click="dispatch()"
            >
              <span class="white-color">
                {{ dispatchStr }}
              </span>
            </va-chip>
          </va-card-content>
        </va-card>
      </div>
    </div>
  </div>
</template>

<script>
import BMap from 'BMap'
import { reactive, onMounted } from 'vue'
import { get } from '@/axios/html'
import { GPS } from '@/util/GPS.js'
import { WS } from '@/util/WebSocket.js'

export default {
  name: 'dashboart-map',
  setup() {
    function translate(lon, lat) {
      let truePoint = GPS.gcj_encrypt(lat,lon)
      truePoint = GPS.bd_encrypt(truePoint.lat, truePoint.lon)
      return truePoint
    }
    const state = reactive({
      map: null,
      orders: new Map(),
      points: new Map(),
      windows: new Map(),
      generateStr: "开启自动生成订单",
      dispatchStr: "开启自动分配订单",
      generateFlag: false,
      dispatchFlag: false,
    })
    onMounted(async() => {

      get("/api/admin/generateFlag", '')
        .then(res => {
          console.log(res.data)
          if(res.data) {
            state.generateFlag = true
            state.generateStr = "停止自动生成订单"
          }
        }).catch(err => {
          console.log(err)
        })

      get("/api/admin/dispatchFlag", '')
        .then(res => {
          console.log(res.data)
          if(res.data) {
            state.dispatchFlag = true
            state.dispatchStr = "停止自动分配订单"
          }
        }).catch(err => {
          console.log(err)
        })

      /* eslint-disable */
      // 创建Map实例
      //坐标转换完之后的回调函数

      var map = new BMap.Map("map")
      const point = translate(119.3968391418457, 32.35741571571183)
      map.centerAndZoom(new BMap.Point(point.lon, point.lat), 14)
      //添加地图类型控件
      map.addControl(new BMap.MapTypeControl({
        mapTypes:[BMAP_NORMAL_MAP, BMAP_HYBRID_MAP]
      }))
      // 设置地图显示的城市 此项是必须设置的
      //开启鼠标滚轮缩放
      map.enableScrollWheelZoom(true)
      /* eslint-enable */
      state.map = map
      get("/api/admin/getOrders", '')
        .then(res => {
          if(res.data.code === "0") {
            res.data.data.forEach(order => {
              console.log(order)
              state.orders.set(order.uuid, order)
              const src = translate(order.src.longitude, order.src.latitude)
              const point = new BMap.Point(src.lon, src.lat)
              const marker = new BMap.Marker(point)
              state.points.set(order.uuid, marker)
              state.map.addOverlay(marker)

              const opts = {
                width: 200,
                height: 100,
                title: order.uuid
              }
              const context = "<p>" + "未被接取" + "</p><br><button type='button' class='delete' onclick='deleteOrder(\"" + order.uuid +  "\")'>删除订单</button>"
              const infoWindow = new BMap.InfoWindow(context, opts)
              state.windows.set(order.uuid, infoWindow)
              marker.addEventListener("click", function() {
                state.map.openInfoWindow(infoWindow, point)
              })
            });
          }
        })
        .catch(err => {
          console.log(err)
        })
      

      WS.client.onmessage = function(event) {
        const data = event.data
        // console.log(datas)
        const json = JSON.parse(data)
        switch(json.flag) {
          case "delete":
            state.map.removeOverlay(state.points.get(json.order_id))
            state.orders.delete(json.order_id)
            break
          case "append": {
            const order = JSON.parse(json.order_id)
            console.log(order.uuid)
            state.orders.set(order.uuid, order)
            const src = translate(order.src.longitude, order.src.latitude)
            const point = new BMap.Point(src.lon, src.lat)
            const marker = new BMap.Marker(point)
            state.points.set(order.uuid, marker)
            state.map.addOverlay(marker)
            const opts = {
                width: 200,
                height: 100,
                title: order.uuid
            }
            const context = "<p>" + "未被接取" + "</p><br><button type='button' class='delete' onclick='deleteOrder(\"" + order.uuid +  "\")'>删除订单</button>"
            const infoWindow = new BMap.InfoWindow(context, opts)
            state.windows.set(order.uuid, infoWindow)
            marker.addEventListener("click", function() {
              state.map.openInfoWindow(infoWindow, point)
            })
            break
          }
          case "accept": {
            const window = state.windows.get(json.order_id)
            window.setContent("<p>" + "已被接取" + "</p><br><button type='button' class='delete' onclick='deleteOrder(\"" + json.order_id +  "\")'>删除订单</button>")
            break
          }
          case "cancel": {
            const window = state.windows.get(json.order_id)
            window.setContent("<p>" + "未被接取" + "</p><br><button type='button' class='delete' onclick='deleteOrder(\"" + json.order_id +  "\")'>删除订单</button>")
            break
          }
        }
      }
    })



    return state
  },
  data () {
    return {
    }
  },
  methods: {
    generate() {
      if(!this.generateFlag) {
        get("/api/admin/startGenerate", '')
          .then(res => {
            console.log(res)
            this.generateFlag = true
            this.generateStr = "停止自动生成订单"
          }).catch(err => {
            console.log(err)
          })
      }else {
        get("/api/admin/stopGenerate", '')
          .then(res => {
            console.log(res)
            this.generateFlag = false
            this.generateStr = "开始自动生成订单"
          }).catch(err => {
            console.log(err)
          })
      }
    },
    dispatch() {
      if(!this.dispatchFlag) {
        get("/api/admin/startDispatch", '')
          .then(res => {
            console.log(res)
            this.dispatchFlag = true
            this.dispatchStr = "停止自动分配订单"
          }).catch(err => {
            console.log(err)
          })
      }else {
        get("/api/admin/stopDispatch", '')
          .then(res => {
            console.log(res)
            this.dispatchFlag = false
            this.dispatchStr = "开始自动分配订单"
          }).catch(err => {
            console.log(err)
          })
      }
    }
  },
}
</script>
<style>
  .BMap_cpyCtrl {
    display:none;
  }
  .anchorBL {
    display:none;
  }
  /* .delete {
  } */
  #map {
    z-index: 223;
    height: 750px;
    width: 1400px;
  }
</style>