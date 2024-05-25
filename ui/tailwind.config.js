import {nextui} from '@nextui-org/theme'

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
    './node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}'
  ],
  theme: {
    extend: {
      fontFamily: {
        satoshi: ["var(--font-satoshi)"],
      },
    },
  },
  darkMode: "class",
  plugins: [
    nextui({
      layout: {
        fontSize: {
          tiny: '0.75rem',
          small: '1rem',
          medium: '1rem',
          large: '1.25rem',
        },
        lineHeight: {
          tiny: '0.5rem',
          small: '.75rem',
          medium: '1rem',
          large: '1.25rem',
        },
        borderWidth: {
          small: '1px',
          medium: '1px',
          large: '1px',
        },
        radius: {
          small: '.25rem',
          medium: '.25rem',
          large: '.25rem',
        }
      },
      themes: {
        light: {
          colors: {}
        },
        dark: {
          colors: {
            primary: {
              50: '#39E2CB',
              100: '#39E2CB',
              200: '#39E2CB',
              300: '#39E2CB',
              400: '#39E2CB',
              500: '#39E2CB',
              600: '#39E2CB',
              700: '#39E2CB',
              800: '#39E2CB',
              900: '#39E2CB',
              foreground: '#000000',
              DEFAULT: '#39E2CB'
            },
            forward: {
              50: '#ffffff',
              100: '#FFFFFF66',
              200: '#FFFFFF59',
              300: '#FFFFFF4D',
              400: '#FFFFFF40',
              500: '#FFFFFF33',
              600: '#FFFFFF26',
              700: '#FFFFFF1A',
              800: '#FFFFFF0D',
              900: '#FFFFFF0D',
            }
          }
        }
      }
    })
  ],
}
