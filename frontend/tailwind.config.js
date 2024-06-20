/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{svelte,js,ts,jsx,tsx}"],
  plugins: [require("daisyui")],
  theme: {
    extend: {
      animation: {
        "spin-slow":
          "spin 1s cubic-bezier(0.65, 0, 0.35, 1) 0s 1 normal forwards",
      },
    },
  },
  daisyui: {
    themes: [
      "emerald",
      "dark",
      "cyberpunk",
      "dracula",
      "pastel",
      "bumblebee",
      "light",
      "forest",
    ],
  },
};
