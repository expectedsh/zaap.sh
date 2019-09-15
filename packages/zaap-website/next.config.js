const withCSS = require('@zeit/next-css')
const { resolve } = require('path')

module.exports = withCSS({
  webpack(config) {
    config.resolve.alias['~'] = resolve(__dirname)
    return config
  }
})
