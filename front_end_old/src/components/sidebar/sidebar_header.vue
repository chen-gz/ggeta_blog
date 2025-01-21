<script setup></script>

<template>
    <header class="profile-wrapper">
        <a href="/" id="avatar" class="rounded-circle">
            <img
                src="https://minio.ggeta.com/blog-public-data/mine_square.jpg"
                width="112"
                height="112"
                alt="avatar"
                onerror="this.style.display='none'"
        /></a>
        <h1 class="title"><a href="/">Guangzong</a></h1>
        <p class="subtitle fst-italic mb-0">Knowledge and Live</p>
    </header>
</template>

<style scoped lang="sass">
//@import "src/assets/variables"
@use "@/assets/sidebar" as *

.profile-wrapper
    margin: 3rem 1.5rem 2.5rem 2.5rem

.rounded-circle
    border-radius: 50%
    overflow: hidden
    display: block
    width: 112px
    height: 112px
    img
        width: 100%
        height: 100%
        object-fit: cover
        object-position: center
        border-radius: 50%
        &:hover
            box-shadow: 0 0 0 2px #007bff

.title
    margin-top: 1rem
    margin-bottom: 0.5rem
    font-weight: $sidebar-title-font-weight
    a
        font-family: $sidebar-title-font-family
        font-size: $sidebar-title-font-size
        font-weight: $sidebar-title-font-weight
        //line-height: $sidebar-title-line-height
        line-height: 34px
        color: $sidebar-title-color
        text-decoration: none

.subtitle
    font-family: $sidebar-subtitle-font-family
    font-size: $sidebar-subtitle-font-size
    font-weight: $sidebar-subtitle-font-weight
    line-height: $sidebar-subtitle-line-height
    color: $sidebar-subtitle-color
    font-style: $sidebar-subtitle-font-style
    margin-bottom: 0
    margin-top: 0
</style>
