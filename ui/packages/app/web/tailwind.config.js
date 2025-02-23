module.exports = {
  purge: {
    enabled: process.env.NODE_ENV === 'production',
    content: ['./src/pages/**/*.{js,ts,jsx,tsx}', './src/components/**/*.{js,ts,jsx,tsx}']
  },
  darkMode: 'class',
  theme: {},
  variants: {},
  plugins: []
}
