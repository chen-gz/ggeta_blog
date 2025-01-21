<script lang="ts" setup>
import { ref } from "vue";
import { SearchPostsRequestV4, searchPostsV4, V4PostData } from "/apiv4.js";
import { formatDate } from "/ui_utils";

let article = ref([] as V4PostData[]);
let param = {} as SearchPostsRequestV4;
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
        <a v-for="post in article" :key="post.id" class="text-decoration-none mb-4"
            style="&:hover backg round-color: #f6f6f6;" @click="$router.push('/post/' + post.url)">

            <s-card clickable="true" style="max-width: none; width: 100%;">
                <div slot="headline">{{ post.title }}</div>
                <div slot="text">{{ post.summary }}</div>
                <div class="home-card-bottom" slot="text" >
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
            </s-card>
        </a>
    </div>
</template>

<style lang="scss" scoped>
</style>
