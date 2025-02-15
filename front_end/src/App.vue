<script setup lang="ts">
// import home from './compon
// ents/About.vue';
import FooterComponent from './components/FooterComponent.vue';
import { newPostV4 } from "/apiv4";
import { useRouter } from "vue-router";
const router = useRouter()
// f3 to create a new post
function new_post() {
    console.log("new post")
    // create a new post and redirect to the new post
    newPostV4().then((response) => {
        if (response.status == "success") {
            console.log(response)
            router.push(`/post_edit/${response.url}`)
        } else {
            // alert("failed to save post, login is required");
            router.push("/login");
        }
    })

}
document.addEventListener("keydown", function (e) {
    // f3 to create a new post
    if (e.key === "F3") {
        e.preventDefault();
        console.log("F3 create new post")
        new_post()
    }
});
function performSearch() {
    console.log("search")
    let search = document.getElementById("search") as HTMLInputElement
    console.log("search: " + search.value)
    if (search.value.length > 0) {
        router.push(`/search/${search.value}`)
    }
}
function navigateToChat() {
    window.location.href = 'https://chat.ggeta.com';
}
</script>

<template>

    <s-appbar style="background-color: white;">
        <!--左侧菜单按钮-->
        <!-- <s-icon-button slot="navigation">
          <s-icon name="menu"></s-icon>
        </s-icon-button> -->
        <!--标题-->
        <div slot="headline" @click="$router.push('/')"> Ggeta </div>
        <s-icon name="dark_mode" slot="headline"></s-icon>
        <!-- <s-icon name="home"></s-icon> -->
        <!--右侧操作按钮-->
        <!-- <s-search></s-search> -->
        <s-search placeholder="search" @keyup.enter="performSearch" @click="performSearch" id="search">
            <s-icon name="search" slot="start"></s-icon>
            <s-icon-button slot="end">
                <s-icon name="close"></s-icon>
            </s-icon-button>
        </s-search>
        <s-button type="text" @click="navigateToChat"> 
         Chat </s-button>
        <s-button type="text" @click="$router.push('/shows/');"> Videos </s-button>
        <s-button type="text" @click="$router.push('/publications/');"> Pubs </s-button>
        <s-button type="text" @click="$router.push('/codeforces/');"> Cf </s-button>
        <s-button type="text" @click="$router.push('/tag/');"> Tags </s-button>
        <s-button type="text" @click="$router.push('/post/');"> Archives </s-button>
        <s-button type="text" @click="$router.push('/about');"> About </s-button>
        <!-- <s-button type="text"> Series </s-button> -->
    </s-appbar>
    <div class="p-4">
        <!-- <home/> -->
        <router-view />
        <FooterComponent />
    </div>
</template>

<style scoped>
/* .logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
} */
</style>
