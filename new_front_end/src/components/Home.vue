<template>
  <div class="home">
    <header>
      <h1>Guangzong Chen</h1>
      <p>Ph.D. Candidate in Electrical and Computer Engineering</p>
    </header>

    <div class="intro">
      <img src="../assets/profile.jpg" alt="Guangzong Chen" class="profile-photo">
      <p>Hello! I'm Guangzong Chen, currently pursuing a Ph.D. in Electrical and Computer Engineering at the University of Pittsburgh.</p>
    </div>
    <section style="display: flex; flex-direction: column">
<!--      <div id="home-inner" class="d-flex flex-column pt-4">-->
        <a v-for="post in article" :key="post.id" class="text-decoration-none"
           style="&:hover { background-color: #f6f6f6; }; margin-left: 150px;" @click="$router.push('/post/' + post.url)">
            <div class="headline">{{ post.title }}</div>
        </a>
<!--      </div>-->
    </section>

  </div>
</template>

<!--<script>-->

<script lang="ts" setup>
// export default {
//   name: 'Home'
// }

import { ref } from "vue";
// import { SearchPostsRequestV4, searchPostsV4, V4PostData } from "apiv4.js";
import { formatDate } from "/ui_utils";
import  { SearchPostsRequestV4, searchPostsV4, V4PostData } from "/apiv4.js";

let article = ref([] as V4PostData[]);
let param = {} as SearchPostsRequestV4;
param.limit = { start: 0, size: 10 };
param.sort = "created_at DESC";
searchPostsV4(param).then((res) => {
  // console.log(res)
  article.value = res.posts;
  console.log(article.value)
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
</script>

<style scoped>
.home {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  color: #333;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

header {
  text-align: center;
  margin-bottom: 20px;
}

header h1 {
  margin: 0;
  font-size: 2.5em;
  color: #2c3e50;
}

header p {
  margin: 0;
  font-size: 1.2em;
  color: #7f8c8d;
}

.intro {
  display: flex;
  align-items: center;
  margin-bottom: 20px;
}

.profile-photo {
  border-radius: 50%;
  width: 150px;
  height: 150px;
  margin-right: 20px;
}

h3 {
  color: #2c3e50;
  margin-top: 20px;
}

ul {
  list-style-type: disc;
  margin-left: 20px;
}

ul li {
  margin-bottom: 10px;
}

</style>