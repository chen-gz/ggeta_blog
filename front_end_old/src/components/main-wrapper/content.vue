<script setup></script>

<template>
    <div id="content-wrapper-inner">
        <router-view></router-view>
    </div>
</template>

<style scoped lang="sass">
#content-wrapper-inner

    padding-left: 1rem
    padding-right: 1rem
    //padding-top: 1rem
    overflow-x: scroll
    overflow-y: scroll
    height: calc(100%)
    //width: 100%
    width: calc(100% - 1rem - 1rem)
    //width: 100%
    scrollbar-width: none
    max-width: 1000px
    margin-left: auto
    margin-right: auto
    @media (max-width: 768px)
        padding-left: 0
        padding-right: 0
        //width: 100%
#content-wrapper-inner::-webkit-scrollbar
    display: none
</style>
