<template>
  <div class="container margin-bottom-20">
    <div class="row error-page align-items-center">
      <div class="col">
        <h1 v-if="error.statusCode === 404">{{ $t('Page not found') }}</h1>
        <h1 v-else>{{ $t('An error occurred') }}</h1>
        <NuxtLink to="/">{{ $t('Home page') }}</NuxtLink>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  layout: 'default',
  props: {
    error: {
      type: Object,
      default: () => ({}),
    },
  },
}
</script>

<style scoped>
.error-page {
  text-align: center;
  height: 80vh;
}
</style>
