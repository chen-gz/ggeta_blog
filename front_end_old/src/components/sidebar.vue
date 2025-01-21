<script setup>
import Sidebar_header from "@/components/sidebar/sidebar_header.vue";
import Sidebar_nav from "@/components/sidebar/sidebar_nav.vue";
import Sidebar_bottom from "@/components/sidebar/sidebar_bottom.vue";
</script>

<template>
    <div id="sidebar-wrapper-inner" class="d-flex flex-column">
        <sidebar_header />
        <sidebar_nav />
        <sidebar_bottom />
    </div>
</template>

<style scoped lang="sass">
// import the variables from the variables.saas file
@use "@/assets/variables" as *
@use "@/assets/typography-light.scss" as *
%sidebar-link-hover
    &:hover
        color: var(--sidebar-active-color)
a
    @extend %sidebar-link-hover
#sidebar-wrapper-inner
    //position: fixed
    @include light-scheme
    //width: calc($sidebar-width - 1px) // remove border size
    border-right: 1px solid var(--sidebar-border-color)
    height: 100%
    background-color: var(--sidebar-bg)
    font-family: $font-family-heading
    // high scroll bar

//#content-wrapper-inner::-webkit-scrollbar
//    display: none
</style>
