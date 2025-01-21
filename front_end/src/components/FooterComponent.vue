<template>
  <!-- <footer class="footer">
    <p>&copy; 2023 Your Company. All rights reserved.</p>
  </footer> -->
  <div class="wrapper">
  <footer>
    <p>Contact me: <a href="mailto:guangzong@pitt.edu">
    guangzong@pitt.edu</a> | <a href="mailto:chen@ggeta.com">chen@ggeta.com</a></p>
      <p>Follow me on <a href="https://www.linkedin.com/in/guangzong-chen">LinkedIn</a></p>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'FooterComponent'
}
</script>

<style scoped>
.wrapper {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  color: #333;
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  padding: auto;
}
.footer {
  background-color: #2c3e50;
  color: #ecf0f1;
  text-align: center;
  padding: 20px 0;
  position: fixed;
  width: 100%;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
}
footer {
  text-align: center;
  margin-top: 20px;
  color: #7f8c8d;
}

footer a {
  color: #3498db;
  text-decoration: none;
}

footer a:hover {
  text-decoration: underline;
}

</style>