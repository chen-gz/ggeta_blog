<script setup>
import Topbar from "@/components/main-wrapper/topbar.vue";
import MainContent from "@/components/main-wrapper/content.vue";
import RightPanel from "@/components/main-wrapper/right-panel.vue";
</script>

<template>
    <div id="main-wrapper-inner">
        <div id="main-wrapper-topbar-wrapper-outer">
            <topbar />
        </div>
        <div id="main-wrapper-main" class="d-flex flex-row">
            <div id="main-content-wrapper-outer">
                <main-content />
            </div>
            <div id="main-wrapper-right-panel-outer">
                <right-panel />
            </div>
        </div>
    </div>
</template>

<style scoped lang="sass">
@use "@/assets/main-wrapper.sass" as *

#main-wrapper-inner
    width: 100%
    padding-left: 0.75rem
    //padding-right: 0.75rem
    height: 100%

#main-wrapper-topbar-wrapper-outer
    height: $topbar-height
    width: 100%
    padding-right: 1rem
    padding-left: 1rem
    //margin-bottom: 8px

#main-wrapper-main
    top: $topbar-height
    overflow: hidden
    height: calc(100% - $topbar-height - 15px)
    width: 100%

#main-content-wrapper-outer
    //width: calc(100% - #{$right-panel-width})
    flex-grow: 1
    height: 100%
    overflow: hidden
    //background-color: #f5f5f5
    z-index: 1
    //flex-grow: 0

#main-wrapper-right-panel-outer
    width: $right-panel-width
    flex-shrink: 0
@media (max-width: 1200px)
    #main-wrapper-right-panel-outer
        width: 0
        display: none
</style>
