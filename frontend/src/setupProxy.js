const { createProxyMiddleware } = require('http-proxy-middleware');

module.exports = function(app) {
  app.use(
    // Должно совпадать с путем в fetch и route на бэке
    '/branches',
    createProxyMiddleware({
      target: 'http://localhost:12121',
      changeOrigin: true,
    })
  );
};