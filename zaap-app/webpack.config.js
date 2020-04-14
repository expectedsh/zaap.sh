const HtmlWebPackPlugin = require("html-webpack-plugin")
const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const {join} = require('path')

const cssLoader = (customLoaders = [], cssLoaderOptions = {}) => [
  process.env.NODE_ENV === 'production'
    ? MiniCssExtractPlugin.loader
    : 'style-loader',
  {
    loader: 'css-loader',
    options: cssLoaderOptions,
  },
  ...customLoaders,
  {
    loader: 'postcss-loader',
    options: {
      ident: 'postcss',
      plugins: [
        // eslint-disable-next-line global-require
        require('postcss-preset-env')(),
        // eslint-disable-next-line global-require
        require('autoprefixer')(),
      ],
    },
  }
]

module.exports = {
  entry: ['core-js', join(__dirname, 'src', 'index.js')],
  context: join(__dirname, 'src'),
  resolve: {
    extensions: ['.js', '.jsx'],
    alias: {
      '~': join(__dirname, 'src'),
      'stylesheets': join(__dirname, 'src', 'assets', 'stylesheets'),
    },
  },
  devServer: {
    hot: true,
    inline: true,
    open: true,
    historyApiFallback: {
      disableDotRule: true
    },
  },
  output: {
    path: join(__dirname, 'build'),
    publicPath: '/',
    filename: '[hash:8].js',
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
          options: {
            presets: [
              "@babel/preset-env",
              "@babel/preset-react"
            ],
            plugins: [
              "@babel/plugin-proposal-nullish-coalescing-operator",
              "@babel/plugin-proposal-optional-chaining",
              [
                "@babel/plugin-transform-runtime",
                {
                  "regenerator": true
                }
              ],
            ],
          },
        }
      },
      {
        test: /\.s[ac]ss$/i,
        use: cssLoader(['sass-loader'], {
          importLoaders: true,
          modules: {
            localIdentName: process.env.NODE_ENV === 'production'
              ? '[hash:base64]'
              : '[path][name]__[local]',
          },
        }),
        include: /\.module\.s[ac]ss$/,
      },
      {
        test: /\.s[ac]ss$/i,
        use: cssLoader(['sass-loader']),
        exclude: /\.module\.s[ac]ss$/,
      },
      {
        test: /\.css$/i,
        use: cssLoader(),
      },
      {
        test: /\.(png|jpe?g|gif|svg|mp4|webm|ogg|mp3|wav|flac|aac|woff2?|eot|ttf|otf)(\?.*)?$/,
        loader: "url-loader",
        options: {
          limit: process.env.NODE_ENV === 'production' ? false : 10000,
          name: 'assets/[hash:8].[ext]',
          fallback: 'file-loader',
        }
      },
    ]
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: join(__dirname, 'public', 'index.html'),
    }),
    new MiniCssExtractPlugin({
      filename: '[hash:8].css',
      chunkFilename: '[hash:8].css',
    }),
  ]
}
