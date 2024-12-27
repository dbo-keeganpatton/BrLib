/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,js}"],
  theme: {
    extend: {
            colors: {
                // Github Dark-esque theme
                'git_black': '#0d1117',
                'git_black_med': '#161b22',
                'git_black_lite': '#21262d',
                'git_grey': '#89929b',
                'git_dark_white': '#c6cdd5',
                'git_white': '#ecf2f8',
                'git_red': '#fa7970',
                'git_yellow': '#faa356',
                'git_green': '#7ce38b',
                'git_lite_blue': '#a2d2fb',
                'git_blue': '#77bdfb',
                'git_purple': '#cea5fb'
            }
        },
  },
  plugins: [],
}

