<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
  <!-- Sidebar -->
  <ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">

    <!-- Sidebar - Brand -->
    <a class="sidebar-brand d-flex align-items-center justify-content-center" href="../../index.html">
      <div class="sidebar-brand-icon rotate-n-15">
        <i class="fas fa-laugh-wink"></i>
      </div>
      <div class="sidebar-brand-text mx-3">Pingo</div>
    </a>

    <!-- Divider -->
    <hr class="sidebar-divider my-0">

    <!-- Nav Item - Dashboard -->
    <li class="nav-item">
      <RouterLink class="nav-link" to="/">
        <i class="fas fa-fw fa-tachometer-alt"></i>
        <span>Dashboard</span>
      </RouterLink>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider">

    <!-- Heading -->
    <div class="sidebar-heading">
      Settings
    </div>

    <!--      TODO: Look for a reason to use a collapsible menu item, it's neat!-->
    <!--      &lt;!&ndash; Nav Item - Pages Collapse Menu &ndash;&gt;-->
    <!--      <li class="nav-item">-->
    <!--        <a class="nav-link collapsed" href="#" data-toggle="collapse" data-target="#collapsePages"-->
    <!--           aria-expanded="true" aria-controls="collapsePages">-->
    <!--          <i class="fas fa-fw fa-folder"></i>-->
    <!--          <span>Pages</span>-->
    <!--        </a>-->
    <!--        <div id="collapsePages" class="collapse" aria-labelledby="headingPages" data-parent="#accordionSidebar">-->
    <!--          <div class="bg-white py-2 collapse-inner rounded">-->
    <!--            <h6 class="collapse-header">Login Screens:</h6>-->
    <!--            <a class="collapse-item" href="login.html">Login</a>-->
    <!--            <a class="collapse-item" href="register.html">Register</a>-->
    <!--            <a class="collapse-item" href="forgot-password.html">Forgot Password</a>-->
    <!--            <div class="collapse-divider"></div>-->
    <!--            <h6 class="collapse-header">Other Pages:</h6>-->
    <!--            <a class="collapse-item" href="404.html">404 Page</a>-->
    <!--            <a class="collapse-item" href="blank.html">Blank Page</a>-->
    <!--          </div>-->
    <!--        </div>-->
    <!--      </li>-->

    <!-- Nav Item - Targets -->
    <li class="nav-item">
      <RouterLink class="nav-link" to="/target">
        <i class="fa-solid fa-bullseye"></i>
        <span>Targets</span>
      </RouterLink>
    </li>

    <!-- Divider -->
    <hr class="sidebar-divider d-none d-md-block">

    <!-- Sidebar Toggler (Sidebar) -->
    <div class="text-center d-none d-md-inline">
      <button class="rounded-circle border-0" id="sidebarToggle"></button>
    </div>

  </ul>
  <!-- End of Sidebar -->
</template>