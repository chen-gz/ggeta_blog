<template>
  <div class="home">
    <header>
      <h1>Guangzong Chen</h1>
      <p>Ph.D. Candidate in Electrical and Computer Engineering</p>
    </header>

    <div class="intro">
      <img src="../assets/profile.jpg" alt="Guangzong Chen" class="profile-photo">
      <p>Hello! I'm Guangzong Chen, currently pursuing a Ph.D. in Electrical and Computer Engineering at the University
        of Pittsburgh. I hold a B.S. in Automation from the Beijing Institute of Technology.</p>
    </div>

    <section>
      <h3>Research and Experience</h3>
      <p>As a Research Assistant at the University of Pittsburgh, I have worked on innovative projects such as:</p>
      <s-card style="display:flex; max-width: none;" class="mb-4">
        <a href="https://arxiv.org/abs/2411.10368" target="_blank" style="margin-right: 20px;">
          <!-- <img src="https://arxiv.org/pdf/2411.10368v1" alt="Paper Image" style="width: 100px; height: 100px;"> -->
          <img src="https://arxiv.org/pdf/2411.10368v1" alt="Paper Thumbnail" style="width:200px;">
        </a>
        <div>
          <strong>Neural Network Analysis:</strong> Investigated the relationship between GAN and Autoencoder for
          enhanced explainability in AI. <br>
          <br>
          <em>
            Generative Adversarial Networks (GANs) are widely used in image-to-image translation. This paper proposes a
            streamlined image-to-image translation network with a simpler architecture than existing models. We
            investigate the relationship between GANs and autoencoders and explain why using only the GAN component for
            image translation is effective. We show that adversarial GAN models yield comparable results to existing
            methods without additional complex loss penalties. We also provide experimental results to validate our
            findings.</em>
        </div>
      </s-card>

      <s-card style="display:flex; max-width: none;" class="mb-4">
        <a href="https://openaccess.thecvf.com/content/CVPR2024W/MTF/html/Chen_Shape-Preserving_Generation_of_Food_Images_for_Automatic_Dietary_Assessment_CVPRW_2024_paper.html" target="_blank" style="margin-right: 20px;">
          <!-- <img src="https://arxiv.org/pdf/2411.10368v1" alt="Paper Image" style="width: 100px; height: 100px;"> -->
          <img src="https://openaccess.thecvf.com/content/CVPR2024W/MTF/papers/Chen_Shape-Preserving_Generation_of_Food_Images_for_Automatic_Dietary_Assessment_CVPRW_2024_paper.pdf" alt="Paper Thumbnail" style="width:200px;">
        </a>
        <div>

        <strong>Constrained Food Image Generation:</strong> Developed a generative model using GANs for automatic
          dietary assessment, creating dataset for same volumn of food.<br>
          <br>
          <em>
            Recognizing foods and estimating volumes from images are key for automatic dietary assessment. However, training images labeled with food names and volumes are unavailable. Recent studies suggest artificially generating training images using GANs. However, conveniently generating large amounts of food images with known volumes remains a challenge. This work presents a simple GAN-based neural network architecture for conditional food image generation. The generated images closely resemble the reference input image, demonstrating realism and shape-preserving capabilities.
          </em>
        </div>
      </s-card>
      <ul>
        <!-- align-items:center;  -->

        <li><strong>Food Image Classification:</strong> Created a system for classifying African food images with 90%
          accuracy using probabilistic inference, SVM, and neural networks.</li>
      </ul>
      <p>During my internship at Nimbus Robotics, I contributed to:</p>
      <ul>
        <li><strong>Motor Controller Development:</strong> Designed a motor controller using Field-Oriented Control
          (FOC) for wearable devices.</li>
        <li><strong>Version Control System:</strong> Collaborated on a Git-based system improving workflow efficiency by
          ~20%.</li>
      </ul>
    </section>

    <section>
      <h3>Technical Skills</h3>
      <ul>
        <li><strong>Programming Languages:</strong> C/C++, Python, MATLAB, Rust</li>
        <li><strong>Specialties:</strong> Embedded Systems, Algorithms, RTOS, Circuit Design, Communication Protocols
          (I2C, SPI, UART, USB, eMMC)</li>
      </ul>
    </section>

    <section>
      <h3>Publications</h3>
      <ul>
        <li> "Mechanisms of Generative Image-to-Image Translation Networks." 10.48550/arXiv.2411.10368., 2024 </li>
        <li>"Shape-Preserving Generation of Food Images for Automatic Dietary Assessment," CVPR, 2024.</li>
        <li>“Image-based volume estimation for food in a bowl,” Journal of Food Engineering, Jul. 2024.</li>
        <li>“Improved Wearable Devices for Dietary Assessment Using a New Camera System,” Sensors, Oct. 2022.</li>
        <li>“Food/Non-Food Classification of Real-Life Egocentric Images in Low- and Middle-Income Countries Based on
          Image Tagging Features,” Frontiers in Artificial Intelligence, Apr. 2021.</li>
      </ul>
    </section>

    <section>
      <h3>Personal Projects</h3>
      <ul>
      </ul>
      <s-card style="display:flex; max-width: none;" class="mb-4">
        <a href="https://github.com/chen-gz/stm32u5lib" target="_blank" style="margin-right: 20px;">
          <img src="https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png" alt="Project Thumbnail"
            style="width:75px;">
        </a>
        <div>
          <strong>STM32 HAL Development:</strong> Designed a user-friendly hardware abstraction layer in Rust for motor
          controllers.
          <!-- <strong>Neural Network Analysis:</strong> Investigated the relationship between GAN and Autoencoder for enhanced explainability in AI. <br> -->
          <br>
        </div>
      </s-card>
      <s-card style="display:flex; max-width: none;" class="mb-4">
        <a href="https://github.com/chen-gz/serial_comm_tool" target="_blank" style="margin-right: 20px;">
          <img src="https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png" alt="Project Thumbnail"
            style="width:75px;">
        </a>
        <div>
          <strong>Serial Communication Tool:</strong> Developed a tool using Rust for USART communication with superior
          stability and speed.
        </div>
      </s-card>
      <!-- https://github.com/chen-gz/serial_comm_tool -->
    </section>
  </div>
</template>

<script>
export default {
  name: 'About'
}
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
