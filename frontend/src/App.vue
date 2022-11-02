<script setup>
import {RouterView} from 'vue-router'
import NavBar from "@/components/NavBar.vue";
import Logout from "@/components/Logout.vue";
import Footer from "@/components/Footer.vue";
import TopBar from "@/components/TopBar.vue";
</script>

<template>

  <body id="page-top">

  <!-- Page Wrapper -->
  <div id="wrapper">

    <NavBar/>

    <!-- Content Wrapper -->
    <div id="content-wrapper" class="d-flex flex-column">

      <TopBar/>

      <!-- VueJS Application Pane -->
      <RouterView/>

      <Footer/>

    </div>
    <!-- End of Content Wrapper -->

  </div>
  <!-- End of Page Wrapper -->

  <!-- Scroll to Top Button-->
  <a class="scroll-to-top rounded" href="#page-top">
    <i class="fas fa-angle-up"></i>
  </a>

  <Logout/>

  </body>

</template>