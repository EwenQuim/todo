module.exports = {
  mode: 'jit',
  purge: {
    enabled: ((process.env.ENV === 'production') ? true : false),
    content: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  },
  darkMode: 'media', // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
