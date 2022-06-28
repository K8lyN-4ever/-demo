import axios from 'axios' // 引用axios
// import config from '@/axios/config.js'

// const _baseUrl = config.baseUrl.dev // 使用到代理
// const apiUrl = _baseUrl
// axios 配置
axios.defaults.timeout = 120000 // 设置接口响应时间
    // axios.defaults.baseURL = 'http://localhost:8080/' // 这是调用数据接口,公共接口url+调用接口名

// 增加token（先写了个固定的token，实际项目中应该是通过接口获取到token）
// axios.defaults.headers.common['X-Access-Token'] = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2MzkxMTYzMzcsInVzZXJuYW1lIjoiYWRtaW4ifQ.YPJ7BV_Pg27NtPVk0FfoYsTXRpR35KXA64mMDibUzHI';

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';



// 封装get请求
export function get(url, params = {}) {
    return new Promise((resolve, reject) => {
        axios
            .get(url, {
                params: params
            })
            .then(response => {
                resolve(response)
            })
            .catch(err => {
                reject(err)
            })
    })
}
// 封装post请求
export function post(url, data = {}) {
    return new Promise((resolve, reject) => {
        axios.post(url, data).then(
            response => {
                // console.log(response.data.code)
                resolve(response.data)
            },
            err => {
                reject(err)
            }
        )
    })
}