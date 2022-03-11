<template>
  <b-navbar toggleable="lg" type="dark" variant="primary">
    <div class="container">
      <nuxt-link to="/" class="navbar-brand">{{ $store.state.appName }}</nuxt-link>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <nuxt-link to="/" class="nav-link">Home</nuxt-link>
        </b-navbar-nav>

        <b-navbar-nav class="ms-auto">
          <template v-if="$auth.user">
            <nuxt-link to="/profile" class="nav-link">Profile</nuxt-link>
            <a href="" class="nav-link" @click.prevent="logout">Logout</a>
          </template>
          <template v-else>
            <nuxt-link to="/login" class="nav-link">Login</nuxt-link>
            <nuxt-link to="/register" class="nav-link">Register</nuxt-link>
          </template>
        </b-navbar-nav>
      </b-collapse>
    </div>
  </b-navbar>
</template>

<script>
export default {
  methods: {
    logout() {
      this.$auth.logout()
    }
  }
}
</script>