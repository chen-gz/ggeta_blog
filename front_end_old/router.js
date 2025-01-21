import {createRouter, createWebHistory} from 'vue-router';
import PostPage from "@/components/main-wrapper/content/post.vue";
import Tags from "@/components/main-wrapper/content/tags.vue";
import PostList from "@/components/main-wrapper/content/post-list.vue";
import Login from "@/components/main-wrapper/content/login.vue";
import Home from "@/components/main-wrapper/content/home.vue";
import PostEdit from "@/components/main-wrapper/content/post-edit.vue";
import Category from "@/components/main-wrapper/content/category.vue";

const routes = [
    {path: '/', name: 'Home', component: Home},
    {path: '/tag/', name: 'Tags', component: Tags},
    {path: '/cate/', name: 'Categories', component: Category},
    {path: '/post/:id', name: 'PostPage', component: PostPage},
    {path: '/post_edit/:id', name: 'PostEdit', component: PostEdit},
    {path: '/login', name: 'Login', component: Login},
    {path: '/tag/:id', name: 'Tag and Category', component: PostList},
    {path: '/post/', name: 'PostList', component: PostList},
    {path: '/search/:id', name: 'Search', component: PostList},
];

const router = createRouter({
    // history: createWebHistory(process.env.BASE_URL),
    history: createWebHistory(),
    routes
});

export default router;

