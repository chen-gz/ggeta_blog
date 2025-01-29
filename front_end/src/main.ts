import { createApp } from 'vue'
import './style.css'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import About from './components/About.vue'
import Post from './components/Post.vue'
import PostList from './components/PostList.vue'
import Tags from './components/Tags.vue'
import Home from './components/Home.vue'
import PostEdit from './components/PostEdit.vue'
import Login from './components/Login.vue'
import Codeforces from './components/Codeforces.vue'
import Pub from './components/Pub.vue'
// import 'sober'

const routes = [
    {path: '/', name: 'Home', component: About},
    {path: '/tag/', name: 'Tags', component: Tags},
    // {path: '/cate/', name: 'Categories', component: Category},
    {path: '/post/:id', name: 'PostPage', component: Post},
    {path: '/post_edit/:id', name: 'PostEdit', component: PostEdit},
    // {path: '/login', name: 'Login', component: Login},
    {path: '/tag/:id', name: 'Tag and Category', component: PostList},
    {path: '/post/', name: 'PostList', component: PostList},
    {path: '/search/:id', name: 'Search', component: PostList},
    {path: '/about', name: 'About', component: About},
    {path: '/login', name: 'Login', component: Login},
    {path: '/codeforces', name: 'Cf', component: Codeforces},
    {path: '/publications', name: 'Pub', component: Pub},

    {path: '/:pathMatch(.*)*', name: 'NotFound', component: Home},

];

const router = createRouter({
    // history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(),
    routes
});
// createApp(App).mount('#app')
createApp(App).use(router).mount('#app')
