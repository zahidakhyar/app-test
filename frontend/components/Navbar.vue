<template>
  <b-navbar toggleable="lg" type="dark" variant="primary">
    <div class="container">
      <nuxt-link to="/" class="navbar-brand">NavBar</nuxt-link>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <nuxt-link to="/" class="nav-link">Home</nuxt-link>
        </b-navbar-nav>

        <b-navbar-nav class="ms-auto">
          <template v-if="$auth.user">
            <b-nav-item-dropdown right>
              <template #button-content>
                <em>User</em>
              </template>
              <nuxt-link to="/auth/profile" class="nav-link">Profile</nuxt-link>
              <nuxt-link to="/auth/logout" class="nav-link">Logout</nuxt-link>
            </b-nav-item-dropdown>
          </template>
          <template v-else>
            <nuxt-link to="/auth/login" class="nav-link">Login</nuxt-link>
          </template>
        </b-navbar-nav>
      </b-collapse>
    </div>
  </b-navbar>
</template>
