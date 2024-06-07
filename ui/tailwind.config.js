import { nextui } from "@nextui-org/theme";

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    letterSpacing: {
      widest: ".15em",
    },
    extend: {
      fontFamily: {
        satoshi: ["var(--font-satoshi)"],
      },
      width: {
        68: "17rem",
      },
    },
  },
  darkMode: "class",
  plugins: [
    nextui({
      defaultTheme: "dark",
      layout: {
        fontSize: {
          tiny: "0.75rem", // text-tiny
          small: "0.875rem", // text-small
          medium: "1rem", // text-medium
          large: "1.125rem", // text-large
        },
        lineHeight: {
          tiny: "1rem", // text-tiny
          small: "1.25rem", // text-small
          medium: "1.5rem", // text-medium
          large: "1.75rem", // text-large
        },
        radius: {
          small: "4px", // rounded-small
          medium: "12px", // rounded-medium
          large: "14px", // rounded-large
        },
        borderWidth: {
          small: "1px", // border-small
          medium: "2px", // border-medium (default)
          large: "3px", // border-large
        },
      },
      themes: {
        dark: {
          colors: {
            primary: {
              // how does forground and background work? Is it always one shade?
              background: "#000000",
              foreground: "#000000",
              DEFAULT: "#39E2CB",
              50: "#39E2CB",
              100: "#39E2CB",
              200: "#39E2CB",
              300: "#39E2CB",
              400: "#39E2CB",
              500: "#39E2CB",
              600: "#39E2CB",
              700: "#39E2CB",
              800: "#39E2CB",
              900: "#39E2CB",
            },
          },
        },
      },
    }),
  ],
};
