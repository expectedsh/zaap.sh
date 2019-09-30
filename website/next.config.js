const withCss = require('@zeit/next-css')
const withSass = require('@zeit/next-sass')
const withImages = require('next-images')
const { resolve } = require('path')

module.exports = withCss(withSass(withImages(
  {
    cssModules: true,
    cssLoaderOptions: {
      importLoaders: 1,
      localIdentName: '[hash:base64:8]',
    },
    webpack(config) {
      config.resolve.alias['~'] = resolve(__dirname, 'src')
      config.resolve.alias['assets'] = resolve(__dirname, 'src', 'assets')
      return config
    }
  }
)))
