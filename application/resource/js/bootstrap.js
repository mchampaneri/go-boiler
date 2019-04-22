// Loading window variables ...
window._ = require('lodash');
window.axios = require('axios');
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
window.axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
window.axios.defaults.headers.common['X-CSRF-Token'] = $("meta[name=token]").attr("content") ;
