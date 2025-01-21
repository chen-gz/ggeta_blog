<script setup>

</script>

<template>
  <div class="sidebar-bottom d-flex flex-row align-items-center">
    <button aria-label="Switch Mode" class="mode-toggle btn" type="button">
      <i class="fas fa-adjust">
      </i>
    </button>
    <a aria-label="github" href="https://github.com/chen-gz"
       rel="noopener noreferrer"
       target="_blank"> <i class="fab fa-github"></i> </a> <a
      aria-label="email" href="javascript:location.href = 'mailto:' + ['chen','ggeta.com'].join('@')">
    <i class="fas fa-envelope"></i>
  </a>
    <a aria-label="rss" href="/feed.xml">
      <i class="fas fa-rss"></i>
    </a>
  </div>

</template>

<style scoped lang="sass">

.sidebar-bottom
  // move to bottom of the sidebar by growing the previous nav
  //margin-bottom: 1.2rem
  padding: 1.5rem 2.5rem 1.5rem 1rem
  z-index: 6
  //background-color: crimson
  display: flex
  justify-content: space-around
  align-items: center

$icon-size: 1.5rem
i
  line-height: $icon-size
  width: $icon-size
  height: $icon-size
  text-align: center
  font-size: $icon-size

$btn-size: 1.75rem
button
  display: flex
  text-align: center
  justify-content: center
  align-items: center
  width: 1.75rem
  height: 1.75rem
  line-height: 1.75rem
  border-radius: 50%
  background-color: inherit
  border: none
a
  text-decoration: none
  color: inherit
  opacity: 0.6

</style>