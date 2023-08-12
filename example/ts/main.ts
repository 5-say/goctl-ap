import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import * as user from './api/user'

// 默认配置
user.http.defaults.baseURL = 'http://localhost:8888'
user.http.defaults.timeout = 3000

// 添加请求拦截器
user.http.interceptors.request.use(function (config) {
  config.headers.Authorization = 'localstorge'
  return config;
}, function (error) {
  return Promise.reject(error);
});

// 添加响应拦截器
user.http.interceptors.response.use(function (response) {
  // 2xx 范围内的状态码都会触发该函数。
  // 更新 newtoken
  return response.data.data;
}, function (error) {
  console.log("响应拦截器 Error: ", error)
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  return Promise.reject(error);
});

user.client_public_signUpAccount({
  account: "string",     // 账号
  password: "string",    // 密码（前端加密后提交）
  captcha_code: "string" // 图形验证码
}).catch((reason)=>{
  console.log("catch", reason.config.headers)
}).then((response)=>{
  console.log(response)
})

createApp(App).mount('#app')
