const path = require('path');

module.exports = {
  entry: './src/main.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'build'),
  },
  module: {
    rules: [
        {
            test: /\.js$/,
            exclude: /(node_modules)/,
            use: ['babel-loader']
        }
    ]
  },
  devServer: {
    static: {
      directory: path.join(__dirname, 'build'),
    },
    compress: true,
    port: 9000,
  },
};
