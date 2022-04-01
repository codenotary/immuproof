const path = require('path');

module.exports = {
  publicPath: '/',
  outputDir: path.resolve(__dirname, '../rest/internal/embed'),
  filenameHashing: false,
  transpileDependencies: [
    'vuetify'
  ]
};
