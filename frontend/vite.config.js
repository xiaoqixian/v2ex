import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import path from "path"

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'), // @ => src
    },
  },
  server: {
    port: 8000,
    proxy: {
      '/api': {
        target: process.env.VUE_ENV === 'online'
          ? 'http://62.234.20.127:8080/'
          : 'http://0.0.0.0:8080/',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      }
    }
  }
  // devServer: {
  //   port: 8000,
  //   proxy: {
  //     '/api': {
  //       target: process.env.VUE_ENV === 'online' 
  //         ? 'http://62.234.20.127:8080/' // 线上服务器地址
  //         : 'http://localhost:8080/',
  //       changeOrigin: true, // 允许跨域
  //       pathRewrite: {
  //         '^/api': '' // 重写路径
  //       }
  //     }
  //   },
  //   client: {
  //     overlay: false // 自适应窗口大小时webpack会报错
  //   },
  // }
})
