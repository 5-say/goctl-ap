import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import { instance as http } from "./api/apiRequester"
import * as api from './api/api'

// 默认配置
http.defaults.baseURL = 'http://localhost:8888'
http.defaults.timeout = 3000

// 添加请求拦截器
http.interceptors.request.use(function (config) {
  config.headers.Authorization = 'localstorge'
  return config;
}, function (error) {
  return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(function (response) {
  // 2xx 范围内的状态码都会触发该函数。
  // 更新 newtoken
  return response.data.data;
}, function (error) {
  console.log("响应拦截器 Error: ", error)
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  return Promise.reject(error);
});

api.admin_public_signInAccount({
  account: "string",     // 账号
  password: "string",    // 密码（前端加密后提交）
  captcha_code: "string" // 图形验证码
}).catch((reason)=>{
  console.log("catch", reason.config.headers)
}).then((response)=>{
  console.log(response)
})

api.user_self_foodList({
  page: 1,                   // 当前页码
  per_page: 10,              // 每页条目
  order_column: "string",    // 排序字段
  order_type: "asc",         // 排序类型，asc、desc
  search_cook_book_id_eq: 1, // 食谱 ID
  search_status_eq: "fully"  // 状态 fully.库存充足 needBuy.需要采购
})

api.user_self_foodDelete({
  ids: [1,2,6],      // 食材 ID
  password: "string" // 账户密码
})

createApp(App).mount('#app')
