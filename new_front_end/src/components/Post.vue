<template>
    <div id="post-inner-wrapper">
        <div class="post_page mt-2">
            <h1>{{ post.title }}</h1>
            <div class="post-meta text-muted">
                <div class="post-header mt-2" style="font-size: 0.8rem">
                    <span> Posted on </span>
                    <span style="font-weight: 400" class="me-2">
                        {{
                            new Date(post.created_at).toLocaleDateString(
                                "en-US",
                                {
                                    year: "numeric",
                                    month: "short",
                                    day: "numeric",
                                },
                            )
                        }}
                    </span>
                    <span> Updated on </span>
                    <span class="before_dot" style="font-weight: 400">
                        {{
                            new Date(post.updated_at).toLocaleDateString(
                                "en-US",
                                {
                                    year: "numeric",
                                    month: "short",
                                    day: "numeric",
                                },
                            )
                        }}</span
                    >
                    <div class="sub-meta">
                        <span>By</span>
                        <span class="author" style="font-weight: 400">{{
                            post.author
                        }}</span>
                        <!--          add edit button-->
                        <button @click="route.push('/post_edit/' + post.url)">
                            edit
                        </button>
                        <button @click="CalldeletePost">delete</button>
                    </div>
                </div>
            </div>
            <div id="markdown_content" v-html="post_content"></div>
        </div>
    </div>
</template>

<style lang="scss" scoped>
/* .post_page {
    margin-top: 1.5rem;
    h1 {
        margin-bottom: 1rem;
        font-weight: 400;
        color: rgb(42, 42, 42);
        -webkit-font-smoothing: antialiased;
    }
} */

/* .post-header {
    margin-top: 2rem;
    margin-bottom: 20px;
    display: block;
    color: rgb(117, 117, 117);
    font-size: 0.8rem;
    font-weight: 500;
} */

.sub-meta {
    margin-top: 8px;
    margin-bottom: 20px;
    display: block;
    color: rgb(117, 117, 117);
    font-size: 0.8rem;
    button {
        width: 4rem;
        background-color: white;
        border: 1px solid #d1d5da;
        //round
        border-radius: 6px;
    }
    & > span,
    button {
        margin-right: 10px;
        &:first-child {
            margin-right: 3px;
        }
        &:last-child {
            margin-right: 0;
        }
    }
}
</style>

<script lang="ts" setup>
import { nextTick, onMounted, ref, watch, watchEffect } from "vue";
import { getPostV4, V4PostData, deletePost } from "/apiv4";
import { useRouter } from "vue-router";

const route = useRouter();
let url = route.currentRoute.value.params.id;
// console.log(url)

let post = ref({} as V4PostData);
let post_content = ref("");
let post_toc = ref("");

function CalldeletePost() {
  deletePost(post.value).then((response) => {
    if (response) {
      // console.log("delete success");
      alert("delete success");
      route.push("/post");
    } else {
      // console.log("delete failed");
      alert("delete failed");
    }
  });
}

getPostV4(url, true).then((response) => {
    // console.log(response)
    post.value = response.post;
    post_content.value = response.html;
    // get toc from post content they are surrended by <nav> tag
    // find the first <nav> tag
    let nav_start = post_content.value.indexOf("<nav");
    let nav_end = post_content.value.indexOf("</nav>");
    if (nav_start == -1 || nav_end == -1) {
        post_toc.value = "";
    } else {
        post_toc.value = post_content.value.substring(nav_start, nav_end + 6);
        post_content.value = post_content.value.substring(nav_end + 6);
    }
    post_toc.value = post_toc.value.substring(16, post_toc.value.length - 19);
    // console.log(post_toc.value)
});

watch(post_content, () => {
    console.log("post changed");
    // rerender whole page 
    // const component = window.sober.getComponent('componentId');

    // 重新渲染该组件
    // component.render();

    nextTick(() => {
    
        // @ts-ignore
        // window.MathJax.Hub.Queue(["Typeset", window.MathJax.Hub]);
        // tocbot.run();
        window.MathJax.startup.defaultReady();

        // @ts-ignore
        window.hljs.highlightAll();
        // @ts-ignore
        if (typeof window.tocbot !== "undefined") {
            // @ts-ignore
            window.tocbot.init({
                tocSelector: "#toc",
                contentSelector: "#markdown_content",
                ignoreSelector: "[data-toc-skip]",
                headingSelector: "h2, h3, h4",
                orderedList: false,
                scrollSmooth: false,
            });
            // @ts-ignore
            window.tocbot.refresh();
            console.log("tocbot is defined");
        } else {
            console.error("tocbot is not defined");
        }

        let code_blocks = document.querySelectorAll("pre code");
        //  = "relative";
        console.log("code block", code_blocks);
        // each code block position should be relative
        code_blocks.forEach((block) => {
            if (block.parentElement) {
                block.parentElement.style.position = "relative";
                block.style.backgroundColor = "#f8f9fa";
                console.log("set position relative for code block");
                console.log(block.parentElement);
                
            }
        });
        code_blocks.forEach((block) => {
            let copy_button = document.createElement("s-tooltips");
            copy_button.style.position = "absolute";
            copy_button.style.right = "10px";
            copy_button.style.top = "10px";
            // set copy icon from font awesome
            // copy_button.innerHTML = '<i class="fa fa-copy"></i>';
            copy_button.innerHTML = '<s-icon-button> <s-icon><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 -960 960 960"> <path d="M360-240q-33 0-56.5-23.5T280-320v-480q0-33 23.5-56.5T360-880h360q33 0 56.5 23.5T800-800v480q0 33-23.5 56.5T720-240H360Zm0-80h360v-480H360v480ZM200-80q-33 0-56.5-23.5T120-160v-560h80v560h440v80H200Zm160-240v-480 480Z"></path> </svg></s-icon> </s-icon-button>'
            copy_button.className = "copy-button";
            copy_button.onclick = () => {
                navigator.clipboard.writeText(block.innerText).then(() => {
                    console.log("copied");
                });
                // update copy button icon to check mark
                // copy_button.innerHTML = '<i class="fa fa-check-circle"></i>';
                // wait some time and change back to copy icon
                // setTimeout(() => {
                //     copy_button.innerHTML = '<i class="fa-solid fa-copy"></i>';
                // }, 2000);
            };
            block.parentElement?.appendChild(copy_button);
        });

    });
});
</script>
