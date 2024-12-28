<script lang="ts" setup>
import { ref } from "vue";
import { SearchPostsRequestV4, searchPostsV4, V4PostData } from "/apiv4.js";
import { formatDate } from "/ui_utils";

let article = ref([] as V4PostData[]);
// get 5 latest articles
// searchPostsV4

let param = {} as SearchPostsRequestV4;
// param.limit = {start: 0, count: 5};
// param.constructor();
// searchPostsV4()
param.limit = { start: 0, size: 10 };
param.sort = "created_at DESC";
searchPostsV4(param).then((res) => {
    // console.log(res)
    article.value = res.posts;
    for (let i = 0; i < article.value.length; i++) {
        if (article.value[i].summary.length == 0) {
            // get summary from content
            article.value[i].summary = article.value[i].content.substring(
                0,
                300,
            );
        }
    }
});

// sort: "created_at DESC"
</script>

<template>
    <div id="home-inner" class="d-flex flex-column pt-4">
        <a
            v-for="post in article"
            :key="post.id"
            class="text-decoration-none mb-4"
            style="&:hover backg round-color: #f6f6f6;"
            @click="$router.push('/post/' + post.url)"
        >
            <div class="card pt-2 px-3 pb-2 border-0 home-card">
                <h4
                    class="card-title home-card-title"
                    style="color: rgb(42, 42, 42)"
                >
                    {{ post.title }}
                </h4>
                <p class="home-summary" style="color: rgb(52, 52, 60)">
                    {{ post.summary }}
                </p>
                <div class="home-card-bottom">
                    <span>
                        <i class="fa-fw fas fa-calendar-alt"></i>
                        {{ formatDate(post.created_at) }}
                    </span>
                    <span v-if="post.category != ''" class="ps-3">
                        <i class="fa-fw fas fa-folder"></i> {{ post.category }}
                    </span>
                    <span v-if="post.tags != ''" class="ps-3">
                        <i class="fa-fw fas fa-tags"></i> {{ post.tags }}</span
                    >
                </div>
            </div>
        </a>
    </div>
</template>

<style lang="scss" scoped>
.home-card {
    border-radius: 0.5rem;
    box-shadow:
        0 0 0 1px rgba(0, 0, 0, 0.125),
        0 2px 3px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s;
    /* &:hover { */
    box-shadow:
        0 0 0 1px rgba(0, 0, 0, 0.125),
        0 4px 6px rgba(0, 0, 0, 0.1);
    /* style="&:hover background-color: #f6f6f6;" */
    &:hover {
        background-color: #f6f6f6;
    }
}
</style>
