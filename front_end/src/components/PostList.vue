<script lang="ts" setup>
import {ref, watch} from "vue";
import {useRouter} from "vue-router";
import {formatDate, SearchPostsRequestV4, SearchPostsResponseV4, searchPostsV4} from "/apiv4";

// let searchParam = SearchPostsRequestV4
const route = useRouter();
// url is tags/:id
// url can be "tag/:id" or "cate/:id" or "post/
// parse the url first
/// watch the url change
let res = ref({} as SearchPostsResponseV4)

function page_init() {
    let tag_cate_post = route.currentRoute.value.path.split("/")[1]
    let url_id = route.currentRoute.value.path.split("/")[2]
    console.log("call list page" + tag_cate_post + " " + url_id)

    let searchParam = {} as SearchPostsRequestV4
    if (tag_cate_post == "tag") {
        searchParam.tags = url_id
        searchParam.sort = "created_at DESC"
    } else if (tag_cate_post == "cate") {
        searchParam.categories = url_id
        searchParam.sort = "created_at DESC"
    } else if (tag_cate_post == "search") {
        searchParam.content = url_id
        searchParam.sort = "created_at DESC"
    } else {
        searchParam.sort = "created_at DESC"
    }
    // searchParam.limit = {start: 0, size: 20}

    let params = searchParam
    console.log(params)
    searchPostsV4(params).then((response) => {
        res.value = response
        console.log(res.value)
        if (res.value.number_of_posts == 0) {
            res.value.posts = []
        }
        for (let i = 0; i < res.value.posts.length; i++) {
            res.value.posts[i].created_at = new Date(res.value.posts[i].created_at);
        }
        console.log(res.value)
    })
}

page_init();

watch(route.currentRoute, (to, from) => {
    page_init();
})

</script>

<template>
    <div class="post-list">
        <h1>Post List</h1>
        <ul>
            <li v-for="result in res.posts" :key="result.id" @click="route.push('/post/' + result.url)">
                <a href="#">
                    {{ result.title }}
                </a>
                <time>
                    {{ formatDate(result.updated_at) }}
                </time>
            </li>
        </ul>
    </div>
</template>

<style lang="sass" scoped>
.post-list
    padding: 1rem
    h1
        font-size: 2rem
        margin-bottom: 1rem
        color: #333

ul
    list-style: none
    padding: 0

li
    display: flex
    align-items: center
    justify-content: space-between
    width: 100%
    margin-bottom: 20px
    padding: 15px
    border-radius: 5px
    transition: background-color 0.3s, box-shadow 0.3s

    &:hover
        background-color: #f9f9f9
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1)

a
    font-size: 1.2rem
    color: #007bff
    text-decoration: none
    font-weight: bold
    transition: color 0.3s

    &:hover
        color: #0056b3
        text-decoration: underline

time
    font-size: 0.9rem
    color: #666
    font-style: italic
</style>