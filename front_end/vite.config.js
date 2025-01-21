import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [
    vue({
      template: {
        compilerOptions: {
          //允许s-开头的自定义组件
          isCustomElement: (tag) => tag.startsWith('s-')
        }
      }
    })
  ]
})