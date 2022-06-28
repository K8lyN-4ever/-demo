// index.js 调用接口的方法
// 引入封装的get/post请求方法
import {get, post } from '@/axios/http'

export const getD = (url, m) => get(url, m)
export const postD = (url, m) => post(url, m)