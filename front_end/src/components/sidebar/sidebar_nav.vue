<script lang="jsx" setup>
// let from_jsx = "from jsx";
// const hello = (<div>hello); // {from_jsx}</div>);
// export hello;
</script>

<template>
    <nav class="flex-column flex-grow-1">
        <div class="nav-item active">
            <a class="nav-link" href="/">
                <i class="fa-fw fas fa-home" /> <span>HOME</span>
            </a>
        </div>
        <!--        <div class="nav-item"> <a class="nav-link" href="/cate/"> <i class="fa-fw fas fa-stream"/> <span>CATEGORIES</span> </a> </div>-->
        <div class="nav-item">
            <a class="nav-link" href="/tag/">
                <i class="fa-fw fas fa-tags"></i> <span>TAGS</span>
            </a>
        </div>
        <div class="nav-item">
            <a class="nav-link" href="/post/">
                <i class="fa-fw fas fa-archive"></i> <span>ARCHIVES</span>
            </a>
        </div>
        <div class="nav-item">
            <a class="nav-link" href="/about/">
                <i class="fa-fw fas fa-info-circle"></i> <span>ABOUT</span>
            </a>
        </div>
    </nav>
</template>

<style lang="sass" scoped>
@import "@/assets/variables"
// define font size and line height
$font-size: 0.9rem
$line-height: 0.8
nav
    width: 100%
    //background-color: aqua
    //z-index: 4
    //padding-left: 2.5rem
    //padding-right: 1.5rem
    padding-bottom: 1.5rem
    //padding-top: 1.2rem
    overflow-y: auto
    font-family: "Droid Sans", sans-serif
    scrollbar-width: none

nav::-webkit-scrollbar
    display: none

.nav-item
    //margin-top: 2rem
    position: relative
    padding-right: 1rem
    padding-left: 1rem
    padding-top: 0.4rem
    padding-bottom: 0.4rem

    //font-size: 1.2rem
    //background-color: #03b303
    align-content: center
    font-weight: 600
    font-size: $font-size
    font-family: $font-family-heading
    //add space between <i> and <span>
    text-decoration: none

    a
        opacity: 0.6
        padding-top: 0.6rem
        padding-bottom: 0.6rem
        display: flex
        align-items: center
        text-decoration: none
        color: var(--sidebar-text-color)
        padding-left: 1.5rem


        span
            letter-spacing: 0.2px
            padding-left: 1.5rem
            line-height: $line-height

    i
        //margin-right: 2rem
        width: $line-height
        height: $line-height
</style>
