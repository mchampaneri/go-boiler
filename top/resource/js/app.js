
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap');

window.Vue = require('vue');

//import vScroll from 'vue-scroll'
//
//Vue.use(VueScroll).

Vue.component('Editor', require('./components/Editor.vue'));

Vue.component('MyPublications', require('./components/MyPublications.vue'));

Vue.component('Search', require('./components/Search.vue'));

Vue.component('Reader', require('./components/Reader.vue'));

Vue.component('Publications', require('./components/Publications.vue'));

Vue.component('Library', require('./components/Library.vue'));

Vue.component('Profilepage', require('./components/Profilepage.vue'));

Vue.component('ProfileActions', require('./components/ProfileActions.vue'));

Vue.component('Followers', require('./components/Followers.vue'));
/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

const app = new Vue({
    el: '#app'
});
