import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import topLevelAwait from "vite-plugin-top-level-await";
import wasm from "vite-plugin-wasm";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), wasm(), topLevelAwait()],
  base: "/sandbox/",
});
