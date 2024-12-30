import { defineConfig } from 'vite';

export default defineConfig({
  server: {
    port: 3000,
    headers: {
      'Cross-Origin-Opener-Policy': 'same-origin',
      'Cross-Origin-Embedder-Policy': 'require-corp',
    },
  },
  optimizeDeps: {
    exclude: ['wasm_exec.js']
  },
  build: {
    rollupOptions: {
      input: {
        main: './index.html'
      }
    }
  }
}); 